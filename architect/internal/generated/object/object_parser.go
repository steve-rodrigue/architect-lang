// Code generated from internal/infrastructure/antlr/grammars/Object.g4 by ANTLR 4.13.2. DO NOT EDIT.

package object // Object
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

type ObjectParser struct {
	*antlr.BaseParser
}

var ObjectParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func objectParserInit() {
	staticData := &ObjectParserStaticData
	staticData.LiteralNames = []string{
		"", "'object'", "'history_of'", "'primary'", "'unique'", "'renamed_from'",
		"'deprecated'", "'{'", "'}'", "", "", "", "'List'", "'input'", "'event'",
		"'result'", "'true'", "'false'", "'('", "')'", "'['", "']'", "'<'",
		"'>'", "','", "'|'", "'+'", "'*'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "OBJECT", "HISTORY_OF", "PRIMARY", "UNIQUE", "RENAMED_FROM", "DEPRECATED",
		"LBRACE", "RBRACE", "LINE_COMMENT", "BLOCK_COMMENT", "WS", "LIST", "INPUT",
		"EVENT", "RESULT", "TRUE", "FALSE", "LPAREN", "RPAREN", "LBRACKET",
		"RBRACKET", "LT", "GT", "COMMA", "PIPE", "PLUS", "STAR", "QUESTION",
		"IDENT", "FLOAT", "INT", "STRING",
	}
	staticData.RuleNames = []string{
		"program", "objectDecl", "objectModifier", "fieldDecl", "defaultValue",
		"fieldModifier", "typeRef", "typeName", "numberConstraint", "optionalMarker",
		"value", "numberValue", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 113, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5,
		1, 33, 8, 1, 10, 1, 12, 1, 36, 9, 1, 1, 1, 1, 1, 5, 1, 40, 8, 1, 10, 1,
		12, 1, 43, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 52, 8,
		2, 1, 3, 1, 3, 1, 3, 3, 3, 57, 8, 3, 1, 3, 5, 3, 60, 8, 3, 10, 3, 12, 3,
		63, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 72, 8, 5, 1,
		6, 1, 6, 3, 6, 76, 8, 6, 1, 6, 3, 6, 79, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 86, 8, 6, 3, 6, 88, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 96, 8, 8, 1, 8, 1, 8, 3, 8, 100, 8, 8, 1, 8, 3, 8, 103, 8, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 0, 0, 13,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 3, 2, 0, 16, 17, 30,
		32, 1, 0, 30, 31, 2, 0, 13, 15, 29, 29, 116, 0, 26, 1, 0, 0, 0, 2, 29,
		1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10,
		71, 1, 0, 0, 0, 12, 87, 1, 0, 0, 0, 14, 89, 1, 0, 0, 0, 16, 102, 1, 0,
		0, 0, 18, 104, 1, 0, 0, 0, 20, 106, 1, 0, 0, 0, 22, 108, 1, 0, 0, 0, 24,
		110, 1, 0, 0, 0, 26, 27, 3, 2, 1, 0, 27, 28, 5, 0, 0, 1, 28, 1, 1, 0, 0,
		0, 29, 30, 5, 1, 0, 0, 30, 34, 5, 29, 0, 0, 31, 33, 3, 4, 2, 0, 32, 31,
		1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0,
		35, 37, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 41, 5, 7, 0, 0, 38, 40, 3,
		6, 3, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41,
		42, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 5, 8, 0,
		0, 45, 3, 1, 0, 0, 0, 46, 47, 5, 2, 0, 0, 47, 52, 5, 29, 0, 0, 48, 49,
		5, 5, 0, 0, 49, 52, 5, 29, 0, 0, 50, 52, 5, 6, 0, 0, 51, 46, 1, 0, 0, 0,
		51, 48, 1, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 5, 1, 0, 0, 0, 53, 54, 5, 29,
		0, 0, 54, 56, 3, 12, 6, 0, 55, 57, 3, 8, 4, 0, 56, 55, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 61, 1, 0, 0, 0, 58, 60, 3, 10, 5, 0, 59, 58, 1, 0,
		0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 7,
		1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 3, 20, 10, 0, 65, 9, 1, 0, 0, 0,
		66, 72, 5, 3, 0, 0, 67, 72, 5, 4, 0, 0, 68, 69, 5, 5, 0, 0, 69, 72, 5,
		29, 0, 0, 70, 72, 5, 6, 0, 0, 71, 66, 1, 0, 0, 0, 71, 67, 1, 0, 0, 0, 71,
		68, 1, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 11, 1, 0, 0, 0, 73, 75, 3, 14,
		7, 0, 74, 76, 3, 16, 8, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76,
		78, 1, 0, 0, 0, 77, 79, 3, 18, 9, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0,
		0, 0, 79, 88, 1, 0, 0, 0, 80, 81, 5, 12, 0, 0, 81, 82, 5, 22, 0, 0, 82,
		83, 3, 14, 7, 0, 83, 85, 5, 23, 0, 0, 84, 86, 3, 18, 9, 0, 85, 84, 1, 0,
		0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 73, 1, 0, 0, 0, 87, 80,
		1, 0, 0, 0, 88, 13, 1, 0, 0, 0, 89, 90, 5, 29, 0, 0, 90, 15, 1, 0, 0, 0,
		91, 103, 5, 26, 0, 0, 92, 103, 5, 27, 0, 0, 93, 95, 5, 20, 0, 0, 94, 96,
		3, 22, 11, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0,
		0, 97, 99, 5, 24, 0, 0, 98, 100, 3, 22, 11, 0, 99, 98, 1, 0, 0, 0, 99,
		100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 5, 21, 0, 0, 102, 91,
		1, 0, 0, 0, 102, 92, 1, 0, 0, 0, 102, 93, 1, 0, 0, 0, 103, 17, 1, 0, 0,
		0, 104, 105, 5, 28, 0, 0, 105, 19, 1, 0, 0, 0, 106, 107, 7, 0, 0, 0, 107,
		21, 1, 0, 0, 0, 108, 109, 7, 1, 0, 0, 109, 23, 1, 0, 0, 0, 110, 111, 7,
		2, 0, 0, 111, 25, 1, 0, 0, 0, 13, 34, 41, 51, 56, 61, 71, 75, 78, 85, 87,
		95, 99, 102,
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

// ObjectParserInit initializes any static state used to implement ObjectParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewObjectParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ObjectParserInit() {
	staticData := &ObjectParserStaticData
	staticData.once.Do(objectParserInit)
}

// NewObjectParser produces a new parser instance for the optional input antlr.TokenStream.
func NewObjectParser(input antlr.TokenStream) *ObjectParser {
	ObjectParserInit()
	this := new(ObjectParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ObjectParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Object.g4"

	return this
}

// ObjectParser tokens.
const (
	ObjectParserEOF           = antlr.TokenEOF
	ObjectParserOBJECT        = 1
	ObjectParserHISTORY_OF    = 2
	ObjectParserPRIMARY       = 3
	ObjectParserUNIQUE        = 4
	ObjectParserRENAMED_FROM  = 5
	ObjectParserDEPRECATED    = 6
	ObjectParserLBRACE        = 7
	ObjectParserRBRACE        = 8
	ObjectParserLINE_COMMENT  = 9
	ObjectParserBLOCK_COMMENT = 10
	ObjectParserWS            = 11
	ObjectParserLIST          = 12
	ObjectParserINPUT         = 13
	ObjectParserEVENT         = 14
	ObjectParserRESULT        = 15
	ObjectParserTRUE          = 16
	ObjectParserFALSE         = 17
	ObjectParserLPAREN        = 18
	ObjectParserRPAREN        = 19
	ObjectParserLBRACKET      = 20
	ObjectParserRBRACKET      = 21
	ObjectParserLT            = 22
	ObjectParserGT            = 23
	ObjectParserCOMMA         = 24
	ObjectParserPIPE          = 25
	ObjectParserPLUS          = 26
	ObjectParserSTAR          = 27
	ObjectParserQUESTION      = 28
	ObjectParserIDENT         = 29
	ObjectParserFLOAT         = 30
	ObjectParserINT           = 31
	ObjectParserSTRING        = 32
)

// ObjectParser rules.
const (
	ObjectParserRULE_program          = 0
	ObjectParserRULE_objectDecl       = 1
	ObjectParserRULE_objectModifier   = 2
	ObjectParserRULE_fieldDecl        = 3
	ObjectParserRULE_defaultValue     = 4
	ObjectParserRULE_fieldModifier    = 5
	ObjectParserRULE_typeRef          = 6
	ObjectParserRULE_typeName         = 7
	ObjectParserRULE_numberConstraint = 8
	ObjectParserRULE_optionalMarker   = 9
	ObjectParserRULE_value            = 10
	ObjectParserRULE_numberValue      = 11
	ObjectParserRULE_identifier       = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ObjectDecl() IObjectDeclContext
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
	p.RuleIndex = ObjectParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ObjectDecl() IObjectDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ObjectParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ObjectParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.ObjectDecl()
	}
	{
		p.SetState(27)
		p.Match(ObjectParserEOF)
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

// IObjectDeclContext is an interface to support dynamic dispatch.
type IObjectDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBJECT() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllObjectModifier() []IObjectModifierContext
	ObjectModifier(i int) IObjectModifierContext
	AllFieldDecl() []IFieldDeclContext
	FieldDecl(i int) IFieldDeclContext

	// IsObjectDeclContext differentiates from other interfaces.
	IsObjectDeclContext()
}

type ObjectDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectDeclContext() *ObjectDeclContext {
	var p = new(ObjectDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_objectDecl
	return p
}

func InitEmptyObjectDeclContext(p *ObjectDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_objectDecl
}

func (*ObjectDeclContext) IsObjectDeclContext() {}

func NewObjectDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectDeclContext {
	var p = new(ObjectDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_objectDecl

	return p
}

func (s *ObjectDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectDeclContext) OBJECT() antlr.TerminalNode {
	return s.GetToken(ObjectParserOBJECT, 0)
}

func (s *ObjectDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ObjectParserIDENT, 0)
}

func (s *ObjectDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ObjectParserLBRACE, 0)
}

func (s *ObjectDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ObjectParserRBRACE, 0)
}

func (s *ObjectDeclContext) AllObjectModifier() []IObjectModifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectModifierContext); ok {
			len++
		}
	}

	tst := make([]IObjectModifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectModifierContext); ok {
			tst[i] = t.(IObjectModifierContext)
			i++
		}
	}

	return tst
}

func (s *ObjectDeclContext) ObjectModifier(i int) IObjectModifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectModifierContext); ok {
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

	return t.(IObjectModifierContext)
}

func (s *ObjectDeclContext) AllFieldDecl() []IFieldDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDeclContext); ok {
			len++
		}
	}

	tst := make([]IFieldDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDeclContext); ok {
			tst[i] = t.(IFieldDeclContext)
			i++
		}
	}

	return tst
}

func (s *ObjectDeclContext) FieldDecl(i int) IFieldDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDeclContext); ok {
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

	return t.(IFieldDeclContext)
}

func (s *ObjectDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitObjectDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) ObjectDecl() (localctx IObjectDeclContext) {
	localctx = NewObjectDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ObjectParserRULE_objectDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(ObjectParserOBJECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Match(ObjectParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&100) != 0 {
		{
			p.SetState(31)
			p.ObjectModifier()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.Match(ObjectParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ObjectParserIDENT {
		{
			p.SetState(38)
			p.FieldDecl()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.Match(ObjectParserRBRACE)
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

// IObjectModifierContext is an interface to support dynamic dispatch.
type IObjectModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HISTORY_OF() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	RENAMED_FROM() antlr.TerminalNode
	DEPRECATED() antlr.TerminalNode

	// IsObjectModifierContext differentiates from other interfaces.
	IsObjectModifierContext()
}

type ObjectModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectModifierContext() *ObjectModifierContext {
	var p = new(ObjectModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_objectModifier
	return p
}

func InitEmptyObjectModifierContext(p *ObjectModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_objectModifier
}

func (*ObjectModifierContext) IsObjectModifierContext() {}

func NewObjectModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectModifierContext {
	var p = new(ObjectModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_objectModifier

	return p
}

func (s *ObjectModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectModifierContext) HISTORY_OF() antlr.TerminalNode {
	return s.GetToken(ObjectParserHISTORY_OF, 0)
}

func (s *ObjectModifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ObjectParserIDENT, 0)
}

func (s *ObjectModifierContext) RENAMED_FROM() antlr.TerminalNode {
	return s.GetToken(ObjectParserRENAMED_FROM, 0)
}

func (s *ObjectModifierContext) DEPRECATED() antlr.TerminalNode {
	return s.GetToken(ObjectParserDEPRECATED, 0)
}

func (s *ObjectModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitObjectModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) ObjectModifier() (localctx IObjectModifierContext) {
	localctx = NewObjectModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ObjectParserRULE_objectModifier)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ObjectParserHISTORY_OF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(ObjectParserHISTORY_OF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(ObjectParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectParserRENAMED_FROM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(ObjectParserRENAMED_FROM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(ObjectParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectParserDEPRECATED:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(ObjectParserDEPRECATED)
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

// IFieldDeclContext is an interface to support dynamic dispatch.
type IFieldDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	TypeRef() ITypeRefContext
	DefaultValue() IDefaultValueContext
	AllFieldModifier() []IFieldModifierContext
	FieldModifier(i int) IFieldModifierContext

	// IsFieldDeclContext differentiates from other interfaces.
	IsFieldDeclContext()
}

type FieldDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclContext() *FieldDeclContext {
	var p = new(FieldDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_fieldDecl
	return p
}

func InitEmptyFieldDeclContext(p *FieldDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_fieldDecl
}

func (*FieldDeclContext) IsFieldDeclContext() {}

func NewFieldDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclContext {
	var p = new(FieldDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_fieldDecl

	return p
}

func (s *FieldDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ObjectParserIDENT, 0)
}

func (s *FieldDeclContext) TypeRef() ITypeRefContext {
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

func (s *FieldDeclContext) DefaultValue() IDefaultValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *FieldDeclContext) AllFieldModifier() []IFieldModifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldModifierContext); ok {
			len++
		}
	}

	tst := make([]IFieldModifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldModifierContext); ok {
			tst[i] = t.(IFieldModifierContext)
			i++
		}
	}

	return tst
}

func (s *FieldDeclContext) FieldModifier(i int) IFieldModifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldModifierContext); ok {
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

	return t.(IFieldModifierContext)
}

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitFieldDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) FieldDecl() (localctx IFieldDeclContext) {
	localctx = NewFieldDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ObjectParserRULE_fieldDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(ObjectParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.TypeRef()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7516389376) != 0 {
		{
			p.SetState(55)
			p.DefaultValue()
		}

	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0 {
		{
			p.SetState(58)
			p.FieldModifier()
		}

		p.SetState(63)
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

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_defaultValue
	return p
}

func InitEmptyDefaultValueContext(p *DefaultValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_defaultValue
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) Value() IValueContext {
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

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitDefaultValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ObjectParserRULE_defaultValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Value()
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

// IFieldModifierContext is an interface to support dynamic dispatch.
type IFieldModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMARY() antlr.TerminalNode
	UNIQUE() antlr.TerminalNode
	RENAMED_FROM() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	DEPRECATED() antlr.TerminalNode

	// IsFieldModifierContext differentiates from other interfaces.
	IsFieldModifierContext()
}

type FieldModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldModifierContext() *FieldModifierContext {
	var p = new(FieldModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_fieldModifier
	return p
}

func InitEmptyFieldModifierContext(p *FieldModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_fieldModifier
}

func (*FieldModifierContext) IsFieldModifierContext() {}

func NewFieldModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldModifierContext {
	var p = new(FieldModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_fieldModifier

	return p
}

func (s *FieldModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldModifierContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(ObjectParserPRIMARY, 0)
}

func (s *FieldModifierContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(ObjectParserUNIQUE, 0)
}

func (s *FieldModifierContext) RENAMED_FROM() antlr.TerminalNode {
	return s.GetToken(ObjectParserRENAMED_FROM, 0)
}

func (s *FieldModifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ObjectParserIDENT, 0)
}

func (s *FieldModifierContext) DEPRECATED() antlr.TerminalNode {
	return s.GetToken(ObjectParserDEPRECATED, 0)
}

func (s *FieldModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitFieldModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) FieldModifier() (localctx IFieldModifierContext) {
	localctx = NewFieldModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ObjectParserRULE_fieldModifier)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ObjectParserPRIMARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(ObjectParserPRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectParserUNIQUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(ObjectParserUNIQUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectParserRENAMED_FROM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Match(ObjectParserRENAMED_FROM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(ObjectParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectParserDEPRECATED:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Match(ObjectParserDEPRECATED)
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
	p.RuleIndex = ObjectParserRULE_typeRef
	return p
}

func InitEmptyTypeRefContext(p *TypeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_typeRef
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_typeRef

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
	return s.GetToken(ObjectParserLIST, 0)
}

func (s *TypeRefContext) LT() antlr.TerminalNode {
	return s.GetToken(ObjectParserLT, 0)
}

func (s *TypeRefContext) GT() antlr.TerminalNode {
	return s.GetToken(ObjectParserGT, 0)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitTypeRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ObjectParserRULE_typeRef)
	var _la int

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ObjectParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.TypeName()
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&202375168) != 0 {
			{
				p.SetState(74)
				p.NumberConstraint()
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ObjectParserQUESTION {
			{
				p.SetState(77)
				p.OptionalMarker()
			}

		}

	case ObjectParserLIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(ObjectParserLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(ObjectParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.TypeName()
		}
		{
			p.SetState(83)
			p.Match(ObjectParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ObjectParserQUESTION {
			{
				p.SetState(84)
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
	p.RuleIndex = ObjectParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ObjectParserIDENT, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ObjectParserRULE_typeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(ObjectParserIDENT)
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
	p.RuleIndex = ObjectParserRULE_numberConstraint
	return p
}

func InitEmptyNumberConstraintContext(p *NumberConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_numberConstraint
}

func (*NumberConstraintContext) IsNumberConstraintContext() {}

func NewNumberConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberConstraintContext {
	var p = new(NumberConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_numberConstraint

	return p
}

func (s *NumberConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberConstraintContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ObjectParserPLUS, 0)
}

func (s *NumberConstraintContext) STAR() antlr.TerminalNode {
	return s.GetToken(ObjectParserSTAR, 0)
}

func (s *NumberConstraintContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ObjectParserLBRACKET, 0)
}

func (s *NumberConstraintContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ObjectParserCOMMA, 0)
}

func (s *NumberConstraintContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ObjectParserRBRACKET, 0)
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
	case ObjectVisitor:
		return t.VisitNumberConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) NumberConstraint() (localctx INumberConstraintContext) {
	localctx = NewNumberConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ObjectParserRULE_numberConstraint)
	var _la int

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ObjectParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Match(ObjectParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(ObjectParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ObjectParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(ObjectParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ObjectParserFLOAT || _la == ObjectParserINT {
			{
				p.SetState(94)
				p.NumberValue()
			}

		}
		{
			p.SetState(97)
			p.Match(ObjectParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ObjectParserFLOAT || _la == ObjectParserINT {
			{
				p.SetState(98)
				p.NumberValue()
			}

		}
		{
			p.SetState(101)
			p.Match(ObjectParserRBRACKET)
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
	p.RuleIndex = ObjectParserRULE_optionalMarker
	return p
}

func InitEmptyOptionalMarkerContext(p *OptionalMarkerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_optionalMarker
}

func (*OptionalMarkerContext) IsOptionalMarkerContext() {}

func NewOptionalMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalMarkerContext {
	var p = new(OptionalMarkerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_optionalMarker

	return p
}

func (s *OptionalMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalMarkerContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ObjectParserQUESTION, 0)
}

func (s *OptionalMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalMarkerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitOptionalMarker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) OptionalMarker() (localctx IOptionalMarkerContext) {
	localctx = NewOptionalMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ObjectParserRULE_optionalMarker)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(ObjectParserQUESTION)
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
	p.RuleIndex = ObjectParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ObjectParserSTRING, 0)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ObjectParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ObjectParserFLOAT, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ObjectParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ObjectParserFALSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ObjectParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7516389376) != 0) {
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
	p.RuleIndex = ObjectParserRULE_numberValue
	return p
}

func InitEmptyNumberValueContext(p *NumberValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_numberValue
}

func (*NumberValueContext) IsNumberValueContext() {}

func NewNumberValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberValueContext {
	var p = new(NumberValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_numberValue

	return p
}

func (s *NumberValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ObjectParserINT, 0)
}

func (s *NumberValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ObjectParserFLOAT, 0)
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitNumberValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) NumberValue() (localctx INumberValueContext) {
	localctx = NewNumberValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ObjectParserRULE_numberValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ObjectParserFLOAT || _la == ObjectParserINT) {
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
	p.RuleIndex = ObjectParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ObjectParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ObjectParserIDENT, 0)
}

func (s *IdentifierContext) INPUT() antlr.TerminalNode {
	return s.GetToken(ObjectParserINPUT, 0)
}

func (s *IdentifierContext) EVENT() antlr.TerminalNode {
	return s.GetToken(ObjectParserEVENT, 0)
}

func (s *IdentifierContext) RESULT() antlr.TerminalNode {
	return s.GetToken(ObjectParserRESULT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ObjectVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ObjectParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ObjectParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536928256) != 0) {
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
