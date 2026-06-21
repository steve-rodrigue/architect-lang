// Code generated from internal/infrastructure/antlr/grammars/Application.g4 by ANTLR 4.13.2. DO NOT EDIT.

package application // Application
import "github.com/antlr4-go/antlr/v4"

type BaseApplicationVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseApplicationVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApplicationVisitor) VisitApplicationDecl(ctx *ApplicationDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApplicationVisitor) VisitApplicationBody(ctx *ApplicationBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApplicationVisitor) VisitPortDecl(ctx *PortDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApplicationVisitor) VisitEndpointsBlock(ctx *EndpointsBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApplicationVisitor) VisitConsumersBlock(ctx *ConsumersBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApplicationVisitor) VisitFileRef(ctx *FileRefContext) interface{} {
	return v.VisitChildren(ctx)
}
