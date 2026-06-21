package antlr

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
	endpoint_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/endpoint"
)

type endpointVisitor struct {
	*endpoint_generated.BaseEndpointVisitor
	common   *endpointCommonVisitor
	workflow *endpointWorkflowVisitor
	err      error
}

func newEndpointVisitor() *endpointVisitor {
	visitor := &endpointVisitor{
		BaseEndpointVisitor: &endpoint_generated.BaseEndpointVisitor{},
	}

	visitor.common = newEndpointCommonVisitor(visitor.setErr)
	visitor.workflow = newEndpointWorkflowVisitor(visitor.setErr)

	return visitor
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
		case workflows.Action:
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

	return ctx.Action_().Accept(v.workflow)
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
	typeRef, ok := ctx.TypeRef().Accept(v.common).(common.TypeReference)
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
