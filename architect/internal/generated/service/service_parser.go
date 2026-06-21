// Code generated from internal/infrastructure/antlr/grammars/Service.g4 by ANTLR 4.13.2. DO NOT EDIT.

package service // Service
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

type ServiceParser struct {
	*antlr.BaseParser
}

var ServiceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func serviceParserInit() {
	staticData := &ServiceParserStaticData
	staticData.LiteralNames = []string{
		"", "'service'", "'version'", "'image'", "'exposes'", "'depends_on'",
		"'application'", "'stores'", "'event'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "SERVICE", "VERSION", "IMAGE", "EXPOSES", "DEPENDS_ON", "APPLICATION",
		"STORES", "EVENT", "LBRACE", "RBRACE", "IDENT", "INT", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "serviceDecl", "serviceKind", "serviceBody", "versionDecl",
		"imageDecl", "exposeDecl", "dependsOnDecl", "applicationDecl", "storeDecl",
		"eventDecl",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 70, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 31, 8, 1, 10,
		1, 12, 1, 34, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 47, 8, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0,
		0, 65, 0, 22, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 46,
		1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 54, 1, 0, 0, 0,
		14, 57, 1, 0, 0, 0, 16, 60, 1, 0, 0, 0, 18, 63, 1, 0, 0, 0, 20, 66, 1,
		0, 0, 0, 22, 23, 3, 2, 1, 0, 23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25,
		26, 5, 1, 0, 0, 26, 27, 5, 11, 0, 0, 27, 28, 3, 4, 2, 0, 28, 32, 5, 9,
		0, 0, 29, 31, 3, 6, 3, 0, 30, 29, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30,
		1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 35, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0,
		35, 36, 5, 10, 0, 0, 36, 3, 1, 0, 0, 0, 37, 38, 5, 11, 0, 0, 38, 5, 1,
		0, 0, 0, 39, 47, 3, 8, 4, 0, 40, 47, 3, 10, 5, 0, 41, 47, 3, 12, 6, 0,
		42, 47, 3, 14, 7, 0, 43, 47, 3, 16, 8, 0, 44, 47, 3, 18, 9, 0, 45, 47,
		3, 20, 10, 0, 46, 39, 1, 0, 0, 0, 46, 40, 1, 0, 0, 0, 46, 41, 1, 0, 0,
		0, 46, 42, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 45,
		1, 0, 0, 0, 47, 7, 1, 0, 0, 0, 48, 49, 5, 2, 0, 0, 49, 50, 5, 13, 0, 0,
		50, 9, 1, 0, 0, 0, 51, 52, 5, 3, 0, 0, 52, 53, 5, 13, 0, 0, 53, 11, 1,
		0, 0, 0, 54, 55, 5, 4, 0, 0, 55, 56, 5, 12, 0, 0, 56, 13, 1, 0, 0, 0, 57,
		58, 5, 5, 0, 0, 58, 59, 5, 11, 0, 0, 59, 15, 1, 0, 0, 0, 60, 61, 5, 6,
		0, 0, 61, 62, 5, 13, 0, 0, 62, 17, 1, 0, 0, 0, 63, 64, 5, 7, 0, 0, 64,
		65, 5, 11, 0, 0, 65, 19, 1, 0, 0, 0, 66, 67, 5, 8, 0, 0, 67, 68, 5, 11,
		0, 0, 68, 21, 1, 0, 0, 0, 2, 32, 46,
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

// ServiceParserInit initializes any static state used to implement ServiceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewServiceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ServiceParserInit() {
	staticData := &ServiceParserStaticData
	staticData.once.Do(serviceParserInit)
}

// NewServiceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewServiceParser(input antlr.TokenStream) *ServiceParser {
	ServiceParserInit()
	this := new(ServiceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ServiceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Service.g4"

	return this
}

// ServiceParser tokens.
const (
	ServiceParserEOF           = antlr.TokenEOF
	ServiceParserSERVICE       = 1
	ServiceParserVERSION       = 2
	ServiceParserIMAGE         = 3
	ServiceParserEXPOSES       = 4
	ServiceParserDEPENDS_ON    = 5
	ServiceParserAPPLICATION   = 6
	ServiceParserSTORES        = 7
	ServiceParserEVENT         = 8
	ServiceParserLBRACE        = 9
	ServiceParserRBRACE        = 10
	ServiceParserIDENT         = 11
	ServiceParserINT           = 12
	ServiceParserSTRING        = 13
	ServiceParserLINE_COMMENT  = 14
	ServiceParserBLOCK_COMMENT = 15
	ServiceParserWS            = 16
)

// ServiceParser rules.
const (
	ServiceParserRULE_program         = 0
	ServiceParserRULE_serviceDecl     = 1
	ServiceParserRULE_serviceKind     = 2
	ServiceParserRULE_serviceBody     = 3
	ServiceParserRULE_versionDecl     = 4
	ServiceParserRULE_imageDecl       = 5
	ServiceParserRULE_exposeDecl      = 6
	ServiceParserRULE_dependsOnDecl   = 7
	ServiceParserRULE_applicationDecl = 8
	ServiceParserRULE_storeDecl       = 9
	ServiceParserRULE_eventDecl       = 10
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ServiceDecl() IServiceDeclContext
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
	p.RuleIndex = ServiceParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ServiceDecl() IServiceDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ServiceParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ServiceParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.ServiceDecl()
	}
	{
		p.SetState(23)
		p.Match(ServiceParserEOF)
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

// IServiceDeclContext is an interface to support dynamic dispatch.
type IServiceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ServiceKind() IServiceKindContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllServiceBody() []IServiceBodyContext
	ServiceBody(i int) IServiceBodyContext

	// IsServiceDeclContext differentiates from other interfaces.
	IsServiceDeclContext()
}

type ServiceDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDeclContext() *ServiceDeclContext {
	var p = new(ServiceDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_serviceDecl
	return p
}

func InitEmptyServiceDeclContext(p *ServiceDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_serviceDecl
}

func (*ServiceDeclContext) IsServiceDeclContext() {}

func NewServiceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDeclContext {
	var p = new(ServiceDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_serviceDecl

	return p
}

func (s *ServiceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDeclContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(ServiceParserSERVICE, 0)
}

func (s *ServiceDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENT, 0)
}

func (s *ServiceDeclContext) ServiceKind() IServiceKindContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceKindContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceKindContext)
}

func (s *ServiceDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ServiceParserLBRACE, 0)
}

func (s *ServiceDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ServiceParserRBRACE, 0)
}

func (s *ServiceDeclContext) AllServiceBody() []IServiceBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceBodyContext); ok {
			len++
		}
	}

	tst := make([]IServiceBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceBodyContext); ok {
			tst[i] = t.(IServiceBodyContext)
			i++
		}
	}

	return tst
}

func (s *ServiceDeclContext) ServiceBody(i int) IServiceBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceBodyContext); ok {
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

	return t.(IServiceBodyContext)
}

func (s *ServiceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitServiceDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) ServiceDecl() (localctx IServiceDeclContext) {
	localctx = NewServiceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ServiceParserRULE_serviceDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(ServiceParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(ServiceParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.ServiceKind()
	}
	{
		p.SetState(28)
		p.Match(ServiceParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&508) != 0 {
		{
			p.SetState(29)
			p.ServiceBody()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Match(ServiceParserRBRACE)
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

// IServiceKindContext is an interface to support dynamic dispatch.
type IServiceKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsServiceKindContext differentiates from other interfaces.
	IsServiceKindContext()
}

type ServiceKindContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceKindContext() *ServiceKindContext {
	var p = new(ServiceKindContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_serviceKind
	return p
}

func InitEmptyServiceKindContext(p *ServiceKindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_serviceKind
}

func (*ServiceKindContext) IsServiceKindContext() {}

func NewServiceKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceKindContext {
	var p = new(ServiceKindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_serviceKind

	return p
}

func (s *ServiceKindContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceKindContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENT, 0)
}

func (s *ServiceKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceKindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitServiceKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) ServiceKind() (localctx IServiceKindContext) {
	localctx = NewServiceKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ServiceParserRULE_serviceKind)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(ServiceParserIDENT)
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

// IServiceBodyContext is an interface to support dynamic dispatch.
type IServiceBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VersionDecl() IVersionDeclContext
	ImageDecl() IImageDeclContext
	ExposeDecl() IExposeDeclContext
	DependsOnDecl() IDependsOnDeclContext
	ApplicationDecl() IApplicationDeclContext
	StoreDecl() IStoreDeclContext
	EventDecl() IEventDeclContext

	// IsServiceBodyContext differentiates from other interfaces.
	IsServiceBodyContext()
}

type ServiceBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceBodyContext() *ServiceBodyContext {
	var p = new(ServiceBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_serviceBody
	return p
}

func InitEmptyServiceBodyContext(p *ServiceBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_serviceBody
}

func (*ServiceBodyContext) IsServiceBodyContext() {}

func NewServiceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceBodyContext {
	var p = new(ServiceBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_serviceBody

	return p
}

func (s *ServiceBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceBodyContext) VersionDecl() IVersionDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionDeclContext)
}

func (s *ServiceBodyContext) ImageDecl() IImageDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImageDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImageDeclContext)
}

func (s *ServiceBodyContext) ExposeDecl() IExposeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExposeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExposeDeclContext)
}

func (s *ServiceBodyContext) DependsOnDecl() IDependsOnDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDependsOnDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDependsOnDeclContext)
}

func (s *ServiceBodyContext) ApplicationDecl() IApplicationDeclContext {
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

func (s *ServiceBodyContext) StoreDecl() IStoreDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStoreDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStoreDeclContext)
}

func (s *ServiceBodyContext) EventDecl() IEventDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventDeclContext)
}

func (s *ServiceBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitServiceBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) ServiceBody() (localctx IServiceBodyContext) {
	localctx = NewServiceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ServiceParserRULE_serviceBody)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ServiceParserVERSION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.VersionDecl()
		}

	case ServiceParserIMAGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.ImageDecl()
		}

	case ServiceParserEXPOSES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(41)
			p.ExposeDecl()
		}

	case ServiceParserDEPENDS_ON:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(42)
			p.DependsOnDecl()
		}

	case ServiceParserAPPLICATION:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.ApplicationDecl()
		}

	case ServiceParserSTORES:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(44)
			p.StoreDecl()
		}

	case ServiceParserEVENT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(45)
			p.EventDecl()
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

// IVersionDeclContext is an interface to support dynamic dispatch.
type IVersionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VERSION() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsVersionDeclContext differentiates from other interfaces.
	IsVersionDeclContext()
}

type VersionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionDeclContext() *VersionDeclContext {
	var p = new(VersionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_versionDecl
	return p
}

func InitEmptyVersionDeclContext(p *VersionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_versionDecl
}

func (*VersionDeclContext) IsVersionDeclContext() {}

func NewVersionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionDeclContext {
	var p = new(VersionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_versionDecl

	return p
}

func (s *VersionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionDeclContext) VERSION() antlr.TerminalNode {
	return s.GetToken(ServiceParserVERSION, 0)
}

func (s *VersionDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(ServiceParserSTRING, 0)
}

func (s *VersionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitVersionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) VersionDecl() (localctx IVersionDeclContext) {
	localctx = NewVersionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ServiceParserRULE_versionDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(ServiceParserVERSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(ServiceParserSTRING)
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

// IImageDeclContext is an interface to support dynamic dispatch.
type IImageDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMAGE() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsImageDeclContext differentiates from other interfaces.
	IsImageDeclContext()
}

type ImageDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImageDeclContext() *ImageDeclContext {
	var p = new(ImageDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_imageDecl
	return p
}

func InitEmptyImageDeclContext(p *ImageDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_imageDecl
}

func (*ImageDeclContext) IsImageDeclContext() {}

func NewImageDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImageDeclContext {
	var p = new(ImageDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_imageDecl

	return p
}

func (s *ImageDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImageDeclContext) IMAGE() antlr.TerminalNode {
	return s.GetToken(ServiceParserIMAGE, 0)
}

func (s *ImageDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(ServiceParserSTRING, 0)
}

func (s *ImageDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImageDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImageDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitImageDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) ImageDecl() (localctx IImageDeclContext) {
	localctx = NewImageDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ServiceParserRULE_imageDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(ServiceParserIMAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(ServiceParserSTRING)
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

// IExposeDeclContext is an interface to support dynamic dispatch.
type IExposeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXPOSES() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsExposeDeclContext differentiates from other interfaces.
	IsExposeDeclContext()
}

type ExposeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExposeDeclContext() *ExposeDeclContext {
	var p = new(ExposeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_exposeDecl
	return p
}

func InitEmptyExposeDeclContext(p *ExposeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_exposeDecl
}

func (*ExposeDeclContext) IsExposeDeclContext() {}

func NewExposeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExposeDeclContext {
	var p = new(ExposeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_exposeDecl

	return p
}

func (s *ExposeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExposeDeclContext) EXPOSES() antlr.TerminalNode {
	return s.GetToken(ServiceParserEXPOSES, 0)
}

func (s *ExposeDeclContext) INT() antlr.TerminalNode {
	return s.GetToken(ServiceParserINT, 0)
}

func (s *ExposeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExposeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExposeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitExposeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) ExposeDecl() (localctx IExposeDeclContext) {
	localctx = NewExposeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ServiceParserRULE_exposeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(ServiceParserEXPOSES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(ServiceParserINT)
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

// IDependsOnDeclContext is an interface to support dynamic dispatch.
type IDependsOnDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEPENDS_ON() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsDependsOnDeclContext differentiates from other interfaces.
	IsDependsOnDeclContext()
}

type DependsOnDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDependsOnDeclContext() *DependsOnDeclContext {
	var p = new(DependsOnDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_dependsOnDecl
	return p
}

func InitEmptyDependsOnDeclContext(p *DependsOnDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_dependsOnDecl
}

func (*DependsOnDeclContext) IsDependsOnDeclContext() {}

func NewDependsOnDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DependsOnDeclContext {
	var p = new(DependsOnDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_dependsOnDecl

	return p
}

func (s *DependsOnDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DependsOnDeclContext) DEPENDS_ON() antlr.TerminalNode {
	return s.GetToken(ServiceParserDEPENDS_ON, 0)
}

func (s *DependsOnDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENT, 0)
}

func (s *DependsOnDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DependsOnDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DependsOnDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitDependsOnDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) DependsOnDecl() (localctx IDependsOnDeclContext) {
	localctx = NewDependsOnDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ServiceParserRULE_dependsOnDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(ServiceParserDEPENDS_ON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(ServiceParserIDENT)
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
	STRING() antlr.TerminalNode

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
	p.RuleIndex = ServiceParserRULE_applicationDecl
	return p
}

func InitEmptyApplicationDeclContext(p *ApplicationDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_applicationDecl
}

func (*ApplicationDeclContext) IsApplicationDeclContext() {}

func NewApplicationDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplicationDeclContext {
	var p = new(ApplicationDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_applicationDecl

	return p
}

func (s *ApplicationDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplicationDeclContext) APPLICATION() antlr.TerminalNode {
	return s.GetToken(ServiceParserAPPLICATION, 0)
}

func (s *ApplicationDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(ServiceParserSTRING, 0)
}

func (s *ApplicationDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicationDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApplicationDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitApplicationDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) ApplicationDecl() (localctx IApplicationDeclContext) {
	localctx = NewApplicationDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ServiceParserRULE_applicationDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(ServiceParserAPPLICATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(ServiceParserSTRING)
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

// IStoreDeclContext is an interface to support dynamic dispatch.
type IStoreDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STORES() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsStoreDeclContext differentiates from other interfaces.
	IsStoreDeclContext()
}

type StoreDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStoreDeclContext() *StoreDeclContext {
	var p = new(StoreDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_storeDecl
	return p
}

func InitEmptyStoreDeclContext(p *StoreDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_storeDecl
}

func (*StoreDeclContext) IsStoreDeclContext() {}

func NewStoreDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreDeclContext {
	var p = new(StoreDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_storeDecl

	return p
}

func (s *StoreDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreDeclContext) STORES() antlr.TerminalNode {
	return s.GetToken(ServiceParserSTORES, 0)
}

func (s *StoreDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENT, 0)
}

func (s *StoreDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitStoreDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) StoreDecl() (localctx IStoreDeclContext) {
	localctx = NewStoreDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ServiceParserRULE_storeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(ServiceParserSTORES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(ServiceParserIDENT)
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

// IEventDeclContext is an interface to support dynamic dispatch.
type IEventDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EVENT() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsEventDeclContext differentiates from other interfaces.
	IsEventDeclContext()
}

type EventDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventDeclContext() *EventDeclContext {
	var p = new(EventDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_eventDecl
	return p
}

func InitEmptyEventDeclContext(p *EventDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ServiceParserRULE_eventDecl
}

func (*EventDeclContext) IsEventDeclContext() {}

func NewEventDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventDeclContext {
	var p = new(EventDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ServiceParserRULE_eventDecl

	return p
}

func (s *EventDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EventDeclContext) EVENT() antlr.TerminalNode {
	return s.GetToken(ServiceParserEVENT, 0)
}

func (s *EventDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ServiceParserIDENT, 0)
}

func (s *EventDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ServiceVisitor:
		return t.VisitEventDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ServiceParser) EventDecl() (localctx IEventDeclContext) {
	localctx = NewEventDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ServiceParserRULE_eventDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(ServiceParserEVENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(ServiceParserIDENT)
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
