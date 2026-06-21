// Code generated from internal/infrastructure/antlr/grammars/Endpoint.g4 by ANTLR 4.13.2. DO NOT EDIT.

package endpoint // Endpoint
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

type EndpointParser struct {
	*antlr.BaseParser
}

var EndpointParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func endpointParserInit() {
	staticData := &EndpointParserStaticData
	staticData.LiteralNames = []string{
		"", "'endpoint'", "'input'", "'fetch'", "'from'", "'where'", "'create'",
		"'store'", "'to'", "'update'", "'emit'", "'return'", "'path'", "'query'",
		"'body'", "'header'", "'cookie'", "'GET'", "'POST'", "'PUT'", "'PATCH'",
		"'DELETE'", "'List'", "'true'", "'false'", "'{'", "'}'", "'('", "')'",
		"'['", "']'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'='", "':'",
		"'.'", "','", "'|'", "'+'", "'*'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "ENDPOINT", "INPUT", "FETCH", "FROM", "WHERE", "CREATE", "STORE",
		"TO", "UPDATE", "EMIT", "RETURN", "PATH", "QUERY", "BODY", "HEADER",
		"COOKIE", "GET", "POST", "PUT", "PATCH", "DELETE", "LIST", "TRUE", "FALSE",
		"LBRACE", "RBRACE", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LT",
		"GT", "LTE", "GTE", "EQ", "NEQ", "ASSIGN", "COLON", "DOT", "COMMA",
		"PIPE", "PLUS", "STAR", "QUESTION", "IDENT", "FLOAT", "INT", "STRING",
		"LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "endpointDecl", "endpointBody", "inputBlock", "inputField",
		"inputSourceRule", "inputSource", "inputSourceKind", "action", "fetchAction",
		"createAction", "storeAction", "updateAction", "emitAction", "returnAction",
		"typedVariable", "assignmentBlock", "assignment", "condition", "comparator",
		"expression", "functionCall", "argumentList", "selector", "identifier",
		"typeRef", "typeName", "numberConstraint", "optionalMarker", "value",
		"numberValue", "httpMethod",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 51, 246, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 74, 8,
		1, 10, 1, 12, 1, 77, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 83, 8, 2, 1, 3,
		1, 3, 1, 3, 5, 3, 88, 8, 3, 10, 3, 12, 3, 91, 9, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 104, 8, 5, 10, 5, 12,
		5, 107, 9, 5, 1, 5, 1, 5, 3, 5, 111, 8, 5, 1, 6, 1, 6, 3, 6, 115, 8, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 125, 8, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		5, 16, 160, 8, 16, 10, 16, 12, 16, 163, 9, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 3, 20, 180, 8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 185, 8, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 5, 22, 192, 8, 22, 10, 22, 12, 22, 195, 9,
		22, 1, 23, 1, 23, 1, 23, 5, 23, 200, 8, 23, 10, 23, 12, 23, 203, 9, 23,
		1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 209, 8, 25, 1, 25, 3, 25, 212, 8, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 219, 8, 25, 3, 25, 221, 8, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 229, 8, 27, 1, 27, 1,
		27, 3, 27, 233, 8, 27, 1, 27, 3, 27, 236, 8, 27, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 0, 0, 32, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 0, 6, 1, 0, 12, 16, 1, 0, 31, 36, 2, 0, 2,
		2, 45, 45, 2, 0, 23, 24, 46, 48, 1, 0, 46, 47, 1, 0, 17, 21, 238, 0, 64,
		1, 0, 0, 0, 2, 67, 1, 0, 0, 0, 4, 82, 1, 0, 0, 0, 6, 84, 1, 0, 0, 0, 8,
		94, 1, 0, 0, 0, 10, 110, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14, 116, 1, 0,
		0, 0, 16, 124, 1, 0, 0, 0, 18, 126, 1, 0, 0, 0, 20, 133, 1, 0, 0, 0, 22,
		137, 1, 0, 0, 0, 24, 142, 1, 0, 0, 0, 26, 146, 1, 0, 0, 0, 28, 150, 1,
		0, 0, 0, 30, 153, 1, 0, 0, 0, 32, 157, 1, 0, 0, 0, 34, 166, 1, 0, 0, 0,
		36, 170, 1, 0, 0, 0, 38, 174, 1, 0, 0, 0, 40, 179, 1, 0, 0, 0, 42, 181,
		1, 0, 0, 0, 44, 188, 1, 0, 0, 0, 46, 196, 1, 0, 0, 0, 48, 204, 1, 0, 0,
		0, 50, 220, 1, 0, 0, 0, 52, 222, 1, 0, 0, 0, 54, 235, 1, 0, 0, 0, 56, 237,
		1, 0, 0, 0, 58, 239, 1, 0, 0, 0, 60, 241, 1, 0, 0, 0, 62, 243, 1, 0, 0,
		0, 64, 65, 3, 2, 1, 0, 65, 66, 5, 0, 0, 1, 66, 1, 1, 0, 0, 0, 67, 68, 5,
		1, 0, 0, 68, 69, 5, 45, 0, 0, 69, 70, 3, 62, 31, 0, 70, 71, 5, 48, 0, 0,
		71, 75, 5, 25, 0, 0, 72, 74, 3, 4, 2, 0, 73, 72, 1, 0, 0, 0, 74, 77, 1,
		0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77,
		75, 1, 0, 0, 0, 78, 79, 5, 26, 0, 0, 79, 3, 1, 0, 0, 0, 80, 83, 3, 6, 3,
		0, 81, 83, 3, 16, 8, 0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 5,
		1, 0, 0, 0, 84, 85, 5, 2, 0, 0, 85, 89, 5, 25, 0, 0, 86, 88, 3, 8, 4, 0,
		87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1,
		0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 26, 0, 0, 93,
		7, 1, 0, 0, 0, 94, 95, 5, 45, 0, 0, 95, 96, 3, 50, 25, 0, 96, 97, 3, 10,
		5, 0, 97, 9, 1, 0, 0, 0, 98, 111, 3, 12, 6, 0, 99, 100, 5, 27, 0, 0, 100,
		105, 3, 12, 6, 0, 101, 102, 5, 41, 0, 0, 102, 104, 3, 12, 6, 0, 103, 101,
		1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0,
		0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109, 5, 28, 0, 0,
		109, 111, 1, 0, 0, 0, 110, 98, 1, 0, 0, 0, 110, 99, 1, 0, 0, 0, 111, 11,
		1, 0, 0, 0, 112, 114, 3, 14, 7, 0, 113, 115, 5, 44, 0, 0, 114, 113, 1,
		0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 13, 1, 0, 0, 0, 116, 117, 7, 0, 0,
		0, 117, 15, 1, 0, 0, 0, 118, 125, 3, 18, 9, 0, 119, 125, 3, 20, 10, 0,
		120, 125, 3, 22, 11, 0, 121, 125, 3, 24, 12, 0, 122, 125, 3, 26, 13, 0,
		123, 125, 3, 28, 14, 0, 124, 118, 1, 0, 0, 0, 124, 119, 1, 0, 0, 0, 124,
		120, 1, 0, 0, 0, 124, 121, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 123,
		1, 0, 0, 0, 125, 17, 1, 0, 0, 0, 126, 127, 5, 3, 0, 0, 127, 128, 3, 30,
		15, 0, 128, 129, 5, 4, 0, 0, 129, 130, 5, 45, 0, 0, 130, 131, 5, 5, 0,
		0, 131, 132, 3, 36, 18, 0, 132, 19, 1, 0, 0, 0, 133, 134, 5, 6, 0, 0, 134,
		135, 3, 30, 15, 0, 135, 136, 3, 32, 16, 0, 136, 21, 1, 0, 0, 0, 137, 138,
		5, 7, 0, 0, 138, 139, 5, 45, 0, 0, 139, 140, 5, 8, 0, 0, 140, 141, 5, 45,
		0, 0, 141, 23, 1, 0, 0, 0, 142, 143, 5, 9, 0, 0, 143, 144, 5, 45, 0, 0,
		144, 145, 3, 32, 16, 0, 145, 25, 1, 0, 0, 0, 146, 147, 5, 10, 0, 0, 147,
		148, 3, 30, 15, 0, 148, 149, 3, 32, 16, 0, 149, 27, 1, 0, 0, 0, 150, 151,
		5, 11, 0, 0, 151, 152, 3, 40, 20, 0, 152, 29, 1, 0, 0, 0, 153, 154, 5,
		45, 0, 0, 154, 155, 5, 38, 0, 0, 155, 156, 3, 50, 25, 0, 156, 31, 1, 0,
		0, 0, 157, 161, 5, 25, 0, 0, 158, 160, 3, 34, 17, 0, 159, 158, 1, 0, 0,
		0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162,
		164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165, 5, 26, 0, 0, 165, 33,
		1, 0, 0, 0, 166, 167, 5, 45, 0, 0, 167, 168, 5, 37, 0, 0, 168, 169, 3,
		40, 20, 0, 169, 35, 1, 0, 0, 0, 170, 171, 3, 40, 20, 0, 171, 172, 3, 38,
		19, 0, 172, 173, 3, 40, 20, 0, 173, 37, 1, 0, 0, 0, 174, 175, 7, 1, 0,
		0, 175, 39, 1, 0, 0, 0, 176, 180, 3, 42, 21, 0, 177, 180, 3, 46, 23, 0,
		178, 180, 3, 58, 29, 0, 179, 176, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179,
		178, 1, 0, 0, 0, 180, 41, 1, 0, 0, 0, 181, 182, 5, 45, 0, 0, 182, 184,
		5, 27, 0, 0, 183, 185, 3, 44, 22, 0, 184, 183, 1, 0, 0, 0, 184, 185, 1,
		0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 5, 28, 0, 0, 187, 43, 1, 0, 0,
		0, 188, 193, 3, 40, 20, 0, 189, 190, 5, 40, 0, 0, 190, 192, 3, 40, 20,
		0, 191, 189, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193,
		194, 1, 0, 0, 0, 194, 45, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 201, 3,
		48, 24, 0, 197, 198, 5, 39, 0, 0, 198, 200, 3, 48, 24, 0, 199, 197, 1,
		0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0,
		0, 202, 47, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 205, 7, 2, 0, 0, 205,
		49, 1, 0, 0, 0, 206, 208, 3, 52, 26, 0, 207, 209, 3, 54, 27, 0, 208, 207,
		1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0, 210, 212, 3, 56,
		28, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 221, 1, 0, 0, 0,
		213, 214, 5, 22, 0, 0, 214, 215, 5, 31, 0, 0, 215, 216, 3, 52, 26, 0, 216,
		218, 5, 32, 0, 0, 217, 219, 3, 56, 28, 0, 218, 217, 1, 0, 0, 0, 218, 219,
		1, 0, 0, 0, 219, 221, 1, 0, 0, 0, 220, 206, 1, 0, 0, 0, 220, 213, 1, 0,
		0, 0, 221, 51, 1, 0, 0, 0, 222, 223, 5, 45, 0, 0, 223, 53, 1, 0, 0, 0,
		224, 236, 5, 42, 0, 0, 225, 236, 5, 43, 0, 0, 226, 228, 5, 29, 0, 0, 227,
		229, 3, 60, 30, 0, 228, 227, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 230,
		1, 0, 0, 0, 230, 232, 5, 40, 0, 0, 231, 233, 3, 60, 30, 0, 232, 231, 1,
		0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 236, 5, 30, 0,
		0, 235, 224, 1, 0, 0, 0, 235, 225, 1, 0, 0, 0, 235, 226, 1, 0, 0, 0, 236,
		55, 1, 0, 0, 0, 237, 238, 5, 44, 0, 0, 238, 57, 1, 0, 0, 0, 239, 240, 7,
		3, 0, 0, 240, 59, 1, 0, 0, 0, 241, 242, 7, 4, 0, 0, 242, 61, 1, 0, 0, 0,
		243, 244, 7, 5, 0, 0, 244, 63, 1, 0, 0, 0, 19, 75, 82, 89, 105, 110, 114,
		124, 161, 179, 184, 193, 201, 208, 211, 218, 220, 228, 232, 235,
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

// EndpointParserInit initializes any static state used to implement EndpointParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEndpointParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EndpointParserInit() {
	staticData := &EndpointParserStaticData
	staticData.once.Do(endpointParserInit)
}

// NewEndpointParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEndpointParser(input antlr.TokenStream) *EndpointParser {
	EndpointParserInit()
	this := new(EndpointParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EndpointParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Endpoint.g4"

	return this
}

// EndpointParser tokens.
const (
	EndpointParserEOF           = antlr.TokenEOF
	EndpointParserENDPOINT      = 1
	EndpointParserINPUT         = 2
	EndpointParserFETCH         = 3
	EndpointParserFROM          = 4
	EndpointParserWHERE         = 5
	EndpointParserCREATE        = 6
	EndpointParserSTORE         = 7
	EndpointParserTO            = 8
	EndpointParserUPDATE        = 9
	EndpointParserEMIT          = 10
	EndpointParserRETURN        = 11
	EndpointParserPATH          = 12
	EndpointParserQUERY         = 13
	EndpointParserBODY          = 14
	EndpointParserHEADER        = 15
	EndpointParserCOOKIE        = 16
	EndpointParserGET           = 17
	EndpointParserPOST          = 18
	EndpointParserPUT           = 19
	EndpointParserPATCH         = 20
	EndpointParserDELETE        = 21
	EndpointParserLIST          = 22
	EndpointParserTRUE          = 23
	EndpointParserFALSE         = 24
	EndpointParserLBRACE        = 25
	EndpointParserRBRACE        = 26
	EndpointParserLPAREN        = 27
	EndpointParserRPAREN        = 28
	EndpointParserLBRACKET      = 29
	EndpointParserRBRACKET      = 30
	EndpointParserLT            = 31
	EndpointParserGT            = 32
	EndpointParserLTE           = 33
	EndpointParserGTE           = 34
	EndpointParserEQ            = 35
	EndpointParserNEQ           = 36
	EndpointParserASSIGN        = 37
	EndpointParserCOLON         = 38
	EndpointParserDOT           = 39
	EndpointParserCOMMA         = 40
	EndpointParserPIPE          = 41
	EndpointParserPLUS          = 42
	EndpointParserSTAR          = 43
	EndpointParserQUESTION      = 44
	EndpointParserIDENT         = 45
	EndpointParserFLOAT         = 46
	EndpointParserINT           = 47
	EndpointParserSTRING        = 48
	EndpointParserLINE_COMMENT  = 49
	EndpointParserBLOCK_COMMENT = 50
	EndpointParserWS            = 51
)

// EndpointParser rules.
const (
	EndpointParserRULE_program          = 0
	EndpointParserRULE_endpointDecl     = 1
	EndpointParserRULE_endpointBody     = 2
	EndpointParserRULE_inputBlock       = 3
	EndpointParserRULE_inputField       = 4
	EndpointParserRULE_inputSourceRule  = 5
	EndpointParserRULE_inputSource      = 6
	EndpointParserRULE_inputSourceKind  = 7
	EndpointParserRULE_action           = 8
	EndpointParserRULE_fetchAction      = 9
	EndpointParserRULE_createAction     = 10
	EndpointParserRULE_storeAction      = 11
	EndpointParserRULE_updateAction     = 12
	EndpointParserRULE_emitAction       = 13
	EndpointParserRULE_returnAction     = 14
	EndpointParserRULE_typedVariable    = 15
	EndpointParserRULE_assignmentBlock  = 16
	EndpointParserRULE_assignment       = 17
	EndpointParserRULE_condition        = 18
	EndpointParserRULE_comparator       = 19
	EndpointParserRULE_expression       = 20
	EndpointParserRULE_functionCall     = 21
	EndpointParserRULE_argumentList     = 22
	EndpointParserRULE_selector         = 23
	EndpointParserRULE_identifier       = 24
	EndpointParserRULE_typeRef          = 25
	EndpointParserRULE_typeName         = 26
	EndpointParserRULE_numberConstraint = 27
	EndpointParserRULE_optionalMarker   = 28
	EndpointParserRULE_value            = 29
	EndpointParserRULE_numberValue      = 30
	EndpointParserRULE_httpMethod       = 31
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EndpointDecl() IEndpointDeclContext
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
	p.RuleIndex = EndpointParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EndpointDecl() IEndpointDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndpointDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(EndpointParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EndpointParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.EndpointDecl()
	}
	{
		p.SetState(65)
		p.Match(EndpointParserEOF)
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

// IEndpointDeclContext is an interface to support dynamic dispatch.
type IEndpointDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENDPOINT() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	HttpMethod() IHttpMethodContext
	STRING() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllEndpointBody() []IEndpointBodyContext
	EndpointBody(i int) IEndpointBodyContext

	// IsEndpointDeclContext differentiates from other interfaces.
	IsEndpointDeclContext()
}

type EndpointDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointDeclContext() *EndpointDeclContext {
	var p = new(EndpointDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_endpointDecl
	return p
}

func InitEmptyEndpointDeclContext(p *EndpointDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_endpointDecl
}

func (*EndpointDeclContext) IsEndpointDeclContext() {}

func NewEndpointDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointDeclContext {
	var p = new(EndpointDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_endpointDecl

	return p
}

func (s *EndpointDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointDeclContext) ENDPOINT() antlr.TerminalNode {
	return s.GetToken(EndpointParserENDPOINT, 0)
}

func (s *EndpointDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *EndpointDeclContext) HttpMethod() IHttpMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpMethodContext)
}

func (s *EndpointDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(EndpointParserSTRING, 0)
}

func (s *EndpointDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(EndpointParserLBRACE, 0)
}

func (s *EndpointDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(EndpointParserRBRACE, 0)
}

func (s *EndpointDeclContext) AllEndpointBody() []IEndpointBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndpointBodyContext); ok {
			len++
		}
	}

	tst := make([]IEndpointBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndpointBodyContext); ok {
			tst[i] = t.(IEndpointBodyContext)
			i++
		}
	}

	return tst
}

func (s *EndpointDeclContext) EndpointBody(i int) IEndpointBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointBodyContext); ok {
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

	return t.(IEndpointBodyContext)
}

func (s *EndpointDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitEndpointDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) EndpointDecl() (localctx IEndpointDeclContext) {
	localctx = NewEndpointDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EndpointParserRULE_endpointDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(EndpointParserENDPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.HttpMethod()
	}
	{
		p.SetState(70)
		p.Match(EndpointParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(EndpointParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3788) != 0 {
		{
			p.SetState(72)
			p.EndpointBody()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(EndpointParserRBRACE)
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

// IEndpointBodyContext is an interface to support dynamic dispatch.
type IEndpointBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InputBlock() IInputBlockContext
	Action_() IActionContext

	// IsEndpointBodyContext differentiates from other interfaces.
	IsEndpointBodyContext()
}

type EndpointBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointBodyContext() *EndpointBodyContext {
	var p = new(EndpointBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_endpointBody
	return p
}

func InitEmptyEndpointBodyContext(p *EndpointBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_endpointBody
}

func (*EndpointBodyContext) IsEndpointBodyContext() {}

func NewEndpointBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointBodyContext {
	var p = new(EndpointBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_endpointBody

	return p
}

func (s *EndpointBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointBodyContext) InputBlock() IInputBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputBlockContext)
}

func (s *EndpointBodyContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *EndpointBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitEndpointBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) EndpointBody() (localctx IEndpointBodyContext) {
	localctx = NewEndpointBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EndpointParserRULE_endpointBody)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EndpointParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.InputBlock()
		}

	case EndpointParserFETCH, EndpointParserCREATE, EndpointParserSTORE, EndpointParserUPDATE, EndpointParserEMIT, EndpointParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Action_()
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

// IInputBlockContext is an interface to support dynamic dispatch.
type IInputBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INPUT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllInputField() []IInputFieldContext
	InputField(i int) IInputFieldContext

	// IsInputBlockContext differentiates from other interfaces.
	IsInputBlockContext()
}

type InputBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputBlockContext() *InputBlockContext {
	var p = new(InputBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputBlock
	return p
}

func InitEmptyInputBlockContext(p *InputBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputBlock
}

func (*InputBlockContext) IsInputBlockContext() {}

func NewInputBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputBlockContext {
	var p = new(InputBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_inputBlock

	return p
}

func (s *InputBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *InputBlockContext) INPUT() antlr.TerminalNode {
	return s.GetToken(EndpointParserINPUT, 0)
}

func (s *InputBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(EndpointParserLBRACE, 0)
}

func (s *InputBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(EndpointParserRBRACE, 0)
}

func (s *InputBlockContext) AllInputField() []IInputFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputFieldContext); ok {
			len++
		}
	}

	tst := make([]IInputFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputFieldContext); ok {
			tst[i] = t.(IInputFieldContext)
			i++
		}
	}

	return tst
}

func (s *InputBlockContext) InputField(i int) IInputFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputFieldContext); ok {
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

	return t.(IInputFieldContext)
}

func (s *InputBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitInputBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) InputBlock() (localctx IInputBlockContext) {
	localctx = NewInputBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EndpointParserRULE_inputBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(EndpointParserINPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(EndpointParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EndpointParserIDENT {
		{
			p.SetState(86)
			p.InputField()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(EndpointParserRBRACE)
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

// IInputFieldContext is an interface to support dynamic dispatch.
type IInputFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	TypeRef() ITypeRefContext
	InputSourceRule() IInputSourceRuleContext

	// IsInputFieldContext differentiates from other interfaces.
	IsInputFieldContext()
}

type InputFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputFieldContext() *InputFieldContext {
	var p = new(InputFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputField
	return p
}

func InitEmptyInputFieldContext(p *InputFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputField
}

func (*InputFieldContext) IsInputFieldContext() {}

func NewInputFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputFieldContext {
	var p = new(InputFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_inputField

	return p
}

func (s *InputFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *InputFieldContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *InputFieldContext) TypeRef() ITypeRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *InputFieldContext) InputSourceRule() IInputSourceRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputSourceRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputSourceRuleContext)
}

func (s *InputFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitInputField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) InputField() (localctx IInputFieldContext) {
	localctx = NewInputFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EndpointParserRULE_inputField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.TypeRef()
	}
	{
		p.SetState(96)
		p.InputSourceRule()
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

// IInputSourceRuleContext is an interface to support dynamic dispatch.
type IInputSourceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInputSource() []IInputSourceContext
	InputSource(i int) IInputSourceContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode

	// IsInputSourceRuleContext differentiates from other interfaces.
	IsInputSourceRuleContext()
}

type InputSourceRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputSourceRuleContext() *InputSourceRuleContext {
	var p = new(InputSourceRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputSourceRule
	return p
}

func InitEmptyInputSourceRuleContext(p *InputSourceRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputSourceRule
}

func (*InputSourceRuleContext) IsInputSourceRuleContext() {}

func NewInputSourceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputSourceRuleContext {
	var p = new(InputSourceRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_inputSourceRule

	return p
}

func (s *InputSourceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *InputSourceRuleContext) AllInputSource() []IInputSourceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputSourceContext); ok {
			len++
		}
	}

	tst := make([]IInputSourceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputSourceContext); ok {
			tst[i] = t.(IInputSourceContext)
			i++
		}
	}

	return tst
}

func (s *InputSourceRuleContext) InputSource(i int) IInputSourceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputSourceContext); ok {
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

	return t.(IInputSourceContext)
}

func (s *InputSourceRuleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EndpointParserLPAREN, 0)
}

func (s *InputSourceRuleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EndpointParserRPAREN, 0)
}

func (s *InputSourceRuleContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(EndpointParserPIPE)
}

func (s *InputSourceRuleContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(EndpointParserPIPE, i)
}

func (s *InputSourceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputSourceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputSourceRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitInputSourceRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) InputSourceRule() (localctx IInputSourceRuleContext) {
	localctx = NewInputSourceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EndpointParserRULE_inputSourceRule)
	var _la int

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EndpointParserPATH, EndpointParserQUERY, EndpointParserBODY, EndpointParserHEADER, EndpointParserCOOKIE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.InputSource()
		}

	case EndpointParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(EndpointParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.InputSource()
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EndpointParserPIPE {
			{
				p.SetState(101)
				p.Match(EndpointParserPIPE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(102)
				p.InputSource()
			}

			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(108)
			p.Match(EndpointParserRPAREN)
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

// IInputSourceContext is an interface to support dynamic dispatch.
type IInputSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InputSourceKind() IInputSourceKindContext
	QUESTION() antlr.TerminalNode

	// IsInputSourceContext differentiates from other interfaces.
	IsInputSourceContext()
}

type InputSourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputSourceContext() *InputSourceContext {
	var p = new(InputSourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputSource
	return p
}

func InitEmptyInputSourceContext(p *InputSourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputSource
}

func (*InputSourceContext) IsInputSourceContext() {}

func NewInputSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputSourceContext {
	var p = new(InputSourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_inputSource

	return p
}

func (s *InputSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *InputSourceContext) InputSourceKind() IInputSourceKindContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputSourceKindContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputSourceKindContext)
}

func (s *InputSourceContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(EndpointParserQUESTION, 0)
}

func (s *InputSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputSourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitInputSource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) InputSource() (localctx IInputSourceContext) {
	localctx = NewInputSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EndpointParserRULE_inputSource)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.InputSourceKind()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EndpointParserQUESTION {
		{
			p.SetState(113)
			p.Match(EndpointParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IInputSourceKindContext is an interface to support dynamic dispatch.
type IInputSourceKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PATH() antlr.TerminalNode
	QUERY() antlr.TerminalNode
	BODY() antlr.TerminalNode
	HEADER() antlr.TerminalNode
	COOKIE() antlr.TerminalNode

	// IsInputSourceKindContext differentiates from other interfaces.
	IsInputSourceKindContext()
}

type InputSourceKindContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputSourceKindContext() *InputSourceKindContext {
	var p = new(InputSourceKindContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputSourceKind
	return p
}

func InitEmptyInputSourceKindContext(p *InputSourceKindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_inputSourceKind
}

func (*InputSourceKindContext) IsInputSourceKindContext() {}

func NewInputSourceKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputSourceKindContext {
	var p = new(InputSourceKindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_inputSourceKind

	return p
}

func (s *InputSourceKindContext) GetParser() antlr.Parser { return s.parser }

func (s *InputSourceKindContext) PATH() antlr.TerminalNode {
	return s.GetToken(EndpointParserPATH, 0)
}

func (s *InputSourceKindContext) QUERY() antlr.TerminalNode {
	return s.GetToken(EndpointParserQUERY, 0)
}

func (s *InputSourceKindContext) BODY() antlr.TerminalNode {
	return s.GetToken(EndpointParserBODY, 0)
}

func (s *InputSourceKindContext) HEADER() antlr.TerminalNode {
	return s.GetToken(EndpointParserHEADER, 0)
}

func (s *InputSourceKindContext) COOKIE() antlr.TerminalNode {
	return s.GetToken(EndpointParserCOOKIE, 0)
}

func (s *InputSourceKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputSourceKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputSourceKindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitInputSourceKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) InputSourceKind() (localctx IInputSourceKindContext) {
	localctx = NewInputSourceKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EndpointParserRULE_inputSourceKind)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126976) != 0) {
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

// IActionContext is an interface to support dynamic dispatch.
type IActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FetchAction() IFetchActionContext
	CreateAction() ICreateActionContext
	StoreAction() IStoreActionContext
	UpdateAction() IUpdateActionContext
	EmitAction() IEmitActionContext
	ReturnAction() IReturnActionContext

	// IsActionContext differentiates from other interfaces.
	IsActionContext()
}

type ActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionContext() *ActionContext {
	var p = new(ActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_action

	return p
}

func (s *ActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionContext) FetchAction() IFetchActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFetchActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFetchActionContext)
}

func (s *ActionContext) CreateAction() ICreateActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateActionContext)
}

func (s *ActionContext) StoreAction() IStoreActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStoreActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStoreActionContext)
}

func (s *ActionContext) UpdateAction() IUpdateActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateActionContext)
}

func (s *ActionContext) EmitAction() IEmitActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmitActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmitActionContext)
}

func (s *ActionContext) ReturnAction() IReturnActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnActionContext)
}

func (s *ActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EndpointParserRULE_action)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EndpointParserFETCH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.FetchAction()
		}

	case EndpointParserCREATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.CreateAction()
		}

	case EndpointParserSTORE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.StoreAction()
		}

	case EndpointParserUPDATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)
			p.UpdateAction()
		}

	case EndpointParserEMIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(122)
			p.EmitAction()
		}

	case EndpointParserRETURN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(123)
			p.ReturnAction()
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

// IFetchActionContext is an interface to support dynamic dispatch.
type IFetchActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FETCH() antlr.TerminalNode
	TypedVariable() ITypedVariableContext
	FROM() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	Condition() IConditionContext

	// IsFetchActionContext differentiates from other interfaces.
	IsFetchActionContext()
}

type FetchActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFetchActionContext() *FetchActionContext {
	var p = new(FetchActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_fetchAction
	return p
}

func InitEmptyFetchActionContext(p *FetchActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_fetchAction
}

func (*FetchActionContext) IsFetchActionContext() {}

func NewFetchActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FetchActionContext {
	var p = new(FetchActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_fetchAction

	return p
}

func (s *FetchActionContext) GetParser() antlr.Parser { return s.parser }

func (s *FetchActionContext) FETCH() antlr.TerminalNode {
	return s.GetToken(EndpointParserFETCH, 0)
}

func (s *FetchActionContext) TypedVariable() ITypedVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedVariableContext)
}

func (s *FetchActionContext) FROM() antlr.TerminalNode {
	return s.GetToken(EndpointParserFROM, 0)
}

func (s *FetchActionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *FetchActionContext) WHERE() antlr.TerminalNode {
	return s.GetToken(EndpointParserWHERE, 0)
}

func (s *FetchActionContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *FetchActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FetchActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FetchActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitFetchAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) FetchAction() (localctx IFetchActionContext) {
	localctx = NewFetchActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EndpointParserRULE_fetchAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(EndpointParserFETCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.TypedVariable()
	}
	{
		p.SetState(128)
		p.Match(EndpointParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(EndpointParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Condition()
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

// ICreateActionContext is an interface to support dynamic dispatch.
type ICreateActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	TypedVariable() ITypedVariableContext
	AssignmentBlock() IAssignmentBlockContext

	// IsCreateActionContext differentiates from other interfaces.
	IsCreateActionContext()
}

type CreateActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateActionContext() *CreateActionContext {
	var p = new(CreateActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_createAction
	return p
}

func InitEmptyCreateActionContext(p *CreateActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_createAction
}

func (*CreateActionContext) IsCreateActionContext() {}

func NewCreateActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateActionContext {
	var p = new(CreateActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_createAction

	return p
}

func (s *CreateActionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateActionContext) CREATE() antlr.TerminalNode {
	return s.GetToken(EndpointParserCREATE, 0)
}

func (s *CreateActionContext) TypedVariable() ITypedVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedVariableContext)
}

func (s *CreateActionContext) AssignmentBlock() IAssignmentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentBlockContext)
}

func (s *CreateActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitCreateAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) CreateAction() (localctx ICreateActionContext) {
	localctx = NewCreateActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, EndpointParserRULE_createAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(EndpointParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.TypedVariable()
	}
	{
		p.SetState(135)
		p.AssignmentBlock()
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

// IStoreActionContext is an interface to support dynamic dispatch.
type IStoreActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STORE() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	TO() antlr.TerminalNode

	// IsStoreActionContext differentiates from other interfaces.
	IsStoreActionContext()
}

type StoreActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStoreActionContext() *StoreActionContext {
	var p = new(StoreActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_storeAction
	return p
}

func InitEmptyStoreActionContext(p *StoreActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_storeAction
}

func (*StoreActionContext) IsStoreActionContext() {}

func NewStoreActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreActionContext {
	var p = new(StoreActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_storeAction

	return p
}

func (s *StoreActionContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreActionContext) STORE() antlr.TerminalNode {
	return s.GetToken(EndpointParserSTORE, 0)
}

func (s *StoreActionContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(EndpointParserIDENT)
}

func (s *StoreActionContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, i)
}

func (s *StoreActionContext) TO() antlr.TerminalNode {
	return s.GetToken(EndpointParserTO, 0)
}

func (s *StoreActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitStoreAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) StoreAction() (localctx IStoreActionContext) {
	localctx = NewStoreActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, EndpointParserRULE_storeAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(EndpointParserSTORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(EndpointParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(EndpointParserIDENT)
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

// IUpdateActionContext is an interface to support dynamic dispatch.
type IUpdateActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UPDATE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	AssignmentBlock() IAssignmentBlockContext

	// IsUpdateActionContext differentiates from other interfaces.
	IsUpdateActionContext()
}

type UpdateActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateActionContext() *UpdateActionContext {
	var p = new(UpdateActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_updateAction
	return p
}

func InitEmptyUpdateActionContext(p *UpdateActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_updateAction
}

func (*UpdateActionContext) IsUpdateActionContext() {}

func NewUpdateActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateActionContext {
	var p = new(UpdateActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_updateAction

	return p
}

func (s *UpdateActionContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateActionContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(EndpointParserUPDATE, 0)
}

func (s *UpdateActionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *UpdateActionContext) AssignmentBlock() IAssignmentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentBlockContext)
}

func (s *UpdateActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitUpdateAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) UpdateAction() (localctx IUpdateActionContext) {
	localctx = NewUpdateActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, EndpointParserRULE_updateAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(EndpointParserUPDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.AssignmentBlock()
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

// IEmitActionContext is an interface to support dynamic dispatch.
type IEmitActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EMIT() antlr.TerminalNode
	TypedVariable() ITypedVariableContext
	AssignmentBlock() IAssignmentBlockContext

	// IsEmitActionContext differentiates from other interfaces.
	IsEmitActionContext()
}

type EmitActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmitActionContext() *EmitActionContext {
	var p = new(EmitActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_emitAction
	return p
}

func InitEmptyEmitActionContext(p *EmitActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_emitAction
}

func (*EmitActionContext) IsEmitActionContext() {}

func NewEmitActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmitActionContext {
	var p = new(EmitActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_emitAction

	return p
}

func (s *EmitActionContext) GetParser() antlr.Parser { return s.parser }

func (s *EmitActionContext) EMIT() antlr.TerminalNode {
	return s.GetToken(EndpointParserEMIT, 0)
}

func (s *EmitActionContext) TypedVariable() ITypedVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedVariableContext)
}

func (s *EmitActionContext) AssignmentBlock() IAssignmentBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentBlockContext)
}

func (s *EmitActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmitActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmitActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitEmitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) EmitAction() (localctx IEmitActionContext) {
	localctx = NewEmitActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, EndpointParserRULE_emitAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(EndpointParserEMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.TypedVariable()
	}
	{
		p.SetState(148)
		p.AssignmentBlock()
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

// IReturnActionContext is an interface to support dynamic dispatch.
type IReturnActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnActionContext differentiates from other interfaces.
	IsReturnActionContext()
}

type ReturnActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnActionContext() *ReturnActionContext {
	var p = new(ReturnActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_returnAction
	return p
}

func InitEmptyReturnActionContext(p *ReturnActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_returnAction
}

func (*ReturnActionContext) IsReturnActionContext() {}

func NewReturnActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnActionContext {
	var p = new(ReturnActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_returnAction

	return p
}

func (s *ReturnActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnActionContext) RETURN() antlr.TerminalNode {
	return s.GetToken(EndpointParserRETURN, 0)
}

func (s *ReturnActionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitReturnAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) ReturnAction() (localctx IReturnActionContext) {
	localctx = NewReturnActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, EndpointParserRULE_returnAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(EndpointParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Expression()
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

// ITypedVariableContext is an interface to support dynamic dispatch.
type ITypedVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeRef() ITypeRefContext

	// IsTypedVariableContext differentiates from other interfaces.
	IsTypedVariableContext()
}

type TypedVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedVariableContext() *TypedVariableContext {
	var p = new(TypedVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_typedVariable
	return p
}

func InitEmptyTypedVariableContext(p *TypedVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_typedVariable
}

func (*TypedVariableContext) IsTypedVariableContext() {}

func NewTypedVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedVariableContext {
	var p = new(TypedVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_typedVariable

	return p
}

func (s *TypedVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedVariableContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *TypedVariableContext) COLON() antlr.TerminalNode {
	return s.GetToken(EndpointParserCOLON, 0)
}

func (s *TypedVariableContext) TypeRef() ITypeRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *TypedVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitTypedVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) TypedVariable() (localctx ITypedVariableContext) {
	localctx = NewTypedVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, EndpointParserRULE_typedVariable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(EndpointParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.TypeRef()
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

// IAssignmentBlockContext is an interface to support dynamic dispatch.
type IAssignmentBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext

	// IsAssignmentBlockContext differentiates from other interfaces.
	IsAssignmentBlockContext()
}

type AssignmentBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentBlockContext() *AssignmentBlockContext {
	var p = new(AssignmentBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_assignmentBlock
	return p
}

func InitEmptyAssignmentBlockContext(p *AssignmentBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_assignmentBlock
}

func (*AssignmentBlockContext) IsAssignmentBlockContext() {}

func NewAssignmentBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentBlockContext {
	var p = new(AssignmentBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_assignmentBlock

	return p
}

func (s *AssignmentBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(EndpointParserLBRACE, 0)
}

func (s *AssignmentBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(EndpointParserRBRACE, 0)
}

func (s *AssignmentBlockContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentBlockContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *AssignmentBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitAssignmentBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) AssignmentBlock() (localctx IAssignmentBlockContext) {
	localctx = NewAssignmentBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, EndpointParserRULE_assignmentBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(EndpointParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EndpointParserIDENT {
		{
			p.SetState(158)
			p.Assignment()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
		p.Match(EndpointParserRBRACE)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(EndpointParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, EndpointParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(EndpointParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Expression()
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Comparator() IComparatorContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ConditionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, EndpointParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Expression()
	}
	{
		p.SetState(171)
		p.Comparator()
	}
	{
		p.SetState(172)
		p.Expression()
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

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(EndpointParserEQ, 0)
}

func (s *ComparatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(EndpointParserNEQ, 0)
}

func (s *ComparatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(EndpointParserLTE, 0)
}

func (s *ComparatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(EndpointParserGTE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(EndpointParserLT, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(EndpointParserGT, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, EndpointParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291469824) != 0) {
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	Selector() ISelectorContext
	Value() IValueContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *ExpressionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, EndpointParserRULE_expression)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Selector()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
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

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EndpointParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EndpointParserRPAREN, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, EndpointParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(EndpointParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(EndpointParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&527765606498308) != 0 {
		{
			p.SetState(183)
			p.ArgumentList()
		}

	}
	{
		p.SetState(186)
		p.Match(EndpointParserRPAREN)
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

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EndpointParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EndpointParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, EndpointParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Expression()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EndpointParserCOMMA {
		{
			p.SetState(189)
			p.Match(EndpointParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Expression()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *SelectorContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *SelectorContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(EndpointParserDOT)
}

func (s *SelectorContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(EndpointParserDOT, i)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, EndpointParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Identifier()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EndpointParserDOT {
		{
			p.SetState(197)
			p.Match(EndpointParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Identifier()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = EndpointParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *IdentifierContext) INPUT() antlr.TerminalNode {
	return s.GetToken(EndpointParserINPUT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, EndpointParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EndpointParserINPUT || _la == EndpointParserIDENT) {
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
	p.RuleIndex = EndpointParserRULE_typeRef
	return p
}

func InitEmptyTypeRefContext(p *TypeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_typeRef
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_typeRef

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
	return s.GetToken(EndpointParserLIST, 0)
}

func (s *TypeRefContext) LT() antlr.TerminalNode {
	return s.GetToken(EndpointParserLT, 0)
}

func (s *TypeRefContext) GT() antlr.TerminalNode {
	return s.GetToken(EndpointParserGT, 0)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitTypeRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, EndpointParserRULE_typeRef)
	var _la int

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EndpointParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.TypeName()
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13194676404224) != 0 {
			{
				p.SetState(207)
				p.NumberConstraint()
			}

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EndpointParserQUESTION {
			{
				p.SetState(210)
				p.OptionalMarker()
			}

		}

	case EndpointParserLIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Match(EndpointParserLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Match(EndpointParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.TypeName()
		}
		{
			p.SetState(216)
			p.Match(EndpointParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EndpointParserQUESTION {
			{
				p.SetState(217)
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
	p.RuleIndex = EndpointParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EndpointParserIDENT, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, EndpointParserRULE_typeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(EndpointParserIDENT)
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
	p.RuleIndex = EndpointParserRULE_numberConstraint
	return p
}

func InitEmptyNumberConstraintContext(p *NumberConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_numberConstraint
}

func (*NumberConstraintContext) IsNumberConstraintContext() {}

func NewNumberConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberConstraintContext {
	var p = new(NumberConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_numberConstraint

	return p
}

func (s *NumberConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberConstraintContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EndpointParserPLUS, 0)
}

func (s *NumberConstraintContext) STAR() antlr.TerminalNode {
	return s.GetToken(EndpointParserSTAR, 0)
}

func (s *NumberConstraintContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(EndpointParserLBRACKET, 0)
}

func (s *NumberConstraintContext) COMMA() antlr.TerminalNode {
	return s.GetToken(EndpointParserCOMMA, 0)
}

func (s *NumberConstraintContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(EndpointParserRBRACKET, 0)
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
	case EndpointVisitor:
		return t.VisitNumberConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) NumberConstraint() (localctx INumberConstraintContext) {
	localctx = NewNumberConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, EndpointParserRULE_numberConstraint)
	var _la int

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EndpointParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.Match(EndpointParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EndpointParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.Match(EndpointParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EndpointParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)
			p.Match(EndpointParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EndpointParserFLOAT || _la == EndpointParserINT {
			{
				p.SetState(227)
				p.NumberValue()
			}

		}
		{
			p.SetState(230)
			p.Match(EndpointParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EndpointParserFLOAT || _la == EndpointParserINT {
			{
				p.SetState(231)
				p.NumberValue()
			}

		}
		{
			p.SetState(234)
			p.Match(EndpointParserRBRACKET)
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
	p.RuleIndex = EndpointParserRULE_optionalMarker
	return p
}

func InitEmptyOptionalMarkerContext(p *OptionalMarkerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_optionalMarker
}

func (*OptionalMarkerContext) IsOptionalMarkerContext() {}

func NewOptionalMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalMarkerContext {
	var p = new(OptionalMarkerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_optionalMarker

	return p
}

func (s *OptionalMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalMarkerContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(EndpointParserQUESTION, 0)
}

func (s *OptionalMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalMarkerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitOptionalMarker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) OptionalMarker() (localctx IOptionalMarkerContext) {
	localctx = NewOptionalMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, EndpointParserRULE_optionalMarker)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(EndpointParserQUESTION)
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
	p.RuleIndex = EndpointParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(EndpointParserSTRING, 0)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(EndpointParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(EndpointParserFLOAT, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(EndpointParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(EndpointParserFALSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, EndpointParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&492581234409472) != 0) {
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
	p.RuleIndex = EndpointParserRULE_numberValue
	return p
}

func InitEmptyNumberValueContext(p *NumberValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_numberValue
}

func (*NumberValueContext) IsNumberValueContext() {}

func NewNumberValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberValueContext {
	var p = new(NumberValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_numberValue

	return p
}

func (s *NumberValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberValueContext) INT() antlr.TerminalNode {
	return s.GetToken(EndpointParserINT, 0)
}

func (s *NumberValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(EndpointParserFLOAT, 0)
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitNumberValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) NumberValue() (localctx INumberValueContext) {
	localctx = NewNumberValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, EndpointParserRULE_numberValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EndpointParserFLOAT || _la == EndpointParserINT) {
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

// IHttpMethodContext is an interface to support dynamic dispatch.
type IHttpMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GET() antlr.TerminalNode
	POST() antlr.TerminalNode
	PUT() antlr.TerminalNode
	PATCH() antlr.TerminalNode
	DELETE() antlr.TerminalNode

	// IsHttpMethodContext differentiates from other interfaces.
	IsHttpMethodContext()
}

type HttpMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpMethodContext() *HttpMethodContext {
	var p = new(HttpMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_httpMethod
	return p
}

func InitEmptyHttpMethodContext(p *HttpMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EndpointParserRULE_httpMethod
}

func (*HttpMethodContext) IsHttpMethodContext() {}

func NewHttpMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpMethodContext {
	var p = new(HttpMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EndpointParserRULE_httpMethod

	return p
}

func (s *HttpMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpMethodContext) GET() antlr.TerminalNode {
	return s.GetToken(EndpointParserGET, 0)
}

func (s *HttpMethodContext) POST() antlr.TerminalNode {
	return s.GetToken(EndpointParserPOST, 0)
}

func (s *HttpMethodContext) PUT() antlr.TerminalNode {
	return s.GetToken(EndpointParserPUT, 0)
}

func (s *HttpMethodContext) PATCH() antlr.TerminalNode {
	return s.GetToken(EndpointParserPATCH, 0)
}

func (s *HttpMethodContext) DELETE() antlr.TerminalNode {
	return s.GetToken(EndpointParserDELETE, 0)
}

func (s *HttpMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EndpointVisitor:
		return t.VisitHttpMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EndpointParser) HttpMethod() (localctx IHttpMethodContext) {
	localctx = NewHttpMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, EndpointParserRULE_httpMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4063232) != 0) {
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
