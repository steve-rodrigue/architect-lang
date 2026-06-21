// Code generated from internal/infrastructure/antlr/grammars/Project.g4 by ANTLR 4.13.2. DO NOT EDIT.

package project // Project
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

type ProjectParser struct {
	*antlr.BaseParser
}

var ProjectParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func projectParserInit() {
	staticData := &ProjectParserStaticData
	staticData.LiteralNames = []string{
		"", "'project'", "'version'", "'next_version'", "'services'", "'objects'",
		"'deployments'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "PROJECT", "VERSION", "NEXT_VERSION", "SERVICES", "OBJECTS", "DEPLOYMENTS",
		"LBRACE", "RBRACE", "IDENT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "projectDecl", "versionDecl", "versionBody", "nextVersionDecl",
		"servicesBlock", "objectsBlock", "deploymentsBlock", "fileRef",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 92, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 4, 1, 26, 8, 1, 11, 1, 12, 1, 27, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39, 9, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 47, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 53,
		8, 4, 10, 4, 12, 4, 56, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 63, 8,
		5, 10, 5, 12, 5, 66, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 73, 8, 6,
		10, 6, 12, 6, 76, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 83, 8, 7, 10,
		7, 12, 7, 86, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 0, 0, 91, 0, 18, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 31,
		1, 0, 0, 0, 6, 46, 1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 59, 1, 0, 0, 0, 12,
		69, 1, 0, 0, 0, 14, 79, 1, 0, 0, 0, 16, 89, 1, 0, 0, 0, 18, 19, 3, 2, 1,
		0, 19, 20, 5, 0, 0, 1, 20, 1, 1, 0, 0, 0, 21, 22, 5, 1, 0, 0, 22, 23, 5,
		9, 0, 0, 23, 25, 5, 7, 0, 0, 24, 26, 3, 4, 2, 0, 25, 24, 1, 0, 0, 0, 26,
		27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 1, 0, 0,
		0, 29, 30, 5, 8, 0, 0, 30, 3, 1, 0, 0, 0, 31, 32, 5, 2, 0, 0, 32, 33, 5,
		10, 0, 0, 33, 37, 5, 7, 0, 0, 34, 36, 3, 6, 3, 0, 35, 34, 1, 0, 0, 0, 36,
		39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0,
		0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 8, 0, 0, 41, 5, 1, 0, 0, 0, 42, 47, 3,
		10, 5, 0, 43, 47, 3, 12, 6, 0, 44, 47, 3, 14, 7, 0, 45, 47, 3, 8, 4, 0,
		46, 42, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 45, 1,
		0, 0, 0, 47, 7, 1, 0, 0, 0, 48, 49, 5, 3, 0, 0, 49, 50, 5, 10, 0, 0, 50,
		54, 5, 7, 0, 0, 51, 53, 3, 6, 3, 0, 52, 51, 1, 0, 0, 0, 53, 56, 1, 0, 0,
		0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 54,
		1, 0, 0, 0, 57, 58, 5, 8, 0, 0, 58, 9, 1, 0, 0, 0, 59, 60, 5, 4, 0, 0,
		60, 64, 5, 7, 0, 0, 61, 63, 3, 16, 8, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1,
		0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66,
		64, 1, 0, 0, 0, 67, 68, 5, 8, 0, 0, 68, 11, 1, 0, 0, 0, 69, 70, 5, 5, 0,
		0, 70, 74, 5, 7, 0, 0, 71, 73, 3, 16, 8, 0, 72, 71, 1, 0, 0, 0, 73, 76,
		1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0,
		76, 74, 1, 0, 0, 0, 77, 78, 5, 8, 0, 0, 78, 13, 1, 0, 0, 0, 79, 80, 5,
		6, 0, 0, 80, 84, 5, 7, 0, 0, 81, 83, 3, 16, 8, 0, 82, 81, 1, 0, 0, 0, 83,
		86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0,
		0, 86, 84, 1, 0, 0, 0, 87, 88, 5, 8, 0, 0, 88, 15, 1, 0, 0, 0, 89, 90,
		5, 10, 0, 0, 90, 17, 1, 0, 0, 0, 7, 27, 37, 46, 54, 64, 74, 84,
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

// ProjectParserInit initializes any static state used to implement ProjectParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProjectParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProjectParserInit() {
	staticData := &ProjectParserStaticData
	staticData.once.Do(projectParserInit)
}

// NewProjectParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProjectParser(input antlr.TokenStream) *ProjectParser {
	ProjectParserInit()
	this := new(ProjectParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ProjectParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Project.g4"

	return this
}

// ProjectParser tokens.
const (
	ProjectParserEOF           = antlr.TokenEOF
	ProjectParserPROJECT       = 1
	ProjectParserVERSION       = 2
	ProjectParserNEXT_VERSION  = 3
	ProjectParserSERVICES      = 4
	ProjectParserOBJECTS       = 5
	ProjectParserDEPLOYMENTS   = 6
	ProjectParserLBRACE        = 7
	ProjectParserRBRACE        = 8
	ProjectParserIDENT         = 9
	ProjectParserSTRING        = 10
	ProjectParserLINE_COMMENT  = 11
	ProjectParserBLOCK_COMMENT = 12
	ProjectParserWS            = 13
)

// ProjectParser rules.
const (
	ProjectParserRULE_program          = 0
	ProjectParserRULE_projectDecl      = 1
	ProjectParserRULE_versionDecl      = 2
	ProjectParserRULE_versionBody      = 3
	ProjectParserRULE_nextVersionDecl  = 4
	ProjectParserRULE_servicesBlock    = 5
	ProjectParserRULE_objectsBlock     = 6
	ProjectParserRULE_deploymentsBlock = 7
	ProjectParserRULE_fileRef          = 8
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ProjectDecl() IProjectDeclContext
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
	p.RuleIndex = ProjectParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ProjectDecl() IProjectDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ProjectParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProjectParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.ProjectDecl()
	}
	{
		p.SetState(19)
		p.Match(ProjectParserEOF)
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

// IProjectDeclContext is an interface to support dynamic dispatch.
type IProjectDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROJECT() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllVersionDecl() []IVersionDeclContext
	VersionDecl(i int) IVersionDeclContext

	// IsProjectDeclContext differentiates from other interfaces.
	IsProjectDeclContext()
}

type ProjectDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectDeclContext() *ProjectDeclContext {
	var p = new(ProjectDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_projectDecl
	return p
}

func InitEmptyProjectDeclContext(p *ProjectDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_projectDecl
}

func (*ProjectDeclContext) IsProjectDeclContext() {}

func NewProjectDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectDeclContext {
	var p = new(ProjectDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_projectDecl

	return p
}

func (s *ProjectDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectDeclContext) PROJECT() antlr.TerminalNode {
	return s.GetToken(ProjectParserPROJECT, 0)
}

func (s *ProjectDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ProjectParserIDENT, 0)
}

func (s *ProjectDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserLBRACE, 0)
}

func (s *ProjectDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBRACE, 0)
}

func (s *ProjectDeclContext) AllVersionDecl() []IVersionDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVersionDeclContext); ok {
			len++
		}
	}

	tst := make([]IVersionDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVersionDeclContext); ok {
			tst[i] = t.(IVersionDeclContext)
			i++
		}
	}

	return tst
}

func (s *ProjectDeclContext) VersionDecl(i int) IVersionDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionDeclContext); ok {
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

	return t.(IVersionDeclContext)
}

func (s *ProjectDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitProjectDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) ProjectDecl() (localctx IProjectDeclContext) {
	localctx = NewProjectDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProjectParserRULE_projectDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		p.Match(ProjectParserPROJECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.Match(ProjectParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(23)
		p.Match(ProjectParserLBRACE)
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

	for ok := true; ok; ok = _la == ProjectParserVERSION {
		{
			p.SetState(24)
			p.VersionDecl()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(29)
		p.Match(ProjectParserRBRACE)
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

// IVersionDeclContext is an interface to support dynamic dispatch.
type IVersionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VERSION() antlr.TerminalNode
	STRING() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllVersionBody() []IVersionBodyContext
	VersionBody(i int) IVersionBodyContext

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
	p.RuleIndex = ProjectParserRULE_versionDecl
	return p
}

func InitEmptyVersionDeclContext(p *VersionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_versionDecl
}

func (*VersionDeclContext) IsVersionDeclContext() {}

func NewVersionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionDeclContext {
	var p = new(VersionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_versionDecl

	return p
}

func (s *VersionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionDeclContext) VERSION() antlr.TerminalNode {
	return s.GetToken(ProjectParserVERSION, 0)
}

func (s *VersionDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserSTRING, 0)
}

func (s *VersionDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserLBRACE, 0)
}

func (s *VersionDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBRACE, 0)
}

func (s *VersionDeclContext) AllVersionBody() []IVersionBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVersionBodyContext); ok {
			len++
		}
	}

	tst := make([]IVersionBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVersionBodyContext); ok {
			tst[i] = t.(IVersionBodyContext)
			i++
		}
	}

	return tst
}

func (s *VersionDeclContext) VersionBody(i int) IVersionBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionBodyContext); ok {
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

	return t.(IVersionBodyContext)
}

func (s *VersionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitVersionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) VersionDecl() (localctx IVersionDeclContext) {
	localctx = NewVersionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProjectParserRULE_versionDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(ProjectParserVERSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(ProjectParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(ProjectParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0 {
		{
			p.SetState(34)
			p.VersionBody()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(ProjectParserRBRACE)
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

// IVersionBodyContext is an interface to support dynamic dispatch.
type IVersionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ServicesBlock() IServicesBlockContext
	ObjectsBlock() IObjectsBlockContext
	DeploymentsBlock() IDeploymentsBlockContext
	NextVersionDecl() INextVersionDeclContext

	// IsVersionBodyContext differentiates from other interfaces.
	IsVersionBodyContext()
}

type VersionBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionBodyContext() *VersionBodyContext {
	var p = new(VersionBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_versionBody
	return p
}

func InitEmptyVersionBodyContext(p *VersionBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_versionBody
}

func (*VersionBodyContext) IsVersionBodyContext() {}

func NewVersionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionBodyContext {
	var p = new(VersionBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_versionBody

	return p
}

func (s *VersionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionBodyContext) ServicesBlock() IServicesBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServicesBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServicesBlockContext)
}

func (s *VersionBodyContext) ObjectsBlock() IObjectsBlockContext {
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

func (s *VersionBodyContext) DeploymentsBlock() IDeploymentsBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeploymentsBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeploymentsBlockContext)
}

func (s *VersionBodyContext) NextVersionDecl() INextVersionDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INextVersionDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INextVersionDeclContext)
}

func (s *VersionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitVersionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) VersionBody() (localctx IVersionBodyContext) {
	localctx = NewVersionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProjectParserRULE_versionBody)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProjectParserSERVICES:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.ServicesBlock()
		}

	case ProjectParserOBJECTS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.ObjectsBlock()
		}

	case ProjectParserDEPLOYMENTS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.DeploymentsBlock()
		}

	case ProjectParserNEXT_VERSION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.NextVersionDecl()
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

// INextVersionDeclContext is an interface to support dynamic dispatch.
type INextVersionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEXT_VERSION() antlr.TerminalNode
	STRING() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllVersionBody() []IVersionBodyContext
	VersionBody(i int) IVersionBodyContext

	// IsNextVersionDeclContext differentiates from other interfaces.
	IsNextVersionDeclContext()
}

type NextVersionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextVersionDeclContext() *NextVersionDeclContext {
	var p = new(NextVersionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_nextVersionDecl
	return p
}

func InitEmptyNextVersionDeclContext(p *NextVersionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_nextVersionDecl
}

func (*NextVersionDeclContext) IsNextVersionDeclContext() {}

func NewNextVersionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextVersionDeclContext {
	var p = new(NextVersionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_nextVersionDecl

	return p
}

func (s *NextVersionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NextVersionDeclContext) NEXT_VERSION() antlr.TerminalNode {
	return s.GetToken(ProjectParserNEXT_VERSION, 0)
}

func (s *NextVersionDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserSTRING, 0)
}

func (s *NextVersionDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserLBRACE, 0)
}

func (s *NextVersionDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBRACE, 0)
}

func (s *NextVersionDeclContext) AllVersionBody() []IVersionBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVersionBodyContext); ok {
			len++
		}
	}

	tst := make([]IVersionBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVersionBodyContext); ok {
			tst[i] = t.(IVersionBodyContext)
			i++
		}
	}

	return tst
}

func (s *NextVersionDeclContext) VersionBody(i int) IVersionBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionBodyContext); ok {
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

	return t.(IVersionBodyContext)
}

func (s *NextVersionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextVersionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextVersionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitNextVersionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) NextVersionDecl() (localctx INextVersionDeclContext) {
	localctx = NewNextVersionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ProjectParserRULE_nextVersionDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(ProjectParserNEXT_VERSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(ProjectParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(ProjectParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0 {
		{
			p.SetState(51)
			p.VersionBody()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(ProjectParserRBRACE)
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

// IServicesBlockContext is an interface to support dynamic dispatch.
type IServicesBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICES() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFileRef() []IFileRefContext
	FileRef(i int) IFileRefContext

	// IsServicesBlockContext differentiates from other interfaces.
	IsServicesBlockContext()
}

type ServicesBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServicesBlockContext() *ServicesBlockContext {
	var p = new(ServicesBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_servicesBlock
	return p
}

func InitEmptyServicesBlockContext(p *ServicesBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_servicesBlock
}

func (*ServicesBlockContext) IsServicesBlockContext() {}

func NewServicesBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServicesBlockContext {
	var p = new(ServicesBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_servicesBlock

	return p
}

func (s *ServicesBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ServicesBlockContext) SERVICES() antlr.TerminalNode {
	return s.GetToken(ProjectParserSERVICES, 0)
}

func (s *ServicesBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserLBRACE, 0)
}

func (s *ServicesBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBRACE, 0)
}

func (s *ServicesBlockContext) AllFileRef() []IFileRefContext {
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

func (s *ServicesBlockContext) FileRef(i int) IFileRefContext {
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

func (s *ServicesBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServicesBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServicesBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitServicesBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) ServicesBlock() (localctx IServicesBlockContext) {
	localctx = NewServicesBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProjectParserRULE_servicesBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(ProjectParserSERVICES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(ProjectParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProjectParserSTRING {
		{
			p.SetState(61)
			p.FileRef()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(ProjectParserRBRACE)
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
	p.RuleIndex = ProjectParserRULE_objectsBlock
	return p
}

func InitEmptyObjectsBlockContext(p *ObjectsBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_objectsBlock
}

func (*ObjectsBlockContext) IsObjectsBlockContext() {}

func NewObjectsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectsBlockContext {
	var p = new(ObjectsBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_objectsBlock

	return p
}

func (s *ObjectsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectsBlockContext) OBJECTS() antlr.TerminalNode {
	return s.GetToken(ProjectParserOBJECTS, 0)
}

func (s *ObjectsBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserLBRACE, 0)
}

func (s *ObjectsBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBRACE, 0)
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
	case ProjectVisitor:
		return t.VisitObjectsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) ObjectsBlock() (localctx IObjectsBlockContext) {
	localctx = NewObjectsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProjectParserRULE_objectsBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(ProjectParserOBJECTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(ProjectParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProjectParserSTRING {
		{
			p.SetState(71)
			p.FileRef()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(ProjectParserRBRACE)
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

// IDeploymentsBlockContext is an interface to support dynamic dispatch.
type IDeploymentsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEPLOYMENTS() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFileRef() []IFileRefContext
	FileRef(i int) IFileRefContext

	// IsDeploymentsBlockContext differentiates from other interfaces.
	IsDeploymentsBlockContext()
}

type DeploymentsBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeploymentsBlockContext() *DeploymentsBlockContext {
	var p = new(DeploymentsBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_deploymentsBlock
	return p
}

func InitEmptyDeploymentsBlockContext(p *DeploymentsBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_deploymentsBlock
}

func (*DeploymentsBlockContext) IsDeploymentsBlockContext() {}

func NewDeploymentsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeploymentsBlockContext {
	var p = new(DeploymentsBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_deploymentsBlock

	return p
}

func (s *DeploymentsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DeploymentsBlockContext) DEPLOYMENTS() antlr.TerminalNode {
	return s.GetToken(ProjectParserDEPLOYMENTS, 0)
}

func (s *DeploymentsBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserLBRACE, 0)
}

func (s *DeploymentsBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBRACE, 0)
}

func (s *DeploymentsBlockContext) AllFileRef() []IFileRefContext {
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

func (s *DeploymentsBlockContext) FileRef(i int) IFileRefContext {
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

func (s *DeploymentsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeploymentsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeploymentsBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitDeploymentsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) DeploymentsBlock() (localctx IDeploymentsBlockContext) {
	localctx = NewDeploymentsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ProjectParserRULE_deploymentsBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(ProjectParserDEPLOYMENTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(ProjectParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProjectParserSTRING {
		{
			p.SetState(81)
			p.FileRef()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(ProjectParserRBRACE)
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
	p.RuleIndex = ProjectParserRULE_fileRef
	return p
}

func InitEmptyFileRefContext(p *FileRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProjectParserRULE_fileRef
}

func (*FileRefContext) IsFileRefContext() {}

func NewFileRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileRefContext {
	var p = new(FileRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_fileRef

	return p
}

func (s *FileRefContext) GetParser() antlr.Parser { return s.parser }

func (s *FileRefContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserSTRING, 0)
}

func (s *FileRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ProjectVisitor:
		return t.VisitFileRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ProjectParser) FileRef() (localctx IFileRefContext) {
	localctx = NewFileRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ProjectParserRULE_fileRef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(ProjectParserSTRING)
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
