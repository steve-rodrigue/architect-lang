package common

import (
	"fmt"
	"strconv"
)

type numberValueBuilder struct{}

func (b *numberValueBuilder) Int(raw string) (NumberValue, error) {
	if _, err := strconv.ParseInt(raw, 10, 64); err != nil {
		return nil, fmt.Errorf("invalid int value %q: %w", raw, err)
	}

	return &numberValue{
		raw:     raw,
		isFloat: false,
	}, nil
}

func (b *numberValueBuilder) Float(raw string) (NumberValue, error) {
	if _, err := strconv.ParseFloat(raw, 64); err != nil {
		return nil, fmt.Errorf("invalid float value %q: %w", raw, err)
	}

	return &numberValue{
		raw:     raw,
		isFloat: true,
	}, nil
}
