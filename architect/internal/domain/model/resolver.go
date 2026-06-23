package model

import (
	"fmt"
	"strings"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model/sources"
)

type resolver struct{}

type versionContext struct {
	objects     map[string]Object
	services    map[string]Service
	deployments map[string]Deployment
	events      map[string]Event
}

func (r *resolver) Resolve(source sources.Project) (Project, error) {
	if source == nil {
		return nil, fmt.Errorf("project source is required")
	}
	if source.Project() == nil {
		return nil, fmt.Errorf("project AST is required")
	}

	builder := NewProjectBuilder().
		Name(source.Project().Name())

	for _, versionSource := range source.Versions() {
		version, err := r.resolveVersion(versionSource)
		if err != nil {
			return nil, err
		}

		builder.AddVersion(version)
	}

	return builder.Build()
}

func (r *resolver) resolveVersion(source sources.Version) (Version, error) {
	if source == nil {
		return nil, fmt.Errorf("version source is required")
	}
	if source.Version() == nil {
		return nil, fmt.Errorf("version AST is required")
	}

	ctx := &versionContext{
		objects:     make(map[string]Object),
		services:    make(map[string]Service),
		deployments: make(map[string]Deployment),
		events:      make(map[string]Event),
	}

	for _, objectAST := range source.Objects() {
		object, err := r.resolveObjectSkeleton(objectAST)
		if err != nil {
			return nil, err
		}
		if _, exists := ctx.objects[object.Name()]; exists {
			return nil, fmt.Errorf("duplicate object %q", object.Name())
		}

		ctx.objects[object.Name()] = object
	}

	for _, objectAST := range source.Objects() {
		object, err := r.resolveObject(objectAST, ctx)
		if err != nil {
			return nil, err
		}

		ctx.objects[object.Name()] = object
	}

	for _, serviceSource := range source.Services() {
		service, err := r.resolveServiceSkeleton(serviceSource)
		if err != nil {
			return nil, err
		}
		if _, exists := ctx.services[service.Name()]; exists {
			return nil, fmt.Errorf("duplicate service %q", service.Name())
		}

		ctx.services[service.Name()] = service
	}

	for _, serviceSource := range source.Services() {
		for _, eventName := range serviceSource.Service().Events() {
			if _, exists := ctx.events[eventName]; exists {
				continue
			}

			event, err := NewEventBuilder().
				Name(eventName).
				DeclaredBy(ctx.services[serviceSource.Service().Name()]).
				Build()
			if err != nil {
				return nil, err
			}

			ctx.events[eventName] = event
		}
	}

	for _, serviceSource := range source.Services() {
		service, err := r.resolveService(serviceSource, ctx)
		if err != nil {
			return nil, err
		}

		ctx.services[service.Name()] = service
	}

	for _, deploymentAST := range source.Deployments() {
		deployment, err := r.resolveDeployment(deploymentAST, ctx)
		if err != nil {
			return nil, err
		}
		if _, exists := ctx.deployments[deployment.Name()]; exists {
			return nil, fmt.Errorf("duplicate deployment %q", deployment.Name())
		}

		ctx.deployments[deployment.Name()] = deployment
	}

	builder := NewVersionBuilder().
		Number(source.Version().Number())

	for _, object := range ctx.objects {
		builder.AddObject(object)
	}
	for _, service := range ctx.services {
		builder.AddService(service)
	}
	for _, deployment := range ctx.deployments {
		builder.AddDeployment(deployment)
	}

	for _, nextVersionAST := range source.Version().NextVersions() {
		nextVersion, err := NewVersionBuilder().
			Number(nextVersionAST.Number()).
			Build()
		if err != nil {
			return nil, err
		}

		builder.AddNextVersion(nextVersion)
	}

	return builder.Build()
}

func (r *resolver) resolveObjectSkeleton(objectAST objects.Object) (Object, error) {
	return NewObjectBuilder().
		Name(objectAST.Name()).
		AST(objectAST).
		Build()
}

func (r *resolver) resolveObject(objectAST objects.Object, ctx *versionContext) (Object, error) {
	builder := NewObjectBuilder().
		Name(objectAST.Name()).
		AST(objectAST)

	if objectAST.HistoryOf() != nil {
		historyOf := ctx.objects[objectAST.HistoryOf().Name()]
		if historyOf == nil {
			return nil, fmt.Errorf("object %q history_of references unknown object %q", objectAST.Name(), objectAST.HistoryOf().Name())
		}

		builder.HistoryOf(historyOf)
	}

	for _, fieldAST := range objectAST.Fields() {
		field, err := r.resolveField(fieldAST, ctx)
		if err != nil {
			return nil, err
		}

		builder.AddField(field)
	}

	return builder.Build()
}

func (r *resolver) resolveField(fieldAST objects.Field, ctx *versionContext) (Field, error) {
	typeRef, err := r.resolveTypeReference(fieldAST.Type(), ctx)
	if err != nil {
		return nil, err
	}

	builder := NewFieldBuilder().
		Name(fieldAST.Name()).
		Type(typeRef).
		AST(fieldAST)

	for _, modifier := range fieldAST.Modifiers() {
		builder.AddModifier(modifier)
	}

	return builder.Build()
}

func (r *resolver) resolveTypeReference(typeAST common.TypeReference, ctx *versionContext) (TypeReference, error) {
	builder := NewTypeReferenceBuilder().
		Name(typeAST.Name()).
		List(typeAST.IsList()).
		Optional(typeAST.IsOptional()).
		NumberConstraint(typeAST.NumberConstraint()).
		DefaultValue(typeAST.DefaultValue()).
		AST(typeAST)

	if object := ctx.objects[typeAST.Name()]; object != nil {
		builder.Object(object)
	}

	return builder.Build()
}

func (r *resolver) resolveServiceSkeleton(source sources.Service) (Service, error) {
	if source == nil {
		return nil, fmt.Errorf("service source is required")
	}
	if source.Service() == nil {
		return nil, fmt.Errorf("service AST is required")
	}

	builder := NewServiceBuilder().
		Name(source.Service().Name()).
		Kind(source.Service().Kind()).
		Version(source.Service().Version()).
		Image(source.Service().Image()).
		AST(source.Service())

	for _, port := range source.Service().ExposedPorts() {
		builder.Expose(port)
	}

	return builder.Build()
}

func (r *resolver) resolveService(source sources.Service, ctx *versionContext) (Service, error) {
	serviceAST := source.Service()

	builder := NewServiceBuilder().
		Name(serviceAST.Name()).
		Kind(serviceAST.Kind()).
		Version(serviceAST.Version()).
		Image(serviceAST.Image()).
		AST(serviceAST)

	for _, port := range serviceAST.ExposedPorts() {
		builder.Expose(port)
	}

	for _, dependencyName := range serviceAST.DependsOn() {
		dependency := ctx.services[dependencyName]
		if dependency == nil {
			return nil, fmt.Errorf("service %q depends on unknown service %q", serviceAST.Name(), dependencyName)
		}

		builder.AddDependency(dependency)
	}

	for _, objectName := range serviceAST.Stores() {
		object := ctx.objects[objectName]
		if object == nil {
			return nil, fmt.Errorf("service %q stores unknown object %q", serviceAST.Name(), objectName)
		}

		builder.AddStoredObject(object)
	}

	for _, eventName := range serviceAST.Events() {
		event := ctx.events[eventName]
		if event == nil {
			return nil, fmt.Errorf("service %q declares unknown event %q", serviceAST.Name(), eventName)
		}

		builder.AddEvent(event)
	}

	if source.HasApplication() {
		application, err := r.resolveApplication(source, ctx)
		if err != nil {
			return nil, err
		}

		builder.Application(application)
	}

	return builder.Build()
}

func (r *resolver) resolveApplication(source sources.Service, ctx *versionContext) (Application, error) {
	applicationAST := source.Application()

	builder := NewApplicationBuilder().
		Name(applicationAST.Name()).
		AST(applicationAST)

	for _, port := range applicationAST.Ports() {
		builder.AddPort(port)
	}

	for _, endpointAST := range source.Endpoints() {
		endpoint, err := r.resolveEndpoint(endpointAST, ctx)
		if err != nil {
			return nil, err
		}

		builder.AddEndpoint(endpoint)
	}

	for _, consumerAST := range source.Consumers() {
		consumer, err := r.resolveConsumer(consumerAST, ctx)
		if err != nil {
			return nil, err
		}

		builder.AddConsumer(consumer)
	}

	return builder.Build()
}

func (r *resolver) resolveEndpoint(endpointAST endpoints.Endpoint, ctx *versionContext) (Endpoint, error) {
	input, err := r.resolveInput(endpointAST.Input(), ctx)
	if err != nil {
		return nil, err
	}

	actions, err := r.resolveActions(endpointAST.Actions(), ctx)
	if err != nil {
		return nil, err
	}

	builder := NewEndpointBuilder().
		Name(endpointAST.Name()).
		Method(endpointAST.Method()).
		Path(endpointAST.Path()).
		Input(input).
		AST(endpointAST)

	for _, action := range actions {
		builder.AddAction(action)
	}

	return builder.Build()
}

func (r *resolver) resolveInput(inputAST endpoints.Input, ctx *versionContext) (Input, error) {
	builder := NewInputBuilder().
		AST(inputAST)

	for _, fieldAST := range inputAST.Fields() {
		typeRef, err := r.resolveTypeReference(fieldAST.Type(), ctx)
		if err != nil {
			return nil, err
		}

		field, err := NewInputFieldBuilder().
			Name(fieldAST.Name()).
			Type(typeRef).
			Sources(fieldAST.Sources()).
			AST(fieldAST).
			Build()
		if err != nil {
			return nil, err
		}

		builder.AddField(field)
	}

	return builder.Build()
}

func (r *resolver) resolveConsumer(consumerAST consumers.Consumer, ctx *versionContext) (Consumer, error) {
	event := ctx.events[consumerAST.EventName()]
	if event == nil {
		created, err := NewEventBuilder().
			Name(consumerAST.EventName()).
			Build()
		if err != nil {
			return nil, err
		}

		event = created
		ctx.events[consumerAST.EventName()] = event
	}

	actions, err := r.resolveActions(consumerAST.Actions(), ctx)
	if err != nil {
		return nil, err
	}

	builder := NewConsumerBuilder().
		Event(event).
		ASTEventName(consumerAST.EventName())

	for _, action := range actions {
		builder.AddAction(action)
	}

	return builder.Build()
}

func (r *resolver) resolveDeployment(deploymentAST deployments.Deployment, ctx *versionContext) (Deployment, error) {
	builder := NewDeploymentBuilder().
		Name(deploymentAST.Name()).
		Vendor(deploymentAST.Vendor()).
		Inventory(deploymentAST.Inventory()).
		AST(deploymentAST)

	for _, serviceDeploymentAST := range deploymentAST.Services() {
		service := ctx.services[serviceDeploymentAST.Name()]
		if service == nil {
			return nil, fmt.Errorf("deployment %q references unknown service %q", deploymentAST.Name(), serviceDeploymentAST.Name())
		}

		serviceDeploymentBuilder := NewServiceDeploymentBuilder().
			Service(service).
			Domain(serviceDeploymentAST.Domain()).
			Host(serviceDeploymentAST.Host()).
			Volume(serviceDeploymentAST.Volume()).
			AST(serviceDeploymentAST)

		if serviceDeploymentAST.HasReplicas() {
			serviceDeploymentBuilder.Replicas(serviceDeploymentAST.Replicas())
		}

		if serviceDeploymentAST.HasBackup() {
			serviceDeploymentBuilder.Backup(serviceDeploymentAST.Backup())
		}

		serviceDeployment, err := serviceDeploymentBuilder.Build()
		if err != nil {
			return nil, err
		}

		builder.AddService(serviceDeployment)
	}

	return builder.Build()
}

func (r *resolver) resolveActions(actions []workflows.Action, ctx *versionContext) ([]Action, error) {
	resolved := make([]Action, 0, len(actions))

	for _, action := range actions {
		resolvedAction, err := r.resolveAction(action, ctx)
		if err != nil {
			return nil, err
		}

		resolved = append(resolved, resolvedAction)
	}

	return resolved, nil
}

func (r *resolver) resolveAction(action workflows.Action, ctx *versionContext) (Action, error) {
	switch value := action.(type) {
	case workflows.FetchAction:
		return r.resolveFetchAction(value, ctx)
	case workflows.CreateAction:
		return r.resolveCreateAction(value, ctx)
	case workflows.StoreAction:
		return r.resolveStoreAction(value, ctx)
	case workflows.UpdateAction:
		return r.resolveUpdateAction(value, ctx)
	case workflows.EmitAction:
		return r.resolveEmitAction(value, ctx)
	case workflows.CallAction:
		return r.resolveCallAction(value, ctx)
	case workflows.ReturnAction:
		return r.resolveReturnAction(value, ctx)
	default:
		return nil, fmt.Errorf("unsupported workflow action kind %q", action.Kind())
	}
}

func (r *resolver) resolveFetchAction(action workflows.FetchAction, ctx *versionContext) (FetchAction, error) {
	variable, err := r.resolveTypedVariable(action.Variable(), ctx)
	if err != nil {
		return nil, err
	}

	source := ctx.services[action.Source()]
	if source == nil {
		return nil, fmt.Errorf("fetch references unknown source service %q", action.Source())
	}

	condition, err := r.resolveCondition(action.Condition(), ctx)
	if err != nil {
		return nil, err
	}

	return NewFetchActionBuilder().
		Variable(variable).
		Source(source).
		Condition(condition).
		AST(action).
		Build()
}

func (r *resolver) resolveCreateAction(action workflows.CreateAction, ctx *versionContext) (CreateAction, error) {
	variable, err := r.resolveTypedVariable(action.Variable(), ctx)
	if err != nil {
		return nil, err
	}

	builder := NewCreateActionBuilder().
		Variable(variable).
		AST(action)

	for _, assignmentAST := range action.Assignments() {
		assignment, err := r.resolveAssignment(assignmentAST, ctx)
		if err != nil {
			return nil, err
		}

		builder.AddAssignment(assignment)
	}

	return builder.Build()
}

func (r *resolver) resolveStoreAction(action workflows.StoreAction, ctx *versionContext) (StoreAction, error) {
	destination := ctx.services[action.Destination()]
	if destination == nil {
		return nil, fmt.Errorf("store references unknown destination service %q", action.Destination())
	}

	return NewStoreActionBuilder().
		VariableName(action.VariableName()).
		Destination(destination).
		AST(action).
		Build()
}

func (r *resolver) resolveUpdateAction(action workflows.UpdateAction, ctx *versionContext) (UpdateAction, error) {
	builder := NewUpdateActionBuilder().
		VariableName(action.VariableName()).
		AST(action)

	for _, assignmentAST := range action.Assignments() {
		assignment, err := r.resolveAssignment(assignmentAST, ctx)
		if err != nil {
			return nil, err
		}

		builder.AddAssignment(assignment)
	}

	return builder.Build()
}

func (r *resolver) resolveEmitAction(action workflows.EmitAction, ctx *versionContext) (EmitAction, error) {
	variable, err := r.resolveTypedVariable(action.Variable(), ctx)
	if err != nil {
		return nil, err
	}

	eventName := action.Variable().Type().Name()
	event := ctx.events[eventName]
	if event == nil {
		created, err := NewEventBuilder().
			Name(eventName).
			Build()
		if err != nil {
			return nil, err
		}

		event = created
		ctx.events[eventName] = event
	}

	builder := NewEmitActionBuilder().
		Event(event).
		Variable(variable).
		AST(action)

	for _, assignmentAST := range action.Assignments() {
		assignment, err := r.resolveAssignment(assignmentAST, ctx)
		if err != nil {
			return nil, err
		}

		builder.AddAssignment(assignment)
	}

	return builder.Build()
}

func (r *resolver) resolveCallAction(action workflows.CallAction, ctx *versionContext) (CallAction, error) {
	serviceName, operation := splitCallTarget(action.Target())

	service := ctx.services[serviceName]
	if service == nil {
		return nil, fmt.Errorf("call references unknown service %q", serviceName)
	}

	builder := NewCallActionBuilder().
		TargetService(service).
		TargetOperation(operation).
		AST(action)

	for _, assignmentAST := range action.Assignments() {
		assignment, err := r.resolveAssignment(assignmentAST, ctx)
		if err != nil {
			return nil, err
		}

		builder.AddAssignment(assignment)
	}

	return builder.Build()
}

func (r *resolver) resolveReturnAction(action workflows.ReturnAction, ctx *versionContext) (ReturnAction, error) {
	expression, err := r.resolveExpression(action.Expression(), ctx)
	if err != nil {
		return nil, err
	}

	return NewReturnActionBuilder().
		Expression(expression).
		AST(action).
		Build()
}

func (r *resolver) resolveTypedVariable(variable workflows.TypedVariable, ctx *versionContext) (TypedVariable, error) {
	typeRef, err := r.resolveTypeReference(variable.Type(), ctx)
	if err != nil {
		return nil, err
	}

	return NewTypedVariableBuilder().
		Name(variable.Name()).
		Type(typeRef).
		AST(variable).
		Build()
}

func (r *resolver) resolveAssignment(assignment workflows.Assignment, ctx *versionContext) (Assignment, error) {
	value, err := r.resolveExpression(assignment.Value(), ctx)
	if err != nil {
		return nil, err
	}

	return NewAssignmentBuilder().
		Name(assignment.Name()).
		Value(value).
		AST(assignment).
		Build()
}

func (r *resolver) resolveCondition(condition workflows.Condition, ctx *versionContext) (Condition, error) {
	left, err := r.resolveExpression(condition.Left(), ctx)
	if err != nil {
		return nil, err
	}

	right, err := r.resolveExpression(condition.Right(), ctx)
	if err != nil {
		return nil, err
	}

	return NewConditionBuilder().
		Left(left).
		Operator(condition.Operator()).
		Right(right).
		AST(condition).
		Build()
}

func (r *resolver) resolveExpression(expression workflows.Expression, ctx *versionContext) (Expression, error) {
	switch value := expression.(type) {
	case workflows.IdentifierExpression:
		return NewIdentifierExpressionBuilder().
			Name(value.Name()).
			AST(value).
			Build()

	case workflows.SelectorExpression:
		builder := NewSelectorExpressionBuilder().
			AST(value)

		for _, part := range value.Parts() {
			builder.AddPart(part)
		}

		return builder.Build()

	case workflows.LiteralExpression:
		return NewLiteralExpressionBuilder().
			Value(value.Value()).
			AST(value).
			Build()

	case workflows.FunctionCallExpression:
		builder := NewFunctionCallExpressionBuilder().
			Name(value.Name()).
			AST(value)

		for _, argumentAST := range value.Arguments() {
			argument, err := r.resolveExpression(argumentAST, ctx)
			if err != nil {
				return nil, err
			}

			builder.AddArgument(argument)
		}

		return builder.Build()

	default:
		return nil, fmt.Errorf("unsupported expression kind %q", expression.Kind())
	}
}

func splitCallTarget(target string) (string, string) {
	parts := strings.SplitN(target, ".", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}

	return parts[0], parts[1]
}
