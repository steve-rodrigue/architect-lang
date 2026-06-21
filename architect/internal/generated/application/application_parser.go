// Code generated from internal/infrastructure/antlr/grammars/Application.g4 by ANTLR 4.13.2. DO NOT EDIT.

package application // Application
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

type ApplicationParser struct {
	*antlr.BaseParser
}

var ApplicationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func applicationParserInit() {
	staticData := &ApplicationParserStaticData
	staticData.LiteralNames = []string{
		"", "'application'", "'emits'", "'listens'", "'on'", "'objects'", "'endpoints'",
		"'consumers'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "APPLICATION", "EMITS", "LISTENS", "ON", "OBJECTS", "ENDPOINTS",
		"CONSUMERS", "LBRACE", "RBRACE", "IDENT", "INT", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "applicationDecl", "applicationBody", "portDecl", "objectsBlock",
		"endpointsBlock", "consumersBlock", "fileRef",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 79, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 45, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 50, 8, 4, 10, 4, 12, 4, 53,
		9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 60, 8, 5, 10, 5, 12, 5, 63, 9,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 70, 8, 6, 10, 6, 12, 6, 73, 9, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 0,
		78, 0, 16, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 34, 1, 0, 0, 0, 6, 44, 1,
		0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 66, 1, 0, 0, 0, 14,
		76, 1, 0, 0, 0, 16, 17, 3, 2, 1, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0,
		0, 19, 20, 5, 1, 0, 0, 20, 21, 5, 10, 0, 0, 21, 25, 5, 8, 0, 0, 22, 24,
		3, 4, 2, 0, 23, 22, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0,
		25, 26, 1, 0, 0, 0, 26, 28, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 5,
		9, 0, 0, 29, 3, 1, 0, 0, 0, 30, 35, 3, 6, 3, 0, 31, 35, 3, 8, 4, 0, 32,
		35, 3, 10, 5, 0, 33, 35, 3, 12, 6, 0, 34, 30, 1, 0, 0, 0, 34, 31, 1, 0,
		0, 0, 34, 32, 1, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35, 5, 1, 0, 0, 0, 36, 37,
		5, 2, 0, 0, 37, 38, 5, 10, 0, 0, 38, 39, 5, 4, 0, 0, 39, 45, 5, 11, 0,
		0, 40, 41, 5, 3, 0, 0, 41, 42, 5, 10, 0, 0, 42, 43, 5, 4, 0, 0, 43, 45,
		5, 11, 0, 0, 44, 36, 1, 0, 0, 0, 44, 40, 1, 0, 0, 0, 45, 7, 1, 0, 0, 0,
		46, 47, 5, 5, 0, 0, 47, 51, 5, 8, 0, 0, 48, 50, 3, 14, 7, 0, 49, 48, 1,
		0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		54, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 5, 9, 0, 0, 55, 9, 1, 0, 0,
		0, 56, 57, 5, 6, 0, 0, 57, 61, 5, 8, 0, 0, 58, 60, 3, 14, 7, 0, 59, 58,
		1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 64, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 5, 9, 0, 0, 65, 11, 1,
		0, 0, 0, 66, 67, 5, 7, 0, 0, 67, 71, 5, 8, 0, 0, 68, 70, 3, 14, 7, 0, 69,
		68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0,
		0, 72, 74, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 9, 0, 0, 75, 13,
		1, 0, 0, 0, 76, 77, 5, 12, 0, 0, 77, 15, 1, 0, 0, 0, 6, 25, 34, 44, 51,
		61, 71,
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

// ApplicationParserInit initializes any static state used to implement ApplicationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewApplicationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ApplicationParserInit() {
	staticData := &ApplicationParserStaticData
	staticData.once.Do(applicationParserInit)
}

// NewApplicationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewApplicationParser(input antlr.TokenStream) *ApplicationParser {
	ApplicationParserInit()
	this := new(ApplicationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ApplicationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Application.g4"

	return this
}

// ApplicationParser tokens.
const (
	ApplicationParserEOF           = antlr.TokenEOF
	ApplicationParserAPPLICATION   = 1
	ApplicationParserEMITS         = 2
	ApplicationParserLISTENS       = 3
	ApplicationParserON            = 4
	ApplicationParserOBJECTS       = 5
	ApplicationParserENDPOINTS     = 6
	ApplicationParserCONSUMERS     = 7
	ApplicationParserLBRACE        = 8
	ApplicationParserRBRACE        = 9
	ApplicationParserIDENT         = 10
	ApplicationParserINT           = 11
	ApplicationParserSTRING        = 12
	ApplicationParserLINE_COMMENT  = 13
	ApplicationParserBLOCK_COMMENT = 14
	ApplicationParserWS            = 15
)

// ApplicationParser rules.
const (
	ApplicationParserRULE_program         = 0
	ApplicationParserRULE_applicationDecl = 1
	ApplicationParserRULE_applicationBody = 2
	ApplicationParserRULE_portDecl        = 3
	ApplicationParserRULE_objectsBlock    = 4
	ApplicationParserRULE_endpointsBlock  = 5
	ApplicationParserRULE_consumersBlock  = 6
	ApplicationParserRULE_fileRef         = 7
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ApplicationDecl() IApplicationDeclContext
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
	p.RuleIndex = ApplicationParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ApplicationDecl() IApplicationDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApplicationDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApplicationDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ApplicationParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ApplicationParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.ApplicationDecl()
	}
	{
		p.SetState(17)
		p.Match(ApplicationParserEOF)
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

// IApplicationDeclContext is an interface to support dynamic dispatch.
type IApplicationDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	APPLICATION() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllApplicationBody() []IApplicationBodyContext
	ApplicationBody(i int) IApplicationBodyContext

	// IsApplicationDeclContext differentiates from other interfaces.
	IsApplicationDeclContext()
}

type ApplicationDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApplicationDeclContext() *ApplicationDeclContext {
	var p = new(ApplicationDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_applicationDecl
	return p
}

func InitEmptyApplicationDeclContext(p *ApplicationDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_applicationDecl
}

func (*ApplicationDeclContext) IsApplicationDeclContext() {}

func NewApplicationDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplicationDeclContext {
	var p = new(ApplicationDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_applicationDecl

	return p
}

func (s *ApplicationDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplicationDeclContext) APPLICATION() antlr.TerminalNode {
	return s.GetToken(ApplicationParserAPPLICATION, 0)
}

func (s *ApplicationDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ApplicationParserIDENT, 0)
}

func (s *ApplicationDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserLBRACE, 0)
}

func (s *ApplicationDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserRBRACE, 0)
}

func (s *ApplicationDeclContext) AllApplicationBody() []IApplicationBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IApplicationBodyContext); ok {
			len++
		}
	}

	tst := make([]IApplicationBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IApplicationBodyContext); ok {
			tst[i] = t.(IApplicationBodyContext)
			i++
		}
	}

	return tst
}

func (s *ApplicationDeclContext) ApplicationBody(i int) IApplicationBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApplicationBodyContext); ok {
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

	return t.(IApplicationBodyContext)
}

func (s *ApplicationDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicationDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApplicationDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitApplicationDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) ApplicationDecl() (localctx IApplicationDeclContext) {
	localctx = NewApplicationDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ApplicationParserRULE_applicationDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(ApplicationParserAPPLICATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Match(ApplicationParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)
		p.Match(ApplicationParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&236) != 0 {
		{
			p.SetState(22)
			p.ApplicationBody()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
		p.Match(ApplicationParserRBRACE)
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

// IApplicationBodyContext is an interface to support dynamic dispatch.
type IApplicationBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PortDecl() IPortDeclContext
	ObjectsBlock() IObjectsBlockContext
	EndpointsBlock() IEndpointsBlockContext
	ConsumersBlock() IConsumersBlockContext

	// IsApplicationBodyContext differentiates from other interfaces.
	IsApplicationBodyContext()
}

type ApplicationBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApplicationBodyContext() *ApplicationBodyContext {
	var p = new(ApplicationBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_applicationBody
	return p
}

func InitEmptyApplicationBodyContext(p *ApplicationBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_applicationBody
}

func (*ApplicationBodyContext) IsApplicationBodyContext() {}

func NewApplicationBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplicationBodyContext {
	var p = new(ApplicationBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_applicationBody

	return p
}

func (s *ApplicationBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplicationBodyContext) PortDecl() IPortDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPortDeclContext)
}

func (s *ApplicationBodyContext) ObjectsBlock() IObjectsBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectsBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectsBlockContext)
}

func (s *ApplicationBodyContext) EndpointsBlock() IEndpointsBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointsBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndpointsBlockContext)
}

func (s *ApplicationBodyContext) ConsumersBlock() IConsumersBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsumersBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsumersBlockContext)
}

func (s *ApplicationBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicationBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApplicationBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitApplicationBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) ApplicationBody() (localctx IApplicationBodyContext) {
	localctx = NewApplicationBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ApplicationParserRULE_applicationBody)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApplicationParserEMITS, ApplicationParserLISTENS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.PortDecl()
		}

	case ApplicationParserOBJECTS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.ObjectsBlock()
		}

	case ApplicationParserENDPOINTS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.EndpointsBlock()
		}

	case ApplicationParserCONSUMERS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
			p.ConsumersBlock()
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

// IPortDeclContext is an interface to support dynamic dispatch.
type IPortDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EMITS() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ON() antlr.TerminalNode
	INT() antlr.TerminalNode
	LISTENS() antlr.TerminalNode

	// IsPortDeclContext differentiates from other interfaces.
	IsPortDeclContext()
}

type PortDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortDeclContext() *PortDeclContext {
	var p = new(PortDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_portDecl
	return p
}

func InitEmptyPortDeclContext(p *PortDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_portDecl
}

func (*PortDeclContext) IsPortDeclContext() {}

func NewPortDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortDeclContext {
	var p = new(PortDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_portDecl

	return p
}

func (s *PortDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PortDeclContext) EMITS() antlr.TerminalNode {
	return s.GetToken(ApplicationParserEMITS, 0)
}

func (s *PortDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ApplicationParserIDENT, 0)
}

func (s *PortDeclContext) ON() antlr.TerminalNode {
	return s.GetToken(ApplicationParserON, 0)
}

func (s *PortDeclContext) INT() antlr.TerminalNode {
	return s.GetToken(ApplicationParserINT, 0)
}

func (s *PortDeclContext) LISTENS() antlr.TerminalNode {
	return s.GetToken(ApplicationParserLISTENS, 0)
}

func (s *PortDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitPortDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) PortDecl() (localctx IPortDeclContext) {
	localctx = NewPortDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ApplicationParserRULE_portDecl)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApplicationParserEMITS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Match(ApplicationParserEMITS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Match(ApplicationParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(ApplicationParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(ApplicationParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ApplicationParserLISTENS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(ApplicationParserLISTENS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(ApplicationParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(ApplicationParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Match(ApplicationParserINT)
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

// IObjectsBlockContext is an interface to support dynamic dispatch.
type IObjectsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBJECTS() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFileRef() []IFileRefContext
	FileRef(i int) IFileRefContext

	// IsObjectsBlockContext differentiates from other interfaces.
	IsObjectsBlockContext()
}

type ObjectsBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectsBlockContext() *ObjectsBlockContext {
	var p = new(ObjectsBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_objectsBlock
	return p
}

func InitEmptyObjectsBlockContext(p *ObjectsBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_objectsBlock
}

func (*ObjectsBlockContext) IsObjectsBlockContext() {}

func NewObjectsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectsBlockContext {
	var p = new(ObjectsBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_objectsBlock

	return p
}

func (s *ObjectsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectsBlockContext) OBJECTS() antlr.TerminalNode {
	return s.GetToken(ApplicationParserOBJECTS, 0)
}

func (s *ObjectsBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserLBRACE, 0)
}

func (s *ObjectsBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserRBRACE, 0)
}

func (s *ObjectsBlockContext) AllFileRef() []IFileRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFileRefContext); ok {
			len++
		}
	}

	tst := make([]IFileRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFileRefContext); ok {
			tst[i] = t.(IFileRefContext)
			i++
		}
	}

	return tst
}

func (s *ObjectsBlockContext) FileRef(i int) IFileRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileRefContext); ok {
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

	return t.(IFileRefContext)
}

func (s *ObjectsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectsBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitObjectsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) ObjectsBlock() (localctx IObjectsBlockContext) {
	localctx = NewObjectsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ApplicationParserRULE_objectsBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(ApplicationParserOBJECTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(ApplicationParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ApplicationParserSTRING {
		{
			p.SetState(48)
			p.FileRef()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(ApplicationParserRBRACE)
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

// IEndpointsBlockContext is an interface to support dynamic dispatch.
type IEndpointsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENDPOINTS() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFileRef() []IFileRefContext
	FileRef(i int) IFileRefContext

	// IsEndpointsBlockContext differentiates from other interfaces.
	IsEndpointsBlockContext()
}

type EndpointsBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointsBlockContext() *EndpointsBlockContext {
	var p = new(EndpointsBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_endpointsBlock
	return p
}

func InitEmptyEndpointsBlockContext(p *EndpointsBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_endpointsBlock
}

func (*EndpointsBlockContext) IsEndpointsBlockContext() {}

func NewEndpointsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointsBlockContext {
	var p = new(EndpointsBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_endpointsBlock

	return p
}

func (s *EndpointsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointsBlockContext) ENDPOINTS() antlr.TerminalNode {
	return s.GetToken(ApplicationParserENDPOINTS, 0)
}

func (s *EndpointsBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserLBRACE, 0)
}

func (s *EndpointsBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserRBRACE, 0)
}

func (s *EndpointsBlockContext) AllFileRef() []IFileRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFileRefContext); ok {
			len++
		}
	}

	tst := make([]IFileRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFileRefContext); ok {
			tst[i] = t.(IFileRefContext)
			i++
		}
	}

	return tst
}

func (s *EndpointsBlockContext) FileRef(i int) IFileRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileRefContext); ok {
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

	return t.(IFileRefContext)
}

func (s *EndpointsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointsBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitEndpointsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) EndpointsBlock() (localctx IEndpointsBlockContext) {
	localctx = NewEndpointsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ApplicationParserRULE_endpointsBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(ApplicationParserENDPOINTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(ApplicationParserLBRACE)
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

	for _la == ApplicationParserSTRING {
		{
			p.SetState(58)
			p.FileRef()
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
		p.Match(ApplicationParserRBRACE)
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

// IConsumersBlockContext is an interface to support dynamic dispatch.
type IConsumersBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONSUMERS() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFileRef() []IFileRefContext
	FileRef(i int) IFileRefContext

	// IsConsumersBlockContext differentiates from other interfaces.
	IsConsumersBlockContext()
}

type ConsumersBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConsumersBlockContext() *ConsumersBlockContext {
	var p = new(ConsumersBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_consumersBlock
	return p
}

func InitEmptyConsumersBlockContext(p *ConsumersBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_consumersBlock
}

func (*ConsumersBlockContext) IsConsumersBlockContext() {}

func NewConsumersBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsumersBlockContext {
	var p = new(ConsumersBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_consumersBlock

	return p
}

func (s *ConsumersBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsumersBlockContext) CONSUMERS() antlr.TerminalNode {
	return s.GetToken(ApplicationParserCONSUMERS, 0)
}

func (s *ConsumersBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserLBRACE, 0)
}

func (s *ConsumersBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ApplicationParserRBRACE, 0)
}

func (s *ConsumersBlockContext) AllFileRef() []IFileRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFileRefContext); ok {
			len++
		}
	}

	tst := make([]IFileRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFileRefContext); ok {
			tst[i] = t.(IFileRefContext)
			i++
		}
	}

	return tst
}

func (s *ConsumersBlockContext) FileRef(i int) IFileRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileRefContext); ok {
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

	return t.(IFileRefContext)
}

func (s *ConsumersBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsumersBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsumersBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitConsumersBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) ConsumersBlock() (localctx IConsumersBlockContext) {
	localctx = NewConsumersBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ApplicationParserRULE_consumersBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(ApplicationParserCONSUMERS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(ApplicationParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ApplicationParserSTRING {
		{
			p.SetState(68)
			p.FileRef()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.Match(ApplicationParserRBRACE)
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

// IFileRefContext is an interface to support dynamic dispatch.
type IFileRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsFileRefContext differentiates from other interfaces.
	IsFileRefContext()
}

type FileRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileRefContext() *FileRefContext {
	var p = new(FileRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_fileRef
	return p
}

func InitEmptyFileRefContext(p *FileRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApplicationParserRULE_fileRef
}

func (*FileRefContext) IsFileRefContext() {}

func NewFileRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileRefContext {
	var p = new(FileRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApplicationParserRULE_fileRef

	return p
}

func (s *FileRefContext) GetParser() antlr.Parser { return s.parser }

func (s *FileRefContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApplicationParserSTRING, 0)
}

func (s *FileRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApplicationVisitor:
		return t.VisitFileRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApplicationParser) FileRef() (localctx IFileRefContext) {
	localctx = NewFileRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ApplicationParserRULE_fileRef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(ApplicationParserSTRING)
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
