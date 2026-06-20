package objects

import (
	"fmt"
	"strconv"
)

type valueBuilder struct{}

func (b *valueBuilder) String(raw string) (Value, error) {
	return &value{
		kind: ValueString,
		raw:  raw,
	}, nil
}

func (b *valueBuilder) Int(raw string) (Value, error) {
	if _, err := strconv.ParseInt(raw, 10, 64); err != nil {
		return nil, fmt.Errorf("invalid int value %q: %w", raw, err)
	}

	return &value{
		kind: ValueInt,
		raw:  raw,
	}, nil
}

func (b *valueBuilder) Float(raw string) (Value, error) {
	if _, err := strconv.ParseFloat(raw, 64); err != nil {
		return nil, fmt.Errorf("invalid float value %q: %w", raw, err)
	}

	return &value{
		kind: ValueFloat,
		raw:  raw,
	}, nil
}

func (b *valueBuilder) Bool(raw string) (Value, error) {
	switch raw {
	case "true", "false":
		return &value{
			kind: ValueBool,
			raw:  raw,
		}, nil
	default:
		return nil, fmt.Errorf("invalid bool value %q", raw)
	}
}
