package antlr

import (
	"fmt"
	"strconv"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
	service_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/service"
)

type serviceVisitor struct {
	*service_generated.BaseServiceVisitor
	err error
}

func newServiceVisitor() *serviceVisitor {
	return &serviceVisitor{
		BaseServiceVisitor: &service_generated.BaseServiceVisitor{},
	}
}

func (v *serviceVisitor) setErr(err error) {
	if err != nil && v.err == nil {
		v.err = err
	}
}

func (v *serviceVisitor) VisitProgram(ctx *service_generated.ProgramContext) interface{} {
	return ctx.ServiceDecl().Accept(v)
}

func (v *serviceVisitor) VisitServiceDecl(ctx *service_generated.ServiceDeclContext) interface{} {
	builder := services.NewServiceBuilder().
		Name(ctx.IDENT().GetText()).
		Kind(services.ServiceKind(ctx.ServiceKind().GetText()))

	for _, bodyCtx := range ctx.AllServiceBody() {
		bodyCtx.Accept(v)
		if v.err != nil {
			return nil
		}

		switch {
		case bodyCtx.VersionDecl() != nil:
			builder.Version(unquote(bodyCtx.VersionDecl().STRING().GetText()))

		case bodyCtx.ImageDecl() != nil:
			builder.Image(unquote(bodyCtx.ImageDecl().STRING().GetText()))

		case bodyCtx.ExposeDecl() != nil:
			port, err := strconv.Atoi(bodyCtx.ExposeDecl().INT().GetText())
			if err != nil {
				v.setErr(err)
				return nil
			}
			builder.Expose(port)

		case bodyCtx.DependsOnDecl() != nil:
			builder.DependsOn(bodyCtx.DependsOnDecl().IDENT().GetText())

		case bodyCtx.ApplicationDecl() != nil:
			builder.Application(unquote(bodyCtx.ApplicationDecl().STRING().GetText()))

		case bodyCtx.StoreDecl() != nil:
			builder.Store(bodyCtx.StoreDecl().IDENT().GetText())

		case bodyCtx.EventDecl() != nil:
			builder.Event(bodyCtx.EventDecl().IDENT().GetText())

		default:
			v.setErr(fmt.Errorf("unknown service body"))
			return nil
		}
	}

	service, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return service
}
