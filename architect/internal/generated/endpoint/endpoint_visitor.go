// Code generated from internal/infrastructure/antlr/grammars/Endpoint.g4 by ANTLR 4.13.2. DO NOT EDIT.

package endpoint // Endpoint
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EndpointParser.
type EndpointVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EndpointParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by EndpointParser#endpointDecl.
	VisitEndpointDecl(ctx *EndpointDeclContext) interface{}

	// Visit a parse tree produced by EndpointParser#endpointBody.
	VisitEndpointBody(ctx *EndpointBodyContext) interface{}

	// Visit a parse tree produced by EndpointParser#inputBlock.
	VisitInputBlock(ctx *InputBlockContext) interface{}

	// Visit a parse tree produced by EndpointParser#inputField.
	VisitInputField(ctx *InputFieldContext) interface{}

	// Visit a parse tree produced by EndpointParser#inputSourceRule.
	VisitInputSourceRule(ctx *InputSourceRuleContext) interface{}

	// Visit a parse tree produced by EndpointParser#inputSource.
	VisitInputSource(ctx *InputSourceContext) interface{}

	// Visit a parse tree produced by EndpointParser#inputSourceKind.
	VisitInputSourceKind(ctx *InputSourceKindContext) interface{}

	// Visit a parse tree produced by EndpointParser#httpMethod.
	VisitHttpMethod(ctx *HttpMethodContext) interface{}

	// Visit a parse tree produced by EndpointParser#action.
	VisitAction(ctx *ActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#fetchAction.
	VisitFetchAction(ctx *FetchActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#createAction.
	VisitCreateAction(ctx *CreateActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#storeAction.
	VisitStoreAction(ctx *StoreActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#updateAction.
	VisitUpdateAction(ctx *UpdateActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#emitAction.
	VisitEmitAction(ctx *EmitActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#callAction.
	VisitCallAction(ctx *CallActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#returnAction.
	VisitReturnAction(ctx *ReturnActionContext) interface{}

	// Visit a parse tree produced by EndpointParser#typedVariable.
	VisitTypedVariable(ctx *TypedVariableContext) interface{}

	// Visit a parse tree produced by EndpointParser#assignmentBlock.
	VisitAssignmentBlock(ctx *AssignmentBlockContext) interface{}

	// Visit a parse tree produced by EndpointParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by EndpointParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by EndpointParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by EndpointParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by EndpointParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by EndpointParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by EndpointParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by EndpointParser#typeRef.
	VisitTypeRef(ctx *TypeRefContext) interface{}

	// Visit a parse tree produced by EndpointParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by EndpointParser#numberConstraint.
	VisitNumberConstraint(ctx *NumberConstraintContext) interface{}

	// Visit a parse tree produced by EndpointParser#optionalMarker.
	VisitOptionalMarker(ctx *OptionalMarkerContext) interface{}

	// Visit a parse tree produced by EndpointParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by EndpointParser#numberValue.
	VisitNumberValue(ctx *NumberValueContext) interface{}

	// Visit a parse tree produced by EndpointParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
