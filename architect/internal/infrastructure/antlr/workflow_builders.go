package antlr

import (
	"strings"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

func buildWorkflowTypedVariable(
	name string,
	typeRef common.TypeReference,
	setErr func(error),
) workflows.TypedVariable {
	variable, err := workflows.NewTypedVariableBuilder().
		Name(name).
		Type(typeRef).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return variable
}

func buildWorkflowFetchAction(
	variable workflows.TypedVariable,
	source string,
	condition workflows.Condition,
	setErr func(error),
) workflows.FetchAction {
	action, err := workflows.NewFetchActionBuilder().
		Variable(variable).
		Source(source).
		Condition(condition).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return action
}

func buildWorkflowCreateAction(
	variable workflows.TypedVariable,
	assignments []workflows.Assignment,
	setErr func(error),
) workflows.CreateAction {
	builder := workflows.NewCreateActionBuilder().Variable(variable)

	for _, assignment := range assignments {
		builder.AddAssignment(assignment)
	}

	action, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return action
}

func buildWorkflowStoreAction(
	variableName string,
	destination string,
	setErr func(error),
) workflows.StoreAction {
	action, err := workflows.NewStoreActionBuilder().
		VariableName(variableName).
		Destination(destination).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return action
}

func buildWorkflowUpdateAction(
	variableName string,
	assignments []workflows.Assignment,
	setErr func(error),
) workflows.UpdateAction {
	builder := workflows.NewUpdateActionBuilder().
		VariableName(variableName)

	for _, assignment := range assignments {
		builder.AddAssignment(assignment)
	}

	action, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return action
}

func buildWorkflowEmitAction(
	variable workflows.TypedVariable,
	assignments []workflows.Assignment,
	setErr func(error),
) workflows.EmitAction {
	builder := workflows.NewEmitActionBuilder().Variable(variable)

	for _, assignment := range assignments {
		builder.AddAssignment(assignment)
	}

	action, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return action
}

func buildWorkflowCallAction(
	target string,
	assignments []workflows.Assignment,
	setErr func(error),
) workflows.CallAction {
	builder := workflows.NewCallActionBuilder().Target(target)

	for _, assignment := range assignments {
		builder.AddAssignment(assignment)
	}

	action, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return action
}

func buildWorkflowReturnAction(
	expression workflows.Expression,
	setErr func(error),
) workflows.ReturnAction {
	action, err := workflows.NewReturnActionBuilder().
		Expression(expression).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return action
}

func buildWorkflowAssignment(
	name string,
	value workflows.Expression,
	setErr func(error),
) workflows.Assignment {
	assignment, err := workflows.NewAssignmentBuilder().
		Name(name).
		Value(value).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return assignment
}

func buildWorkflowCondition(
	left workflows.Expression,
	operator workflows.Comparator,
	right workflows.Expression,
	setErr func(error),
) workflows.Condition {
	condition, err := workflows.NewConditionBuilder().
		Left(left).
		Operator(operator).
		Right(right).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return condition
}

func buildWorkflowIdentifierExpression(
	name string,
	setErr func(error),
) workflows.IdentifierExpression {
	expression, err := workflows.NewIdentifierExpressionBuilder().
		Name(name).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return expression
}

func buildWorkflowSelectorExpression(
	parts []string,
	setErr func(error),
) workflows.SelectorExpression {
	builder := workflows.NewSelectorExpressionBuilder()

	for _, part := range parts {
		builder.AddPart(part)
	}

	expression, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return expression
}

func buildWorkflowLiteralExpression(
	value common.Value,
	setErr func(error),
) workflows.LiteralExpression {
	expression, err := workflows.NewLiteralExpressionBuilder().
		Value(value).
		Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return expression
}

func buildWorkflowFunctionCallExpression(
	name string,
	args []workflows.Expression,
	setErr func(error),
) workflows.FunctionCallExpression {
	builder := workflows.NewFunctionCallExpressionBuilder().Name(name)

	for _, arg := range args {
		builder.AddArgument(arg)
	}

	expression, err := builder.Build()
	if err != nil {
		setErr(err)
		return nil
	}

	return expression
}

func splitSelector(raw string) []string {
	return strings.Split(raw, ".")
}
