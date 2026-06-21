// Code generated from internal/infrastructure/antlr/grammars/Consumer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package consumer // Consumer
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

type ConsumerParser struct {
	*antlr.BaseParser
}

var ConsumerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func consumerParserInit() {
	staticData := &ConsumerParserStaticData
	staticData.LiteralNames = []string{
		"", "'consumes'", "'fetch'", "'from'", "'where'", "'create'", "'store'",
		"'to'", "'update'", "'emit'", "'call'", "'return'", "'('", "')'", "'<='",
		"'>='", "'=='", "'!='", "'='", "':'", "'.'", "'List'", "'input'", "'event'",
		"'result'", "'true'", "'false'", "'{'", "'}'", "'['", "']'", "'<'",
		"'>'", "','", "'|'", "'+'", "'*'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "CONSUMES", "FETCH", "FROM", "WHERE", "CREATE", "STORE", "TO", "UPDATE",
		"EMIT", "CALL", "RETURN", "LPAREN", "RPAREN", "LTE", "GTE", "EQ", "NEQ",
		"ASSIGN", "COLON", "DOT", "LIST", "INPUT", "EVENT", "RESULT", "TRUE",
		"FALSE", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "LT", "GT", "COMMA",
		"PIPE", "PLUS", "STAR", "QUESTION", "IDENT", "FLOAT", "INT", "STRING",
		"LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "consumerDecl", "action", "fetchAction", "createAction",
		"storeAction", "updateAction", "emitAction", "callAction", "returnAction",
		"typedVariable", "assignmentBlock", "assignment", "condition", "comparator",
		"expression", "functionCall", "argumentList", "selector", "typeRef",
		"typeName", "numberConstraint", "optionalMarker", "value", "numberValue",
		"identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 197, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1, 63, 9,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 74, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 5, 11, 113, 8, 11, 10, 11, 12, 11, 116, 9, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 3, 15, 133, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 138, 8,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 145, 8, 17, 10, 17, 12, 17,
		148, 9, 17, 1, 18, 1, 18, 1, 18, 5, 18, 153, 8, 18, 10, 18, 12, 18, 156,
		9, 18, 1, 19, 1, 19, 3, 19, 160, 8, 19, 1, 19, 3, 19, 163, 8, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 170, 8, 19, 3, 19, 172, 8, 19, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 180, 8, 21, 1, 21, 1, 21, 3,
		21, 184, 8, 21, 1, 21, 3, 21, 187, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 0, 0, 26, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 0,
		4, 2, 0, 14, 17, 31, 32, 2, 0, 25, 26, 39, 41, 1, 0, 39, 40, 2, 0, 22,
		24, 38, 38, 191, 0, 52, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4, 73, 1, 0, 0,
		0, 6, 75, 1, 0, 0, 0, 8, 82, 1, 0, 0, 0, 10, 86, 1, 0, 0, 0, 12, 91, 1,
		0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 99, 1, 0, 0, 0, 18, 103, 1, 0, 0, 0, 20,
		106, 1, 0, 0, 0, 22, 110, 1, 0, 0, 0, 24, 119, 1, 0, 0, 0, 26, 123, 1,
		0, 0, 0, 28, 127, 1, 0, 0, 0, 30, 132, 1, 0, 0, 0, 32, 134, 1, 0, 0, 0,
		34, 141, 1, 0, 0, 0, 36, 149, 1, 0, 0, 0, 38, 171, 1, 0, 0, 0, 40, 173,
		1, 0, 0, 0, 42, 186, 1, 0, 0, 0, 44, 188, 1, 0, 0, 0, 46, 190, 1, 0, 0,
		0, 48, 192, 1, 0, 0, 0, 50, 194, 1, 0, 0, 0, 52, 53, 3, 2, 1, 0, 53, 54,
		5, 0, 0, 1, 54, 1, 1, 0, 0, 0, 55, 56, 5, 1, 0, 0, 56, 57, 3, 50, 25, 0,
		57, 61, 5, 27, 0, 0, 58, 60, 3, 4, 2, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1,
		0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0, 63,
		61, 1, 0, 0, 0, 64, 65, 5, 28, 0, 0, 65, 3, 1, 0, 0, 0, 66, 74, 3, 6, 3,
		0, 67, 74, 3, 8, 4, 0, 68, 74, 3, 10, 5, 0, 69, 74, 3, 12, 6, 0, 70, 74,
		3, 14, 7, 0, 71, 74, 3, 16, 8, 0, 72, 74, 3, 18, 9, 0, 73, 66, 1, 0, 0,
		0, 73, 67, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 73, 70,
		1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74, 5, 1, 0, 0, 0,
		75, 76, 5, 2, 0, 0, 76, 77, 3, 20, 10, 0, 77, 78, 5, 3, 0, 0, 78, 79, 5,
		38, 0, 0, 79, 80, 5, 4, 0, 0, 80, 81, 3, 26, 13, 0, 81, 7, 1, 0, 0, 0,
		82, 83, 5, 5, 0, 0, 83, 84, 3, 20, 10, 0, 84, 85, 3, 22, 11, 0, 85, 9,
		1, 0, 0, 0, 86, 87, 5, 6, 0, 0, 87, 88, 5, 38, 0, 0, 88, 89, 5, 7, 0, 0,
		89, 90, 5, 38, 0, 0, 90, 11, 1, 0, 0, 0, 91, 92, 5, 8, 0, 0, 92, 93, 5,
		38, 0, 0, 93, 94, 3, 22, 11, 0, 94, 13, 1, 0, 0, 0, 95, 96, 5, 9, 0, 0,
		96, 97, 3, 20, 10, 0, 97, 98, 3, 22, 11, 0, 98, 15, 1, 0, 0, 0, 99, 100,
		5, 10, 0, 0, 100, 101, 3, 36, 18, 0, 101, 102, 3, 22, 11, 0, 102, 17, 1,
		0, 0, 0, 103, 104, 5, 11, 0, 0, 104, 105, 3, 30, 15, 0, 105, 19, 1, 0,
		0, 0, 106, 107, 3, 50, 25, 0, 107, 108, 5, 19, 0, 0, 108, 109, 3, 38, 19,
		0, 109, 21, 1, 0, 0, 0, 110, 114, 5, 27, 0, 0, 111, 113, 3, 24, 12, 0,
		112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118,
		5, 28, 0, 0, 118, 23, 1, 0, 0, 0, 119, 120, 3, 50, 25, 0, 120, 121, 5,
		18, 0, 0, 121, 122, 3, 30, 15, 0, 122, 25, 1, 0, 0, 0, 123, 124, 3, 30,
		15, 0, 124, 125, 3, 28, 14, 0, 125, 126, 3, 30, 15, 0, 126, 27, 1, 0, 0,
		0, 127, 128, 7, 0, 0, 0, 128, 29, 1, 0, 0, 0, 129, 133, 3, 32, 16, 0, 130,
		133, 3, 36, 18, 0, 131, 133, 3, 46, 23, 0, 132, 129, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 31, 1, 0, 0, 0, 134, 135, 5, 38,
		0, 0, 135, 137, 5, 12, 0, 0, 136, 138, 3, 34, 17, 0, 137, 136, 1, 0, 0,
		0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 5, 13, 0, 0, 140,
		33, 1, 0, 0, 0, 141, 146, 3, 30, 15, 0, 142, 143, 5, 33, 0, 0, 143, 145,
		3, 30, 15, 0, 144, 142, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1,
		0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 35, 1, 0, 0, 0, 148, 146, 1, 0, 0,
		0, 149, 154, 3, 50, 25, 0, 150, 151, 5, 20, 0, 0, 151, 153, 3, 50, 25,
		0, 152, 150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 37, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 159, 3,
		40, 20, 0, 158, 160, 3, 42, 21, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0,
		0, 0, 160, 162, 1, 0, 0, 0, 161, 163, 3, 44, 22, 0, 162, 161, 1, 0, 0,
		0, 162, 163, 1, 0, 0, 0, 163, 172, 1, 0, 0, 0, 164, 165, 5, 21, 0, 0, 165,
		166, 5, 31, 0, 0, 166, 167, 3, 40, 20, 0, 167, 169, 5, 32, 0, 0, 168, 170,
		3, 44, 22, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 172, 1,
		0, 0, 0, 171, 157, 1, 0, 0, 0, 171, 164, 1, 0, 0, 0, 172, 39, 1, 0, 0,
		0, 173, 174, 5, 38, 0, 0, 174, 41, 1, 0, 0, 0, 175, 187, 5, 35, 0, 0, 176,
		187, 5, 36, 0, 0, 177, 179, 5, 29, 0, 0, 178, 180, 3, 48, 24, 0, 179, 178,
		1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 183, 5, 33,
		0, 0, 182, 184, 3, 48, 24, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0,
		0, 184, 185, 1, 0, 0, 0, 185, 187, 5, 30, 0, 0, 186, 175, 1, 0, 0, 0, 186,
		176, 1, 0, 0, 0, 186, 177, 1, 0, 0, 0, 187, 43, 1, 0, 0, 0, 188, 189, 5,
		37, 0, 0, 189, 45, 1, 0, 0, 0, 190, 191, 7, 1, 0, 0, 191, 47, 1, 0, 0,
		0, 192, 193, 7, 2, 0, 0, 193, 49, 1, 0, 0, 0, 194, 195, 7, 3, 0, 0, 195,
		51, 1, 0, 0, 0, 14, 61, 73, 114, 132, 137, 146, 154, 159, 162, 169, 171,
		179, 183, 186,
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

// ConsumerParserInit initializes any static state used to implement ConsumerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConsumerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConsumerParserInit() {
	staticData := &ConsumerParserStaticData
	staticData.once.Do(consumerParserInit)
}

// NewConsumerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConsumerParser(input antlr.TokenStream) *ConsumerParser {
	ConsumerParserInit()
	this := new(ConsumerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ConsumerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Consumer.g4"

	return this
}

// ConsumerParser tokens.
const (
	ConsumerParserEOF           = antlr.TokenEOF
	ConsumerParserCONSUMES      = 1
	ConsumerParserFETCH         = 2
	ConsumerParserFROM          = 3
	ConsumerParserWHERE         = 4
	ConsumerParserCREATE        = 5
	ConsumerParserSTORE         = 6
	ConsumerParserTO            = 7
	ConsumerParserUPDATE        = 8
	ConsumerParserEMIT          = 9
	ConsumerParserCALL          = 10
	ConsumerParserRETURN        = 11
	ConsumerParserLPAREN        = 12
	ConsumerParserRPAREN        = 13
	ConsumerParserLTE           = 14
	ConsumerParserGTE           = 15
	ConsumerParserEQ            = 16
	ConsumerParserNEQ           = 17
	ConsumerParserASSIGN        = 18
	ConsumerParserCOLON         = 19
	ConsumerParserDOT           = 20
	ConsumerParserLIST          = 21
	ConsumerParserINPUT         = 22
	ConsumerParserEVENT         = 23
	ConsumerParserRESULT        = 24
	ConsumerParserTRUE          = 25
	ConsumerParserFALSE         = 26
	ConsumerParserLBRACE        = 27
	ConsumerParserRBRACE        = 28
	ConsumerParserLBRACKET      = 29
	ConsumerParserRBRACKET      = 30
	ConsumerParserLT            = 31
	ConsumerParserGT            = 32
	ConsumerParserCOMMA         = 33
	ConsumerParserPIPE          = 34
	ConsumerParserPLUS          = 35
	ConsumerParserSTAR          = 36
	ConsumerParserQUESTION      = 37
	ConsumerParserIDENT         = 38
	ConsumerParserFLOAT         = 39
	ConsumerParserINT           = 40
	ConsumerParserSTRING        = 41
	ConsumerParserLINE_COMMENT  = 42
	ConsumerParserBLOCK_COMMENT = 43
	ConsumerParserWS            = 44
)

// ConsumerParser rules.
const (
	ConsumerParserRULE_program          = 0
	ConsumerParserRULE_consumerDecl     = 1
	ConsumerParserRULE_action           = 2
	ConsumerParserRULE_fetchAction      = 3
	ConsumerParserRULE_createAction     = 4
	ConsumerParserRULE_storeAction      = 5
	ConsumerParserRULE_updateAction     = 6
	ConsumerParserRULE_emitAction       = 7
	ConsumerParserRULE_callAction       = 8
	ConsumerParserRULE_returnAction     = 9
	ConsumerParserRULE_typedVariable    = 10
	ConsumerParserRULE_assignmentBlock  = 11
	ConsumerParserRULE_assignment       = 12
	ConsumerParserRULE_condition        = 13
	ConsumerParserRULE_comparator       = 14
	ConsumerParserRULE_expression       = 15
	ConsumerParserRULE_functionCall     = 16
	ConsumerParserRULE_argumentList     = 17
	ConsumerParserRULE_selector         = 18
	ConsumerParserRULE_typeRef          = 19
	ConsumerParserRULE_typeName         = 20
	ConsumerParserRULE_numberConstraint = 21
	ConsumerParserRULE_optionalMarker   = 22
	ConsumerParserRULE_value            = 23
	ConsumerParserRULE_numberValue      = 24
	ConsumerParserRULE_identifier       = 25
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConsumerDecl() IConsumerDeclContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ConsumerDecl() IConsumerDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsumerDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsumerDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConsumerParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConsumerParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.ConsumerDecl()
	}
	{
		p.SetState(53)
		p.Match(ConsumerParserEOF)
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

// IConsumerDeclContext is an interface to support dynamic dispatch.
type IConsumerDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONSUMES() antlr.TerminalNode
	Identifier() IIdentifierContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllAction_() []IActionContext
	Action_(i int) IActionContext

	// IsConsumerDeclContext differentiates from other interfaces.
	IsConsumerDeclContext()
}

type ConsumerDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConsumerDeclContext() *ConsumerDeclContext {
	var p = new(ConsumerDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_consumerDecl
	return p
}

func InitEmptyConsumerDeclContext(p *ConsumerDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_consumerDecl
}

func (*ConsumerDeclContext) IsConsumerDeclContext() {}

func NewConsumerDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsumerDeclContext {
	var p = new(ConsumerDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_consumerDecl

	return p
}

func (s *ConsumerDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsumerDeclContext) CONSUMES() antlr.TerminalNode {
	return s.GetToken(ConsumerParserCONSUMES, 0)
}

func (s *ConsumerDeclContext) Identifier() IIdentifierContext {
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

func (s *ConsumerDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserLBRACE, 0)
}

func (s *ConsumerDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserRBRACE, 0)
}

func (s *ConsumerDeclContext) AllAction_() []IActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionContext); ok {
			len++
		}
	}

	tst := make([]IActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionContext); ok {
			tst[i] = t.(IActionContext)
			i++
		}
	}

	return tst
}

func (s *ConsumerDeclContext) Action_(i int) IActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
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

	return t.(IActionContext)
}

func (s *ConsumerDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsumerDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsumerDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitConsumerDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) ConsumerDecl() (localctx IConsumerDeclContext) {
	localctx = NewConsumerDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConsumerParserRULE_consumerDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(ConsumerParserCONSUMES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Identifier()
	}
	{
		p.SetState(57)
		p.Match(ConsumerParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3940) != 0 {
		{
			p.SetState(58)
			p.Action_()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(ConsumerParserRBRACE)
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
	p.RuleIndex = ConsumerParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_action

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
	case ConsumerVisitor:
		return t.VisitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConsumerParserRULE_action)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ConsumerParserFETCH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.FetchAction()
		}

	case ConsumerParserCREATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.CreateAction()
		}

	case ConsumerParserSTORE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.StoreAction()
		}

	case ConsumerParserUPDATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.UpdateAction()
		}

	case ConsumerParserEMIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(70)
			p.EmitAction()
		}

	case ConsumerParserCALL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(71)
			p.CallAction()
		}

	case ConsumerParserRETURN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(72)
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
	p.RuleIndex = ConsumerParserRULE_fetchAction
	return p
}

func InitEmptyFetchActionContext(p *FetchActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_fetchAction
}

func (*FetchActionContext) IsFetchActionContext() {}

func NewFetchActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FetchActionContext {
	var p = new(FetchActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_fetchAction

	return p
}

func (s *FetchActionContext) GetParser() antlr.Parser { return s.parser }

func (s *FetchActionContext) FETCH() antlr.TerminalNode {
	return s.GetToken(ConsumerParserFETCH, 0)
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
	return s.GetToken(ConsumerParserFROM, 0)
}

func (s *FetchActionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserIDENT, 0)
}

func (s *FetchActionContext) WHERE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserWHERE, 0)
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
	case ConsumerVisitor:
		return t.VisitFetchAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) FetchAction() (localctx IFetchActionContext) {
	localctx = NewFetchActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConsumerParserRULE_fetchAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(ConsumerParserFETCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.TypedVariable()
	}
	{
		p.SetState(77)
		p.Match(ConsumerParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(ConsumerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(ConsumerParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
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
	p.RuleIndex = ConsumerParserRULE_createAction
	return p
}

func InitEmptyCreateActionContext(p *CreateActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_createAction
}

func (*CreateActionContext) IsCreateActionContext() {}

func NewCreateActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateActionContext {
	var p = new(CreateActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_createAction

	return p
}

func (s *CreateActionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateActionContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserCREATE, 0)
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
	case ConsumerVisitor:
		return t.VisitCreateAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) CreateAction() (localctx ICreateActionContext) {
	localctx = NewCreateActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConsumerParserRULE_createAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(ConsumerParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.TypedVariable()
	}
	{
		p.SetState(84)
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
	p.RuleIndex = ConsumerParserRULE_storeAction
	return p
}

func InitEmptyStoreActionContext(p *StoreActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_storeAction
}

func (*StoreActionContext) IsStoreActionContext() {}

func NewStoreActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreActionContext {
	var p = new(StoreActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_storeAction

	return p
}

func (s *StoreActionContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreActionContext) STORE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserSTORE, 0)
}

func (s *StoreActionContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(ConsumerParserIDENT)
}

func (s *StoreActionContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(ConsumerParserIDENT, i)
}

func (s *StoreActionContext) TO() antlr.TerminalNode {
	return s.GetToken(ConsumerParserTO, 0)
}

func (s *StoreActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitStoreAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) StoreAction() (localctx IStoreActionContext) {
	localctx = NewStoreActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConsumerParserRULE_storeAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(ConsumerParserSTORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(ConsumerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(ConsumerParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(ConsumerParserIDENT)
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
	p.RuleIndex = ConsumerParserRULE_updateAction
	return p
}

func InitEmptyUpdateActionContext(p *UpdateActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_updateAction
}

func (*UpdateActionContext) IsUpdateActionContext() {}

func NewUpdateActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateActionContext {
	var p = new(UpdateActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_updateAction

	return p
}

func (s *UpdateActionContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateActionContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserUPDATE, 0)
}

func (s *UpdateActionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserIDENT, 0)
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
	case ConsumerVisitor:
		return t.VisitUpdateAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) UpdateAction() (localctx IUpdateActionContext) {
	localctx = NewUpdateActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConsumerParserRULE_updateAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(ConsumerParserUPDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(ConsumerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
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
	p.RuleIndex = ConsumerParserRULE_emitAction
	return p
}

func InitEmptyEmitActionContext(p *EmitActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_emitAction
}

func (*EmitActionContext) IsEmitActionContext() {}

func NewEmitActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmitActionContext {
	var p = new(EmitActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_emitAction

	return p
}

func (s *EmitActionContext) GetParser() antlr.Parser { return s.parser }

func (s *EmitActionContext) EMIT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserEMIT, 0)
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
	case ConsumerVisitor:
		return t.VisitEmitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) EmitAction() (localctx IEmitActionContext) {
	localctx = NewEmitActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConsumerParserRULE_emitAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(ConsumerParserEMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.TypedVariable()
	}
	{
		p.SetState(97)
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
	p.RuleIndex = ConsumerParserRULE_callAction
	return p
}

func InitEmptyCallActionContext(p *CallActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_callAction
}

func (*CallActionContext) IsCallActionContext() {}

func NewCallActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallActionContext {
	var p = new(CallActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_callAction

	return p
}

func (s *CallActionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallActionContext) CALL() antlr.TerminalNode {
	return s.GetToken(ConsumerParserCALL, 0)
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
	case ConsumerVisitor:
		return t.VisitCallAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) CallAction() (localctx ICallActionContext) {
	localctx = NewCallActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ConsumerParserRULE_callAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(ConsumerParserCALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Selector()
	}
	{
		p.SetState(101)
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
	p.RuleIndex = ConsumerParserRULE_returnAction
	return p
}

func InitEmptyReturnActionContext(p *ReturnActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_returnAction
}

func (*ReturnActionContext) IsReturnActionContext() {}

func NewReturnActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnActionContext {
	var p = new(ReturnActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_returnAction

	return p
}

func (s *ReturnActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnActionContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ConsumerParserRETURN, 0)
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
	case ConsumerVisitor:
		return t.VisitReturnAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) ReturnAction() (localctx IReturnActionContext) {
	localctx = NewReturnActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ConsumerParserRULE_returnAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(ConsumerParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
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
	p.RuleIndex = ConsumerParserRULE_typedVariable
	return p
}

func InitEmptyTypedVariableContext(p *TypedVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_typedVariable
}

func (*TypedVariableContext) IsTypedVariableContext() {}

func NewTypedVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedVariableContext {
	var p = new(TypedVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_typedVariable

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
	return s.GetToken(ConsumerParserCOLON, 0)
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
	case ConsumerVisitor:
		return t.VisitTypedVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) TypedVariable() (localctx ITypedVariableContext) {
	localctx = NewTypedVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ConsumerParserRULE_typedVariable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Identifier()
	}
	{
		p.SetState(107)
		p.Match(ConsumerParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
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
	p.RuleIndex = ConsumerParserRULE_assignmentBlock
	return p
}

func InitEmptyAssignmentBlockContext(p *AssignmentBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_assignmentBlock
}

func (*AssignmentBlockContext) IsAssignmentBlockContext() {}

func NewAssignmentBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentBlockContext {
	var p = new(AssignmentBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_assignmentBlock

	return p
}

func (s *AssignmentBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserLBRACE, 0)
}

func (s *AssignmentBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserRBRACE, 0)
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
	case ConsumerVisitor:
		return t.VisitAssignmentBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) AssignmentBlock() (localctx IAssignmentBlockContext) {
	localctx = NewAssignmentBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ConsumerParserRULE_assignmentBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(ConsumerParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274907267072) != 0 {
		{
			p.SetState(111)
			p.Assignment()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(ConsumerParserRBRACE)
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
	p.RuleIndex = ConsumerParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_assignment

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
	return s.GetToken(ConsumerParserASSIGN, 0)
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
	case ConsumerVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ConsumerParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Identifier()
	}
	{
		p.SetState(120)
		p.Match(ConsumerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
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
	p.RuleIndex = ConsumerParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_condition

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
	case ConsumerVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ConsumerParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Expression()
	}
	{
		p.SetState(124)
		p.Comparator()
	}
	{
		p.SetState(125)
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
	p.RuleIndex = ConsumerParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(ConsumerParserEQ, 0)
}

func (s *ComparatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ConsumerParserNEQ, 0)
}

func (s *ComparatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserLTE, 0)
}

func (s *ComparatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserGTE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserLT, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserGT, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ConsumerParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6442696704) != 0) {
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
	p.RuleIndex = ConsumerParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_expression

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
	case ConsumerVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ConsumerParserRULE_expression)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Selector()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
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
	p.RuleIndex = ConsumerParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserIDENT, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ConsumerParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ConsumerParserRPAREN, 0)
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
	case ConsumerVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ConsumerParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(ConsumerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(ConsumerParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4123298627584) != 0 {
		{
			p.SetState(136)
			p.ArgumentList()
		}

	}
	{
		p.SetState(139)
		p.Match(ConsumerParserRPAREN)
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
	p.RuleIndex = ConsumerParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_argumentList

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
	return s.GetTokens(ConsumerParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ConsumerParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ConsumerParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Expression()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConsumerParserCOMMA {
		{
			p.SetState(142)
			p.Match(ConsumerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Expression()
		}

		p.SetState(148)
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
	p.RuleIndex = ConsumerParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_selector

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
	return s.GetTokens(ConsumerParserDOT)
}

func (s *SelectorContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ConsumerParserDOT, i)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ConsumerParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Identifier()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConsumerParserDOT {
		{
			p.SetState(150)
			p.Match(ConsumerParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Identifier()
		}

		p.SetState(156)
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
	p.RuleIndex = ConsumerParserRULE_typeRef
	return p
}

func InitEmptyTypeRefContext(p *TypeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_typeRef
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_typeRef

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
	return s.GetToken(ConsumerParserLIST, 0)
}

func (s *TypeRefContext) LT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserLT, 0)
}

func (s *TypeRefContext) GT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserGT, 0)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitTypeRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ConsumerParserRULE_typeRef)
	var _la int

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ConsumerParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.TypeName()
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&103616086016) != 0 {
			{
				p.SetState(158)
				p.NumberConstraint()
			}

		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ConsumerParserQUESTION {
			{
				p.SetState(161)
				p.OptionalMarker()
			}

		}

	case ConsumerParserLIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(ConsumerParserLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(ConsumerParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.TypeName()
		}
		{
			p.SetState(167)
			p.Match(ConsumerParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ConsumerParserQUESTION {
			{
				p.SetState(168)
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
	p.RuleIndex = ConsumerParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserIDENT, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ConsumerParserRULE_typeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(ConsumerParserIDENT)
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
	p.RuleIndex = ConsumerParserRULE_numberConstraint
	return p
}

func InitEmptyNumberConstraintContext(p *NumberConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_numberConstraint
}

func (*NumberConstraintContext) IsNumberConstraintContext() {}

func NewNumberConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberConstraintContext {
	var p = new(NumberConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_numberConstraint

	return p
}

func (s *NumberConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberConstraintContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ConsumerParserPLUS, 0)
}

func (s *NumberConstraintContext) STAR() antlr.TerminalNode {
	return s.GetToken(ConsumerParserSTAR, 0)
}

func (s *NumberConstraintContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ConsumerParserLBRACKET, 0)
}

func (s *NumberConstraintContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ConsumerParserCOMMA, 0)
}

func (s *NumberConstraintContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ConsumerParserRBRACKET, 0)
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
	case ConsumerVisitor:
		return t.VisitNumberConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) NumberConstraint() (localctx INumberConstraintContext) {
	localctx = NewNumberConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ConsumerParserRULE_numberConstraint)
	var _la int

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ConsumerParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Match(ConsumerParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ConsumerParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Match(ConsumerParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ConsumerParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.Match(ConsumerParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ConsumerParserFLOAT || _la == ConsumerParserINT {
			{
				p.SetState(178)
				p.NumberValue()
			}

		}
		{
			p.SetState(181)
			p.Match(ConsumerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ConsumerParserFLOAT || _la == ConsumerParserINT {
			{
				p.SetState(182)
				p.NumberValue()
			}

		}
		{
			p.SetState(185)
			p.Match(ConsumerParserRBRACKET)
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
	p.RuleIndex = ConsumerParserRULE_optionalMarker
	return p
}

func InitEmptyOptionalMarkerContext(p *OptionalMarkerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_optionalMarker
}

func (*OptionalMarkerContext) IsOptionalMarkerContext() {}

func NewOptionalMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalMarkerContext {
	var p = new(OptionalMarkerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_optionalMarker

	return p
}

func (s *OptionalMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalMarkerContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ConsumerParserQUESTION, 0)
}

func (s *OptionalMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalMarkerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitOptionalMarker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) OptionalMarker() (localctx IOptionalMarkerContext) {
	localctx = NewOptionalMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ConsumerParserRULE_optionalMarker)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(ConsumerParserQUESTION)
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
	p.RuleIndex = ConsumerParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ConsumerParserSTRING, 0)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserFLOAT, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ConsumerParserFALSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ConsumerParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848391360512) != 0) {
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
	p.RuleIndex = ConsumerParserRULE_numberValue
	return p
}

func InitEmptyNumberValueContext(p *NumberValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_numberValue
}

func (*NumberValueContext) IsNumberValueContext() {}

func NewNumberValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberValueContext {
	var p = new(NumberValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_numberValue

	return p
}

func (s *NumberValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserINT, 0)
}

func (s *NumberValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserFLOAT, 0)
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitNumberValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) NumberValue() (localctx INumberValueContext) {
	localctx = NewNumberValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ConsumerParserRULE_numberValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConsumerParserFLOAT || _la == ConsumerParserINT) {
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
	p.RuleIndex = ConsumerParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConsumerParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumerParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserIDENT, 0)
}

func (s *IdentifierContext) INPUT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserINPUT, 0)
}

func (s *IdentifierContext) EVENT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserEVENT, 0)
}

func (s *IdentifierContext) RESULT() antlr.TerminalNode {
	return s.GetToken(ConsumerParserRESULT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConsumerVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConsumerParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ConsumerParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274907267072) != 0) {
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
