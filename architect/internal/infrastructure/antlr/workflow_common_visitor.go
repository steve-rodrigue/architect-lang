package antlr

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	workflow_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/workflow"
)

type workflowCommonVisitor struct {
	*workflow_generated.BaseWorkflowVisitor
	setErr func(error)
}

func newWorkflowCommonVisitor(setErr func(error)) *workflowCommonVisitor {
	return &workflowCommonVisitor{
		BaseWorkflowVisitor: &workflow_generated.BaseWorkflowVisitor{},
		setErr:              setErr,
	}
}

func (v *workflowCommonVisitor) VisitTypeRef(ctx *workflow_generated.TypeRefContext) interface{} {
	var constraint common.NumberConstraint

	if ctx.NumberConstraint() != nil {
		parsed, ok := ctx.NumberConstraint().Accept(v).(common.NumberConstraint)
		if !ok || parsed == nil {
			v.setErr(fmt.Errorf("failed to parse number constraint for type %q", ctx.TypeName().GetText()))
			return nil
		}

		constraint = parsed
	}

	return buildCommonTypeReference(
		ctx.TypeName().GetText(),
		ctx.LIST() != nil,
		ctx.OptionalMarker() != nil,
		constraint,
		nil,
		false,
		v.setErr,
	)
}

func (v *workflowCommonVisitor) VisitNumberConstraint(ctx *workflow_generated.NumberConstraintContext) interface{} {
	return buildCommonNumberConstraint(
		ctx.GetText(),
		ctx.PLUS() != nil,
		ctx.STAR() != nil,
		v.setErr,
	)
}

func (v *workflowCommonVisitor) VisitValue(ctx *workflow_generated.ValueContext) interface{} {
	switch {
	case ctx.STRING() != nil:
		return buildCommonValue("string", ctx.STRING().GetText(), v.setErr)

	case ctx.INT() != nil:
		return buildCommonValue("int", ctx.INT().GetText(), v.setErr)

	case ctx.FLOAT() != nil:
		return buildCommonValue("float", ctx.FLOAT().GetText(), v.setErr)

	case ctx.TRUE() != nil:
		return buildCommonValue("bool", "true", v.setErr)

	case ctx.FALSE() != nil:
		return buildCommonValue("bool", "false", v.setErr)
	}

	v.setErr(fmt.Errorf("unknown value: %s", ctx.GetText()))
	return nil
}
