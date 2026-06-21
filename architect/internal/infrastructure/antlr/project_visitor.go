package antlr

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/projects"
	project_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/project"
)

type serviceFiles []string
type deploymentFiles []string

type projectVisitor struct {
	*project_generated.BaseProjectVisitor
	err error
}

func newProjectVisitor() *projectVisitor {
	return &projectVisitor{
		BaseProjectVisitor: &project_generated.BaseProjectVisitor{},
	}
}

func (v *projectVisitor) setErr(err error) {
	if err != nil && v.err == nil {
		v.err = err
	}
}

func (v *projectVisitor) VisitProgram(ctx *project_generated.ProgramContext) interface{} {
	return ctx.ProjectDecl().Accept(v)
}

func (v *projectVisitor) VisitProjectDecl(ctx *project_generated.ProjectDeclContext) interface{} {
	builder := projects.NewProjectBuilder().
		Name(ctx.IDENT().GetText())

	for _, versionCtx := range ctx.AllVersionDecl() {
		result := versionCtx.Accept(v)
		if v.err != nil {
			return nil
		}

		version, ok := result.(projects.Version)
		if !ok || version == nil {
			v.setErr(fmt.Errorf("failed to parse project version"))
			return nil
		}

		builder.AddVersion(version)
	}

	project, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return project
}

func (v *projectVisitor) VisitVersionDecl(ctx *project_generated.VersionDeclContext) interface{} {
	return v.visitVersion(
		unquote(ctx.STRING().GetText()),
		ctx.AllVersionBody(),
	)
}

func (v *projectVisitor) VisitNextVersionDecl(ctx *project_generated.NextVersionDeclContext) interface{} {
	return v.visitVersion(
		unquote(ctx.STRING().GetText()),
		ctx.AllVersionBody(),
	)
}

func (v *projectVisitor) visitVersion(
	number string,
	bodyContexts []project_generated.IVersionBodyContext,
) interface{} {
	builder := projects.NewVersionBuilder().
		Number(number)

	for _, bodyCtx := range bodyContexts {
		result := bodyCtx.Accept(v)
		if v.err != nil {
			return nil
		}

		switch value := result.(type) {
		case serviceFiles:
			for _, file := range value {
				builder.AddServiceFile(file)
			}

		case objectFiles:
			for _, file := range value {
				builder.AddObjectFile(file)
			}

		case deploymentFiles:
			for _, file := range value {
				builder.AddDeploymentFile(file)
			}

		case projects.Version:
			builder.AddNextVersion(value)

		default:
			v.setErr(fmt.Errorf("unknown project version body"))
			return nil
		}
	}

	version, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return version
}

func (v *projectVisitor) VisitVersionBody(ctx *project_generated.VersionBodyContext) interface{} {
	switch {
	case ctx.ServicesBlock() != nil:
		return ctx.ServicesBlock().Accept(v)

	case ctx.ObjectsBlock() != nil:
		return ctx.ObjectsBlock().Accept(v)

	case ctx.DeploymentsBlock() != nil:
		return ctx.DeploymentsBlock().Accept(v)

	case ctx.NextVersionDecl() != nil:
		return ctx.NextVersionDecl().Accept(v)
	}

	v.setErr(fmt.Errorf("unknown project version body"))
	return nil
}

func (v *projectVisitor) VisitServicesBlock(ctx *project_generated.ServicesBlockContext) interface{} {
	files := make(serviceFiles, 0)

	for _, fileRefCtx := range ctx.AllFileRef() {
		files = append(files, unquote(fileRefCtx.STRING().GetText()))
	}

	return files
}

func (v *projectVisitor) VisitObjectsBlock(ctx *project_generated.ObjectsBlockContext) interface{} {
	files := make(objectFiles, 0)

	for _, fileRefCtx := range ctx.AllFileRef() {
		files = append(files, unquote(fileRefCtx.STRING().GetText()))
	}

	return files
}

func (v *projectVisitor) VisitDeploymentsBlock(ctx *project_generated.DeploymentsBlockContext) interface{} {
	files := make(deploymentFiles, 0)

	for _, fileRefCtx := range ctx.AllFileRef() {
		files = append(files, unquote(fileRefCtx.STRING().GetText()))
	}

	return files
}
