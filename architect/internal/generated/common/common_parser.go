// Code generated from internal/infrastructure/antlr/grammars/Common.g4 by ANTLR 4.13.2. DO NOT EDIT.

package object // Common
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

type CommonParser struct {
	*antlr.BaseParser
}

var CommonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func commonParserInit() {
	staticData := &CommonParserStaticData
	staticData.LiteralNames = []string{
		"", "'List'", "'input'", "'event'", "'result'", "'true'", "'false'",
		"'{'", "'}'", "'('", "')'", "'['", "']'", "'<'", "'>'", "','", "'|'",
		"'+'", "'*'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "LIST", "INPUT", "EVENT", "RESULT", "TRUE", "FALSE", "LBRACE", "RBRACE",
		"LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LT", "GT", "COMMA", "PIPE",
		"PLUS", "STAR", "QUESTION", "IDENT", "FLOAT", "INT", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"typeRef", "typeName", "numberConstraint", "optionalMarker", "value",
		"numberValue", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 54, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 3, 0, 17, 8, 0, 1, 0, 3, 0, 20,
		8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 27, 8, 0, 3, 0, 29, 8, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 2, 1, 2, 3, 2, 41,
		8, 2, 1, 2, 3, 2, 44, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 3, 2, 0, 5, 6, 21, 23, 1, 0,
		21, 22, 2, 0, 2, 4, 20, 20, 54, 0, 28, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4,
		43, 1, 0, 0, 0, 6, 45, 1, 0, 0, 0, 8, 47, 1, 0, 0, 0, 10, 49, 1, 0, 0,
		0, 12, 51, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 17, 3, 4, 2, 0, 16, 15,
		1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 19, 1, 0, 0, 0, 18, 20, 3, 6, 3, 0,
		19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 29, 1, 0, 0, 0, 21, 22, 5,
		1, 0, 0, 22, 23, 5, 13, 0, 0, 23, 24, 3, 2, 1, 0, 24, 26, 5, 14, 0, 0,
		25, 27, 3, 6, 3, 0, 26, 25, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 29, 1,
		0, 0, 0, 28, 14, 1, 0, 0, 0, 28, 21, 1, 0, 0, 0, 29, 1, 1, 0, 0, 0, 30,
		31, 5, 20, 0, 0, 31, 3, 1, 0, 0, 0, 32, 44, 5, 17, 0, 0, 33, 44, 5, 18,
		0, 0, 34, 36, 5, 11, 0, 0, 35, 37, 3, 10, 5, 0, 36, 35, 1, 0, 0, 0, 36,
		37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 5, 15, 0, 0, 39, 41, 3, 10,
		5, 0, 40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44,
		5, 12, 0, 0, 43, 32, 1, 0, 0, 0, 43, 33, 1, 0, 0, 0, 43, 34, 1, 0, 0, 0,
		44, 5, 1, 0, 0, 0, 45, 46, 5, 19, 0, 0, 46, 7, 1, 0, 0, 0, 47, 48, 7, 0,
		0, 0, 48, 9, 1, 0, 0, 0, 49, 50, 7, 1, 0, 0, 50, 11, 1, 0, 0, 0, 51, 52,
		7, 2, 0, 0, 52, 13, 1, 0, 0, 0, 7, 16, 19, 26, 28, 36, 40, 43,
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

// CommonParserInit initializes any static state used to implement CommonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCommonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CommonParserInit() {
	staticData := &CommonParserStaticData
	staticData.once.Do(commonParserInit)
}

// NewCommonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCommonParser(input antlr.TokenStream) *CommonParser {
	CommonParserInit()
	this := new(CommonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CommonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Common.g4"

	return this
}

// CommonParser tokens.
const (
	CommonParserEOF           = antlr.TokenEOF
	CommonParserLIST          = 1
	CommonParserINPUT         = 2
	CommonParserEVENT         = 3
	CommonParserRESULT        = 4
	CommonParserTRUE          = 5
	CommonParserFALSE         = 6
	CommonParserLBRACE        = 7
	CommonParserRBRACE        = 8
	CommonParserLPAREN        = 9
	CommonParserRPAREN        = 10
	CommonParserLBRACKET      = 11
	CommonParserRBRACKET      = 12
	CommonParserLT            = 13
	CommonParserGT            = 14
	CommonParserCOMMA         = 15
	CommonParserPIPE          = 16
	CommonParserPLUS          = 17
	CommonParserSTAR          = 18
	CommonParserQUESTION      = 19
	CommonParserIDENT         = 20
	CommonParserFLOAT         = 21
	CommonParserINT           = 22
	CommonParserSTRING        = 23
	CommonParserLINE_COMMENT  = 24
	CommonParserBLOCK_COMMENT = 25
	CommonParserWS            = 26
)

// CommonParser rules.
const (
	CommonParserRULE_typeRef          = 0
	CommonParserRULE_typeName         = 1
	CommonParserRULE_numberConstraint = 2
	CommonParserRULE_optionalMarker   = 3
	CommonParserRULE_value            = 4
	CommonParserRULE_numberValue      = 5
	CommonParserRULE_identifier       = 6
)

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
	p.RuleIndex = CommonParserRULE_typeRef
	return p
}

func InitEmptyTypeRefContext(p *TypeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommonParserRULE_typeRef
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommonParserRULE_typeRef

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
	return s.GetToken(CommonParserLIST, 0)
}

func (s *TypeRefContext) LT() antlr.TerminalNode {
	return s.GetToken(CommonParserLT, 0)
}

func (s *TypeRefContext) GT() antlr.TerminalNode {
	return s.GetToken(CommonParserGT, 0)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CommonVisitor:
		return t.VisitTypeRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CommonParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CommonParserRULE_typeRef)
	var _la int

	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CommonParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.TypeName()
		}
		p.SetState(16)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&395264) != 0 {
			{
				p.SetState(15)
				p.NumberConstraint()
			}

		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CommonParserQUESTION {
			{
				p.SetState(18)
				p.OptionalMarker()
			}

		}

	case CommonParserLIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Match(CommonParserLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.Match(CommonParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.TypeName()
		}
		{
			p.SetState(24)
			p.Match(CommonParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CommonParserQUESTION {
			{
				p.SetState(25)
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
	p.RuleIndex = CommonParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommonParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommonParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CommonParserIDENT, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CommonVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CommonParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CommonParserRULE_typeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(CommonParserIDENT)
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
	p.RuleIndex = CommonParserRULE_numberConstraint
	return p
}

func InitEmptyNumberConstraintContext(p *NumberConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommonParserRULE_numberConstraint
}

func (*NumberConstraintContext) IsNumberConstraintContext() {}

func NewNumberConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberConstraintContext {
	var p = new(NumberConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommonParserRULE_numberConstraint

	return p
}

func (s *NumberConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberConstraintContext) PLUS() antlr.TerminalNode {
	return s.GetToken(CommonParserPLUS, 0)
}

func (s *NumberConstraintContext) STAR() antlr.TerminalNode {
	return s.GetToken(CommonParserSTAR, 0)
}

func (s *NumberConstraintContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(CommonParserLBRACKET, 0)
}

func (s *NumberConstraintContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CommonParserCOMMA, 0)
}

func (s *NumberConstraintContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(CommonParserRBRACKET, 0)
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
	case CommonVisitor:
		return t.VisitNumberConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CommonParser) NumberConstraint() (localctx INumberConstraintContext) {
	localctx = NewNumberConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CommonParserRULE_numberConstraint)
	var _la int

	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CommonParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Match(CommonParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CommonParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Match(CommonParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CommonParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(CommonParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CommonParserFLOAT || _la == CommonParserINT {
			{
				p.SetState(35)
				p.NumberValue()
			}

		}
		{
			p.SetState(38)
			p.Match(CommonParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CommonParserFLOAT || _la == CommonParserINT {
			{
				p.SetState(39)
				p.NumberValue()
			}

		}
		{
			p.SetState(42)
			p.Match(CommonParserRBRACKET)
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
	p.RuleIndex = CommonParserRULE_optionalMarker
	return p
}

func InitEmptyOptionalMarkerContext(p *OptionalMarkerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommonParserRULE_optionalMarker
}

func (*OptionalMarkerContext) IsOptionalMarkerContext() {}

func NewOptionalMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalMarkerContext {
	var p = new(OptionalMarkerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommonParserRULE_optionalMarker

	return p
}

func (s *OptionalMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalMarkerContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(CommonParserQUESTION, 0)
}

func (s *OptionalMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalMarkerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CommonVisitor:
		return t.VisitOptionalMarker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CommonParser) OptionalMarker() (localctx IOptionalMarkerContext) {
	localctx = NewOptionalMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CommonParserRULE_optionalMarker)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(CommonParserQUESTION)
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
	p.RuleIndex = CommonParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommonParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommonParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(CommonParserSTRING, 0)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(CommonParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CommonParserFLOAT, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(CommonParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(CommonParserFALSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CommonVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CommonParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CommonParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14680160) != 0) {
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
	p.RuleIndex = CommonParserRULE_numberValue
	return p
}

func InitEmptyNumberValueContext(p *NumberValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommonParserRULE_numberValue
}

func (*NumberValueContext) IsNumberValueContext() {}

func NewNumberValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberValueContext {
	var p = new(NumberValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommonParserRULE_numberValue

	return p
}

func (s *NumberValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberValueContext) INT() antlr.TerminalNode {
	return s.GetToken(CommonParserINT, 0)
}

func (s *NumberValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CommonParserFLOAT, 0)
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CommonVisitor:
		return t.VisitNumberValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CommonParser) NumberValue() (localctx INumberValueContext) {
	localctx = NewNumberValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CommonParserRULE_numberValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CommonParserFLOAT || _la == CommonParserINT) {
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
	p.RuleIndex = CommonParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommonParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommonParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(CommonParserIDENT, 0)
}

func (s *IdentifierContext) INPUT() antlr.TerminalNode {
	return s.GetToken(CommonParserINPUT, 0)
}

func (s *IdentifierContext) EVENT() antlr.TerminalNode {
	return s.GetToken(CommonParserEVENT, 0)
}

func (s *IdentifierContext) RESULT() antlr.TerminalNode {
	return s.GetToken(CommonParserRESULT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CommonVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CommonParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CommonParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1048604) != 0) {
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
