// Code generated from internal/infrastructure/antlr/grammars/Common.g4 by ANTLR 4.13.2. DO NOT EDIT.

package object

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

type CommonLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CommonLexerLexerStaticData struct {
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

func commonlexerLexerInit() {
	staticData := &CommonLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"LIST", "INPUT", "EVENT", "RESULT", "TRUE", "FALSE", "LBRACE", "RBRACE",
		"LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LT", "GT", "COMMA", "PIPE",
		"PLUS", "STAR", "QUESTION", "IDENT", "FLOAT", "INT", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 180, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 117, 8, 19, 10, 19, 12, 19, 120,
		9, 19, 1, 20, 4, 20, 123, 8, 20, 11, 20, 12, 20, 124, 1, 20, 1, 20, 4,
		20, 129, 8, 20, 11, 20, 12, 20, 130, 1, 21, 4, 21, 134, 8, 21, 11, 21,
		12, 21, 135, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 142, 8, 22, 10, 22, 12,
		22, 145, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 153, 8,
		23, 10, 23, 12, 23, 156, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24,
		5, 24, 164, 8, 24, 10, 24, 12, 24, 167, 9, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 4, 25, 175, 8, 25, 11, 25, 12, 25, 176, 1, 25, 1, 25,
		1, 165, 0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 1, 0, 6,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 2, 0, 34, 34, 92, 92, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13,
		13, 32, 32, 188, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		1, 53, 1, 0, 0, 0, 3, 58, 1, 0, 0, 0, 5, 64, 1, 0, 0, 0, 7, 70, 1, 0, 0,
		0, 9, 77, 1, 0, 0, 0, 11, 82, 1, 0, 0, 0, 13, 88, 1, 0, 0, 0, 15, 90, 1,
		0, 0, 0, 17, 92, 1, 0, 0, 0, 19, 94, 1, 0, 0, 0, 21, 96, 1, 0, 0, 0, 23,
		98, 1, 0, 0, 0, 25, 100, 1, 0, 0, 0, 27, 102, 1, 0, 0, 0, 29, 104, 1, 0,
		0, 0, 31, 106, 1, 0, 0, 0, 33, 108, 1, 0, 0, 0, 35, 110, 1, 0, 0, 0, 37,
		112, 1, 0, 0, 0, 39, 114, 1, 0, 0, 0, 41, 122, 1, 0, 0, 0, 43, 133, 1,
		0, 0, 0, 45, 137, 1, 0, 0, 0, 47, 148, 1, 0, 0, 0, 49, 159, 1, 0, 0, 0,
		51, 174, 1, 0, 0, 0, 53, 54, 5, 76, 0, 0, 54, 55, 5, 105, 0, 0, 55, 56,
		5, 115, 0, 0, 56, 57, 5, 116, 0, 0, 57, 2, 1, 0, 0, 0, 58, 59, 5, 105,
		0, 0, 59, 60, 5, 110, 0, 0, 60, 61, 5, 112, 0, 0, 61, 62, 5, 117, 0, 0,
		62, 63, 5, 116, 0, 0, 63, 4, 1, 0, 0, 0, 64, 65, 5, 101, 0, 0, 65, 66,
		5, 118, 0, 0, 66, 67, 5, 101, 0, 0, 67, 68, 5, 110, 0, 0, 68, 69, 5, 116,
		0, 0, 69, 6, 1, 0, 0, 0, 70, 71, 5, 114, 0, 0, 71, 72, 5, 101, 0, 0, 72,
		73, 5, 115, 0, 0, 73, 74, 5, 117, 0, 0, 74, 75, 5, 108, 0, 0, 75, 76, 5,
		116, 0, 0, 76, 8, 1, 0, 0, 0, 77, 78, 5, 116, 0, 0, 78, 79, 5, 114, 0,
		0, 79, 80, 5, 117, 0, 0, 80, 81, 5, 101, 0, 0, 81, 10, 1, 0, 0, 0, 82,
		83, 5, 102, 0, 0, 83, 84, 5, 97, 0, 0, 84, 85, 5, 108, 0, 0, 85, 86, 5,
		115, 0, 0, 86, 87, 5, 101, 0, 0, 87, 12, 1, 0, 0, 0, 88, 89, 5, 123, 0,
		0, 89, 14, 1, 0, 0, 0, 90, 91, 5, 125, 0, 0, 91, 16, 1, 0, 0, 0, 92, 93,
		5, 40, 0, 0, 93, 18, 1, 0, 0, 0, 94, 95, 5, 41, 0, 0, 95, 20, 1, 0, 0,
		0, 96, 97, 5, 91, 0, 0, 97, 22, 1, 0, 0, 0, 98, 99, 5, 93, 0, 0, 99, 24,
		1, 0, 0, 0, 100, 101, 5, 60, 0, 0, 101, 26, 1, 0, 0, 0, 102, 103, 5, 62,
		0, 0, 103, 28, 1, 0, 0, 0, 104, 105, 5, 44, 0, 0, 105, 30, 1, 0, 0, 0,
		106, 107, 5, 124, 0, 0, 107, 32, 1, 0, 0, 0, 108, 109, 5, 43, 0, 0, 109,
		34, 1, 0, 0, 0, 110, 111, 5, 42, 0, 0, 111, 36, 1, 0, 0, 0, 112, 113, 5,
		63, 0, 0, 113, 38, 1, 0, 0, 0, 114, 118, 7, 0, 0, 0, 115, 117, 7, 1, 0,
		0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118,
		119, 1, 0, 0, 0, 119, 40, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 123, 7,
		2, 0, 0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 122, 1, 0, 0,
		0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 5, 46, 0, 0, 127,
		129, 7, 2, 0, 0, 128, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 128,
		1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 42, 1, 0, 0, 0, 132, 134, 7, 2,
		0, 0, 133, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0,
		135, 136, 1, 0, 0, 0, 136, 44, 1, 0, 0, 0, 137, 143, 5, 34, 0, 0, 138,
		142, 8, 3, 0, 0, 139, 140, 5, 92, 0, 0, 140, 142, 9, 0, 0, 0, 141, 138,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0,
		146, 147, 5, 34, 0, 0, 147, 46, 1, 0, 0, 0, 148, 149, 5, 47, 0, 0, 149,
		150, 5, 47, 0, 0, 150, 154, 1, 0, 0, 0, 151, 153, 8, 4, 0, 0, 152, 151,
		1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 6, 23, 0, 0,
		158, 48, 1, 0, 0, 0, 159, 160, 5, 47, 0, 0, 160, 161, 5, 42, 0, 0, 161,
		165, 1, 0, 0, 0, 162, 164, 9, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 167,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 168, 1, 0,
		0, 0, 167, 165, 1, 0, 0, 0, 168, 169, 5, 42, 0, 0, 169, 170, 5, 47, 0,
		0, 170, 171, 1, 0, 0, 0, 171, 172, 6, 24, 0, 0, 172, 50, 1, 0, 0, 0, 173,
		175, 7, 5, 0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 174,
		1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 6, 25,
		0, 0, 179, 52, 1, 0, 0, 0, 10, 0, 118, 124, 130, 135, 141, 143, 154, 165,
		176, 1, 6, 0, 0,
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

// CommonLexerInit initializes any static state used to implement CommonLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCommonLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CommonLexerInit() {
	staticData := &CommonLexerLexerStaticData
	staticData.once.Do(commonlexerLexerInit)
}

// NewCommonLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCommonLexer(input antlr.CharStream) *CommonLexer {
	CommonLexerInit()
	l := new(CommonLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CommonLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Common.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CommonLexer tokens.
const (
	CommonLexerLIST          = 1
	CommonLexerINPUT         = 2
	CommonLexerEVENT         = 3
	CommonLexerRESULT        = 4
	CommonLexerTRUE          = 5
	CommonLexerFALSE         = 6
	CommonLexerLBRACE        = 7
	CommonLexerRBRACE        = 8
	CommonLexerLPAREN        = 9
	CommonLexerRPAREN        = 10
	CommonLexerLBRACKET      = 11
	CommonLexerRBRACKET      = 12
	CommonLexerLT            = 13
	CommonLexerGT            = 14
	CommonLexerCOMMA         = 15
	CommonLexerPIPE          = 16
	CommonLexerPLUS          = 17
	CommonLexerSTAR          = 18
	CommonLexerQUESTION      = 19
	CommonLexerIDENT         = 20
	CommonLexerFLOAT         = 21
	CommonLexerINT           = 22
	CommonLexerSTRING        = 23
	CommonLexerLINE_COMMENT  = 24
	CommonLexerBLOCK_COMMENT = 25
	CommonLexerWS            = 26
)
