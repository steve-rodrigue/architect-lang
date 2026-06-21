package model

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

// NewProjectBuilder creates a new resolved project builder.
func NewProjectBuilder() ProjectBuilder {
	return &projectBuilder{
		versions: make([]Version, 0),
	}
}

// NewVersionBuilder creates a new resolved project version builder.
func NewVersionBuilder() VersionBuilder {
	return &versionBuilder{
		objects:      make([]Object, 0),
		services:     make([]Service, 0),
		deployments:  make([]Deployment, 0),
		nextVersions: make([]Version, 0),
	}
}

// NewObjectBuilder creates a new resolved object builder.
func NewObjectBuilder() ObjectBuilder {
	return &objectBuilder{
		fields: make([]Field, 0),
	}
}

// NewFieldBuilder creates a new resolved field builder.
func NewFieldBuilder() FieldBuilder {
	return &fieldBuilder{
		modifiers: make([]objects.FieldModifier, 0),
	}
}

// NewTypeReferenceBuilder creates a new resolved type reference builder.
func NewTypeReferenceBuilder() TypeReferenceBuilder {
	return &typeReferenceBuilder{}
}

// NewServiceBuilder creates a new resolved service builder.
func NewServiceBuilder() ServiceBuilder {
	return &serviceBuilder{
		exposedPorts:  make([]int, 0),
		dependencies:  make([]Service, 0),
		storedObjects: make([]Object, 0),
		events:        make([]Event, 0),
	}
}

// NewApplicationBuilder creates a new resolved application builder.
func NewApplicationBuilder() ApplicationBuilder {
	return &applicationBuilder{
		ports:     make([]applications.Port, 0),
		endpoints: make([]Endpoint, 0),
		consumers: make([]Consumer, 0),
	}
}

// NewEndpointBuilder creates a new resolved endpoint builder.
func NewEndpointBuilder() EndpointBuilder {
	return &endpointBuilder{
		actions: make([]Action, 0),
	}
}

// NewInputBuilder creates a new resolved input builder.
func NewInputBuilder() InputBuilder {
	return &inputBuilder{
		fields: make([]InputField, 0),
	}
}

// NewInputFieldBuilder creates a new resolved input field builder.
func NewInputFieldBuilder() InputFieldBuilder {
	return &inputFieldBuilder{}
}

// NewConsumerBuilder creates a new resolved consumer builder.
func NewConsumerBuilder() ConsumerBuilder {
	return &consumerBuilder{
		actions: make([]Action, 0),
	}
}

// NewEventBuilder creates a new event builder.
func NewEventBuilder() EventBuilder {
	return &eventBuilder{}
}

// NewDeploymentBuilder creates a new resolved deployment builder.
func NewDeploymentBuilder() DeploymentBuilder {
	return &deploymentBuilder{
		services: make([]ServiceDeployment, 0),
	}
}

// NewServiceDeploymentBuilder creates a new resolved service deployment builder.
func NewServiceDeploymentBuilder() ServiceDeploymentBuilder {
	return &serviceDeploymentBuilder{}
}

// NewFetchActionBuilder creates a new resolved fetch action builder.
func NewFetchActionBuilder() FetchActionBuilder {
	return &fetchActionBuilder{}
}

// NewCreateActionBuilder creates a new resolved create action builder.
func NewCreateActionBuilder() CreateActionBuilder {
	return &createActionBuilder{
		assignments: make([]Assignment, 0),
	}
}

// NewStoreActionBuilder creates a new resolved store action builder.
func NewStoreActionBuilder() StoreActionBuilder {
	return &storeActionBuilder{}
}

// NewUpdateActionBuilder creates a new resolved update action builder.
func NewUpdateActionBuilder() UpdateActionBuilder {
	return &updateActionBuilder{
		assignments: make([]Assignment, 0),
	}
}

// NewEmitActionBuilder creates a new resolved emit action builder.
func NewEmitActionBuilder() EmitActionBuilder {
	return &emitActionBuilder{
		assignments: make([]Assignment, 0),
	}
}

// NewCallActionBuilder creates a new resolved call action builder.
func NewCallActionBuilder() CallActionBuilder {
	return &callActionBuilder{
		assignments: make([]Assignment, 0),
	}
}

// NewReturnActionBuilder creates a new resolved return action builder.
func NewReturnActionBuilder() ReturnActionBuilder {
	return &returnActionBuilder{}
}

// NewTypedVariableBuilder creates a new resolved typed variable builder.
func NewTypedVariableBuilder() TypedVariableBuilder {
	return &typedVariableBuilder{}
}

// NewAssignmentBuilder creates a new resolved assignment builder.
func NewAssignmentBuilder() AssignmentBuilder {
	return &assignmentBuilder{}
}

// NewConditionBuilder creates a new resolved condition builder.
func NewConditionBuilder() ConditionBuilder {
	return &conditionBuilder{}
}

// NewIdentifierExpressionBuilder creates a new resolved identifier expression builder.
func NewIdentifierExpressionBuilder() IdentifierExpressionBuilder {
	return &identifierExpressionBuilder{}
}

// NewSelectorExpressionBuilder creates a new resolved selector expression builder.
func NewSelectorExpressionBuilder() SelectorExpressionBuilder {
	return &selectorExpressionBuilder{
		parts: make([]string, 0),
	}
}

// NewLiteralExpressionBuilder creates a new resolved literal expression builder.
func NewLiteralExpressionBuilder() LiteralExpressionBuilder {
	return &literalExpressionBuilder{}
}

// NewFunctionCallExpressionBuilder creates a new resolved function call expression builder.
func NewFunctionCallExpressionBuilder() FunctionCallExpressionBuilder {
	return &functionCallExpressionBuilder{
		arguments: make([]Expression, 0),
	}
}

// ProjectBuilder represents a resolved project builder.
type ProjectBuilder interface {
	Name(name string) ProjectBuilder
	AddVersion(version Version) ProjectBuilder
	Build() (Project, error)
}

// Project represents a fully loaded project model.
type Project interface {
	Name() string
	Versions() []Version
}

// VersionBuilder represents a resolved project version builder.
type VersionBuilder interface {
	Number(number string) VersionBuilder
	AddObject(object Object) VersionBuilder
	AddService(service Service) VersionBuilder
	AddDeployment(deployment Deployment) VersionBuilder
	AddNextVersion(version Version) VersionBuilder
	Build() (Version, error)
}

// Version represents a fully loaded project version.
type Version interface {
	Number() string
	Objects() []Object
	Services() []Service
	Deployments() []Deployment
	NextVersions() []Version

	ObjectByName(name string) Object
	ServiceByName(name string) Service
	DeploymentByName(name string) Deployment
	EventByName(name string) Event
}

// ObjectBuilder represents a resolved object builder.
type ObjectBuilder interface {
	Name(name string) ObjectBuilder
	HistoryOf(object Object) ObjectBuilder
	AddField(field Field) ObjectBuilder
	AST(ast objects.Object) ObjectBuilder
	Build() (Object, error)
}

// Object represents a resolved domain object.
type Object interface {
	Name() string
	HistoryOf() Object
	Fields() []Field
	AST() objects.Object
}

// FieldBuilder represents a resolved field builder.
type FieldBuilder interface {
	Name(name string) FieldBuilder
	Type(typeRef TypeReference) FieldBuilder
	AddModifier(modifier objects.FieldModifier) FieldBuilder
	AST(ast objects.Field) FieldBuilder
	Build() (Field, error)
}

// Field represents a resolved object field.
type Field interface {
	Name() string
	Type() TypeReference
	Modifiers() []objects.FieldModifier
	AST() objects.Field
}

// TypeReferenceBuilder represents a resolved type reference builder.
type TypeReferenceBuilder interface {
	Name(name string) TypeReferenceBuilder
	Object(object Object) TypeReferenceBuilder
	List(isList bool) TypeReferenceBuilder
	Optional(isOptional bool) TypeReferenceBuilder
	NumberConstraint(constraint common.NumberConstraint) TypeReferenceBuilder
	DefaultValue(value common.Value) TypeReferenceBuilder
	AST(ast common.TypeReference) TypeReferenceBuilder
	Build() (TypeReference, error)
}

// TypeReference represents a resolved type reference.
type TypeReference interface {
	Name() string
	Object() Object
	IsObject() bool
	IsList() bool
	IsOptional() bool
	NumberConstraint() common.NumberConstraint
	DefaultValue() common.Value
	HasDefaultValue() bool
	AST() common.TypeReference
}

// ServiceBuilder represents a resolved service builder.
type ServiceBuilder interface {
	Name(name string) ServiceBuilder
	Kind(kind services.ServiceKind) ServiceBuilder
	Version(version string) ServiceBuilder
	Image(image string) ServiceBuilder
	Expose(port int) ServiceBuilder
	Application(application Application) ServiceBuilder
	AddDependency(service Service) ServiceBuilder
	AddStoredObject(object Object) ServiceBuilder
	AddEvent(event Event) ServiceBuilder
	AST(ast services.Service) ServiceBuilder
	Build() (Service, error)
}

// Service represents a resolved runtime service.
type Service interface {
	Name() string
	Kind() services.ServiceKind
	Version() string
	Image() string
	ExposedPorts() []int

	Application() Application
	HasApplication() bool

	Dependencies() []Service
	StoredObjects() []Object
	Events() []Event

	AST() services.Service
}

// ApplicationBuilder represents a resolved application builder.
type ApplicationBuilder interface {
	Name(name string) ApplicationBuilder
	AddPort(port applications.Port) ApplicationBuilder
	AddEndpoint(endpoint Endpoint) ApplicationBuilder
	AddConsumer(consumer Consumer) ApplicationBuilder
	AST(ast applications.Application) ApplicationBuilder
	Build() (Application, error)
}

// Application represents a resolved service application.
type Application interface {
	Name() string
	Ports() []applications.Port
	Endpoints() []Endpoint
	Consumers() []Consumer
	AST() applications.Application
}

// EndpointBuilder represents a resolved endpoint builder.
type EndpointBuilder interface {
	Name(name string) EndpointBuilder
	Method(method endpoints.HTTPMethod) EndpointBuilder
	Path(path string) EndpointBuilder
	Input(input Input) EndpointBuilder
	AddAction(action Action) EndpointBuilder
	AST(ast endpoints.Endpoint) EndpointBuilder
	Build() (Endpoint, error)
}

// Endpoint represents a resolved HTTP endpoint.
type Endpoint interface {
	Name() string
	Method() endpoints.HTTPMethod
	Path() string
	Input() Input
	Actions() []Action
	AST() endpoints.Endpoint
}

// InputBuilder represents a resolved input builder.
type InputBuilder interface {
	AddField(field InputField) InputBuilder
	AST(ast endpoints.Input) InputBuilder
	Build() (Input, error)
}

// Input represents resolved endpoint input.
type Input interface {
	Fields() []InputField
	AST() endpoints.Input
}

// InputFieldBuilder represents a resolved input field builder.
type InputFieldBuilder interface {
	Name(name string) InputFieldBuilder
	Type(typeRef TypeReference) InputFieldBuilder
	Sources(rule endpoints.InputSourceRule) InputFieldBuilder
	AST(ast endpoints.InputField) InputFieldBuilder
	Build() (InputField, error)
}

// InputField represents a resolved endpoint input field.
type InputField interface {
	Name() string
	Type() TypeReference
	Sources() endpoints.InputSourceRule
	AST() endpoints.InputField
}

// ConsumerBuilder represents a resolved consumer builder.
type ConsumerBuilder interface {
	Event(event Event) ConsumerBuilder
	AddAction(action Action) ConsumerBuilder
	ASTEventName(name string) ConsumerBuilder
	Build() (Consumer, error)
}

// Consumer represents a resolved event consumer.
type Consumer interface {
	Event() Event
	Actions() []Action
	ASTEventName() string
}

// EventBuilder represents an event builder.
type EventBuilder interface {
	Name(name string) EventBuilder
	DeclaredBy(service Service) EventBuilder
	Build() (Event, error)
}

// Event represents a declared or inferred event.
type Event interface {
	Name() string
	DeclaredBy() Service
}

// DeploymentBuilder represents a resolved deployment builder.
type DeploymentBuilder interface {
	Name(name string) DeploymentBuilder
	Vendor(vendor deployments.Vendor) DeploymentBuilder
	Inventory(file string) DeploymentBuilder
	AddService(service ServiceDeployment) DeploymentBuilder
	AST(ast deployments.Deployment) DeploymentBuilder
	Build() (Deployment, error)
}

// Deployment represents a resolved deployment environment.
type Deployment interface {
	Name() string
	Vendor() deployments.Vendor
	Inventory() string
	Services() []ServiceDeployment
	AST() deployments.Deployment
}

// ServiceDeploymentBuilder represents a resolved service deployment builder.
type ServiceDeploymentBuilder interface {
	Service(service Service) ServiceDeploymentBuilder
	Replicas(count int) ServiceDeploymentBuilder
	Domain(domain string) ServiceDeploymentBuilder
	Host(host string) ServiceDeploymentBuilder
	Volume(volume string) ServiceDeploymentBuilder
	Backup(enabled bool) ServiceDeploymentBuilder
	AST(ast deployments.ServiceDeployment) ServiceDeploymentBuilder
	Build() (ServiceDeployment, error)
}

// ServiceDeployment represents deployment settings for one resolved service.
type ServiceDeployment interface {
	Service() Service
	Replicas() int
	HasReplicas() bool
	Domain() string
	Host() string
	Volume() string
	Backup() bool
	HasBackup() bool
	AST() deployments.ServiceDeployment
}

// Action represents a resolved workflow action.
type Action interface {
	Kind() workflows.ActionKind
	AST() workflows.Action
}

// FetchActionBuilder represents a resolved fetch action builder.
type FetchActionBuilder interface {
	Variable(variable TypedVariable) FetchActionBuilder
	Source(service Service) FetchActionBuilder
	Condition(condition Condition) FetchActionBuilder
	AST(ast workflows.FetchAction) FetchActionBuilder
	Build() (FetchAction, error)
}

// FetchAction represents a resolved fetch action.
type FetchAction interface {
	Action
	Variable() TypedVariable
	Source() Service
	Condition() Condition
}

// CreateActionBuilder represents a resolved create action builder.
type CreateActionBuilder interface {
	Variable(variable TypedVariable) CreateActionBuilder
	AddAssignment(assignment Assignment) CreateActionBuilder
	AST(ast workflows.CreateAction) CreateActionBuilder
	Build() (CreateAction, error)
}

// CreateAction represents a resolved create action.
type CreateAction interface {
	Action
	Variable() TypedVariable
	Assignments() []Assignment
}

// StoreActionBuilder represents a resolved store action builder.
type StoreActionBuilder interface {
	VariableName(name string) StoreActionBuilder
	Destination(service Service) StoreActionBuilder
	AST(ast workflows.StoreAction) StoreActionBuilder
	Build() (StoreAction, error)
}

// StoreAction represents a resolved store action.
type StoreAction interface {
	Action
	VariableName() string
	Destination() Service
}

// UpdateActionBuilder represents a resolved update action builder.
type UpdateActionBuilder interface {
	VariableName(name string) UpdateActionBuilder
	AddAssignment(assignment Assignment) UpdateActionBuilder
	AST(ast workflows.UpdateAction) UpdateActionBuilder
	Build() (UpdateAction, error)
}

// UpdateAction represents a resolved update action.
type UpdateAction interface {
	Action
	VariableName() string
	Assignments() []Assignment
}

// EmitActionBuilder represents a resolved emit action builder.
type EmitActionBuilder interface {
	Event(event Event) EmitActionBuilder
	Variable(variable TypedVariable) EmitActionBuilder
	AddAssignment(assignment Assignment) EmitActionBuilder
	AST(ast workflows.EmitAction) EmitActionBuilder
	Build() (EmitAction, error)
}

// EmitAction represents a resolved emit action.
type EmitAction interface {
	Action
	Event() Event
	Variable() TypedVariable
	Assignments() []Assignment
}

// CallActionBuilder represents a resolved call action builder.
type CallActionBuilder interface {
	TargetService(service Service) CallActionBuilder
	TargetOperation(operation string) CallActionBuilder
	AddAssignment(assignment Assignment) CallActionBuilder
	AST(ast workflows.CallAction) CallActionBuilder
	Build() (CallAction, error)
}

// CallAction represents a resolved call action.
type CallAction interface {
	Action
	TargetService() Service
	TargetOperation() string
	Assignments() []Assignment
}

// ReturnActionBuilder represents a resolved return action builder.
type ReturnActionBuilder interface {
	Expression(expression Expression) ReturnActionBuilder
	AST(ast workflows.ReturnAction) ReturnActionBuilder
	Build() (ReturnAction, error)
}

// ReturnAction represents a resolved return action.
type ReturnAction interface {
	Action
	Expression() Expression
}

// TypedVariableBuilder represents a resolved typed variable builder.
type TypedVariableBuilder interface {
	Name(name string) TypedVariableBuilder
	Type(typeRef TypeReference) TypedVariableBuilder
	AST(ast workflows.TypedVariable) TypedVariableBuilder
	Build() (TypedVariable, error)
}

// TypedVariable represents a resolved typed variable.
type TypedVariable interface {
	Name() string
	Type() TypeReference
	AST() workflows.TypedVariable
}

// AssignmentBuilder represents a resolved assignment builder.
type AssignmentBuilder interface {
	Name(name string) AssignmentBuilder
	Value(value Expression) AssignmentBuilder
	AST(ast workflows.Assignment) AssignmentBuilder
	Build() (Assignment, error)
}

// Assignment represents a resolved assignment.
type Assignment interface {
	Name() string
	Value() Expression
	AST() workflows.Assignment
}

// ConditionBuilder represents a resolved condition builder.
type ConditionBuilder interface {
	Left(expression Expression) ConditionBuilder
	Operator(operator workflows.Comparator) ConditionBuilder
	Right(expression Expression) ConditionBuilder
	AST(ast workflows.Condition) ConditionBuilder
	Build() (Condition, error)
}

// Condition represents a resolved condition.
type Condition interface {
	Left() Expression
	Operator() workflows.Comparator
	Right() Expression
	AST() workflows.Condition
}

// Expression represents a resolved expression.
type Expression interface {
	Kind() workflows.ExpressionKind
	AST() workflows.Expression
}

// IdentifierExpressionBuilder represents a resolved identifier expression builder.
type IdentifierExpressionBuilder interface {
	Name(name string) IdentifierExpressionBuilder
	AST(ast workflows.IdentifierExpression) IdentifierExpressionBuilder
	Build() (IdentifierExpression, error)
}

// IdentifierExpression represents a resolved identifier expression
type IdentifierExpression interface {
	Expression
	Name() string
}

// SelectorExpressionBuilder represents a resolved selector expression builder.
type SelectorExpressionBuilder interface {
	AddPart(part string) SelectorExpressionBuilder
	AST(ast workflows.SelectorExpression) SelectorExpressionBuilder
	Build() (SelectorExpression, error)
}

// SelectorExpression represents a resolved selector expression.
type SelectorExpression interface {
	Expression
	Parts() []string
}

// LiteralExpressionBuilder represents a resolved literal expression builder.
type LiteralExpressionBuilder interface {
	Value(value common.Value) LiteralExpressionBuilder
	AST(ast workflows.LiteralExpression) LiteralExpressionBuilder
	Build() (LiteralExpression, error)
}

// LiteralExpression represents a resolved literal expression
type LiteralExpression interface {
	Expression
	Value() common.Value
}

// FunctionCallExpressionBuilder represents a resolved function call expression builder.
type FunctionCallExpressionBuilder interface {
	Name(name string) FunctionCallExpressionBuilder
	AddArgument(argument Expression) FunctionCallExpressionBuilder
	AST(ast workflows.FunctionCallExpression) FunctionCallExpressionBuilder
	Build() (FunctionCallExpression, error)
}

// FunctionCallExpression represents a resolved function call expression
type FunctionCallExpression interface {
	Expression
	Name() string
	Arguments() []Expression
}
