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
		"", "'application'", "'emits'", "'listens'", "'on'", "'objects'", "'endpoints'",
		"'consumers'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "APPLICATION", "EMITS", "LISTENS", "ON", "OBJECTS", "ENDPOINTS",
		"CONSUMERS", "LBRACE", "RBRACE", "IDENT", "INT", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"APPLICATION", "EMITS", "LISTENS", "ON", "OBJECTS", "ENDPOINTS", "CONSUMERS",
		"LBRACE", "RBRACE", "IDENT", "INT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 147, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 5, 9, 95, 8, 9, 10, 9, 12, 9, 98, 9, 9, 1, 10, 4, 10, 101,
		8, 10, 11, 10, 12, 10, 102, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 109, 8,
		11, 10, 11, 12, 11, 112, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		5, 12, 120, 8, 12, 10, 12, 12, 12, 123, 9, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 5, 13, 131, 8, 13, 10, 13, 12, 13, 134, 9, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 4, 14, 142, 8, 14, 11, 14, 12, 14, 143,
		1, 14, 1, 14, 1, 132, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 1, 0, 6,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 2, 0, 34, 34, 92, 92, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13,
		13, 32, 32, 153, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 49, 1, 0, 0, 0, 7, 57,
		1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13, 78, 1, 0, 0, 0,
		15, 88, 1, 0, 0, 0, 17, 90, 1, 0, 0, 0, 19, 92, 1, 0, 0, 0, 21, 100, 1,
		0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 115, 1, 0, 0, 0, 27, 126, 1, 0, 0, 0,
		29, 141, 1, 0, 0, 0, 31, 32, 5, 97, 0, 0, 32, 33, 5, 112, 0, 0, 33, 34,
		5, 112, 0, 0, 34, 35, 5, 108, 0, 0, 35, 36, 5, 105, 0, 0, 36, 37, 5, 99,
		0, 0, 37, 38, 5, 97, 0, 0, 38, 39, 5, 116, 0, 0, 39, 40, 5, 105, 0, 0,
		40, 41, 5, 111, 0, 0, 41, 42, 5, 110, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44,
		5, 101, 0, 0, 44, 45, 5, 109, 0, 0, 45, 46, 5, 105, 0, 0, 46, 47, 5, 116,
		0, 0, 47, 48, 5, 115, 0, 0, 48, 4, 1, 0, 0, 0, 49, 50, 5, 108, 0, 0, 50,
		51, 5, 105, 0, 0, 51, 52, 5, 115, 0, 0, 52, 53, 5, 116, 0, 0, 53, 54, 5,
		101, 0, 0, 54, 55, 5, 110, 0, 0, 55, 56, 5, 115, 0, 0, 56, 6, 1, 0, 0,
		0, 57, 58, 5, 111, 0, 0, 58, 59, 5, 110, 0, 0, 59, 8, 1, 0, 0, 0, 60, 61,
		5, 111, 0, 0, 61, 62, 5, 98, 0, 0, 62, 63, 5, 106, 0, 0, 63, 64, 5, 101,
		0, 0, 64, 65, 5, 99, 0, 0, 65, 66, 5, 116, 0, 0, 66, 67, 5, 115, 0, 0,
		67, 10, 1, 0, 0, 0, 68, 69, 5, 101, 0, 0, 69, 70, 5, 110, 0, 0, 70, 71,
		5, 100, 0, 0, 71, 72, 5, 112, 0, 0, 72, 73, 5, 111, 0, 0, 73, 74, 5, 105,
		0, 0, 74, 75, 5, 110, 0, 0, 75, 76, 5, 116, 0, 0, 76, 77, 5, 115, 0, 0,
		77, 12, 1, 0, 0, 0, 78, 79, 5, 99, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81,
		5, 110, 0, 0, 81, 82, 5, 115, 0, 0, 82, 83, 5, 117, 0, 0, 83, 84, 5, 109,
		0, 0, 84, 85, 5, 101, 0, 0, 85, 86, 5, 114, 0, 0, 86, 87, 5, 115, 0, 0,
		87, 14, 1, 0, 0, 0, 88, 89, 5, 123, 0, 0, 89, 16, 1, 0, 0, 0, 90, 91, 5,
		125, 0, 0, 91, 18, 1, 0, 0, 0, 92, 96, 7, 0, 0, 0, 93, 95, 7, 1, 0, 0,
		94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1,
		0, 0, 0, 97, 20, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 101, 7, 2, 0, 0, 100,
		99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1,
		0, 0, 0, 103, 22, 1, 0, 0, 0, 104, 110, 5, 34, 0, 0, 105, 109, 8, 3, 0,
		0, 106, 107, 5, 92, 0, 0, 107, 109, 9, 0, 0, 0, 108, 105, 1, 0, 0, 0, 108,
		106, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111,
		1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 114, 5, 34,
		0, 0, 114, 24, 1, 0, 0, 0, 115, 116, 5, 47, 0, 0, 116, 117, 5, 47, 0, 0,
		117, 121, 1, 0, 0, 0, 118, 120, 8, 4, 0, 0, 119, 118, 1, 0, 0, 0, 120,
		123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124,
		1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 125, 6, 12, 0, 0, 125, 26, 1, 0,
		0, 0, 126, 127, 5, 47, 0, 0, 127, 128, 5, 42, 0, 0, 128, 132, 1, 0, 0,
		0, 129, 131, 9, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 135, 136, 5, 42, 0, 0, 136, 137, 5, 47, 0, 0, 137, 138, 1,
		0, 0, 0, 138, 139, 6, 13, 0, 0, 139, 28, 1, 0, 0, 0, 140, 142, 7, 5, 0,
		0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 6, 14, 0, 0, 146, 30,
		1, 0, 0, 0, 8, 0, 96, 102, 108, 110, 121, 132, 143, 1, 6, 0, 0,
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
	ApplicationLexerOBJECTS       = 5
	ApplicationLexerENDPOINTS     = 6
	ApplicationLexerCONSUMERS     = 7
	ApplicationLexerLBRACE        = 8
	ApplicationLexerRBRACE        = 9
	ApplicationLexerIDENT         = 10
	ApplicationLexerINT           = 11
	ApplicationLexerSTRING        = 12
	ApplicationLexerLINE_COMMENT  = 13
	ApplicationLexerBLOCK_COMMENT = 14
	ApplicationLexerWS            = 15
)
