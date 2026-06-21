package endpoints

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"

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

type ActionKind string

const (
	ActionKindFetch  ActionKind = "fetch"
	ActionKindCreate ActionKind = "create"
	ActionKindStore  ActionKind = "store"
	ActionKindUpdate ActionKind = "update"
	ActionKindEmit   ActionKind = "emit"
	ActionKindReturn ActionKind = "return"
)

type Comparator string

const (
	ComparatorEqual              Comparator = "=="
	ComparatorNotEqual           Comparator = "!="
	ComparatorLessThan           Comparator = "<"
	ComparatorLessThanOrEqual    Comparator = "<="
	ComparatorGreaterThan        Comparator = ">"
	ComparatorGreaterThanOrEqual Comparator = ">="
)

type ExpressionKind string

const (
	ExpressionKindIdentifier   ExpressionKind = "identifier"
	ExpressionKindSelector     ExpressionKind = "selector"
	ExpressionKindLiteral      ExpressionKind = "literal"
	ExpressionKindFunctionCall ExpressionKind = "function_call"
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

// NewTypedVariableBuilder creates a new typed variable builder
func NewTypedVariableBuilder() TypedVariableBuilder {
	return &typedVariableBuilder{}
}

// NewFetchActionBuilder creates a new fetch action builder
func NewFetchActionBuilder() FetchActionBuilder {
	return &fetchActionBuilder{}
}

// NewCreateActionBuilder creates a new create action builder
func NewCreateActionBuilder() CreateActionBuilder {
	return &createActionBuilder{}
}

// NewStoreActionBuilder creates a new store action builder
func NewStoreActionBuilder() StoreActionBuilder {
	return &storeActionBuilder{}
}

// NewUpdateActionBuilder creates a new update action builder
func NewUpdateActionBuilder() UpdateActionBuilder {
	return &updateActionBuilder{}
}

// NewEmitActionBuilder creates a new emit action builder
func NewEmitActionBuilder() EmitActionBuilder {
	return &emitActionBuilder{}
}

// NewReturnActionBuilder creates a new return action builder
func NewReturnActionBuilder() ReturnActionBuilder {
	return &returnActionBuilder{}
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return &assignmentBuilder{}
}

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	return &conditionBuilder{}
}

// NewIdentifierExpressionBuilder creates a new identifier expression builder
func NewIdentifierExpressionBuilder() IdentifierExpressionBuilder {
	return &identifierExpressionBuilder{}
}

// NewSelectorExpressionBuilder creates a new selector expression builder
func NewSelectorExpressionBuilder() SelectorExpressionBuilder {
	return &selectorExpressionBuilder{}
}

// NewLiteralExpressionBuilder creates a new literal expression builder
func NewLiteralExpressionBuilder() LiteralExpressionBuilder {
	return &literalExpressionBuilder{}
}

// NewFunctionCallExpressionBuilder creates a new function call expression builder
func NewFunctionCallExpressionBuilder() FunctionCallExpressionBuilder {
	return &functionCallExpressionBuilder{}
}

// EndpointBuilder represents an endpoint builder
type EndpointBuilder interface {
	Name(name string) EndpointBuilder
	Method(method HTTPMethod) EndpointBuilder
	Path(path string) EndpointBuilder
	Input(input Input) EndpointBuilder
	AddAction(action Action) EndpointBuilder
	Build() (Endpoint, error)
}

// Endpoint represents an endpoint
type Endpoint interface {
	Name() string
	Method() HTTPMethod
	Path() string
	Input() Input
	Actions() []Action
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
	Type(typeRef objects.TypeReference) InputFieldBuilder
	Sources(rule InputSourceRule) InputFieldBuilder
	Build() (InputField, error)
}

// InputField represents an input field
type InputField interface {
	Name() string
	Type() objects.TypeReference
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

// Action represents an action
type Action interface {
	Kind() ActionKind
}

// TypedVariableBuilder represents a typed variable builder
type TypedVariableBuilder interface {
	Name(name string) TypedVariableBuilder
	Type(typeRef objects.TypeReference) TypedVariableBuilder
	Build() (TypedVariable, error)
}

// TypedVariable represents a typed variable
type TypedVariable interface {
	Name() string
	Type() objects.TypeReference
}

// FetchActionBuilder represents a fetch action builder
type FetchActionBuilder interface {
	Variable(variable TypedVariable) FetchActionBuilder
	Source(source string) FetchActionBuilder
	Condition(condition Condition) FetchActionBuilder
	Build() (FetchAction, error)
}

// FetchAction represents a fetch action
type FetchAction interface {
	Action
	Variable() TypedVariable
	Source() string
	Condition() Condition
}

// CreateActionBuilder represents a create action builder
type CreateActionBuilder interface {
	Variable(variable TypedVariable) CreateActionBuilder
	AddAssignment(assignment Assignment) CreateActionBuilder
	Build() (CreateAction, error)
}

// CreateAction represents a create action
type CreateAction interface {
	Action
	Variable() TypedVariable
	Assignments() []Assignment
}

// StoreActionBuilder represents a store action builder
type StoreActionBuilder interface {
	VariableName(name string) StoreActionBuilder
	Destination(destination string) StoreActionBuilder
	Build() (StoreAction, error)
}

// StoreAction represents a store action
type StoreAction interface {
	Action
	VariableName() string
	Destination() string
}

// UpdateActionBuilder represents an update action builder
type UpdateActionBuilder interface {
	VariableName(name string) UpdateActionBuilder
	AddAssignment(assignment Assignment) UpdateActionBuilder
	Build() (UpdateAction, error)
}

// UpdateAction represents an update action
type UpdateAction interface {
	Action
	VariableName() string
	Assignments() []Assignment
}

// EmitActionBuilder represents an emit action builder
type EmitActionBuilder interface {
	Variable(variable TypedVariable) EmitActionBuilder
	AddAssignment(assignment Assignment) EmitActionBuilder
	Build() (EmitAction, error)
}

// EmitAction represents an emit action
type EmitAction interface {
	Action
	Variable() TypedVariable
	Assignments() []Assignment
}

// ReturnActionBuilder represents a return action builder
type ReturnActionBuilder interface {
	Expression(expression Expression) ReturnActionBuilder
	Build() (ReturnAction, error)
}

// ReturnAction represents a return action
type ReturnAction interface {
	Action
	Expression() Expression
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Name(name string) AssignmentBuilder
	Value(value Expression) AssignmentBuilder
	Build() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Value() Expression
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Left(expression Expression) ConditionBuilder
	Operator(operator Comparator) ConditionBuilder
	Right(expression Expression) ConditionBuilder
	Build() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Left() Expression
	Operator() Comparator
	Right() Expression
}

// Expression represents an expression
type Expression interface {
	Kind() ExpressionKind
}

// IdentifierExpressionBuilder represents an identifier expression builder
type IdentifierExpressionBuilder interface {
	Name(name string) IdentifierExpressionBuilder
	Build() (IdentifierExpression, error)
}

// IdentifierExpression represents an identifier expression
type IdentifierExpression interface {
	Expression
	Name() string
}

// SelectorExpressionBuilder represents a selector expression builder
type SelectorExpressionBuilder interface {
	AddPart(part string) SelectorExpressionBuilder
	Build() (SelectorExpression, error)
}

// SelectorExpression represents a selector expression
type SelectorExpression interface {
	Expression
	Parts() []string
}

// LiteralExpressionBuilder represents a literal expression builder
type LiteralExpressionBuilder interface {
	Value(value objects.Value) LiteralExpressionBuilder
	Build() (LiteralExpression, error)
}

// LiteralExpression represents a literal expression
type LiteralExpression interface {
	Expression
	Value() objects.Value
}

// FunctionCallExpressionBuilder represents a function call expression builder
type FunctionCallExpressionBuilder interface {
	Name(name string) FunctionCallExpressionBuilder
	AddArgument(argument Expression) FunctionCallExpressionBuilder
	Build() (FunctionCallExpression, error)
}

// FunctionCallExpression represents a function call expression
type FunctionCallExpression interface {
	Expression
	Name() string
	Arguments() []Expression
}
