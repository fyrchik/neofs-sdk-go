// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SorterLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var sorterlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sorterlexerLexerInit() {
	staticData := &sorterlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'CASE'", "'WHEN'", "'END'", "", "", "'0'",
	}
	staticData.symbolicNames = []string{
		"", "CASE", "WHEN", "END", "IDENT", "NUMBER1", "ZERO", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"CASE", "WHEN", "END", "IDENT", "Digit", "Nondigit", "NUMBER1", "ZERO",
		"STRING", "ESC", "UNICODE", "HEX", "SAFECODEPOINTSINGLE", "SAFECODEPOINTDOUBLE",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 110, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 49, 8, 3, 10, 3, 12, 3, 52, 9, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 60, 8, 6, 10, 6, 12, 6, 63, 9, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 70, 8, 8, 10, 8, 12, 8, 73, 9, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 5, 8, 79, 8, 8, 10, 8, 12, 8, 82, 9, 8, 1, 8, 3, 8, 85,
		8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 90, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 4, 14, 105,
		8, 14, 11, 14, 12, 14, 106, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 0, 11, 0, 13, 5, 15, 6, 17, 7, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0,
		29, 8, 1, 0, 8, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 49,
		57, 9, 0, 34, 34, 39, 39, 47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114,
		114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 0, 31, 39, 39, 92,
		92, 3, 0, 0, 31, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 112, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1,
		31, 1, 0, 0, 0, 3, 36, 1, 0, 0, 0, 5, 41, 1, 0, 0, 0, 7, 45, 1, 0, 0, 0,
		9, 53, 1, 0, 0, 0, 11, 55, 1, 0, 0, 0, 13, 57, 1, 0, 0, 0, 15, 64, 1, 0,
		0, 0, 17, 84, 1, 0, 0, 0, 19, 86, 1, 0, 0, 0, 21, 91, 1, 0, 0, 0, 23, 97,
		1, 0, 0, 0, 25, 99, 1, 0, 0, 0, 27, 101, 1, 0, 0, 0, 29, 104, 1, 0, 0,
		0, 31, 32, 5, 67, 0, 0, 32, 33, 5, 65, 0, 0, 33, 34, 5, 83, 0, 0, 34, 35,
		5, 69, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 87, 0, 0, 37, 38, 5, 72, 0,
		0, 38, 39, 5, 69, 0, 0, 39, 40, 5, 78, 0, 0, 40, 4, 1, 0, 0, 0, 41, 42,
		5, 69, 0, 0, 42, 43, 5, 78, 0, 0, 43, 44, 5, 68, 0, 0, 44, 6, 1, 0, 0,
		0, 45, 50, 3, 11, 5, 0, 46, 49, 3, 9, 4, 0, 47, 49, 3, 11, 5, 0, 48, 46,
		1, 0, 0, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0,
		50, 51, 1, 0, 0, 0, 51, 8, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 7, 0,
		0, 0, 54, 10, 1, 0, 0, 0, 55, 56, 7, 1, 0, 0, 56, 12, 1, 0, 0, 0, 57, 61,
		7, 2, 0, 0, 58, 60, 3, 9, 4, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0,
		61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 14, 1, 0, 0, 0, 63, 61, 1,
		0, 0, 0, 64, 65, 5, 48, 0, 0, 65, 16, 1, 0, 0, 0, 66, 71, 5, 34, 0, 0,
		67, 70, 3, 19, 9, 0, 68, 70, 3, 27, 13, 0, 69, 67, 1, 0, 0, 0, 69, 68,
		1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0,
		72, 74, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 85, 5, 34, 0, 0, 75, 80, 5,
		39, 0, 0, 76, 79, 3, 19, 9, 0, 77, 79, 3, 25, 12, 0, 78, 76, 1, 0, 0, 0,
		78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 85, 5, 39, 0, 0, 84,
		66, 1, 0, 0, 0, 84, 75, 1, 0, 0, 0, 85, 18, 1, 0, 0, 0, 86, 89, 5, 92,
		0, 0, 87, 90, 7, 3, 0, 0, 88, 90, 3, 21, 10, 0, 89, 87, 1, 0, 0, 0, 89,
		88, 1, 0, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92, 5, 117, 0, 0, 92, 93, 3, 23,
		11, 0, 93, 94, 3, 23, 11, 0, 94, 95, 3, 23, 11, 0, 95, 96, 3, 23, 11, 0,
		96, 22, 1, 0, 0, 0, 97, 98, 7, 4, 0, 0, 98, 24, 1, 0, 0, 0, 99, 100, 8,
		5, 0, 0, 100, 26, 1, 0, 0, 0, 101, 102, 8, 6, 0, 0, 102, 28, 1, 0, 0, 0,
		103, 105, 7, 7, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109,
		6, 14, 0, 0, 109, 30, 1, 0, 0, 0, 11, 0, 48, 50, 61, 69, 71, 78, 80, 84,
		89, 106, 1, 6, 0, 0,
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

// SorterLexerInit initializes any static state used to implement SorterLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSorterLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SorterLexerInit() {
	staticData := &sorterlexerLexerStaticData
	staticData.once.Do(sorterlexerLexerInit)
}

// NewSorterLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSorterLexer(input antlr.CharStream) *SorterLexer {
	SorterLexerInit()
	l := new(SorterLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &sorterlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SorterLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SorterLexer tokens.
const (
	SorterLexerCASE    = 1
	SorterLexerWHEN    = 2
	SorterLexerEND     = 3
	SorterLexerIDENT   = 4
	SorterLexerNUMBER1 = 5
	SorterLexerZERO    = 6
	SorterLexerSTRING  = 7
	SorterLexerWS      = 8
)
