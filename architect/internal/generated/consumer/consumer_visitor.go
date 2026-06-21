// Code generated from internal/infrastructure/antlr/grammars/Consumer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package consumer // Consumer
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ConsumerParser.
type ConsumerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConsumerParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ConsumerParser#consumerDecl.
	VisitConsumerDecl(ctx *ConsumerDeclContext) interface{}

	// Visit a parse tree produced by ConsumerParser#action.
	VisitAction(ctx *ActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#fetchAction.
	VisitFetchAction(ctx *FetchActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#createAction.
	VisitCreateAction(ctx *CreateActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#storeAction.
	VisitStoreAction(ctx *StoreActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#updateAction.
	VisitUpdateAction(ctx *UpdateActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#emitAction.
	VisitEmitAction(ctx *EmitActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#callAction.
	VisitCallAction(ctx *CallActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#returnAction.
	VisitReturnAction(ctx *ReturnActionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#typedVariable.
	VisitTypedVariable(ctx *TypedVariableContext) interface{}

	// Visit a parse tree produced by ConsumerParser#assignmentBlock.
	VisitAssignmentBlock(ctx *AssignmentBlockContext) interface{}

	// Visit a parse tree produced by ConsumerParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by ConsumerParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by ConsumerParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ConsumerParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by ConsumerParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by ConsumerParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by ConsumerParser#typeRef.
	VisitTypeRef(ctx *TypeRefContext) interface{}

	// Visit a parse tree produced by ConsumerParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by ConsumerParser#numberConstraint.
	VisitNumberConstraint(ctx *NumberConstraintContext) interface{}

	// Visit a parse tree produced by ConsumerParser#optionalMarker.
	VisitOptionalMarker(ctx *OptionalMarkerContext) interface{}

	// Visit a parse tree produced by ConsumerParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by ConsumerParser#numberValue.
	VisitNumberValue(ctx *NumberValueContext) interface{}

	// Visit a parse tree produced by ConsumerParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
