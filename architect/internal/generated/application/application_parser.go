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
		"", "'application'", "'emits'", "'listens'", "'on'", "'endpoints'",
		"'consumers'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "APPLICATION", "EMITS", "LISTENS", "ON", "ENDPOINTS", "CONSUMERS",
		"LBRACE", "RBRACE", "IDENT", "INT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "applicationDecl", "applicationBody", "portDecl", "endpointsBlock",
		"consumersBlock", "fileRef",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 66, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 22, 8, 1, 10, 1, 12, 1, 25, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2,
		32, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 42, 8,
		3, 1, 4, 1, 4, 1, 4, 5, 4, 47, 8, 4, 10, 4, 12, 4, 50, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 5, 5, 57, 8, 5, 10, 5, 12, 5, 60, 9, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 64, 0, 14, 1, 0, 0,
		0, 2, 17, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 43, 1, 0,
		0, 0, 10, 53, 1, 0, 0, 0, 12, 63, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 16,
		5, 0, 0, 1, 16, 1, 1, 0, 0, 0, 17, 18, 5, 1, 0, 0, 18, 19, 5, 9, 0, 0,
		19, 23, 5, 7, 0, 0, 20, 22, 3, 4, 2, 0, 21, 20, 1, 0, 0, 0, 22, 25, 1,
		0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 26, 1, 0, 0, 0, 25,
		23, 1, 0, 0, 0, 26, 27, 5, 8, 0, 0, 27, 3, 1, 0, 0, 0, 28, 32, 3, 6, 3,
		0, 29, 32, 3, 8, 4, 0, 30, 32, 3, 10, 5, 0, 31, 28, 1, 0, 0, 0, 31, 29,
		1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 5, 1, 0, 0, 0, 33, 34, 5, 2, 0, 0,
		34, 35, 5, 9, 0, 0, 35, 36, 5, 4, 0, 0, 36, 42, 5, 10, 0, 0, 37, 38, 5,
		3, 0, 0, 38, 39, 5, 9, 0, 0, 39, 40, 5, 4, 0, 0, 40, 42, 5, 10, 0, 0, 41,
		33, 1, 0, 0, 0, 41, 37, 1, 0, 0, 0, 42, 7, 1, 0, 0, 0, 43, 44, 5, 5, 0,
		0, 44, 48, 5, 7, 0, 0, 45, 47, 3, 12, 6, 0, 46, 45, 1, 0, 0, 0, 47, 50,
		1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0,
		50, 48, 1, 0, 0, 0, 51, 52, 5, 8, 0, 0, 52, 9, 1, 0, 0, 0, 53, 54, 5, 6,
		0, 0, 54, 58, 5, 7, 0, 0, 55, 57, 3, 12, 6, 0, 56, 55, 1, 0, 0, 0, 57,
		60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 61, 1, 0, 0,
		0, 60, 58, 1, 0, 0, 0, 61, 62, 5, 8, 0, 0, 62, 11, 1, 0, 0, 0, 63, 64,
		5, 11, 0, 0, 64, 13, 1, 0, 0, 0, 5, 23, 31, 41, 48, 58,
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
	ApplicationParserENDPOINTS     = 5
	ApplicationParserCONSUMERS     = 6
	ApplicationParserLBRACE        = 7
	ApplicationParserRBRACE        = 8
	ApplicationParserIDENT         = 9
	ApplicationParserINT           = 10
	ApplicationParserSTRING        = 11
	ApplicationParserLINE_COMMENT  = 12
	ApplicationParserBLOCK_COMMENT = 13
	ApplicationParserWS            = 14
)

// ApplicationParser rules.
const (
	ApplicationParserRULE_program         = 0
	ApplicationParserRULE_applicationDecl = 1
	ApplicationParserRULE_applicationBody = 2
	ApplicationParserRULE_portDecl        = 3
	ApplicationParserRULE_endpointsBlock  = 4
	ApplicationParserRULE_consumersBlock  = 5
	ApplicationParserRULE_fileRef         = 6
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
		p.SetState(14)
		p.ApplicationDecl()
	}
	{
		p.SetState(15)
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
		p.SetState(17)
		p.Match(ApplicationParserAPPLICATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(18)
		p.Match(ApplicationParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Match(ApplicationParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&108) != 0 {
		{
			p.SetState(20)
			p.ApplicationBody()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApplicationParserEMITS, ApplicationParserLISTENS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.PortDecl()
		}

	case ApplicationParserENDPOINTS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.EndpointsBlock()
		}

	case ApplicationParserCONSUMERS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApplicationParserEMITS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(ApplicationParserEMITS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(ApplicationParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Match(ApplicationParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(ApplicationParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ApplicationParserLISTENS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(ApplicationParserLISTENS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(ApplicationParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(ApplicationParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
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
	p.EnterRule(localctx, 8, ApplicationParserRULE_endpointsBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(ApplicationParserENDPOINTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(ApplicationParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ApplicationParserSTRING {
		{
			p.SetState(45)
			p.FileRef()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
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
	p.EnterRule(localctx, 10, ApplicationParserRULE_consumersBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(ApplicationParserCONSUMERS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(ApplicationParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ApplicationParserSTRING {
		{
			p.SetState(55)
			p.FileRef()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
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
	p.EnterRule(localctx, 12, ApplicationParserRULE_fileRef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
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
