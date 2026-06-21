// Code generated from internal/infrastructure/antlr/grammars/Project.g4 by ANTLR 4.13.2. DO NOT EDIT.

package project // Project
import "github.com/antlr4-go/antlr/v4"

type BaseProjectVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseProjectVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitProjectDecl(ctx *ProjectDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitVersionDecl(ctx *VersionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitVersionBody(ctx *VersionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitNextVersionDecl(ctx *NextVersionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitServicesBlock(ctx *ServicesBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitObjectsBlock(ctx *ObjectsBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitDeploymentsBlock(ctx *DeploymentsBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProjectVisitor) VisitFileRef(ctx *FileRefContext) interface{} {
	return v.VisitChildren(ctx)
}
