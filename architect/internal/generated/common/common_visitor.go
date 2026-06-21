// Code generated from internal/infrastructure/antlr/grammars/Common.g4 by ANTLR 4.13.2. DO NOT EDIT.

package object // Common
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CommonParser.
type CommonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CommonParser#typeRef.
	VisitTypeRef(ctx *TypeRefContext) interface{}

	// Visit a parse tree produced by CommonParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by CommonParser#numberConstraint.
	VisitNumberConstraint(ctx *NumberConstraintContext) interface{}

	// Visit a parse tree produced by CommonParser#optionalMarker.
	VisitOptionalMarker(ctx *OptionalMarkerContext) interface{}

	// Visit a parse tree produced by CommonParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by CommonParser#numberValue.
	VisitNumberValue(ctx *NumberValueContext) interface{}

	// Visit a parse tree produced by CommonParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
