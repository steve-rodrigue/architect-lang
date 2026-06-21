// Code generated from internal/infrastructure/antlr/grammars/Deployment.g4 by ANTLR 4.13.2. DO NOT EDIT.

package deployment // Deployment
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

type DeploymentParser struct {
	*antlr.BaseParser
}

var DeploymentParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func deploymentParserInit() {
	staticData := &DeploymentParserStaticData
	staticData.LiteralNames = []string{
		"", "'deployment'", "'vendor'", "'inventory'", "'service'", "'replicas'",
		"'domain'", "'host'", "'volume'", "'backup'", "'true'", "'false'", "'{'",
		"'}'",
	}
	staticData.SymbolicNames = []string{
		"", "DEPLOYMENT", "VENDOR", "INVENTORY", "SERVICE", "REPLICAS", "DOMAIN",
		"HOST", "VOLUME", "BACKUP", "TRUE", "FALSE", "LBRACE", "RBRACE", "IDENT",
		"INT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "deploymentDecl", "deploymentBody", "vendorDecl", "inventoryDecl",
		"serviceDeploymentDecl", "serviceDeploymentBody", "replicasDecl", "domainDecl",
		"hostDecl", "volumeDecl", "backupDecl", "booleanValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 87, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 34, 8, 1, 10, 1, 12, 1, 37, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		3, 2, 44, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 5, 5, 56, 8, 5, 10, 5, 12, 5, 59, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 68, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 1, 1, 0, 10,
		11, 81, 0, 26, 1, 0, 0, 0, 2, 29, 1, 0, 0, 0, 4, 43, 1, 0, 0, 0, 6, 45,
		1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 67, 1, 0, 0, 0,
		14, 69, 1, 0, 0, 0, 16, 72, 1, 0, 0, 0, 18, 75, 1, 0, 0, 0, 20, 78, 1,
		0, 0, 0, 22, 81, 1, 0, 0, 0, 24, 84, 1, 0, 0, 0, 26, 27, 3, 2, 1, 0, 27,
		28, 5, 0, 0, 1, 28, 1, 1, 0, 0, 0, 29, 30, 5, 1, 0, 0, 30, 31, 5, 14, 0,
		0, 31, 35, 5, 12, 0, 0, 32, 34, 3, 4, 2, 0, 33, 32, 1, 0, 0, 0, 34, 37,
		1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0,
		37, 35, 1, 0, 0, 0, 38, 39, 5, 13, 0, 0, 39, 3, 1, 0, 0, 0, 40, 44, 3,
		6, 3, 0, 41, 44, 3, 8, 4, 0, 42, 44, 3, 10, 5, 0, 43, 40, 1, 0, 0, 0, 43,
		41, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 5, 1, 0, 0, 0, 45, 46, 5, 2, 0,
		0, 46, 47, 5, 14, 0, 0, 47, 7, 1, 0, 0, 0, 48, 49, 5, 3, 0, 0, 49, 50,
		5, 16, 0, 0, 50, 9, 1, 0, 0, 0, 51, 52, 5, 4, 0, 0, 52, 53, 5, 14, 0, 0,
		53, 57, 5, 12, 0, 0, 54, 56, 3, 12, 6, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1,
		0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59,
		57, 1, 0, 0, 0, 60, 61, 5, 13, 0, 0, 61, 11, 1, 0, 0, 0, 62, 68, 3, 14,
		7, 0, 63, 68, 3, 16, 8, 0, 64, 68, 3, 18, 9, 0, 65, 68, 3, 20, 10, 0, 66,
		68, 3, 22, 11, 0, 67, 62, 1, 0, 0, 0, 67, 63, 1, 0, 0, 0, 67, 64, 1, 0,
		0, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 13, 1, 0, 0, 0, 69, 70,
		5, 5, 0, 0, 70, 71, 5, 15, 0, 0, 71, 15, 1, 0, 0, 0, 72, 73, 5, 6, 0, 0,
		73, 74, 5, 16, 0, 0, 74, 17, 1, 0, 0, 0, 75, 76, 5, 7, 0, 0, 76, 77, 5,
		16, 0, 0, 77, 19, 1, 0, 0, 0, 78, 79, 5, 8, 0, 0, 79, 80, 5, 16, 0, 0,
		80, 21, 1, 0, 0, 0, 81, 82, 5, 9, 0, 0, 82, 83, 3, 24, 12, 0, 83, 23, 1,
		0, 0, 0, 84, 85, 7, 0, 0, 0, 85, 25, 1, 0, 0, 0, 4, 35, 43, 57, 67,
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

// DeploymentParserInit initializes any static state used to implement DeploymentParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDeploymentParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DeploymentParserInit() {
	staticData := &DeploymentParserStaticData
	staticData.once.Do(deploymentParserInit)
}

// NewDeploymentParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDeploymentParser(input antlr.TokenStream) *DeploymentParser {
	DeploymentParserInit()
	this := new(DeploymentParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DeploymentParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Deployment.g4"

	return this
}

// DeploymentParser tokens.
const (
	DeploymentParserEOF           = antlr.TokenEOF
	DeploymentParserDEPLOYMENT    = 1
	DeploymentParserVENDOR        = 2
	DeploymentParserINVENTORY     = 3
	DeploymentParserSERVICE       = 4
	DeploymentParserREPLICAS      = 5
	DeploymentParserDOMAIN        = 6
	DeploymentParserHOST          = 7
	DeploymentParserVOLUME        = 8
	DeploymentParserBACKUP        = 9
	DeploymentParserTRUE          = 10
	DeploymentParserFALSE         = 11
	DeploymentParserLBRACE        = 12
	DeploymentParserRBRACE        = 13
	DeploymentParserIDENT         = 14
	DeploymentParserINT           = 15
	DeploymentParserSTRING        = 16
	DeploymentParserLINE_COMMENT  = 17
	DeploymentParserBLOCK_COMMENT = 18
	DeploymentParserWS            = 19
)

// DeploymentParser rules.
const (
	DeploymentParserRULE_program               = 0
	DeploymentParserRULE_deploymentDecl        = 1
	DeploymentParserRULE_deploymentBody        = 2
	DeploymentParserRULE_vendorDecl            = 3
	DeploymentParserRULE_inventoryDecl         = 4
	DeploymentParserRULE_serviceDeploymentDecl = 5
	DeploymentParserRULE_serviceDeploymentBody = 6
	DeploymentParserRULE_replicasDecl          = 7
	DeploymentParserRULE_domainDecl            = 8
	DeploymentParserRULE_hostDecl              = 9
	DeploymentParserRULE_volumeDecl            = 10
	DeploymentParserRULE_backupDecl            = 11
	DeploymentParserRULE_booleanValue          = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeploymentDecl() IDeploymentDeclContext
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
	p.RuleIndex = DeploymentParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) DeploymentDecl() IDeploymentDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeploymentDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeploymentDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(DeploymentParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DeploymentParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.DeploymentDecl()
	}
	{
		p.SetState(27)
		p.Match(DeploymentParserEOF)
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

// IDeploymentDeclContext is an interface to support dynamic dispatch.
type IDeploymentDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEPLOYMENT() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDeploymentBody() []IDeploymentBodyContext
	DeploymentBody(i int) IDeploymentBodyContext

	// IsDeploymentDeclContext differentiates from other interfaces.
	IsDeploymentDeclContext()
}

type DeploymentDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeploymentDeclContext() *DeploymentDeclContext {
	var p = new(DeploymentDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_deploymentDecl
	return p
}

func InitEmptyDeploymentDeclContext(p *DeploymentDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_deploymentDecl
}

func (*DeploymentDeclContext) IsDeploymentDeclContext() {}

func NewDeploymentDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeploymentDeclContext {
	var p = new(DeploymentDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_deploymentDecl

	return p
}

func (s *DeploymentDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeploymentDeclContext) DEPLOYMENT() antlr.TerminalNode {
	return s.GetToken(DeploymentParserDEPLOYMENT, 0)
}

func (s *DeploymentDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DeploymentParserIDENT, 0)
}

func (s *DeploymentDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DeploymentParserLBRACE, 0)
}

func (s *DeploymentDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DeploymentParserRBRACE, 0)
}

func (s *DeploymentDeclContext) AllDeploymentBody() []IDeploymentBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeploymentBodyContext); ok {
			len++
		}
	}

	tst := make([]IDeploymentBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeploymentBodyContext); ok {
			tst[i] = t.(IDeploymentBodyContext)
			i++
		}
	}

	return tst
}

func (s *DeploymentDeclContext) DeploymentBody(i int) IDeploymentBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeploymentBodyContext); ok {
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

	return t.(IDeploymentBodyContext)
}

func (s *DeploymentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeploymentDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeploymentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitDeploymentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) DeploymentDecl() (localctx IDeploymentDeclContext) {
	localctx = NewDeploymentDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DeploymentParserRULE_deploymentDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(DeploymentParserDEPLOYMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Match(DeploymentParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(DeploymentParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28) != 0 {
		{
			p.SetState(32)
			p.DeploymentBody()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(DeploymentParserRBRACE)
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

// IDeploymentBodyContext is an interface to support dynamic dispatch.
type IDeploymentBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VendorDecl() IVendorDeclContext
	InventoryDecl() IInventoryDeclContext
	ServiceDeploymentDecl() IServiceDeploymentDeclContext

	// IsDeploymentBodyContext differentiates from other interfaces.
	IsDeploymentBodyContext()
}

type DeploymentBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeploymentBodyContext() *DeploymentBodyContext {
	var p = new(DeploymentBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_deploymentBody
	return p
}

func InitEmptyDeploymentBodyContext(p *DeploymentBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_deploymentBody
}

func (*DeploymentBodyContext) IsDeploymentBodyContext() {}

func NewDeploymentBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeploymentBodyContext {
	var p = new(DeploymentBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_deploymentBody

	return p
}

func (s *DeploymentBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *DeploymentBodyContext) VendorDecl() IVendorDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVendorDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVendorDeclContext)
}

func (s *DeploymentBodyContext) InventoryDecl() IInventoryDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInventoryDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInventoryDeclContext)
}

func (s *DeploymentBodyContext) ServiceDeploymentDecl() IServiceDeploymentDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceDeploymentDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceDeploymentDeclContext)
}

func (s *DeploymentBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeploymentBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeploymentBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitDeploymentBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) DeploymentBody() (localctx IDeploymentBodyContext) {
	localctx = NewDeploymentBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DeploymentParserRULE_deploymentBody)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DeploymentParserVENDOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.VendorDecl()
		}

	case DeploymentParserINVENTORY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.InventoryDecl()
		}

	case DeploymentParserSERVICE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.ServiceDeploymentDecl()
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

// IVendorDeclContext is an interface to support dynamic dispatch.
type IVendorDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VENDOR() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsVendorDeclContext differentiates from other interfaces.
	IsVendorDeclContext()
}

type VendorDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVendorDeclContext() *VendorDeclContext {
	var p = new(VendorDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_vendorDecl
	return p
}

func InitEmptyVendorDeclContext(p *VendorDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_vendorDecl
}

func (*VendorDeclContext) IsVendorDeclContext() {}

func NewVendorDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VendorDeclContext {
	var p = new(VendorDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_vendorDecl

	return p
}

func (s *VendorDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VendorDeclContext) VENDOR() antlr.TerminalNode {
	return s.GetToken(DeploymentParserVENDOR, 0)
}

func (s *VendorDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DeploymentParserIDENT, 0)
}

func (s *VendorDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VendorDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VendorDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitVendorDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) VendorDecl() (localctx IVendorDeclContext) {
	localctx = NewVendorDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DeploymentParserRULE_vendorDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(DeploymentParserVENDOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(DeploymentParserIDENT)
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

// IInventoryDeclContext is an interface to support dynamic dispatch.
type IInventoryDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INVENTORY() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsInventoryDeclContext differentiates from other interfaces.
	IsInventoryDeclContext()
}

type InventoryDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInventoryDeclContext() *InventoryDeclContext {
	var p = new(InventoryDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_inventoryDecl
	return p
}

func InitEmptyInventoryDeclContext(p *InventoryDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_inventoryDecl
}

func (*InventoryDeclContext) IsInventoryDeclContext() {}

func NewInventoryDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InventoryDeclContext {
	var p = new(InventoryDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_inventoryDecl

	return p
}

func (s *InventoryDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *InventoryDeclContext) INVENTORY() antlr.TerminalNode {
	return s.GetToken(DeploymentParserINVENTORY, 0)
}

func (s *InventoryDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(DeploymentParserSTRING, 0)
}

func (s *InventoryDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InventoryDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InventoryDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitInventoryDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) InventoryDecl() (localctx IInventoryDeclContext) {
	localctx = NewInventoryDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DeploymentParserRULE_inventoryDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(DeploymentParserINVENTORY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(DeploymentParserSTRING)
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

// IServiceDeploymentDeclContext is an interface to support dynamic dispatch.
type IServiceDeploymentDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllServiceDeploymentBody() []IServiceDeploymentBodyContext
	ServiceDeploymentBody(i int) IServiceDeploymentBodyContext

	// IsServiceDeploymentDeclContext differentiates from other interfaces.
	IsServiceDeploymentDeclContext()
}

type ServiceDeploymentDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDeploymentDeclContext() *ServiceDeploymentDeclContext {
	var p = new(ServiceDeploymentDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_serviceDeploymentDecl
	return p
}

func InitEmptyServiceDeploymentDeclContext(p *ServiceDeploymentDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_serviceDeploymentDecl
}

func (*ServiceDeploymentDeclContext) IsServiceDeploymentDeclContext() {}

func NewServiceDeploymentDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDeploymentDeclContext {
	var p = new(ServiceDeploymentDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_serviceDeploymentDecl

	return p
}

func (s *ServiceDeploymentDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDeploymentDeclContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(DeploymentParserSERVICE, 0)
}

func (s *ServiceDeploymentDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DeploymentParserIDENT, 0)
}

func (s *ServiceDeploymentDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DeploymentParserLBRACE, 0)
}

func (s *ServiceDeploymentDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DeploymentParserRBRACE, 0)
}

func (s *ServiceDeploymentDeclContext) AllServiceDeploymentBody() []IServiceDeploymentBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceDeploymentBodyContext); ok {
			len++
		}
	}

	tst := make([]IServiceDeploymentBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceDeploymentBodyContext); ok {
			tst[i] = t.(IServiceDeploymentBodyContext)
			i++
		}
	}

	return tst
}

func (s *ServiceDeploymentDeclContext) ServiceDeploymentBody(i int) IServiceDeploymentBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceDeploymentBodyContext); ok {
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

	return t.(IServiceDeploymentBodyContext)
}

func (s *ServiceDeploymentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDeploymentDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDeploymentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitServiceDeploymentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) ServiceDeploymentDecl() (localctx IServiceDeploymentDeclContext) {
	localctx = NewServiceDeploymentDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DeploymentParserRULE_serviceDeploymentDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(DeploymentParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(DeploymentParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(DeploymentParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&992) != 0 {
		{
			p.SetState(54)
			p.ServiceDeploymentBody()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(DeploymentParserRBRACE)
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

// IServiceDeploymentBodyContext is an interface to support dynamic dispatch.
type IServiceDeploymentBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ReplicasDecl() IReplicasDeclContext
	DomainDecl() IDomainDeclContext
	HostDecl() IHostDeclContext
	VolumeDecl() IVolumeDeclContext
	BackupDecl() IBackupDeclContext

	// IsServiceDeploymentBodyContext differentiates from other interfaces.
	IsServiceDeploymentBodyContext()
}

type ServiceDeploymentBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDeploymentBodyContext() *ServiceDeploymentBodyContext {
	var p = new(ServiceDeploymentBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_serviceDeploymentBody
	return p
}

func InitEmptyServiceDeploymentBodyContext(p *ServiceDeploymentBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_serviceDeploymentBody
}

func (*ServiceDeploymentBodyContext) IsServiceDeploymentBodyContext() {}

func NewServiceDeploymentBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDeploymentBodyContext {
	var p = new(ServiceDeploymentBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_serviceDeploymentBody

	return p
}

func (s *ServiceDeploymentBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDeploymentBodyContext) ReplicasDecl() IReplicasDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReplicasDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReplicasDeclContext)
}

func (s *ServiceDeploymentBodyContext) DomainDecl() IDomainDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDomainDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDomainDeclContext)
}

func (s *ServiceDeploymentBodyContext) HostDecl() IHostDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHostDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHostDeclContext)
}

func (s *ServiceDeploymentBodyContext) VolumeDecl() IVolumeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVolumeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVolumeDeclContext)
}

func (s *ServiceDeploymentBodyContext) BackupDecl() IBackupDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBackupDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBackupDeclContext)
}

func (s *ServiceDeploymentBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDeploymentBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDeploymentBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitServiceDeploymentBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) ServiceDeploymentBody() (localctx IServiceDeploymentBodyContext) {
	localctx = NewServiceDeploymentBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DeploymentParserRULE_serviceDeploymentBody)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DeploymentParserREPLICAS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.ReplicasDecl()
		}

	case DeploymentParserDOMAIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.DomainDecl()
		}

	case DeploymentParserHOST:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.HostDecl()
		}

	case DeploymentParserVOLUME:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.VolumeDecl()
		}

	case DeploymentParserBACKUP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(66)
			p.BackupDecl()
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

// IReplicasDeclContext is an interface to support dynamic dispatch.
type IReplicasDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REPLICAS() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsReplicasDeclContext differentiates from other interfaces.
	IsReplicasDeclContext()
}

type ReplicasDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplicasDeclContext() *ReplicasDeclContext {
	var p = new(ReplicasDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_replicasDecl
	return p
}

func InitEmptyReplicasDeclContext(p *ReplicasDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_replicasDecl
}

func (*ReplicasDeclContext) IsReplicasDeclContext() {}

func NewReplicasDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplicasDeclContext {
	var p = new(ReplicasDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_replicasDecl

	return p
}

func (s *ReplicasDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplicasDeclContext) REPLICAS() antlr.TerminalNode {
	return s.GetToken(DeploymentParserREPLICAS, 0)
}

func (s *ReplicasDeclContext) INT() antlr.TerminalNode {
	return s.GetToken(DeploymentParserINT, 0)
}

func (s *ReplicasDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplicasDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplicasDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitReplicasDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) ReplicasDecl() (localctx IReplicasDeclContext) {
	localctx = NewReplicasDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DeploymentParserRULE_replicasDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(DeploymentParserREPLICAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(DeploymentParserINT)
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

// IDomainDeclContext is an interface to support dynamic dispatch.
type IDomainDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOMAIN() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsDomainDeclContext differentiates from other interfaces.
	IsDomainDeclContext()
}

type DomainDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainDeclContext() *DomainDeclContext {
	var p = new(DomainDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_domainDecl
	return p
}

func InitEmptyDomainDeclContext(p *DomainDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_domainDecl
}

func (*DomainDeclContext) IsDomainDeclContext() {}

func NewDomainDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainDeclContext {
	var p = new(DomainDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_domainDecl

	return p
}

func (s *DomainDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainDeclContext) DOMAIN() antlr.TerminalNode {
	return s.GetToken(DeploymentParserDOMAIN, 0)
}

func (s *DomainDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(DeploymentParserSTRING, 0)
}

func (s *DomainDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitDomainDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) DomainDecl() (localctx IDomainDeclContext) {
	localctx = NewDomainDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DeploymentParserRULE_domainDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(DeploymentParserDOMAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(DeploymentParserSTRING)
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

// IHostDeclContext is an interface to support dynamic dispatch.
type IHostDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HOST() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsHostDeclContext differentiates from other interfaces.
	IsHostDeclContext()
}

type HostDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostDeclContext() *HostDeclContext {
	var p = new(HostDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_hostDecl
	return p
}

func InitEmptyHostDeclContext(p *HostDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_hostDecl
}

func (*HostDeclContext) IsHostDeclContext() {}

func NewHostDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostDeclContext {
	var p = new(HostDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_hostDecl

	return p
}

func (s *HostDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *HostDeclContext) HOST() antlr.TerminalNode {
	return s.GetToken(DeploymentParserHOST, 0)
}

func (s *HostDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(DeploymentParserSTRING, 0)
}

func (s *HostDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitHostDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) HostDecl() (localctx IHostDeclContext) {
	localctx = NewHostDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DeploymentParserRULE_hostDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(DeploymentParserHOST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(DeploymentParserSTRING)
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

// IVolumeDeclContext is an interface to support dynamic dispatch.
type IVolumeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VOLUME() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsVolumeDeclContext differentiates from other interfaces.
	IsVolumeDeclContext()
}

type VolumeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVolumeDeclContext() *VolumeDeclContext {
	var p = new(VolumeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_volumeDecl
	return p
}

func InitEmptyVolumeDeclContext(p *VolumeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_volumeDecl
}

func (*VolumeDeclContext) IsVolumeDeclContext() {}

func NewVolumeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VolumeDeclContext {
	var p = new(VolumeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_volumeDecl

	return p
}

func (s *VolumeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VolumeDeclContext) VOLUME() antlr.TerminalNode {
	return s.GetToken(DeploymentParserVOLUME, 0)
}

func (s *VolumeDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(DeploymentParserSTRING, 0)
}

func (s *VolumeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VolumeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VolumeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitVolumeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) VolumeDecl() (localctx IVolumeDeclContext) {
	localctx = NewVolumeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DeploymentParserRULE_volumeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(DeploymentParserVOLUME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(DeploymentParserSTRING)
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

// IBackupDeclContext is an interface to support dynamic dispatch.
type IBackupDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BACKUP() antlr.TerminalNode
	BooleanValue() IBooleanValueContext

	// IsBackupDeclContext differentiates from other interfaces.
	IsBackupDeclContext()
}

type BackupDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBackupDeclContext() *BackupDeclContext {
	var p = new(BackupDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_backupDecl
	return p
}

func InitEmptyBackupDeclContext(p *BackupDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_backupDecl
}

func (*BackupDeclContext) IsBackupDeclContext() {}

func NewBackupDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BackupDeclContext {
	var p = new(BackupDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_backupDecl

	return p
}

func (s *BackupDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *BackupDeclContext) BACKUP() antlr.TerminalNode {
	return s.GetToken(DeploymentParserBACKUP, 0)
}

func (s *BackupDeclContext) BooleanValue() IBooleanValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *BackupDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BackupDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BackupDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitBackupDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) BackupDecl() (localctx IBackupDeclContext) {
	localctx = NewBackupDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DeploymentParserRULE_backupDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(DeploymentParserBACKUP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.BooleanValue()
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

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_booleanValue
	return p
}

func InitEmptyBooleanValueContext(p *BooleanValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DeploymentParserRULE_booleanValue
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeploymentParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(DeploymentParserTRUE, 0)
}

func (s *BooleanValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(DeploymentParserFALSE, 0)
}

func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DeploymentVisitor:
		return t.VisitBooleanValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DeploymentParser) BooleanValue() (localctx IBooleanValueContext) {
	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DeploymentParserRULE_booleanValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DeploymentParserTRUE || _la == DeploymentParserFALSE) {
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
