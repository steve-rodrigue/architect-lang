// Code generated from internal/infrastructure/antlr/grammars/Endpoint.g4 by ANTLR 4.13.2. DO NOT EDIT.

package endpoint // Endpoint
import "github.com/antlr4-go/antlr/v4"

type BaseEndpointVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEndpointVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitEndpointDecl(ctx *EndpointDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitEndpointBody(ctx *EndpointBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitInputBlock(ctx *InputBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitInputField(ctx *InputFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitInputSourceRule(ctx *InputSourceRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitInputSource(ctx *InputSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitInputSourceKind(ctx *InputSourceKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitAction(ctx *ActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitFetchAction(ctx *FetchActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitCreateAction(ctx *CreateActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitStoreAction(ctx *StoreActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitUpdateAction(ctx *UpdateActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitEmitAction(ctx *EmitActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitReturnAction(ctx *ReturnActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitTypedVariable(ctx *TypedVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitAssignmentBlock(ctx *AssignmentBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitTypeRef(ctx *TypeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitNumberConstraint(ctx *NumberConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitOptionalMarker(ctx *OptionalMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitNumberValue(ctx *NumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEndpointVisitor) VisitHttpMethod(ctx *HttpMethodContext) interface{} {
	return v.VisitChildren(ctx)
}
