package antlr

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/object"
)

type objectVisitor struct {
	*generated.BaseObjectVisitor
	common *objectCommonVisitor
	err    error
}

func newObjectVisitor() *objectVisitor {
	visitor := &objectVisitor{
		BaseObjectVisitor: &generated.BaseObjectVisitor{},
	}

	visitor.common = newObjectCommonVisitor(visitor.setErr)

	return visitor
}

func (v *objectVisitor) setErr(err error) {
	if err != nil && v.err == nil {
		v.err = err
	}
}

func (v *objectVisitor) VisitProgram(ctx *generated.ProgramContext) interface{} {
	return ctx.ObjectDecl().Accept(v)
}

func (v *objectVisitor) VisitObjectDecl(ctx *generated.ObjectDeclContext) interface{} {
	builder := objects.NewObjectBuilder().
		Name(ctx.IDENT().GetText())

	for _, modifier := range ctx.AllObjectModifier() {
		if modifier.HISTORY_OF() != nil {
			historyOf, err := common.NewTypeReferenceBuilder().
				Name(modifier.IDENT().GetText()).
				Build()
			if err != nil {
				v.setErr(err)
				return nil
			}

			builder.HistoryOf(historyOf)
		}
	}

	for _, fieldCtx := range ctx.AllFieldDecl() {
		fieldResult := fieldCtx.Accept(v)
		if v.err != nil {
			return nil
		}

		field, ok := fieldResult.(objects.Field)
		if !ok || field == nil {
			v.setErr(fmt.Errorf("failed to parse field in object %q", ctx.IDENT().GetText()))
			return nil
		}

		builder.AddField(field)
	}

	obj, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return obj
}

func (v *objectVisitor) VisitFieldDecl(ctx *generated.FieldDeclContext) interface{} {
	typeRef, ok := ctx.TypeRef().Accept(v.common).(common.TypeReference)
	if !ok || typeRef == nil {
		v.setErr(fmt.Errorf("failed to parse type reference for field %q", ctx.IDENT().GetText()))
		return nil
	}

	if ctx.DefaultValue() != nil {
		if !typeRef.IsOptional() {
			v.setErr(fmt.Errorf("default value can only be used with optional type: %s", ctx.GetText()))
			return nil
		}

		value, ok := ctx.DefaultValue().Accept(v.common).(common.Value)
		if !ok || value == nil {
			v.setErr(fmt.Errorf("failed to parse default value for field %q", ctx.IDENT().GetText()))
			return nil
		}

		var err error
		typeRef, err = common.NewTypeReferenceBuilder().
			Name(typeRef.Name()).
			List(typeRef.IsList()).
			Optional(typeRef.IsOptional()).
			NumberConstraint(typeRef.NumberConstraint()).
			DefaultValue(value).
			Build()
		if err != nil {
			v.setErr(err)
			return nil
		}
	}

	builder := objects.NewFieldBuilder().
		Name(ctx.IDENT().GetText()).
		Type(typeRef)

	for _, modifier := range ctx.AllFieldModifier() {
		modifierBuilder := objects.NewFieldModifierBuilder()

		switch {
		case modifier.PRIMARY() != nil:
			fieldModifier, err := modifierBuilder.Primary()
			if err != nil {
				v.setErr(err)
				return nil
			}
			builder.AddModifier(fieldModifier)

		case modifier.UNIQUE() != nil:
			fieldModifier, err := modifierBuilder.Unique()
			if err != nil {
				v.setErr(err)
				return nil
			}
			builder.AddModifier(fieldModifier)

		case modifier.RENAMED_FROM() != nil:
			fieldModifier, err := modifierBuilder.RenamedFrom(modifier.IDENT().GetText())
			if err != nil {
				v.setErr(err)
				return nil
			}
			builder.AddModifier(fieldModifier)

		case modifier.DEPRECATED() != nil:
			fieldModifier, err := modifierBuilder.Deprecated()
			if err != nil {
				v.setErr(err)
				return nil
			}
			builder.AddModifier(fieldModifier)
		}
	}

	field, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return field
}
