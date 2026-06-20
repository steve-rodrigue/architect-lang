package antlr

import (
	"fmt"
	"strings"

	antlr4 "github.com/antlr4-go/antlr/v4"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/object"
)

type parserApplication struct{}

func (p *parserApplication) Object(script string) (objects.Object, error) {
	input := antlr4.NewInputStream(script)

	lexer := generated.NewObjectLexer(input)
	tokens := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	parser := generated.NewObjectParser(tokens)
	parser.BuildParseTrees = true

	tree := parser.Program()

	visitor := &objectVisitor{
		BaseObjectVisitor: &generated.BaseObjectVisitor{},
	}

	result := tree.Accept(visitor)

	if visitor.err != nil {
		return nil, visitor.err
	}

	obj, ok := result.(objects.Object)
	if !ok || obj == nil {
		return nil, fmt.Errorf("failed to parse object")
	}

	return obj, nil
}

type objectVisitor struct {
	*generated.BaseObjectVisitor
	err error
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
			historyOf, err := objects.NewTypeReferenceBuilder().
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
	typeRefResult := ctx.TypeRef().Accept(v)
	if v.err != nil {
		return nil
	}

	typeRef, ok := typeRefResult.(objects.TypeReference)
	if !ok || typeRef == nil {
		v.setErr(fmt.Errorf("failed to parse type reference for field %q", ctx.IDENT().GetText()))
		return nil
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

func (v *objectVisitor) VisitTypeRef(ctx *generated.TypeRefContext) interface{} {
	builder := objects.NewTypeReferenceBuilder()

	if ctx.LIST() != nil {
		builder.
			Name(ctx.TypeName().GetText()).
			List(true)
	} else {
		builder.Name(ctx.TypeName().GetText())
	}

	if ctx.NumberConstraint() != nil {
		constraintResult := ctx.NumberConstraint().Accept(v)
		if v.err != nil {
			return nil
		}

		constraint, ok := constraintResult.(objects.NumberConstraint)
		if !ok || constraint == nil {
			v.setErr(fmt.Errorf("failed to parse number constraint for type %q", ctx.TypeName().GetText()))
			return nil
		}

		builder.NumberConstraint(constraint)
	} else {
		constraint, err := objects.NewNumberConstraintBuilder().
			Kind(objects.NumberConstraintNone).
			Build()
		if err != nil {
			v.setErr(err)
			return nil
		}

		builder.NumberConstraint(constraint)
	}

	if ctx.OptionalMarker() != nil {
		builder.Optional(true)
	}

	if ctx.DefaultValue() != nil {
		if ctx.OptionalMarker() == nil {
			v.setErr(fmt.Errorf("default value can only be used with optional type: %s", ctx.GetText()))
			return nil
		}

		valueResult := ctx.DefaultValue().Accept(v)
		if v.err != nil {
			return nil
		}

		value, ok := valueResult.(objects.Value)
		if !ok || value == nil {
			v.setErr(fmt.Errorf("failed to parse default value for type %q", ctx.TypeName().GetText()))
			return nil
		}

		builder.DefaultValue(value)
	}

	typeRef, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return typeRef
}

func (v *objectVisitor) VisitNumberConstraint(ctx *generated.NumberConstraintContext) interface{} {
	builder := objects.NewNumberConstraintBuilder()

	switch {
	case ctx.PLUS() != nil:
		constraint, err := builder.
			Kind(objects.NumberConstraintOneOrMore).
			Build()
		if err != nil {
			v.setErr(err)
			return nil
		}

		return constraint

	case ctx.STAR() != nil:
		constraint, err := builder.
			Kind(objects.NumberConstraintZeroOrMore).
			Build()
		if err != nil {
			v.setErr(err)
			return nil
		}

		return constraint
	}

	builder.Kind(objects.NumberConstraintRange)

	text := strings.TrimSpace(ctx.GetText())
	text = strings.TrimPrefix(text, "[")
	text = strings.TrimSuffix(text, "]")

	parts := strings.Split(text, ",")
	if len(parts) != 2 {
		v.setErr(fmt.Errorf("invalid number range constraint: %s", ctx.GetText()))
		return nil
	}

	minRaw := strings.TrimSpace(parts[0])
	maxRaw := strings.TrimSpace(parts[1])

	if minRaw != "" {
		min, err := parseNumberValue(minRaw)
		if err != nil {
			v.setErr(err)
			return nil
		}

		builder.Min(min)
	}

	if maxRaw != "" {
		max, err := parseNumberValue(maxRaw)
		if err != nil {
			v.setErr(err)
			return nil
		}

		builder.Max(max)
	}

	constraint, err := builder.Build()
	if err != nil {
		v.setErr(err)
		return nil
	}

	return constraint
}

func (v *objectVisitor) VisitDefaultValue(ctx *generated.DefaultValueContext) interface{} {
	return ctx.Value().Accept(v)
}

func (v *objectVisitor) VisitValue(ctx *generated.ValueContext) interface{} {
	builder := objects.NewValueBuilder()

	switch {
	case ctx.STRING() != nil:
		value, err := builder.String(unquote(ctx.STRING().GetText()))
		if err != nil {
			v.setErr(err)
			return nil
		}

		return value

	case ctx.INT() != nil:
		value, err := builder.Int(ctx.INT().GetText())
		if err != nil {
			v.setErr(err)
			return nil
		}

		return value

	case ctx.FLOAT() != nil:
		value, err := builder.Float(ctx.FLOAT().GetText())
		if err != nil {
			v.setErr(err)
			return nil
		}

		return value

	case ctx.TRUE() != nil:
		value, err := builder.Bool("true")
		if err != nil {
			v.setErr(err)
			return nil
		}

		return value

	case ctx.FALSE() != nil:
		value, err := builder.Bool("false")
		if err != nil {
			v.setErr(err)
			return nil
		}

		return value
	}

	v.setErr(fmt.Errorf("unknown value: %s", ctx.GetText()))
	return nil
}

func parseNumberValue(raw string) (objects.NumberValue, error) {
	builder := objects.NewNumberValueBuilder()

	if strings.Contains(raw, ".") {
		return builder.Float(raw)
	}

	return builder.Int(raw)
}

func unquote(value string) string {
	return strings.Trim(value, `"`)
}
