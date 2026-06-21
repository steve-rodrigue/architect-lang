// Code generated from internal/infrastructure/antlr/grammars/Project.g4 by ANTLR 4.13.2. DO NOT EDIT.

package project // Project
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ProjectParser.
type ProjectVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ProjectParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ProjectParser#projectDecl.
	VisitProjectDecl(ctx *ProjectDeclContext) interface{}

	// Visit a parse tree produced by ProjectParser#versionDecl.
	VisitVersionDecl(ctx *VersionDeclContext) interface{}

	// Visit a parse tree produced by ProjectParser#versionBody.
	VisitVersionBody(ctx *VersionBodyContext) interface{}

	// Visit a parse tree produced by ProjectParser#nextVersionDecl.
	VisitNextVersionDecl(ctx *NextVersionDeclContext) interface{}

	// Visit a parse tree produced by ProjectParser#servicesBlock.
	VisitServicesBlock(ctx *ServicesBlockContext) interface{}

	// Visit a parse tree produced by ProjectParser#objectsBlock.
	VisitObjectsBlock(ctx *ObjectsBlockContext) interface{}

	// Visit a parse tree produced by ProjectParser#deploymentsBlock.
	VisitDeploymentsBlock(ctx *DeploymentsBlockContext) interface{}

	// Visit a parse tree produced by ProjectParser#fileRef.
	VisitFileRef(ctx *FileRefContext) interface{}
}
