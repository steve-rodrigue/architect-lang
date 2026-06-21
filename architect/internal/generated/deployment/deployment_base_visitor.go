// Code generated from internal/infrastructure/antlr/grammars/Deployment.g4 by ANTLR 4.13.2. DO NOT EDIT.

package deployment // Deployment
import "github.com/antlr4-go/antlr/v4"

type BaseDeploymentVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDeploymentVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitDeploymentDecl(ctx *DeploymentDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitDeploymentBody(ctx *DeploymentBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitVendorDecl(ctx *VendorDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitInventoryDecl(ctx *InventoryDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitServiceDeploymentDecl(ctx *ServiceDeploymentDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitServiceDeploymentBody(ctx *ServiceDeploymentBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitReplicasDecl(ctx *ReplicasDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitDomainDecl(ctx *DomainDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitHostDecl(ctx *HostDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitVolumeDecl(ctx *VolumeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitBackupDecl(ctx *BackupDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDeploymentVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}
