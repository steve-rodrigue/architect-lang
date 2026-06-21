// Code generated from internal/infrastructure/antlr/grammars/Object.g4 by ANTLR 4.13.2. DO NOT EDIT.

package object // Object
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ObjectParser.
type ObjectVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ObjectParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ObjectParser#objectDecl.
	VisitObjectDecl(ctx *ObjectDeclContext) interface{}

	// Visit a parse tree produced by ObjectParser#objectModifier.
	VisitObjectModifier(ctx *ObjectModifierContext) interface{}

	// Visit a parse tree produced by ObjectParser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by ObjectParser#typeRef.
	VisitTypeRef(ctx *TypeRefContext) interface{}

	// Visit a parse tree produced by ObjectParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by ObjectParser#numberConstraint.
	VisitNumberConstraint(ctx *NumberConstraintContext) interface{}

	// Visit a parse tree produced by ObjectParser#optionalMarker.
	VisitOptionalMarker(ctx *OptionalMarkerContext) interface{}

	// Visit a parse tree produced by ObjectParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by ObjectParser#fieldModifier.
	VisitFieldModifier(ctx *FieldModifierContext) interface{}

	// Visit a parse tree produced by ObjectParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by ObjectParser#numberValue.
	VisitNumberValue(ctx *NumberValueContext) interface{}
}
