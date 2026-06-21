// Code generated from internal/infrastructure/antlr/grammars/Service.g4 by ANTLR 4.13.2. DO NOT EDIT.

package service // Service
import "github.com/antlr4-go/antlr/v4"

type BaseServiceVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseServiceVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitServiceDecl(ctx *ServiceDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitServiceKind(ctx *ServiceKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitServiceBody(ctx *ServiceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitVersionDecl(ctx *VersionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitImageDecl(ctx *ImageDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitExposeDecl(ctx *ExposeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitDependsOnDecl(ctx *DependsOnDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitApplicationDecl(ctx *ApplicationDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitStoreDecl(ctx *StoreDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServiceVisitor) VisitEventDecl(ctx *EventDeclContext) interface{} {
	return v.VisitChildren(ctx)
}
