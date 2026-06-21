// Code generated from internal/infrastructure/antlr/grammars/Deployment.g4 by ANTLR 4.13.2. DO NOT EDIT.

package deployment

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DeploymentLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DeploymentLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func deploymentlexerLexerInit() {
	staticData := &DeploymentLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"DEPLOYMENT", "VENDOR", "INVENTORY", "SERVICE", "REPLICAS", "DOMAIN",
		"HOST", "VOLUME", "BACKUP", "TRUE", "FALSE", "LBRACE", "RBRACE", "IDENT",
		"INT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 180, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 5, 13, 128, 8, 13, 10, 13, 12, 13, 131, 9, 13, 1, 14, 4, 14,
		134, 8, 14, 11, 14, 12, 14, 135, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 142,
		8, 15, 10, 15, 12, 15, 145, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 5, 16, 153, 8, 16, 10, 16, 12, 16, 156, 9, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 5, 17, 164, 8, 17, 10, 17, 12, 17, 167, 9, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 4, 18, 175, 8, 18, 11, 18, 12, 18,
		176, 1, 18, 1, 18, 1, 165, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92,
		2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 186, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3,
		50, 1, 0, 0, 0, 5, 57, 1, 0, 0, 0, 7, 67, 1, 0, 0, 0, 9, 75, 1, 0, 0, 0,
		11, 84, 1, 0, 0, 0, 13, 91, 1, 0, 0, 0, 15, 96, 1, 0, 0, 0, 17, 103, 1,
		0, 0, 0, 19, 110, 1, 0, 0, 0, 21, 115, 1, 0, 0, 0, 23, 121, 1, 0, 0, 0,
		25, 123, 1, 0, 0, 0, 27, 125, 1, 0, 0, 0, 29, 133, 1, 0, 0, 0, 31, 137,
		1, 0, 0, 0, 33, 148, 1, 0, 0, 0, 35, 159, 1, 0, 0, 0, 37, 174, 1, 0, 0,
		0, 39, 40, 5, 100, 0, 0, 40, 41, 5, 101, 0, 0, 41, 42, 5, 112, 0, 0, 42,
		43, 5, 108, 0, 0, 43, 44, 5, 111, 0, 0, 44, 45, 5, 121, 0, 0, 45, 46, 5,
		109, 0, 0, 46, 47, 5, 101, 0, 0, 47, 48, 5, 110, 0, 0, 48, 49, 5, 116,
		0, 0, 49, 2, 1, 0, 0, 0, 50, 51, 5, 118, 0, 0, 51, 52, 5, 101, 0, 0, 52,
		53, 5, 110, 0, 0, 53, 54, 5, 100, 0, 0, 54, 55, 5, 111, 0, 0, 55, 56, 5,
		114, 0, 0, 56, 4, 1, 0, 0, 0, 57, 58, 5, 105, 0, 0, 58, 59, 5, 110, 0,
		0, 59, 60, 5, 118, 0, 0, 60, 61, 5, 101, 0, 0, 61, 62, 5, 110, 0, 0, 62,
		63, 5, 116, 0, 0, 63, 64, 5, 111, 0, 0, 64, 65, 5, 114, 0, 0, 65, 66, 5,
		121, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68, 5, 115, 0, 0, 68, 69, 5, 101, 0,
		0, 69, 70, 5, 114, 0, 0, 70, 71, 5, 118, 0, 0, 71, 72, 5, 105, 0, 0, 72,
		73, 5, 99, 0, 0, 73, 74, 5, 101, 0, 0, 74, 8, 1, 0, 0, 0, 75, 76, 5, 114,
		0, 0, 76, 77, 5, 101, 0, 0, 77, 78, 5, 112, 0, 0, 78, 79, 5, 108, 0, 0,
		79, 80, 5, 105, 0, 0, 80, 81, 5, 99, 0, 0, 81, 82, 5, 97, 0, 0, 82, 83,
		5, 115, 0, 0, 83, 10, 1, 0, 0, 0, 84, 85, 5, 100, 0, 0, 85, 86, 5, 111,
		0, 0, 86, 87, 5, 109, 0, 0, 87, 88, 5, 97, 0, 0, 88, 89, 5, 105, 0, 0,
		89, 90, 5, 110, 0, 0, 90, 12, 1, 0, 0, 0, 91, 92, 5, 104, 0, 0, 92, 93,
		5, 111, 0, 0, 93, 94, 5, 115, 0, 0, 94, 95, 5, 116, 0, 0, 95, 14, 1, 0,
		0, 0, 96, 97, 5, 118, 0, 0, 97, 98, 5, 111, 0, 0, 98, 99, 5, 108, 0, 0,
		99, 100, 5, 117, 0, 0, 100, 101, 5, 109, 0, 0, 101, 102, 5, 101, 0, 0,
		102, 16, 1, 0, 0, 0, 103, 104, 5, 98, 0, 0, 104, 105, 5, 97, 0, 0, 105,
		106, 5, 99, 0, 0, 106, 107, 5, 107, 0, 0, 107, 108, 5, 117, 0, 0, 108,
		109, 5, 112, 0, 0, 109, 18, 1, 0, 0, 0, 110, 111, 5, 116, 0, 0, 111, 112,
		5, 114, 0, 0, 112, 113, 5, 117, 0, 0, 113, 114, 5, 101, 0, 0, 114, 20,
		1, 0, 0, 0, 115, 116, 5, 102, 0, 0, 116, 117, 5, 97, 0, 0, 117, 118, 5,
		108, 0, 0, 118, 119, 5, 115, 0, 0, 119, 120, 5, 101, 0, 0, 120, 22, 1,
		0, 0, 0, 121, 122, 5, 123, 0, 0, 122, 24, 1, 0, 0, 0, 123, 124, 5, 125,
		0, 0, 124, 26, 1, 0, 0, 0, 125, 129, 7, 0, 0, 0, 126, 128, 7, 1, 0, 0,
		127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129,
		130, 1, 0, 0, 0, 130, 28, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 134, 7,
		2, 0, 0, 133, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 133, 1, 0, 0,
		0, 135, 136, 1, 0, 0, 0, 136, 30, 1, 0, 0, 0, 137, 143, 5, 34, 0, 0, 138,
		142, 8, 3, 0, 0, 139, 140, 5, 92, 0, 0, 140, 142, 9, 0, 0, 0, 141, 138,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0,
		146, 147, 5, 34, 0, 0, 147, 32, 1, 0, 0, 0, 148, 149, 5, 47, 0, 0, 149,
		150, 5, 47, 0, 0, 150, 154, 1, 0, 0, 0, 151, 153, 8, 4, 0, 0, 152, 151,
		1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 6, 16, 0, 0,
		158, 34, 1, 0, 0, 0, 159, 160, 5, 47, 0, 0, 160, 161, 5, 42, 0, 0, 161,
		165, 1, 0, 0, 0, 162, 164, 9, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 167,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 168, 1, 0,
		0, 0, 167, 165, 1, 0, 0, 0, 168, 169, 5, 42, 0, 0, 169, 170, 5, 47, 0,
		0, 170, 171, 1, 0, 0, 0, 171, 172, 6, 17, 0, 0, 172, 36, 1, 0, 0, 0, 173,
		175, 7, 5, 0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 174,
		1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 6, 18,
		0, 0, 179, 38, 1, 0, 0, 0, 8, 0, 129, 135, 141, 143, 154, 165, 176, 1,
		6, 0, 0,
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

// DeploymentLexerInit initializes any static state used to implement DeploymentLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDeploymentLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DeploymentLexerInit() {
	staticData := &DeploymentLexerLexerStaticData
	staticData.once.Do(deploymentlexerLexerInit)
}

// NewDeploymentLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDeploymentLexer(input antlr.CharStream) *DeploymentLexer {
	DeploymentLexerInit()
	l := new(DeploymentLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DeploymentLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Deployment.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DeploymentLexer tokens.
const (
	DeploymentLexerDEPLOYMENT    = 1
	DeploymentLexerVENDOR        = 2
	DeploymentLexerINVENTORY     = 3
	DeploymentLexerSERVICE       = 4
	DeploymentLexerREPLICAS      = 5
	DeploymentLexerDOMAIN        = 6
	DeploymentLexerHOST          = 7
	DeploymentLexerVOLUME        = 8
	DeploymentLexerBACKUP        = 9
	DeploymentLexerTRUE          = 10
	DeploymentLexerFALSE         = 11
	DeploymentLexerLBRACE        = 12
	DeploymentLexerRBRACE        = 13
	DeploymentLexerIDENT         = 14
	DeploymentLexerINT           = 15
	DeploymentLexerSTRING        = 16
	DeploymentLexerLINE_COMMENT  = 17
	DeploymentLexerBLOCK_COMMENT = 18
	DeploymentLexerWS            = 19
)
