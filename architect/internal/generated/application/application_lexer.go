// Code generated from internal/infrastructure/antlr/grammars/Application.g4 by ANTLR 4.13.2. DO NOT EDIT.

package application

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

type ApplicationLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ApplicationLexerLexerStaticData struct {
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

func applicationlexerLexerInit() {
	staticData := &ApplicationLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"APPLICATION", "EMITS", "LISTENS", "ON", "ENDPOINTS", "CONSUMERS", "LBRACE",
		"RBRACE", "IDENT", "INT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 137, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 5, 8, 85, 8, 8, 10, 8, 12, 8, 88, 9, 8, 1, 9, 4,
		9, 91, 8, 9, 11, 9, 12, 9, 92, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 99, 8,
		10, 10, 10, 12, 10, 102, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		5, 11, 110, 8, 11, 10, 11, 12, 11, 113, 9, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 5, 12, 121, 8, 12, 10, 12, 12, 12, 124, 9, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 4, 13, 132, 8, 13, 11, 13, 12, 13, 133,
		1, 13, 1, 13, 1, 122, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 1, 0, 6, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57,
		2, 0, 34, 34, 92, 92, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32,
		143, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 1, 29, 1, 0, 0, 0, 3,
		41, 1, 0, 0, 0, 5, 47, 1, 0, 0, 0, 7, 55, 1, 0, 0, 0, 9, 58, 1, 0, 0, 0,
		11, 68, 1, 0, 0, 0, 13, 78, 1, 0, 0, 0, 15, 80, 1, 0, 0, 0, 17, 82, 1,
		0, 0, 0, 19, 90, 1, 0, 0, 0, 21, 94, 1, 0, 0, 0, 23, 105, 1, 0, 0, 0, 25,
		116, 1, 0, 0, 0, 27, 131, 1, 0, 0, 0, 29, 30, 5, 97, 0, 0, 30, 31, 5, 112,
		0, 0, 31, 32, 5, 112, 0, 0, 32, 33, 5, 108, 0, 0, 33, 34, 5, 105, 0, 0,
		34, 35, 5, 99, 0, 0, 35, 36, 5, 97, 0, 0, 36, 37, 5, 116, 0, 0, 37, 38,
		5, 105, 0, 0, 38, 39, 5, 111, 0, 0, 39, 40, 5, 110, 0, 0, 40, 2, 1, 0,
		0, 0, 41, 42, 5, 101, 0, 0, 42, 43, 5, 109, 0, 0, 43, 44, 5, 105, 0, 0,
		44, 45, 5, 116, 0, 0, 45, 46, 5, 115, 0, 0, 46, 4, 1, 0, 0, 0, 47, 48,
		5, 108, 0, 0, 48, 49, 5, 105, 0, 0, 49, 50, 5, 115, 0, 0, 50, 51, 5, 116,
		0, 0, 51, 52, 5, 101, 0, 0, 52, 53, 5, 110, 0, 0, 53, 54, 5, 115, 0, 0,
		54, 6, 1, 0, 0, 0, 55, 56, 5, 111, 0, 0, 56, 57, 5, 110, 0, 0, 57, 8, 1,
		0, 0, 0, 58, 59, 5, 101, 0, 0, 59, 60, 5, 110, 0, 0, 60, 61, 5, 100, 0,
		0, 61, 62, 5, 112, 0, 0, 62, 63, 5, 111, 0, 0, 63, 64, 5, 105, 0, 0, 64,
		65, 5, 110, 0, 0, 65, 66, 5, 116, 0, 0, 66, 67, 5, 115, 0, 0, 67, 10, 1,
		0, 0, 0, 68, 69, 5, 99, 0, 0, 69, 70, 5, 111, 0, 0, 70, 71, 5, 110, 0,
		0, 71, 72, 5, 115, 0, 0, 72, 73, 5, 117, 0, 0, 73, 74, 5, 109, 0, 0, 74,
		75, 5, 101, 0, 0, 75, 76, 5, 114, 0, 0, 76, 77, 5, 115, 0, 0, 77, 12, 1,
		0, 0, 0, 78, 79, 5, 123, 0, 0, 79, 14, 1, 0, 0, 0, 80, 81, 5, 125, 0, 0,
		81, 16, 1, 0, 0, 0, 82, 86, 7, 0, 0, 0, 83, 85, 7, 1, 0, 0, 84, 83, 1,
		0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87,
		18, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 91, 7, 2, 0, 0, 90, 89, 1, 0, 0,
		0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 20,
		1, 0, 0, 0, 94, 100, 5, 34, 0, 0, 95, 99, 8, 3, 0, 0, 96, 97, 5, 92, 0,
		0, 97, 99, 9, 0, 0, 0, 98, 95, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 102,
		1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0,
		0, 0, 102, 100, 1, 0, 0, 0, 103, 104, 5, 34, 0, 0, 104, 22, 1, 0, 0, 0,
		105, 106, 5, 47, 0, 0, 106, 107, 5, 47, 0, 0, 107, 111, 1, 0, 0, 0, 108,
		110, 8, 4, 0, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111, 1, 0,
		0, 0, 114, 115, 6, 11, 0, 0, 115, 24, 1, 0, 0, 0, 116, 117, 5, 47, 0, 0,
		117, 118, 5, 42, 0, 0, 118, 122, 1, 0, 0, 0, 119, 121, 9, 0, 0, 0, 120,
		119, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 122, 120,
		1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 126, 5, 42,
		0, 0, 126, 127, 5, 47, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 6, 12, 0,
		0, 129, 26, 1, 0, 0, 0, 130, 132, 7, 5, 0, 0, 131, 130, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135,
		1, 0, 0, 0, 135, 136, 6, 13, 0, 0, 136, 28, 1, 0, 0, 0, 8, 0, 86, 92, 98,
		100, 111, 122, 133, 1, 6, 0, 0,
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

// ApplicationLexerInit initializes any static state used to implement ApplicationLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewApplicationLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ApplicationLexerInit() {
	staticData := &ApplicationLexerLexerStaticData
	staticData.once.Do(applicationlexerLexerInit)
}

// NewApplicationLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewApplicationLexer(input antlr.CharStream) *ApplicationLexer {
	ApplicationLexerInit()
	l := new(ApplicationLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ApplicationLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Application.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ApplicationLexer tokens.
const (
	ApplicationLexerAPPLICATION   = 1
	ApplicationLexerEMITS         = 2
	ApplicationLexerLISTENS       = 3
	ApplicationLexerON            = 4
	ApplicationLexerENDPOINTS     = 5
	ApplicationLexerCONSUMERS     = 6
	ApplicationLexerLBRACE        = 7
	ApplicationLexerRBRACE        = 8
	ApplicationLexerIDENT         = 9
	ApplicationLexerINT           = 10
	ApplicationLexerSTRING        = 11
	ApplicationLexerLINE_COMMENT  = 12
	ApplicationLexerBLOCK_COMMENT = 13
	ApplicationLexerWS            = 14
)
