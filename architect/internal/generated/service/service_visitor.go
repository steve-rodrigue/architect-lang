// Code generated from internal/infrastructure/antlr/grammars/Service.g4 by ANTLR 4.13.2. DO NOT EDIT.

package service // Service
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ServiceParser.
type ServiceVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ServiceParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ServiceParser#serviceDecl.
	VisitServiceDecl(ctx *ServiceDeclContext) interface{}

	// Visit a parse tree produced by ServiceParser#serviceKind.
	VisitServiceKind(ctx *ServiceKindContext) interface{}

	// Visit a parse tree produced by ServiceParser#serviceBody.
	VisitServiceBody(ctx *ServiceBodyContext) interface{}

	// Visit a parse tree produced by ServiceParser#versionDecl.
	VisitVersionDecl(ctx *VersionDeclContext) interface{}

	// Visit a parse tree produced by ServiceParser#imageDecl.
	VisitImageDecl(ctx *ImageDeclContext) interface{}

	// Visit a parse tree produced by ServiceParser#exposeDecl.
	VisitExposeDecl(ctx *ExposeDeclContext) interface{}

	// Visit a parse tree produced by ServiceParser#dependsOnDecl.
	VisitDependsOnDecl(ctx *DependsOnDeclContext) interface{}

	// Visit a parse tree produced by ServiceParser#applicationDecl.
	VisitApplicationDecl(ctx *ApplicationDeclContext) interface{}

	// Visit a parse tree produced by ServiceParser#storeDecl.
	VisitStoreDecl(ctx *StoreDeclContext) interface{}

	// Visit a parse tree produced by ServiceParser#eventDecl.
	VisitEventDecl(ctx *EventDeclContext) interface{}
}
