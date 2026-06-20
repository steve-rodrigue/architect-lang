// Code generated from internal/infrastructure/antlr/Object.g4 by ANTLR 4.13.2. DO NOT EDIT.

package object // Object
import "github.com/antlr4-go/antlr/v4"

type BaseObjectVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseObjectVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitObjectDecl(ctx *ObjectDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitObjectModifier(ctx *ObjectModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitFieldDecl(ctx *FieldDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitTypeRef(ctx *TypeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitNumberConstraint(ctx *NumberConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitOptionalMarker(ctx *OptionalMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitFieldModifier(ctx *FieldModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjectVisitor) VisitNumberValue(ctx *NumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}
