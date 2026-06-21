package antlr

import (
	"fmt"
	"strconv"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	deployment_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/deployment"
)

type inventoryFile string

type deploymentVisitor struct {
	*deployment_generated.BaseDeploymentVisitor
	err error
}

func newDeploymentVisitor() *deploymentVisitor {
	return &deploymentVisitor{
		BaseDeploymentVisitor: &deployment_generated.BaseDeploymentVisitor{},
	}
}

func (v *deploymentVisitor) setErr(err error) {
	if err != nil && v.err == nil {
		v.err = err
	}
}

func (v *deploymentVisitor) VisitProgram(ctx *deployment_generated.ProgramContext) interface{} {
	return ctx.DeploymentDecl().Accept(v)
}

func (v *deploymentVisitor) VisitDeploymentDecl(ctx *deployment_generated.DeploymentDeclContext) interface{} {
	builder := deployments.NewDeploymentBuilder().
		Name(ctx.IDENT().GetText())

	for _, bodyCtx := range ctx.AllDeploymentBody() {
		result := bodyCtx.Accept(v)
		if v.err != nil {
			return nil
		}

		switch value := result.(type) {
		case deployments.Vendor:
			builder.Vendor(value)

		case inventoryFile:
			builder.Inventory(string(value))

		case deployments.ServiceDeployment:
			builder.AddService(value)

		default:
			v.setErr(fmt.Errorf("unknown deployment body"))
			return nil
		}
	}

	deployment, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return deployment
}

func (v *deploymentVisitor) VisitDeploymentBody(ctx *deployment_generated.DeploymentBodyContext) interface{} {
	switch {
	case ctx.VendorDecl() != nil:
		return ctx.VendorDecl().Accept(v)

	case ctx.InventoryDecl() != nil:
		return ctx.InventoryDecl().Accept(v)

	case ctx.ServiceDeploymentDecl() != nil:
		return ctx.ServiceDeploymentDecl().Accept(v)
	}

	v.setErr(fmt.Errorf("unknown deployment body"))
	return nil
}

func (v *deploymentVisitor) VisitVendorDecl(ctx *deployment_generated.VendorDeclContext) interface{} {
	return deployments.Vendor(ctx.IDENT().GetText())
}

func (v *deploymentVisitor) VisitInventoryDecl(ctx *deployment_generated.InventoryDeclContext) interface{} {
	return inventoryFile(unquote(ctx.STRING().GetText()))
}

func (v *deploymentVisitor) VisitServiceDeploymentDecl(ctx *deployment_generated.ServiceDeploymentDeclContext) interface{} {
	builder := deployments.NewServiceDeploymentBuilder().
		Name(ctx.IDENT().GetText())

	for _, bodyCtx := range ctx.AllServiceDeploymentBody() {
		bodyCtx.Accept(v)

		if v.err != nil {
			return nil
		}

		switch {
		case bodyCtx.ReplicasDecl() != nil:
			count, err := strconv.Atoi(bodyCtx.ReplicasDecl().INT().GetText())
			if err != nil {
				v.setErr(err)
				return nil
			}
			builder.Replicas(count)

		case bodyCtx.DomainDecl() != nil:
			builder.Domain(unquote(bodyCtx.DomainDecl().STRING().GetText()))

		case bodyCtx.HostDecl() != nil:
			builder.Host(unquote(bodyCtx.HostDecl().STRING().GetText()))

		case bodyCtx.VolumeDecl() != nil:
			builder.Volume(unquote(bodyCtx.VolumeDecl().STRING().GetText()))

		case bodyCtx.BackupDecl() != nil:
			builder.Backup(bodyCtx.BackupDecl().BooleanValue().TRUE() != nil)

		default:
			v.setErr(fmt.Errorf("unknown service deployment body"))
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
