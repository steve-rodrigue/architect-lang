package endpoints

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type HTTPMethod string

const (
	HTTPMethodGet    HTTPMethod = "GET"
	HTTPMethodPost   HTTPMethod = "POST"
	HTTPMethodPut    HTTPMethod = "PUT"
	HTTPMethodPatch  HTTPMethod = "PATCH"
	HTTPMethodDelete HTTPMethod = "DELETE"
)

type InputSourceRuleKind string

const (
	InputSourceRuleSingle       InputSourceRuleKind = "single"
	InputSourceRuleExactlyOneOf InputSourceRuleKind = "exactly_one_of"
)

type InputSourceKind string

const (
	InputSourcePath   InputSourceKind = "path"
	InputSourceQuery  InputSourceKind = "query"
	InputSourceBody   InputSourceKind = "body"
	InputSourceHeader InputSourceKind = "header"
	InputSourceCookie InputSourceKind = "cookie"
)

// NewEndpointBuilder creates a new endpoint builder
func NewEndpointBuilder() EndpointBuilder {
	return &endpointBuilder{}
}

// NewInputBuilder creates a new input builder
func NewInputBuilder() InputBuilder {
	return &inputBuilder{}
}

// NewInputFieldBuilder creates a new input field builder
func NewInputFieldBuilder() InputFieldBuilder {
	return &inputFieldBuilder{}
}

// NewInputSourceRuleBuilder creates a new input source rule builder
func NewInputSourceRuleBuilder() InputSourceRuleBuilder {
	return &inputSourceRuleBuilder{}
}

// NewInputSourceBuilder creates a new input source builder
func NewInputSourceBuilder() InputSourceBuilder {
	return &inputSourceBuilder{}
}

// EndpointBuilder represents an endpoint builder
type EndpointBuilder interface {
	Name(name string) EndpointBuilder
	Method(method HTTPMethod) EndpointBuilder
	Path(path string) EndpointBuilder
	Input(input Input) EndpointBuilder
	AddAction(action workflows.Action) EndpointBuilder
	Build() (Endpoint, error)
}

// Endpoint represents an endpoint
type Endpoint interface {
	Name() string
	Method() HTTPMethod
	Path() string
	Input() Input
	Actions() []workflows.Action
}

// InputBuilder represents an input builder
type InputBuilder interface {
	AddField(field InputField) InputBuilder
	Build() (Input, error)
}

// Input represents an input
type Input interface {
	Fields() []InputField
}

// InputFieldBuilder represents an input field builder
type InputFieldBuilder interface {
	Name(name string) InputFieldBuilder
	Type(typeRef common.TypeReference) InputFieldBuilder
	Sources(rule InputSourceRule) InputFieldBuilder
	Build() (InputField, error)
}

// InputField represents an input field
type InputField interface {
	Name() string
	Type() common.TypeReference
	Sources() InputSourceRule
}

// InputSourceRuleBuilder represents an input source rule builder
type InputSourceRuleBuilder interface {
	Kind(kind InputSourceRuleKind) InputSourceRuleBuilder
	AddSource(source InputSource) InputSourceRuleBuilder
	Build() (InputSourceRule, error)
}

// InputSourceRule represents an input source rule
type InputSourceRule interface {
	Kind() InputSourceRuleKind
	Sources() []InputSource
}

// InputSourceBuilder represents an input source builder
type InputSourceBuilder interface {
	Kind(kind InputSourceKind) InputSourceBuilder
	Optional(optional bool) InputSourceBuilder
	Build() (InputSource, error)
}

// InputSource represents an input source
type InputSource interface {
	Kind() InputSourceKind
	IsOptional() bool
}
