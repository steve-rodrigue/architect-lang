// Code generated from internal/infrastructure/antlr/grammars/Common.g4 by ANTLR 4.13.2. DO NOT EDIT.

package object // Common
import "github.com/antlr4-go/antlr/v4"

type BaseCommonVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCommonVisitor) VisitTypeRef(ctx *TypeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCommonVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCommonVisitor) VisitNumberConstraint(ctx *NumberConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCommonVisitor) VisitOptionalMarker(ctx *OptionalMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCommonVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCommonVisitor) VisitNumberValue(ctx *NumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCommonVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
