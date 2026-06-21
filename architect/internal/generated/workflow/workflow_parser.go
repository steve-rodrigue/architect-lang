// Code generated from internal/infrastructure/antlr/grammars/Workflow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workflow // Workflow
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type WorkflowParser struct {
	*antlr.BaseParser
}

var WorkflowParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func workflowParserInit() {
	staticData := &WorkflowParserStaticData
	staticData.LiteralNames = []string{
		"", "'fetch'", "'from'", "'where'", "'create'", "'store'", "'to'", "'update'",
		"'emit'", "'call'", "'return'", "'('", "')'", "'<='", "'>='", "'=='",
		"'!='", "'='", "':'", "'.'", "'List'", "'input'", "'event'", "'result'",
		"'true'", "'false'", "'{'", "'}'", "'['", "']'", "'<'", "'>'", "','",
		"'|'", "'+'", "'*'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "FETCH", "FROM", "WHERE", "CREATE", "STORE", "TO", "UPDATE", "EMIT",
		"CALL", "RETURN", "LPAREN", "RPAREN", "LTE", "GTE", "EQ", "NEQ", "ASSIGN",
		"COLON", "DOT", "LIST", "INPUT", "EVENT", "RESULT", "TRUE", "FALSE",
		"LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "LT", "GT", "COMMA", "PIPE",
		"PLUS", "STAR", "QUESTION", "IDENT", "FLOAT", "INT", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"action", "fetchAction", "createAction", "storeAction", "updateAction",
		"emitAction", "callAction", "returnAction", "typedVariable", "assignmentBlock",
		"assignment", "condition", "comparator", "expression", "functionCall",
		"argumentList", "selector", "typeRef", "typeName", "numberConstraint",
		"optionalMarker", "value", "numberValue", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 179, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 3, 0, 56, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 95, 8, 9, 10, 9, 12, 9, 98, 9, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 115, 8, 13, 1, 14, 1, 14, 1, 14, 3,
		14, 120, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 127, 8, 15, 10,
		15, 12, 15, 130, 9, 15, 1, 16, 1, 16, 1, 16, 5, 16, 135, 8, 16, 10, 16,
		12, 16, 138, 9, 16, 1, 17, 1, 17, 3, 17, 142, 8, 17, 1, 17, 3, 17, 145,
		8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 152, 8, 17, 3, 17, 154,
		8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 162, 8, 19, 1,
		19, 1, 19, 3, 19, 166, 8, 19, 1, 19, 3, 19, 169, 8, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 0, 0, 24, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 0, 4, 2, 0, 13, 16, 30, 31, 2, 0, 24, 25, 38, 40, 1, 0, 38, 39, 2,
		0, 21, 23, 37, 37, 174, 0, 55, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4, 64, 1,
		0, 0, 0, 6, 68, 1, 0, 0, 0, 8, 73, 1, 0, 0, 0, 10, 77, 1, 0, 0, 0, 12,
		81, 1, 0, 0, 0, 14, 85, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18, 92, 1, 0, 0,
		0, 20, 101, 1, 0, 0, 0, 22, 105, 1, 0, 0, 0, 24, 109, 1, 0, 0, 0, 26, 114,
		1, 0, 0, 0, 28, 116, 1, 0, 0, 0, 30, 123, 1, 0, 0, 0, 32, 131, 1, 0, 0,
		0, 34, 153, 1, 0, 0, 0, 36, 155, 1, 0, 0, 0, 38, 168, 1, 0, 0, 0, 40, 170,
		1, 0, 0, 0, 42, 172, 1, 0, 0, 0, 44, 174, 1, 0, 0, 0, 46, 176, 1, 0, 0,
		0, 48, 56, 3, 2, 1, 0, 49, 56, 3, 4, 2, 0, 50, 56, 3, 6, 3, 0, 51, 56,
		3, 8, 4, 0, 52, 56, 3, 10, 5, 0, 53, 56, 3, 12, 6, 0, 54, 56, 3, 14, 7,
		0, 55, 48, 1, 0, 0, 0, 55, 49, 1, 0, 0, 0, 55, 50, 1, 0, 0, 0, 55, 51,
		1, 0, 0, 0, 55, 52, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0,
		56, 1, 1, 0, 0, 0, 57, 58, 5, 1, 0, 0, 58, 59, 3, 16, 8, 0, 59, 60, 5,
		2, 0, 0, 60, 61, 5, 37, 0, 0, 61, 62, 5, 3, 0, 0, 62, 63, 3, 22, 11, 0,
		63, 3, 1, 0, 0, 0, 64, 65, 5, 4, 0, 0, 65, 66, 3, 16, 8, 0, 66, 67, 3,
		18, 9, 0, 67, 5, 1, 0, 0, 0, 68, 69, 5, 5, 0, 0, 69, 70, 5, 37, 0, 0, 70,
		71, 5, 6, 0, 0, 71, 72, 5, 37, 0, 0, 72, 7, 1, 0, 0, 0, 73, 74, 5, 7, 0,
		0, 74, 75, 5, 37, 0, 0, 75, 76, 3, 18, 9, 0, 76, 9, 1, 0, 0, 0, 77, 78,
		5, 8, 0, 0, 78, 79, 3, 16, 8, 0, 79, 80, 3, 18, 9, 0, 80, 11, 1, 0, 0,
		0, 81, 82, 5, 9, 0, 0, 82, 83, 3, 32, 16, 0, 83, 84, 3, 18, 9, 0, 84, 13,
		1, 0, 0, 0, 85, 86, 5, 10, 0, 0, 86, 87, 3, 26, 13, 0, 87, 15, 1, 0, 0,
		0, 88, 89, 3, 46, 23, 0, 89, 90, 5, 18, 0, 0, 90, 91, 3, 34, 17, 0, 91,
		17, 1, 0, 0, 0, 92, 96, 5, 26, 0, 0, 93, 95, 3, 20, 10, 0, 94, 93, 1, 0,
		0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99,
		1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5, 27, 0, 0, 100, 19, 1, 0, 0,
		0, 101, 102, 3, 46, 23, 0, 102, 103, 5, 17, 0, 0, 103, 104, 3, 26, 13,
		0, 104, 21, 1, 0, 0, 0, 105, 106, 3, 26, 13, 0, 106, 107, 3, 24, 12, 0,
		107, 108, 3, 26, 13, 0, 108, 23, 1, 0, 0, 0, 109, 110, 7, 0, 0, 0, 110,
		25, 1, 0, 0, 0, 111, 115, 3, 28, 14, 0, 112, 115, 3, 32, 16, 0, 113, 115,
		3, 42, 21, 0, 114, 111, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1,
		0, 0, 0, 115, 27, 1, 0, 0, 0, 116, 117, 5, 37, 0, 0, 117, 119, 5, 11, 0,
		0, 118, 120, 3, 30, 15, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0,
		120, 121, 1, 0, 0, 0, 121, 122, 5, 12, 0, 0, 122, 29, 1, 0, 0, 0, 123,
		128, 3, 26, 13, 0, 124, 125, 5, 32, 0, 0, 125, 127, 3, 26, 13, 0, 126,
		124, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129,
		1, 0, 0, 0, 129, 31, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 136, 3, 46,
		23, 0, 132, 133, 5, 19, 0, 0, 133, 135, 3, 46, 23, 0, 134, 132, 1, 0, 0,
		0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		33, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 141, 3, 36, 18, 0, 140, 142,
		3, 38, 19, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 1,
		0, 0, 0, 143, 145, 3, 40, 20, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0,
		0, 0, 145, 154, 1, 0, 0, 0, 146, 147, 5, 20, 0, 0, 147, 148, 5, 30, 0,
		0, 148, 149, 3, 36, 18, 0, 149, 151, 5, 31, 0, 0, 150, 152, 3, 40, 20,
		0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153,
		139, 1, 0, 0, 0, 153, 146, 1, 0, 0, 0, 154, 35, 1, 0, 0, 0, 155, 156, 5,
		37, 0, 0, 156, 37, 1, 0, 0, 0, 157, 169, 5, 34, 0, 0, 158, 169, 5, 35,
		0, 0, 159, 161, 5, 28, 0, 0, 160, 162, 3, 44, 22, 0, 161, 160, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 5, 32, 0, 0, 164,
		166, 3, 44, 22, 0, 165, 164, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167,
		1, 0, 0, 0, 167, 169, 5, 29, 0, 0, 168, 157, 1, 0, 0, 0, 168, 158, 1, 0,
		0, 0, 168, 159, 1, 0, 0, 0, 169, 39, 1, 0, 0, 0, 170, 171, 5, 36, 0, 0,
		171, 41, 1, 0, 0, 0, 172, 173, 7, 1, 0, 0, 173, 43, 1, 0, 0, 0, 174, 175,
		7, 2, 0, 0, 175, 45, 1, 0, 0, 0, 176, 177, 7, 3, 0, 0, 177, 47, 1, 0, 0,
		0, 13, 55, 96, 114, 119, 128, 136, 141, 144, 151, 153, 161, 165, 168,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// WorkflowParserInit initializes any static state used to implement WorkflowParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWorkflowParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WorkflowParserInit() {
	staticData := &WorkflowParserStaticData
	staticData.once.Do(workflowParserInit)
}

// NewWorkflowParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWorkflowParser(input antlr.TokenStream) *WorkflowParser {
	WorkflowParserInit()
	this := new(WorkflowParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &WorkflowParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Workflow.g4"

	return this
}

// WorkflowParser tokens.
const (
	WorkflowParserEOF           = antlr.TokenEOF
	WorkflowParserFETCH         = 1
	WorkflowParserFROM          = 2
	WorkflowParserWHERE         = 3
	WorkflowParserCREATE        = 4
	WorkflowParserSTORE         = 5
	WorkflowParserTO            = 6
	WorkflowParserUPDATE        = 7
	WorkflowParserEMIT          = 8
	WorkflowParserCALL          = 9
	WorkflowParserRETURN        = 10
	WorkflowParserLPAREN        = 11
	WorkflowParserRPAREN        = 12
	WorkflowParserLTE           = 13
	WorkflowParserGTE           = 14
	WorkflowParserEQ            = 15
	WorkflowParserNEQ           = 16
	WorkflowParserASSIGN        = 17
	WorkflowParserCOLON         = 18
	WorkflowParserDOT           = 19
	WorkflowParserLIST          = 20
	WorkflowParserINPUT         = 21
	WorkflowParserEVENT         = 22
	WorkflowParserRESULT        = 23
	WorkflowParserTRUE          = 24
	WorkflowParserFALSE         = 25
	WorkflowParserLBRACE        = 26
	WorkflowParserRBRACE        = 27
	WorkflowParserLBRACKET      = 28
	WorkflowParserRBRACKET      = 29
	WorkflowParserLT            = 30
	WorkflowParserGT            = 31
	WorkflowParserCOMMA         = 32
	WorkflowParserPIPE          = 33
	WorkflowParserPLUS          = 34
	WorkflowParserSTAR          = 35
	WorkflowParserQUESTION      = 36
	WorkflowParserIDENT         = 37
	WorkflowParserFLOAT         = 38
	WorkflowParserINT           = 39
	WorkflowParserSTRING        = 40
	WorkflowParserLINE_COMMENT  = 41
	WorkflowParserBLOCK_COMMENT = 42
	WorkflowParserWS            = 43
)

// WorkflowParser rules.
const (
	WorkflowParserRULE_action           = 0
	WorkflowParserRULE_fetchAction      = 1
	WorkflowParserRULE_createAction     = 2
	WorkflowParserRULE_storeAction      = 3
	WorkflowParserRULE_updateAction     = 4
	WorkflowParserRULE_emitAction       = 5
	WorkflowParserRULE_callAction       = 6
	WorkflowParserRULE_returnAction     = 7
	WorkflowParserRULE_typedVariable    = 8
	WorkflowParserRULE_assignmentBlock  = 9
	WorkflowParserRULE_assignment       = 10
	WorkflowParserRULE_condition        = 11
	WorkflowParserRULE_comparator       = 12
	WorkflowParserRULE_expression       = 13
	WorkflowParserRULE_functionCall     = 14
	WorkflowParserRULE_argumentList     = 15
	WorkflowParserRULE_selector         = 16
	WorkflowParserRULE_typeRef          = 17
	WorkflowParserRULE_typeName         = 18
	WorkflowParserRULE_numberConstraint = 19
	WorkflowParserRULE_optionalMarker   = 20
	WorkflowParserRULE_value            = 21
	WorkflowParserRULE_numberValue      = 22
	WorkflowParserRULE_identifier       = 23
)

// IActionContext is an interface to support dynamic dispatch.
type IActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FetchAction() IFetchActionContext
	CreateAction() ICreateActionContext
	StoreAction() IStoreActionContext
	UpdateAction() IUpdateActionContext
	EmitAction() IEmitActionContext
	CallAction() ICallActionContext
	ReturnAction() IReturnActionContext

	// IsActionContext differentiates from other interfaces.
	IsActionContext()
}

type ActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionContext() *ActionContext {
	var p = new(ActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_action

	return p
}

func (s *ActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionContext) FetchAction() IFetchActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFetchActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFetchActionContext)
}

func (s *ActionContext) CreateAction() ICreateActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateActionContext)
}

func (s *ActionContext) StoreAction() IStoreActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStoreActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStoreActionContext)
}

func (s *ActionContext) UpdateAction() IUpdateActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateActionContext)
}

func (s *ActionContext) EmitAction() IEmitActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmitActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmitActionContext)
}

func (s *ActionContext) CallAction() ICallActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallActionContext)
}

func (s *ActionContext) ReturnAction() IReturnActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnActionContext)
}

func (s *ActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WorkflowParserRULE_action)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case WorkflowParserFETCH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.FetchAction()
		}

	case WorkflowParserCREATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.CreateAction()
		}

	case WorkflowParserSTORE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.StoreAction()
		}

	case WorkflowParserUPDATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.UpdateAction()
		}

	case WorkflowParserEMIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(52)
			p.EmitAction()
		}

	case WorkflowParserCALL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.CallAction()
		}

	case WorkflowParserRETURN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(54)
			p.ReturnAction()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFetchActionContext is an interface to support dynamic dispatch.
type IFetchActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FETCH() antlr.TerminalNode
	TypedVariable() ITypedVariableContext
	FROM() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	Condition() IConditionContext

	// IsFetchActionContext differentiates from other interfaces.
	IsFetchActionContext()
}

type FetchActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFetchActionContext() *FetchActionContext {
	var p = new(FetchActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_fetchAction
	return p
}

func InitEmptyFetchActionContext(p *FetchActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_fetchAction
}

func (*FetchActionContext) IsFetchActionContext() {}

func NewFetchActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FetchActionContext {
	var p = new(FetchActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_fetchAction

	return p
}

func (s *FetchActionContext) GetParser() antlr.Parser { return s.parser }

func (s *FetchActionContext) FETCH() antlr.TerminalNode {
	return s.GetToken(WorkflowParserFETCH, 0)
}

func (s *FetchActionContext) TypedVariable() ITypedVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedVariableContext)
}

func (s *FetchActionContext) FROM() antlr.TerminalNode {
	return s.GetToken(WorkflowParserFROM, 0)
}

func (s *FetchActionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserIDENT, 0)
}

func (s *FetchActionContext) WHERE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserWHERE, 0)
}

func (s *FetchActionContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *FetchActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FetchActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FetchActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitFetchAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) FetchAction() (localctx IFetchActionContext) {
	localctx = NewFetchActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WorkflowParserRULE_fetchAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(WorkflowParserFETCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.TypedVariable()
	}
	{
		p.SetState(59)
		p.Match(WorkflowParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(WorkflowParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(WorkflowParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Condition()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateActionContext is an interface to support dynamic dispatch.
type ICreateActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	TypedVariable() ITypedVariableContext
	AssignmentBlock() IAssignmentBlockContext

	// IsCreateActionContext differentiates from other interfaces.
	IsCreateActionContext()
}

type CreateActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateActionContext() *CreateActionContext {
	var p = new(CreateActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_createAction
	return p
}

func InitEmptyCreateActionContext(p *CreateActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_createAction
}

func (*CreateActionContext) IsCreateActionContext() {}

func NewCreateActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateActionContext {
	var p = new(CreateActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_createAction

	return p
}

func (s *CreateActionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateActionContext) CREATE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserCREATE, 0)
}

func (s *CreateActionContext) TypedVariable() ITypedVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedVariableContext)
}

func (s *CreateActionContext) AssignmentBlock() IAssignmentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentBlockContext)
}

func (s *CreateActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitCreateAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) CreateAction() (localctx ICreateActionContext) {
	localctx = NewCreateActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WorkflowParserRULE_createAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(WorkflowParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.TypedVariable()
	}
	{
		p.SetState(66)
		p.AssignmentBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStoreActionContext is an interface to support dynamic dispatch.
type IStoreActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STORE() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	TO() antlr.TerminalNode

	// IsStoreActionContext differentiates from other interfaces.
	IsStoreActionContext()
}

type StoreActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStoreActionContext() *StoreActionContext {
	var p = new(StoreActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_storeAction
	return p
}

func InitEmptyStoreActionContext(p *StoreActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_storeAction
}

func (*StoreActionContext) IsStoreActionContext() {}

func NewStoreActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreActionContext {
	var p = new(StoreActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_storeAction

	return p
}

func (s *StoreActionContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreActionContext) STORE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserSTORE, 0)
}

func (s *StoreActionContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(WorkflowParserIDENT)
}

func (s *StoreActionContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowParserIDENT, i)
}

func (s *StoreActionContext) TO() antlr.TerminalNode {
	return s.GetToken(WorkflowParserTO, 0)
}

func (s *StoreActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitStoreAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) StoreAction() (localctx IStoreActionContext) {
	localctx = NewStoreActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WorkflowParserRULE_storeAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(WorkflowParserSTORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(WorkflowParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(WorkflowParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(WorkflowParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdateActionContext is an interface to support dynamic dispatch.
type IUpdateActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UPDATE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	AssignmentBlock() IAssignmentBlockContext

	// IsUpdateActionContext differentiates from other interfaces.
	IsUpdateActionContext()
}

type UpdateActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateActionContext() *UpdateActionContext {
	var p = new(UpdateActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_updateAction
	return p
}

func InitEmptyUpdateActionContext(p *UpdateActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_updateAction
}

func (*UpdateActionContext) IsUpdateActionContext() {}

func NewUpdateActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateActionContext {
	var p = new(UpdateActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_updateAction

	return p
}

func (s *UpdateActionContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateActionContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserUPDATE, 0)
}

func (s *UpdateActionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserIDENT, 0)
}

func (s *UpdateActionContext) AssignmentBlock() IAssignmentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentBlockContext)
}

func (s *UpdateActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitUpdateAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) UpdateAction() (localctx IUpdateActionContext) {
	localctx = NewUpdateActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WorkflowParserRULE_updateAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(WorkflowParserUPDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(WorkflowParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.AssignmentBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmitActionContext is an interface to support dynamic dispatch.
type IEmitActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EMIT() antlr.TerminalNode
	TypedVariable() ITypedVariableContext
	AssignmentBlock() IAssignmentBlockContext

	// IsEmitActionContext differentiates from other interfaces.
	IsEmitActionContext()
}

type EmitActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmitActionContext() *EmitActionContext {
	var p = new(EmitActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_emitAction
	return p
}

func InitEmptyEmitActionContext(p *EmitActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_emitAction
}

func (*EmitActionContext) IsEmitActionContext() {}

func NewEmitActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmitActionContext {
	var p = new(EmitActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_emitAction

	return p
}

func (s *EmitActionContext) GetParser() antlr.Parser { return s.parser }

func (s *EmitActionContext) EMIT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserEMIT, 0)
}

func (s *EmitActionContext) TypedVariable() ITypedVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedVariableContext)
}

func (s *EmitActionContext) AssignmentBlock() IAssignmentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentBlockContext)
}

func (s *EmitActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmitActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmitActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitEmitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) EmitAction() (localctx IEmitActionContext) {
	localctx = NewEmitActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WorkflowParserRULE_emitAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(WorkflowParserEMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.TypedVariable()
	}
	{
		p.SetState(79)
		p.AssignmentBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallActionContext is an interface to support dynamic dispatch.
type ICallActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CALL() antlr.TerminalNode
	Selector() ISelectorContext
	AssignmentBlock() IAssignmentBlockContext

	// IsCallActionContext differentiates from other interfaces.
	IsCallActionContext()
}

type CallActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallActionContext() *CallActionContext {
	var p = new(CallActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_callAction
	return p
}

func InitEmptyCallActionContext(p *CallActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_callAction
}

func (*CallActionContext) IsCallActionContext() {}

func NewCallActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallActionContext {
	var p = new(CallActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_callAction

	return p
}

func (s *CallActionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallActionContext) CALL() antlr.TerminalNode {
	return s.GetToken(WorkflowParserCALL, 0)
}

func (s *CallActionContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *CallActionContext) AssignmentBlock() IAssignmentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentBlockContext)
}

func (s *CallActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitCallAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) CallAction() (localctx ICallActionContext) {
	localctx = NewCallActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WorkflowParserRULE_callAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(WorkflowParserCALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Selector()
	}
	{
		p.SetState(83)
		p.AssignmentBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnActionContext is an interface to support dynamic dispatch.
type IReturnActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnActionContext differentiates from other interfaces.
	IsReturnActionContext()
}

type ReturnActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnActionContext() *ReturnActionContext {
	var p = new(ReturnActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_returnAction
	return p
}

func InitEmptyReturnActionContext(p *ReturnActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_returnAction
}

func (*ReturnActionContext) IsReturnActionContext() {}

func NewReturnActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnActionContext {
	var p = new(ReturnActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_returnAction

	return p
}

func (s *ReturnActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnActionContext) RETURN() antlr.TerminalNode {
	return s.GetToken(WorkflowParserRETURN, 0)
}

func (s *ReturnActionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitReturnAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) ReturnAction() (localctx IReturnActionContext) {
	localctx = NewReturnActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WorkflowParserRULE_returnAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(WorkflowParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypedVariableContext is an interface to support dynamic dispatch.
type ITypedVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	COLON() antlr.TerminalNode
	TypeRef() ITypeRefContext

	// IsTypedVariableContext differentiates from other interfaces.
	IsTypedVariableContext()
}

type TypedVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedVariableContext() *TypedVariableContext {
	var p = new(TypedVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_typedVariable
	return p
}

func InitEmptyTypedVariableContext(p *TypedVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_typedVariable
}

func (*TypedVariableContext) IsTypedVariableContext() {}

func NewTypedVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedVariableContext {
	var p = new(TypedVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_typedVariable

	return p
}

func (s *TypedVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedVariableContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TypedVariableContext) COLON() antlr.TerminalNode {
	return s.GetToken(WorkflowParserCOLON, 0)
}

func (s *TypedVariableContext) TypeRef() ITypeRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *TypedVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitTypedVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) TypedVariable() (localctx ITypedVariableContext) {
	localctx = NewTypedVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, WorkflowParserRULE_typedVariable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Identifier()
	}
	{
		p.SetState(89)
		p.Match(WorkflowParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.TypeRef()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentBlockContext is an interface to support dynamic dispatch.
type IAssignmentBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext

	// IsAssignmentBlockContext differentiates from other interfaces.
	IsAssignmentBlockContext()
}

type AssignmentBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentBlockContext() *AssignmentBlockContext {
	var p = new(AssignmentBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_assignmentBlock
	return p
}

func InitEmptyAssignmentBlockContext(p *AssignmentBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_assignmentBlock
}

func (*AssignmentBlockContext) IsAssignmentBlockContext() {}

func NewAssignmentBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentBlockContext {
	var p = new(AssignmentBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_assignmentBlock

	return p
}

func (s *AssignmentBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserLBRACE, 0)
}

func (s *AssignmentBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserRBRACE, 0)
}

func (s *AssignmentBlockContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentBlockContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignmentBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitAssignmentBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) AssignmentBlock() (localctx IAssignmentBlockContext) {
	localctx = NewAssignmentBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, WorkflowParserRULE_assignmentBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(WorkflowParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137453633536) != 0 {
		{
			p.SetState(93)
			p.Assignment()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(WorkflowParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(WorkflowParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, WorkflowParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Identifier()
	}
	{
		p.SetState(102)
		p.Match(WorkflowParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Comparator() IComparatorContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, WorkflowParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Expression()
	}
	{
		p.SetState(106)
		p.Comparator()
	}
	{
		p.SetState(107)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(WorkflowParserEQ, 0)
}

func (s *ComparatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(WorkflowParserNEQ, 0)
}

func (s *ComparatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserLTE, 0)
}

func (s *ComparatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserGTE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserLT, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserGT, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, WorkflowParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3221348352) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	Selector() ISelectorContext
	Value() IValueContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *ExpressionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, WorkflowParserRULE_expression)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Selector()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserIDENT, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowParserRPAREN, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, WorkflowParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(WorkflowParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(WorkflowParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2061649313792) != 0 {
		{
			p.SetState(118)
			p.ArgumentList()
		}

	}
	{
		p.SetState(121)
		p.Match(WorkflowParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(WorkflowParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, WorkflowParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Expression()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WorkflowParserCOMMA {
		{
			p.SetState(124)
			p.Match(WorkflowParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Expression()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *SelectorContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SelectorContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(WorkflowParserDOT)
}

func (s *SelectorContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowParserDOT, i)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, WorkflowParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Identifier()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WorkflowParserDOT {
		{
			p.SetState(132)
			p.Match(WorkflowParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Identifier()
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeRefContext is an interface to support dynamic dispatch.
type ITypeRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeName() ITypeNameContext
	NumberConstraint() INumberConstraintContext
	OptionalMarker() IOptionalMarkerContext
	LIST() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode

	// IsTypeRefContext differentiates from other interfaces.
	IsTypeRefContext()
}

type TypeRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRefContext() *TypeRefContext {
	var p = new(TypeRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_typeRef
	return p
}

func InitEmptyTypeRefContext(p *TypeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_typeRef
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_typeRef

	return p
}

func (s *TypeRefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRefContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeRefContext) NumberConstraint() INumberConstraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberConstraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberConstraintContext)
}

func (s *TypeRefContext) OptionalMarker() IOptionalMarkerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionalMarkerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionalMarkerContext)
}

func (s *TypeRefContext) LIST() antlr.TerminalNode {
	return s.GetToken(WorkflowParserLIST, 0)
}

func (s *TypeRefContext) LT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserLT, 0)
}

func (s *TypeRefContext) GT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserGT, 0)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitTypeRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, WorkflowParserRULE_typeRef)
	var _la int

	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case WorkflowParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.TypeName()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51808043008) != 0 {
			{
				p.SetState(140)
				p.NumberConstraint()
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkflowParserQUESTION {
			{
				p.SetState(143)
				p.OptionalMarker()
			}

		}

	case WorkflowParserLIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(WorkflowParserLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(WorkflowParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.TypeName()
		}
		{
			p.SetState(149)
			p.Match(WorkflowParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkflowParserQUESTION {
			{
				p.SetState(150)
				p.OptionalMarker()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserIDENT, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, WorkflowParserRULE_typeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(WorkflowParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberConstraintContext is an interface to support dynamic dispatch.
type INumberConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllNumberValue() []INumberValueContext
	NumberValue(i int) INumberValueContext

	// IsNumberConstraintContext differentiates from other interfaces.
	IsNumberConstraintContext()
}

type NumberConstraintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberConstraintContext() *NumberConstraintContext {
	var p = new(NumberConstraintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_numberConstraint
	return p
}

func InitEmptyNumberConstraintContext(p *NumberConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_numberConstraint
}

func (*NumberConstraintContext) IsNumberConstraintContext() {}

func NewNumberConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberConstraintContext {
	var p = new(NumberConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_numberConstraint

	return p
}

func (s *NumberConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberConstraintContext) PLUS() antlr.TerminalNode {
	return s.GetToken(WorkflowParserPLUS, 0)
}

func (s *NumberConstraintContext) STAR() antlr.TerminalNode {
	return s.GetToken(WorkflowParserSTAR, 0)
}

func (s *NumberConstraintContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(WorkflowParserLBRACKET, 0)
}

func (s *NumberConstraintContext) COMMA() antlr.TerminalNode {
	return s.GetToken(WorkflowParserCOMMA, 0)
}

func (s *NumberConstraintContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(WorkflowParserRBRACKET, 0)
}

func (s *NumberConstraintContext) AllNumberValue() []INumberValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberValueContext); ok {
			len++
		}
	}

	tst := make([]INumberValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberValueContext); ok {
			tst[i] = t.(INumberValueContext)
			i++
		}
	}

	return tst
}

func (s *NumberConstraintContext) NumberValue(i int) INumberValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberValueContext)
}

func (s *NumberConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitNumberConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) NumberConstraint() (localctx INumberConstraintContext) {
	localctx = NewNumberConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, WorkflowParserRULE_numberConstraint)
	var _la int

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case WorkflowParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(WorkflowParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(WorkflowParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.Match(WorkflowParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkflowParserFLOAT || _la == WorkflowParserINT {
			{
				p.SetState(160)
				p.NumberValue()
			}

		}
		{
			p.SetState(163)
			p.Match(WorkflowParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkflowParserFLOAT || _la == WorkflowParserINT {
			{
				p.SetState(164)
				p.NumberValue()
			}

		}
		{
			p.SetState(167)
			p.Match(WorkflowParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionalMarkerContext is an interface to support dynamic dispatch.
type IOptionalMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUESTION() antlr.TerminalNode

	// IsOptionalMarkerContext differentiates from other interfaces.
	IsOptionalMarkerContext()
}

type OptionalMarkerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalMarkerContext() *OptionalMarkerContext {
	var p = new(OptionalMarkerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_optionalMarker
	return p
}

func InitEmptyOptionalMarkerContext(p *OptionalMarkerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_optionalMarker
}

func (*OptionalMarkerContext) IsOptionalMarkerContext() {}

func NewOptionalMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalMarkerContext {
	var p = new(OptionalMarkerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_optionalMarker

	return p
}

func (s *OptionalMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalMarkerContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(WorkflowParserQUESTION, 0)
}

func (s *OptionalMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalMarkerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitOptionalMarker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) OptionalMarker() (localctx IOptionalMarkerContext) {
	localctx = NewOptionalMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, WorkflowParserRULE_optionalMarker)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(WorkflowParserQUESTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(WorkflowParserSTRING, 0)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserFLOAT, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(WorkflowParserFALSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, WorkflowParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924195680256) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberValueContext is an interface to support dynamic dispatch.
type INumberValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsNumberValueContext differentiates from other interfaces.
	IsNumberValueContext()
}

type NumberValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberValueContext() *NumberValueContext {
	var p = new(NumberValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_numberValue
	return p
}

func InitEmptyNumberValueContext(p *NumberValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_numberValue
}

func (*NumberValueContext) IsNumberValueContext() {}

func NewNumberValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberValueContext {
	var p = new(NumberValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_numberValue

	return p
}

func (s *NumberValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberValueContext) INT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserINT, 0)
}

func (s *NumberValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserFLOAT, 0)
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitNumberValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) NumberValue() (localctx INumberValueContext) {
	localctx = NewNumberValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, WorkflowParserRULE_numberValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WorkflowParserFLOAT || _la == WorkflowParserINT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	INPUT() antlr.TerminalNode
	EVENT() antlr.TerminalNode
	RESULT() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserIDENT, 0)
}

func (s *IdentifierContext) INPUT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserINPUT, 0)
}

func (s *IdentifierContext) EVENT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserEVENT, 0)
}

func (s *IdentifierContext) RESULT() antlr.TerminalNode {
	return s.GetToken(WorkflowParserRESULT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkflowVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkflowParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, WorkflowParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137453633536) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
