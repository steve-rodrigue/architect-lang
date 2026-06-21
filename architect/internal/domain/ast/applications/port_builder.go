package applications

import "fmt"

func (b *portBuilder) Direction(direction PortDirection) PortBuilder {
	b.direction = direction
	return b
}

func (b *portBuilder) Kind(kind PortKind) PortBuilder {
	b.kind = kind
	return b
}

func (b *portBuilder) Number(number int) PortBuilder {
	b.number = number
	return b
}

func (b *portBuilder) Build() (Port, error) {
	if b.direction == "" {
		return nil, fmt.Errorf("port direction is required")
	}

	if b.kind == "" {
		return nil, fmt.Errorf("port kind is required")
	}

	if b.number <= 0 {
		return nil, fmt.Errorf("port number must be greater than 0")
	}

	return &port{
		direction: b.direction,
		kind:      b.kind,
		number:    b.number,
	}, nil
}
