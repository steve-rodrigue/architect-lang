// Code generated from internal/infrastructure/antlr/grammars/Deployment.g4 by ANTLR 4.13.2. DO NOT EDIT.

package deployment // Deployment
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DeploymentParser.
type DeploymentVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DeploymentParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by DeploymentParser#deploymentDecl.
	VisitDeploymentDecl(ctx *DeploymentDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#deploymentBody.
	VisitDeploymentBody(ctx *DeploymentBodyContext) interface{}

	// Visit a parse tree produced by DeploymentParser#vendorDecl.
	VisitVendorDecl(ctx *VendorDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#inventoryDecl.
	VisitInventoryDecl(ctx *InventoryDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#serviceDeploymentDecl.
	VisitServiceDeploymentDecl(ctx *ServiceDeploymentDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#serviceDeploymentBody.
	VisitServiceDeploymentBody(ctx *ServiceDeploymentBodyContext) interface{}

	// Visit a parse tree produced by DeploymentParser#replicasDecl.
	VisitReplicasDecl(ctx *ReplicasDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#domainDecl.
	VisitDomainDecl(ctx *DomainDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#hostDecl.
	VisitHostDecl(ctx *HostDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#volumeDecl.
	VisitVolumeDecl(ctx *VolumeDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#backupDecl.
	VisitBackupDecl(ctx *BackupDeclContext) interface{}

	// Visit a parse tree produced by DeploymentParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}
}
