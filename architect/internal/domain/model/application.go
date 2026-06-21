package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"

type application struct {
	name      string
	ports     []applications.Port
	endpoints []Endpoint
	consumers []Consumer
	ast       applications.Application
}

func (a *application) Name() string {
	return a.name
}

func (a *application) Ports() []applications.Port {
	return append([]applications.Port(nil), a.ports...)
}

func (a *application) Endpoints() []Endpoint {
	return append([]Endpoint(nil), a.endpoints...)
}

func (a *application) Consumers() []Consumer {
	return append([]Consumer(nil), a.consumers...)
}

func (a *application) AST() applications.Application {
	return a.ast
}
