package antlr

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
	consumer_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/consumer"
)

type consumerWorkflowVisitor struct {
	*consumer_generated.BaseConsumerVisitor
	common *consumerCommonVisitor
	setErr func(error)
}

func newConsumerWorkflowVisitor(setErr func(error)) *consumerWorkflowVisitor {
	visitor := &consumerWorkflowVisitor{
		BaseConsumerVisitor: &consumer_generated.BaseConsumerVisitor{},
		setErr:              setErr,
	}

	visitor.common = newConsumerCommonVisitor(setErr)

	return visitor
}

func (v *consumerWorkflowVisitor) VisitAction(ctx *consumer_generated.ActionContext) interface{} {
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
	case ctx.CallAction() != nil:
		return ctx.CallAction().Accept(v)
	case ctx.ReturnAction() != nil:
		return ctx.ReturnAction().Accept(v)
	}

	v.setErr(fmt.Errorf("unknown action"))
	return nil
}

func (v *consumerWorkflowVisitor) VisitFetchAction(ctx *consumer_generated.FetchActionContext) interface{} {
	variable, ok := ctx.TypedVariable().Accept(v).(workflows.TypedVariable)
	if !ok || variable == nil {
		v.setErr(fmt.Errorf("failed to parse fetch variable"))
		return nil
	}

	condition, ok := ctx.Condition().Accept(v).(workflows.Condition)
	if !ok || condition == nil {
		v.setErr(fmt.Errorf("failed to parse fetch condition"))
		return nil
	}

	return buildWorkflowFetchAction(variable, ctx.IDENT().GetText(), condition, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitCreateAction(ctx *consumer_generated.CreateActionContext) interface{} {
	variable, ok := ctx.TypedVariable().Accept(v).(workflows.TypedVariable)
	if !ok || variable == nil {
		v.setErr(fmt.Errorf("failed to parse create variable"))
		return nil
	}

	assignments, ok := ctx.AssignmentBlock().Accept(v).([]workflows.Assignment)
	if !ok {
		v.setErr(fmt.Errorf("failed to parse create assignments"))
		return nil
	}

	return buildWorkflowCreateAction(variable, assignments, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitStoreAction(ctx *consumer_generated.StoreActionContext) interface{} {
	return buildWorkflowStoreAction(ctx.IDENT(0).GetText(), ctx.IDENT(1).GetText(), v.setErr)
}

func (v *consumerWorkflowVisitor) VisitUpdateAction(ctx *consumer_generated.UpdateActionContext) interface{} {
	assignments, ok := ctx.AssignmentBlock().Accept(v).([]workflows.Assignment)
	if !ok {
		v.setErr(fmt.Errorf("failed to parse update assignments"))
		return nil
	}

	return buildWorkflowUpdateAction(ctx.IDENT().GetText(), assignments, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitEmitAction(ctx *consumer_generated.EmitActionContext) interface{} {
	variable, ok := ctx.TypedVariable().Accept(v).(workflows.TypedVariable)
	if !ok || variable == nil {
		v.setErr(fmt.Errorf("failed to parse emit variable"))
		return nil
	}

	assignments, ok := ctx.AssignmentBlock().Accept(v).([]workflows.Assignment)
	if !ok {
		v.setErr(fmt.Errorf("failed to parse emit assignments"))
		return nil
	}

	return buildWorkflowEmitAction(variable, assignments, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitCallAction(ctx *consumer_generated.CallActionContext) interface{} {
	assignments, ok := ctx.AssignmentBlock().Accept(v).([]workflows.Assignment)
	if !ok {
		v.setErr(fmt.Errorf("failed to parse call assignments"))
		return nil
	}

	return buildWorkflowCallAction(ctx.Selector().GetText(), assignments, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitReturnAction(ctx *consumer_generated.ReturnActionContext) interface{} {
	expression, ok := ctx.Expression().Accept(v).(workflows.Expression)
	if !ok || expression == nil {
		v.setErr(fmt.Errorf("failed to parse return expression"))
		return nil
	}

	return buildWorkflowReturnAction(expression, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitTypedVariable(ctx *consumer_generated.TypedVariableContext) interface{} {
	typeRef, ok := ctx.TypeRef().Accept(v.common).(common.TypeReference)
	if !ok || typeRef == nil {
		v.setErr(fmt.Errorf("failed to parse typed variable type"))
		return nil
	}

	return buildWorkflowTypedVariable(ctx.Identifier().GetText(), typeRef, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitAssignmentBlock(ctx *consumer_generated.AssignmentBlockContext) interface{} {
	assignments := make([]workflows.Assignment, 0)

	for _, assignmentCtx := range ctx.AllAssignment() {
		assignment, ok := assignmentCtx.Accept(v).(workflows.Assignment)
		if !ok || assignment == nil {
			v.setErr(fmt.Errorf("failed to parse assignment"))
			return nil
		}

		assignments = append(assignments, assignment)
	}

	return assignments
}

func (v *consumerWorkflowVisitor) VisitAssignment(ctx *consumer_generated.AssignmentContext) interface{} {
	expression, ok := ctx.Expression().Accept(v).(workflows.Expression)
	if !ok || expression == nil {
		v.setErr(fmt.Errorf("failed to parse assignment expression"))
		return nil
	}

	return buildWorkflowAssignment(ctx.Identifier().GetText(), expression, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitCondition(ctx *consumer_generated.ConditionContext) interface{} {
	left, ok := ctx.Expression(0).Accept(v).(workflows.Expression)
	if !ok || left == nil {
		v.setErr(fmt.Errorf("failed to parse condition left expression"))
		return nil
	}

	right, ok := ctx.Expression(1).Accept(v).(workflows.Expression)
	if !ok || right == nil {
		v.setErr(fmt.Errorf("failed to parse condition right expression"))
		return nil
	}

	return buildWorkflowCondition(left, parseComparator(ctx.Comparator().GetText()), right, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitExpression(ctx *consumer_generated.ExpressionContext) interface{} {
	if ctx.FunctionCall() != nil {
		return ctx.FunctionCall().Accept(v)
	}

	if ctx.Selector() != nil {
		return ctx.Selector().Accept(v)
	}

	if ctx.Value() != nil {
		value, ok := ctx.Value().Accept(v.common).(common.Value)
		if !ok || value == nil {
			v.setErr(fmt.Errorf("failed to parse literal value"))
			return nil
		}

		return buildWorkflowLiteralExpression(value, v.setErr)
	}

	v.setErr(fmt.Errorf("invalid expression: %s", ctx.GetText()))
	return nil
}

func (v *consumerWorkflowVisitor) VisitFunctionCall(ctx *consumer_generated.FunctionCallContext) interface{} {
	args := make([]workflows.Expression, 0)

	if ctx.ArgumentList() != nil {
		parsed, ok := ctx.ArgumentList().Accept(v).([]workflows.Expression)
		if !ok {
			v.setErr(fmt.Errorf("failed to parse function call arguments"))
			return nil
		}

		args = parsed
	}

	return buildWorkflowFunctionCallExpression(ctx.IDENT().GetText(), args, v.setErr)
}

func (v *consumerWorkflowVisitor) VisitArgumentList(ctx *consumer_generated.ArgumentListContext) interface{} {
	args := make([]workflows.Expression, 0)

	for _, expressionCtx := range ctx.AllExpression() {
		expression, ok := expressionCtx.Accept(v).(workflows.Expression)
		if !ok || expression == nil {
			v.setErr(fmt.Errorf("failed to parse function call argument"))
			return nil
		}

		args = append(args, expression)
	}

	return args
}

func (v *consumerWorkflowVisitor) VisitSelector(ctx *consumer_generated.SelectorContext) interface{} {
	parts := splitSelector(ctx.GetText())

	if len(parts) == 1 {
		return buildWorkflowIdentifierExpression(parts[0], v.setErr)
	}

	return buildWorkflowSelectorExpression(parts, v.setErr)
}
