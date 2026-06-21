// Code generated from internal/infrastructure/antlr/grammars/Consumer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package consumer // Consumer
import "github.com/antlr4-go/antlr/v4"

type BaseConsumerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseConsumerVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitConsumerDecl(ctx *ConsumerDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitAction(ctx *ActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitFetchAction(ctx *FetchActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitCreateAction(ctx *CreateActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitStoreAction(ctx *StoreActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitUpdateAction(ctx *UpdateActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitEmitAction(ctx *EmitActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitCallAction(ctx *CallActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitReturnAction(ctx *ReturnActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitTypedVariable(ctx *TypedVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitAssignmentBlock(ctx *AssignmentBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitTypeRef(ctx *TypeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitNumberConstraint(ctx *NumberConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitOptionalMarker(ctx *OptionalMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitNumberValue(ctx *NumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConsumerVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
