package antlr

import (
	"fmt"
	"strings"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
)

func buildCommonTypeReference(
	name string,
	isList bool,
	isOptional bool,
	constraint common.NumberConstraint,
	defaultValue common.Value,
	hasDefaultValue bool,
	setErr func(error),
) common.TypeReference {
	builder := common.NewTypeReferenceBuilder().
		Name(name).
		List(isList).
		Optional(isOptional)

	if constraint == nil {
		var err error
		constraint, err = common.NewNumberConstraintBuilder().
			Kind(common.NumberConstraintNone).
			Build()
		if err != nil {
			setErr(err)
			return nil
		}
	}

	builder.NumberConstraint(constraint)

	if hasDefaultValue {
		builder.DefaultValue(defaultValue)
	}

	typeRef, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return typeRef
}

func buildCommonNumberConstraint(text string, hasPlus bool, hasStar bool, setErr func(error)) common.NumberConstraint {
	builder := common.NewNumberConstraintBuilder()

	switch {
	case hasPlus:
		constraint, err := builder.Kind(common.NumberConstraintOneOrMore).Build()
		if err != nil {
			setErr(err)
			return nil
		}
		return constraint

	case hasStar:
		constraint, err := builder.Kind(common.NumberConstraintZeroOrMore).Build()
		if err != nil {
			setErr(err)
			return nil
		}
		return constraint
	}

	builder.Kind(common.NumberConstraintRange)

	text = strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(text), "["), "]")
	parts := strings.Split(text, ",")
	if len(parts) != 2 {
		setErr(fmt.Errorf("invalid number range constraint: %s", text))
		return nil
	}

	if minRaw := strings.TrimSpace(parts[0]); minRaw != "" {
		min, err := parseCommonNumberValue(minRaw)
		if err != nil {
			setErr(err)
			return nil
		}
		builder.Min(min)
	}

	if maxRaw := strings.TrimSpace(parts[1]); maxRaw != "" {
		max, err := parseCommonNumberValue(maxRaw)
		if err != nil {
			setErr(err)
			return nil
		}
		builder.Max(max)
	}

	constraint, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return constraint
}

func buildCommonValue(kind string, raw string, setErr func(error)) common.Value {
	builder := common.NewValueBuilder()

	var (
		value common.Value
		err   error
	)

	switch kind {
	case "string":
		value, err = builder.String(unquote(raw))
	case "int":
		value, err = builder.Int(raw)
	case "float":
		value, err = builder.Float(raw)
	case "bool":
		value, err = builder.Bool(raw)
	default:
		setErr(fmt.Errorf("unknown value kind %q", kind))
		return nil
	}

	if err != nil {
		setErr(err)
		return nil
	}

	return value
}

func parseCommonNumberValue(raw string) (common.NumberValue, error) {
	builder := common.NewNumberValueBuilder()

	if strings.Contains(raw, ".") {
		return builder.Float(raw)
	}

	return builder.Int(raw)
}
