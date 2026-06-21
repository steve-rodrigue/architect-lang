// Code generated from internal/infrastructure/antlr/grammars/Application.g4 by ANTLR 4.13.2. DO NOT EDIT.

package application // Application
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ApplicationParser.
type ApplicationVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ApplicationParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ApplicationParser#applicationDecl.
	VisitApplicationDecl(ctx *ApplicationDeclContext) interface{}

	// Visit a parse tree produced by ApplicationParser#applicationBody.
	VisitApplicationBody(ctx *ApplicationBodyContext) interface{}

	// Visit a parse tree produced by ApplicationParser#portDecl.
	VisitPortDecl(ctx *PortDeclContext) interface{}

	// Visit a parse tree produced by ApplicationParser#endpointsBlock.
	VisitEndpointsBlock(ctx *EndpointsBlockContext) interface{}

	// Visit a parse tree produced by ApplicationParser#consumersBlock.
	VisitConsumersBlock(ctx *ConsumersBlockContext) interface{}

	// Visit a parse tree produced by ApplicationParser#fileRef.
	VisitFileRef(ctx *FileRefContext) interface{}
}
