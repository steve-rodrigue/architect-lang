package applications

type port struct {
	direction PortDirection
	kind      PortKind
	number    int
}

func (p *port) Direction() PortDirection {
	return p.direction
}

func (p *port) Kind() PortKind {
	return p.kind
}

func (p *port) Number() int {
	return p.number
}

type portBuilder struct {
	direction PortDirection
	kind      PortKind
	number    int
}
