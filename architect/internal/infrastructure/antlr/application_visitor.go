package antlr

import (
	"fmt"
	"strconv"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	application_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/application"
)

type objectFiles []string
type endpointFiles []string
type consumerFiles []string

type applicationVisitor struct {
	*application_generated.BaseApplicationVisitor
	err error
}

func newApplicationVisitor() *applicationVisitor {
	return &applicationVisitor{
		BaseApplicationVisitor: &application_generated.BaseApplicationVisitor{},
	}
}

func (v *applicationVisitor) setErr(err error) {
	if err != nil && v.err == nil {
		v.err = err
	}
}

func (v *applicationVisitor) VisitProgram(ctx *application_generated.ProgramContext) interface{} {
	return ctx.ApplicationDecl().Accept(v)
}

func (v *applicationVisitor) VisitApplicationDecl(ctx *application_generated.ApplicationDeclContext) interface{} {
	builder := applications.NewApplicationBuilder().
		Name(ctx.IDENT().GetText())

	for _, bodyCtx := range ctx.AllApplicationBody() {
		result := bodyCtx.Accept(v)
		if v.err != nil {
			return nil
		}

		switch value := result.(type) {
		case applications.Port:
			builder.AddPort(value)

		case endpointFiles:
			for _, file := range value {
				builder.AddEndpointFile(file)
			}

		case consumerFiles:
			for _, file := range value {
				builder.AddConsumerFile(file)
			}
		}
	}

	application, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return application
}

func (v *applicationVisitor) VisitApplicationBody(ctx *application_generated.ApplicationBodyContext) interface{} {
	switch {
	case ctx.PortDecl() != nil:
		return ctx.PortDecl().Accept(v)

	case ctx.EndpointsBlock() != nil:
		return ctx.EndpointsBlock().Accept(v)

	case ctx.ConsumersBlock() != nil:
		return ctx.ConsumersBlock().Accept(v)
	}

	v.setErr(fmt.Errorf("unknown application body"))
	return nil
}

func (v *applicationVisitor) VisitPortDecl(ctx *application_generated.PortDeclContext) interface{} {
	number, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		v.setErr(err)
		return nil
	}

	direction := applications.PortDirectionEmit
	if ctx.LISTENS() != nil {
		direction = applications.PortDirectionListen
	}

	port, err := applications.NewPortBuilder().
		Direction(direction).
		Kind(applications.PortKind(ctx.IDENT().GetText())).
		Number(number).
		Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return port
}

func (v *applicationVisitor) VisitEndpointsBlock(ctx *application_generated.EndpointsBlockContext) interface{} {
	files := make(endpointFiles, 0)

	for _, fileRefCtx := range ctx.AllFileRef() {
		files = append(files, unquote(fileRefCtx.STRING().GetText()))
	}

	return files
}

func (v *applicationVisitor) VisitConsumersBlock(ctx *application_generated.ConsumersBlockContext) interface{} {
	files := make(consumerFiles, 0)

	for _, fileRefCtx := range ctx.AllFileRef() {
		files = append(files, unquote(fileRefCtx.STRING().GetText()))
	}

	return files
}
