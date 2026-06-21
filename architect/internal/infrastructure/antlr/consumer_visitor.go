package antlr

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
	consumer_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/consumer"
)

type consumerVisitor struct {
	*consumer_generated.BaseConsumerVisitor
	workflow *consumerWorkflowVisitor
	err      error
}

func newConsumerVisitor() *consumerVisitor {
	visitor := &consumerVisitor{
		BaseConsumerVisitor: &consumer_generated.BaseConsumerVisitor{},
	}

	visitor.workflow = newConsumerWorkflowVisitor(visitor.setErr)

	return visitor
}

func (v *consumerVisitor) setErr(err error) {
	if err != nil && v.err == nil {
		v.err = err
	}
}

func (v *consumerVisitor) VisitProgram(ctx *consumer_generated.ProgramContext) interface{} {
	return ctx.ConsumerDecl().Accept(v)
}

func (v *consumerVisitor) VisitConsumerDecl(ctx *consumer_generated.ConsumerDeclContext) interface{} {
	builder := consumers.NewConsumerBuilder().
		EventName(ctx.Identifier().GetText())

	for _, actionCtx := range ctx.AllAction_() {
		result := actionCtx.Accept(v.workflow)
		if v.err != nil {
			return nil
		}

		action, ok := result.(workflows.Action)
		if !ok || action == nil {
			v.setErr(fmt.Errorf("failed to parse consumer action"))
			return nil
		}

		builder.AddAction(action)
	}

	consumer, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return consumer
}
