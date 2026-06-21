// Code generated from internal/infrastructure/antlr/grammars/Object.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type ObjectLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ObjectLexerLexerStaticData struct {
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

func objectlexerLexerInit() {
	staticData := &ObjectLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'object'", "'history_of'", "'primary'", "'unique'", "'renamed_from'",
		"'deprecated'", "'List'", "'true'", "'false'", "'{'", "'}'", "'['",
		"']'", "'<'", "'>'", "','", "'+'", "'*'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "OBJECT", "HISTORY_OF", "PRIMARY", "UNIQUE", "RENAMED_FROM", "DEPRECATED",
		"LIST", "TRUE", "FALSE", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET",
		"LT", "GT", "COMMA", "PLUS", "STAR", "QUESTION", "IDENT", "FLOAT", "INT",
		"STRING", "LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"OBJECT", "HISTORY_OF", "PRIMARY", "UNIQUE", "RENAMED_FROM", "DEPRECATED",
		"LIST", "TRUE", "FALSE", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET",
		"LT", "GT", "COMMA", "PLUS", "STAR", "QUESTION", "IDENT", "FLOAT", "INT",
		"STRING", "LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 212, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 5, 19, 149, 8, 19, 10, 19, 12, 19, 152, 9, 19, 1, 20, 4, 20,
		155, 8, 20, 11, 20, 12, 20, 156, 1, 20, 1, 20, 4, 20, 161, 8, 20, 11, 20,
		12, 20, 162, 1, 21, 4, 21, 166, 8, 21, 11, 21, 12, 21, 167, 1, 22, 1, 22,
		1, 22, 1, 22, 5, 22, 174, 8, 22, 10, 22, 12, 22, 177, 9, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 185, 8, 23, 10, 23, 12, 23, 188,
		9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 196, 8, 24, 10,
		24, 12, 24, 199, 9, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 4, 25,
		207, 8, 25, 11, 25, 12, 25, 208, 1, 25, 1, 25, 1, 197, 0, 26, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 1, 0, 6, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 34,
		34, 92, 92, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 220, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 53, 1, 0, 0, 0, 3,
		60, 1, 0, 0, 0, 5, 71, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0, 9, 86, 1, 0, 0, 0,
		11, 99, 1, 0, 0, 0, 13, 110, 1, 0, 0, 0, 15, 115, 1, 0, 0, 0, 17, 120,
		1, 0, 0, 0, 19, 126, 1, 0, 0, 0, 21, 128, 1, 0, 0, 0, 23, 130, 1, 0, 0,
		0, 25, 132, 1, 0, 0, 0, 27, 134, 1, 0, 0, 0, 29, 136, 1, 0, 0, 0, 31, 138,
		1, 0, 0, 0, 33, 140, 1, 0, 0, 0, 35, 142, 1, 0, 0, 0, 37, 144, 1, 0, 0,
		0, 39, 146, 1, 0, 0, 0, 41, 154, 1, 0, 0, 0, 43, 165, 1, 0, 0, 0, 45, 169,
		1, 0, 0, 0, 47, 180, 1, 0, 0, 0, 49, 191, 1, 0, 0, 0, 51, 206, 1, 0, 0,
		0, 53, 54, 5, 111, 0, 0, 54, 55, 5, 98, 0, 0, 55, 56, 5, 106, 0, 0, 56,
		57, 5, 101, 0, 0, 57, 58, 5, 99, 0, 0, 58, 59, 5, 116, 0, 0, 59, 2, 1,
		0, 0, 0, 60, 61, 5, 104, 0, 0, 61, 62, 5, 105, 0, 0, 62, 63, 5, 115, 0,
		0, 63, 64, 5, 116, 0, 0, 64, 65, 5, 111, 0, 0, 65, 66, 5, 114, 0, 0, 66,
		67, 5, 121, 0, 0, 67, 68, 5, 95, 0, 0, 68, 69, 5, 111, 0, 0, 69, 70, 5,
		102, 0, 0, 70, 4, 1, 0, 0, 0, 71, 72, 5, 112, 0, 0, 72, 73, 5, 114, 0,
		0, 73, 74, 5, 105, 0, 0, 74, 75, 5, 109, 0, 0, 75, 76, 5, 97, 0, 0, 76,
		77, 5, 114, 0, 0, 77, 78, 5, 121, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 117,
		0, 0, 80, 81, 5, 110, 0, 0, 81, 82, 5, 105, 0, 0, 82, 83, 5, 113, 0, 0,
		83, 84, 5, 117, 0, 0, 84, 85, 5, 101, 0, 0, 85, 8, 1, 0, 0, 0, 86, 87,
		5, 114, 0, 0, 87, 88, 5, 101, 0, 0, 88, 89, 5, 110, 0, 0, 89, 90, 5, 97,
		0, 0, 90, 91, 5, 109, 0, 0, 91, 92, 5, 101, 0, 0, 92, 93, 5, 100, 0, 0,
		93, 94, 5, 95, 0, 0, 94, 95, 5, 102, 0, 0, 95, 96, 5, 114, 0, 0, 96, 97,
		5, 111, 0, 0, 97, 98, 5, 109, 0, 0, 98, 10, 1, 0, 0, 0, 99, 100, 5, 100,
		0, 0, 100, 101, 5, 101, 0, 0, 101, 102, 5, 112, 0, 0, 102, 103, 5, 114,
		0, 0, 103, 104, 5, 101, 0, 0, 104, 105, 5, 99, 0, 0, 105, 106, 5, 97, 0,
		0, 106, 107, 5, 116, 0, 0, 107, 108, 5, 101, 0, 0, 108, 109, 5, 100, 0,
		0, 109, 12, 1, 0, 0, 0, 110, 111, 5, 76, 0, 0, 111, 112, 5, 105, 0, 0,
		112, 113, 5, 115, 0, 0, 113, 114, 5, 116, 0, 0, 114, 14, 1, 0, 0, 0, 115,
		116, 5, 116, 0, 0, 116, 117, 5, 114, 0, 0, 117, 118, 5, 117, 0, 0, 118,
		119, 5, 101, 0, 0, 119, 16, 1, 0, 0, 0, 120, 121, 5, 102, 0, 0, 121, 122,
		5, 97, 0, 0, 122, 123, 5, 108, 0, 0, 123, 124, 5, 115, 0, 0, 124, 125,
		5, 101, 0, 0, 125, 18, 1, 0, 0, 0, 126, 127, 5, 123, 0, 0, 127, 20, 1,
		0, 0, 0, 128, 129, 5, 125, 0, 0, 129, 22, 1, 0, 0, 0, 130, 131, 5, 91,
		0, 0, 131, 24, 1, 0, 0, 0, 132, 133, 5, 93, 0, 0, 133, 26, 1, 0, 0, 0,
		134, 135, 5, 60, 0, 0, 135, 28, 1, 0, 0, 0, 136, 137, 5, 62, 0, 0, 137,
		30, 1, 0, 0, 0, 138, 139, 5, 44, 0, 0, 139, 32, 1, 0, 0, 0, 140, 141, 5,
		43, 0, 0, 141, 34, 1, 0, 0, 0, 142, 143, 5, 42, 0, 0, 143, 36, 1, 0, 0,
		0, 144, 145, 5, 63, 0, 0, 145, 38, 1, 0, 0, 0, 146, 150, 7, 0, 0, 0, 147,
		149, 7, 1, 0, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 40, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 153, 155, 7, 2, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0,
		156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		160, 5, 46, 0, 0, 159, 161, 7, 2, 0, 0, 160, 159, 1, 0, 0, 0, 161, 162,
		1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 42, 1, 0,
		0, 0, 164, 166, 7, 2, 0, 0, 165, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0,
		167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 44, 1, 0, 0, 0, 169, 175,
		5, 34, 0, 0, 170, 174, 8, 3, 0, 0, 171, 172, 5, 92, 0, 0, 172, 174, 9,
		0, 0, 0, 173, 170, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 177, 1, 0, 0,
		0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0, 0, 177,
		175, 1, 0, 0, 0, 178, 179, 5, 34, 0, 0, 179, 46, 1, 0, 0, 0, 180, 181,
		5, 47, 0, 0, 181, 182, 5, 47, 0, 0, 182, 186, 1, 0, 0, 0, 183, 185, 8,
		4, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0,
		0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189,
		190, 6, 23, 0, 0, 190, 48, 1, 0, 0, 0, 191, 192, 5, 47, 0, 0, 192, 193,
		5, 42, 0, 0, 193, 197, 1, 0, 0, 0, 194, 196, 9, 0, 0, 0, 195, 194, 1, 0,
		0, 0, 196, 199, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0,
		198, 200, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 201, 5, 42, 0, 0, 201,
		202, 5, 47, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 6, 24, 0, 0, 204, 50,
		1, 0, 0, 0, 205, 207, 7, 5, 0, 0, 206, 205, 1, 0, 0, 0, 207, 208, 1, 0,
		0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0,
		210, 211, 6, 25, 0, 0, 211, 52, 1, 0, 0, 0, 10, 0, 150, 156, 162, 167,
		173, 175, 186, 197, 208, 1, 6, 0, 0,
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

// ObjectLexerInit initializes any static state used to implement ObjectLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewObjectLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ObjectLexerInit() {
	staticData := &ObjectLexerLexerStaticData
	staticData.once.Do(objectlexerLexerInit)
}

// NewObjectLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewObjectLexer(input antlr.CharStream) *ObjectLexer {
	ObjectLexerInit()
	l := new(ObjectLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ObjectLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Object.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ObjectLexer tokens.
const (
	ObjectLexerOBJECT        = 1
	ObjectLexerHISTORY_OF    = 2
	ObjectLexerPRIMARY       = 3
	ObjectLexerUNIQUE        = 4
	ObjectLexerRENAMED_FROM  = 5
	ObjectLexerDEPRECATED    = 6
	ObjectLexerLIST          = 7
	ObjectLexerTRUE          = 8
	ObjectLexerFALSE         = 9
	ObjectLexerLBRACE        = 10
	ObjectLexerRBRACE        = 11
	ObjectLexerLBRACKET      = 12
	ObjectLexerRBRACKET      = 13
	ObjectLexerLT            = 14
	ObjectLexerGT            = 15
	ObjectLexerCOMMA         = 16
	ObjectLexerPLUS          = 17
	ObjectLexerSTAR          = 18
	ObjectLexerQUESTION      = 19
	ObjectLexerIDENT         = 20
	ObjectLexerFLOAT         = 21
	ObjectLexerINT           = 22
	ObjectLexerSTRING        = 23
	ObjectLexerLINE_COMMENT  = 24
	ObjectLexerBLOCK_COMMENT = 25
	ObjectLexerWS            = 26
)
