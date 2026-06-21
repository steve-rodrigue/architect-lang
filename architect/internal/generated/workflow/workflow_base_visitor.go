// Code generated from internal/infrastructure/antlr/grammars/Workflow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workflow // Workflow
import "github.com/antlr4-go/antlr/v4"

type BaseWorkflowVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseWorkflowVisitor) VisitAction(ctx *ActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitFetchAction(ctx *FetchActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitCreateAction(ctx *CreateActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitStoreAction(ctx *StoreActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitUpdateAction(ctx *UpdateActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitEmitAction(ctx *EmitActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitCallAction(ctx *CallActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitReturnAction(ctx *ReturnActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitTypedVariable(ctx *TypedVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitAssignmentBlock(ctx *AssignmentBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitTypeRef(ctx *TypeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitNumberConstraint(ctx *NumberConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitOptionalMarker(ctx *OptionalMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitNumberValue(ctx *NumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkflowVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
