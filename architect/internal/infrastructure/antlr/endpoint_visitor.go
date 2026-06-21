package antlr

import (
	"fmt"
	"strings"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	endpoint_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/endpoint"
)

type endpointVisitor struct {
	*endpoint_generated.BaseEndpointVisitor
	err error
}

func (v *endpointVisitor) setErr(err error) {
	if err != nil && v.err == nil {
		v.err = err
	}
}

func (v *endpointVisitor) VisitProgram(ctx *endpoint_generated.ProgramContext) interface{} {
	return ctx.EndpointDecl().Accept(v)
}

func (v *endpointVisitor) VisitEndpointDecl(ctx *endpoint_generated.EndpointDeclContext) interface{} {
	builder := endpoints.NewEndpointBuilder().
		Name(ctx.IDENT().GetText()).
		Method(parseHTTPMethod(ctx.HttpMethod().GetText())).
		Path(unquote(ctx.STRING().GetText()))

	for _, body := range ctx.AllEndpointBody() {
		result := body.Accept(v)
		if v.err != nil {
			return nil
		}

		switch value := result.(type) {
		case endpoints.Input:
			builder.Input(value)

		case endpoints.Action:
			builder.AddAction(value)
		}
	}

	endpoint, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return endpoint
}

func (v *endpointVisitor) VisitEndpointBody(ctx *endpoint_generated.EndpointBodyContext) interface{} {
	if ctx.InputBlock() != nil {
		return ctx.InputBlock().Accept(v)
	}

	return ctx.Action_().Accept(v)
}

func (v *endpointVisitor) VisitInputBlock(ctx *endpoint_generated.InputBlockContext) interface{} {
	builder := endpoints.NewInputBuilder()

	for _, fieldCtx := range ctx.AllInputField() {
		field, ok := fieldCtx.Accept(v).(endpoints.InputField)
		if !ok || field == nil {
			v.setErr(fmt.Errorf("failed to parse input field"))
			return nil
		}

		builder.AddField(field)
	}

	input, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return input
}

func (v *endpointVisitor) VisitInputField(ctx *endpoint_generated.InputFieldContext) interface{} {
	typeRef, ok := ctx.TypeRef().Accept(v).(objects.TypeReference)
	if !ok || typeRef == nil {
		v.setErr(fmt.Errorf("failed to parse input field type"))
		return nil
	}

	rule, ok := ctx.InputSourceRule().Accept(v).(endpoints.InputSourceRule)
	if !ok || rule == nil {
		v.setErr(fmt.Errorf("failed to parse input source rule"))
		return nil
	}

	field, err := endpoints.NewInputFieldBuilder().
		Name(ctx.IDENT().GetText()).
		Type(typeRef).
		Sources(rule).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return field
}

func (v *endpointVisitor) VisitInputSourceRule(ctx *endpoint_generated.InputSourceRuleContext) interface{} {
	builder := endpoints.NewInputSourceRuleBuilder()

	if ctx.LPAREN() != nil {
		builder.Kind(endpoints.InputSourceRuleExactlyOneOf)
	} else {
		builder.Kind(endpoints.InputSourceRuleSingle)
	}

	for _, sourceCtx := range ctx.AllInputSource() {
		source, ok := sourceCtx.Accept(v).(endpoints.InputSource)
		if !ok || source == nil {
			v.setErr(fmt.Errorf("failed to parse input source"))
			return nil
		}

		builder.AddSource(source)
	}

	rule, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return rule
}

func (v *endpointVisitor) VisitInputSource(ctx *endpoint_generated.InputSourceContext) interface{} {
	source, err := endpoints.NewInputSourceBuilder().
		Kind(parseInputSourceKind(ctx.InputSourceKind().GetText())).
		Optional(ctx.QUESTION() != nil).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return source
}

func (v *endpointVisitor) VisitAction(ctx *endpoint_generated.ActionContext) interface{} {
	switch {
	case ctx.FetchAction() != nil:
		return ctx.FetchAction().Accept(v)
	case ctx.CreateAction() != nil:
		return ctx.CreateAction().Accept(v)
	case ctx.StoreAction() != nil:
		return ctx.StoreAction().Accept(v)
	case ctx.UpdateAction() != nil:
		return ctx.UpdateAction().Accept(v)
	case ctx.EmitAction() != nil:
		return ctx.EmitAction().Accept(v)
	case ctx.ReturnAction() != nil:
		return ctx.ReturnAction().Accept(v)
	}

	v.setErr(fmt.Errorf("unknown action"))
	return nil
}

func (v *endpointVisitor) VisitFetchAction(ctx *endpoint_generated.FetchActionContext) interface{} {
	variable := ctx.TypedVariable().Accept(v).(endpoints.TypedVariable)
	condition := ctx.Condition().Accept(v).(endpoints.Condition)

	action, err := endpoints.NewFetchActionBuilder().
		Variable(variable).
		Source(ctx.IDENT().GetText()).
		Condition(condition).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return action
}

func (v *endpointVisitor) VisitCreateAction(ctx *endpoint_generated.CreateActionContext) interface{} {
	variable := ctx.TypedVariable().Accept(v).(endpoints.TypedVariable)
	assignments := ctx.AssignmentBlock().Accept(v).([]endpoints.Assignment)

	builder := endpoints.NewCreateActionBuilder().Variable(variable)
	for _, assignment := range assignments {
		builder.AddAssignment(assignment)
	}

	action, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return action
}

func (v *endpointVisitor) VisitStoreAction(ctx *endpoint_generated.StoreActionContext) interface{} {
	action, err := endpoints.NewStoreActionBuilder().
		VariableName(ctx.IDENT(0).GetText()).
		Destination(ctx.IDENT(1).GetText()).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return action
}

func (v *endpointVisitor) VisitUpdateAction(ctx *endpoint_generated.UpdateActionContext) interface{} {
	assignments := ctx.AssignmentBlock().Accept(v).([]endpoints.Assignment)

	builder := endpoints.NewUpdateActionBuilder().
		VariableName(ctx.IDENT().GetText())

	for _, assignment := range assignments {
		builder.AddAssignment(assignment)
	}

	action, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return action
}

func (v *endpointVisitor) VisitEmitAction(ctx *endpoint_generated.EmitActionContext) interface{} {
	variable := ctx.TypedVariable().Accept(v).(endpoints.TypedVariable)
	assignments := ctx.AssignmentBlock().Accept(v).([]endpoints.Assignment)

	builder := endpoints.NewEmitActionBuilder().Variable(variable)
	for _, assignment := range assignments {
		builder.AddAssignment(assignment)
	}

	action, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return action
}

func (v *endpointVisitor) VisitReturnAction(ctx *endpoint_generated.ReturnActionContext) interface{} {
	expression := ctx.Expression().Accept(v).(endpoints.Expression)

	action, err := endpoints.NewReturnActionBuilder().
		Expression(expression).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return action
}

func (v *endpointVisitor) VisitTypedVariable(ctx *endpoint_generated.TypedVariableContext) interface{} {
	typeRef := ctx.TypeRef().Accept(v).(objects.TypeReference)

	variable, err := endpoints.NewTypedVariableBuilder().
		Name(ctx.IDENT().GetText()).
		Type(typeRef).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return variable
}

func (v *endpointVisitor) VisitAssignmentBlock(ctx *endpoint_generated.AssignmentBlockContext) interface{} {
	assignments := make([]endpoints.Assignment, 0)

	for _, assignmentCtx := range ctx.AllAssignment() {
		assignment := assignmentCtx.Accept(v).(endpoints.Assignment)
		assignments = append(assignments, assignment)
	}

	return assignments
}

func (v *endpointVisitor) VisitAssignment(ctx *endpoint_generated.AssignmentContext) interface{} {
	expression := ctx.Expression().Accept(v).(endpoints.Expression)

	assignment, err := endpoints.NewAssignmentBuilder().
		Name(ctx.IDENT().GetText()).
		Value(expression).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return assignment
}

func (v *endpointVisitor) VisitCondition(ctx *endpoint_generated.ConditionContext) interface{} {
	left := ctx.Expression(0).Accept(v).(endpoints.Expression)
	right := ctx.Expression(1).Accept(v).(endpoints.Expression)

	condition, err := endpoints.NewConditionBuilder().
		Left(left).
		Operator(parseComparator(ctx.Comparator().GetText())).
		Right(right).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return condition
}

func (v *endpointVisitor) VisitExpression(ctx *endpoint_generated.ExpressionContext) interface{} {
	if ctx.FunctionCall() != nil {
		return ctx.FunctionCall().Accept(v)
	}

	if ctx.Selector() != nil {
		return ctx.Selector().Accept(v)
	}

	value := ctx.Value().Accept(v).(objects.Value)

	expression, err := endpoints.NewLiteralExpressionBuilder().
		Value(value).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return expression
}

func (v *endpointVisitor) VisitFunctionCall(ctx *endpoint_generated.FunctionCallContext) interface{} {
	builder := endpoints.NewFunctionCallExpressionBuilder().
		Name(ctx.IDENT().GetText())

	if ctx.ArgumentList() != nil {
		args := ctx.ArgumentList().Accept(v).([]endpoints.Expression)
		for _, arg := range args {
			builder.AddArgument(arg)
		}
	}

	expression, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return expression
}

func (v *endpointVisitor) VisitArgumentList(ctx *endpoint_generated.ArgumentListContext) interface{} {
	args := make([]endpoints.Expression, 0)

	for _, expressionCtx := range ctx.AllExpression() {
		args = append(args, expressionCtx.Accept(v).(endpoints.Expression))
	}

	return args
}

func (v *endpointVisitor) VisitSelector(ctx *endpoint_generated.SelectorContext) interface{} {
	parts := strings.Split(ctx.GetText(), ".")

	if len(parts) == 1 {
		expression, err := endpoints.NewIdentifierExpressionBuilder().
			Name(parts[0]).
			Build()
		if err != nil {
			v.setErr(err)
			return nil
		}
		return expression
	}

	builder := endpoints.NewSelectorExpressionBuilder()
	for _, part := range parts {
		builder.AddPart(part)
	}

	expression, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return expression
}

func (v *endpointVisitor) VisitTypeRef(ctx *endpoint_generated.TypeRefContext) interface{} {
	builder := objects.NewTypeReferenceBuilder().Name(ctx.TypeName().GetText())

	if ctx.LIST() != nil {
		builder.List(true)
	}

	if ctx.NumberConstraint() != nil {
		constraint := ctx.NumberConstraint().Accept(v).(objects.NumberConstraint)
		builder.NumberConstraint(constraint)
	} else {
		constraint, err := objects.NewNumberConstraintBuilder().
			Kind(objects.NumberConstraintNone).
			Build()
		if err != nil {
			v.setErr(err)
			return nil
		}
		builder.NumberConstraint(constraint)
	}

	if ctx.OptionalMarker() != nil {
		builder.Optional(true)
	}

	typeRef, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return typeRef
}

func (v *endpointVisitor) VisitNumberConstraint(ctx *endpoint_generated.NumberConstraintContext) interface{} {
	builder := objects.NewNumberConstraintBuilder()

	switch {
	case ctx.PLUS() != nil:
		constraint, err := builder.Kind(objects.NumberConstraintOneOrMore).Build()
		if err != nil {
			v.setErr(err)
			return nil
		}
		return constraint

	case ctx.STAR() != nil:
		constraint, err := builder.Kind(objects.NumberConstraintZeroOrMore).Build()
		if err != nil {
			v.setErr(err)
			return nil
		}
		return constraint
	}

	builder.Kind(objects.NumberConstraintRange)

	text := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(ctx.GetText()), "["), "]")
	parts := strings.Split(text, ",")

	if len(parts) != 2 {
		v.setErr(fmt.Errorf("invalid number constraint %q", ctx.GetText()))
		return nil
	}

	if min := strings.TrimSpace(parts[0]); min != "" {
		value, err := parseNumberValue(min)
		if err != nil {
			v.setErr(err)
			return nil
		}
		builder.Min(value)
	}

	if max := strings.TrimSpace(parts[1]); max != "" {
		value, err := parseNumberValue(max)
		if err != nil {
			v.setErr(err)
			return nil
		}
		builder.Max(value)
	}

	constraint, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return constraint
}

func (v *endpointVisitor) VisitValue(ctx *endpoint_generated.ValueContext) interface{} {
	builder := objects.NewValueBuilder()

	switch {
	case ctx.STRING() != nil:
		value, err := builder.String(unquote(ctx.STRING().GetText()))
		if err != nil {
			v.setErr(err)
			return nil
		}
		return value

	case ctx.INT() != nil:
		value, err := builder.Int(ctx.INT().GetText())
		if err != nil {
			v.setErr(err)
			return nil
		}
		return value

	case ctx.FLOAT() != nil:
		value, err := builder.Float(ctx.FLOAT().GetText())
		if err != nil {
			v.setErr(err)
			return nil
		}
		return value

	case ctx.TRUE() != nil:
		value, err := builder.Bool("true")
		if err != nil {
			v.setErr(err)
			return nil
		}
		return value

	case ctx.FALSE() != nil:
		value, err := builder.Bool("false")
		if err != nil {
			v.setErr(err)
			return nil
		}
		return value
	}

	v.setErr(fmt.Errorf("unknown value %q", ctx.GetText()))
	return nil
}

func parseHTTPMethod(value string) endpoints.HTTPMethod {
	switch value {
	case "GET":
		return endpoints.HTTPMethodGet
	case "POST":
		return endpoints.HTTPMethodPost
	case "PUT":
		return endpoints.HTTPMethodPut
	case "PATCH":
		return endpoints.HTTPMethodPatch
	case "DELETE":
		return endpoints.HTTPMethodDelete
	default:
		return endpoints.HTTPMethod(value)
	}
}

func parseInputSourceKind(value string) endpoints.InputSourceKind {
	switch value {
	case "path":
		return endpoints.InputSourcePath
	case "query":
		return endpoints.InputSourceQuery
	case "body":
		return endpoints.InputSourceBody
	case "header":
		return endpoints.InputSourceHeader
	case "cookie":
		return endpoints.InputSourceCookie
	default:
		return endpoints.InputSourceKind(value)
	}
}

func parseComparator(value string) endpoints.Comparator {
	switch value {
	case "==":
		return endpoints.ComparatorEqual

	case "!=":
		return endpoints.ComparatorNotEqual

	case "<":
		return endpoints.ComparatorLessThan

	case "<=":
		return endpoints.ComparatorLessThanOrEqual

	case ">":
		return endpoints.ComparatorGreaterThan

	case ">=":
		return endpoints.ComparatorGreaterThanOrEqual

	default:
		return endpoints.Comparator(value)
	}
}
