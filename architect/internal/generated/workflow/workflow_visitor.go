// Code generated from internal/infrastructure/antlr/grammars/Workflow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workflow // Workflow
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by WorkflowParser.
type WorkflowVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by WorkflowParser#action.
	VisitAction(ctx *ActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#fetchAction.
	VisitFetchAction(ctx *FetchActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#createAction.
	VisitCreateAction(ctx *CreateActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#storeAction.
	VisitStoreAction(ctx *StoreActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#updateAction.
	VisitUpdateAction(ctx *UpdateActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#emitAction.
	VisitEmitAction(ctx *EmitActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#callAction.
	VisitCallAction(ctx *CallActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#returnAction.
	VisitReturnAction(ctx *ReturnActionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#typedVariable.
	VisitTypedVariable(ctx *TypedVariableContext) interface{}

	// Visit a parse tree produced by WorkflowParser#assignmentBlock.
	VisitAssignmentBlock(ctx *AssignmentBlockContext) interface{}

	// Visit a parse tree produced by WorkflowParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by WorkflowParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by WorkflowParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by WorkflowParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by WorkflowParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by WorkflowParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by WorkflowParser#typeRef.
	VisitTypeRef(ctx *TypeRefContext) interface{}

	// Visit a parse tree produced by WorkflowParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by WorkflowParser#numberConstraint.
	VisitNumberConstraint(ctx *NumberConstraintContext) interface{}

	// Visit a parse tree produced by WorkflowParser#optionalMarker.
	VisitOptionalMarker(ctx *OptionalMarkerContext) interface{}

	// Visit a parse tree produced by WorkflowParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by WorkflowParser#numberValue.
	VisitNumberValue(ctx *NumberValueContext) interface{}

	// Visit a parse tree produced by WorkflowParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
