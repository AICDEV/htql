// Code generated from /Users/aicdev/coding/htql/kotlin/app/src/main/antlr/Htql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package htql

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

type HtqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var HtqlLexerLexerStaticData struct {
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

func htqllexerLexerInit() {
	staticData := &HtqlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'attributes'", "'SELECT'", "'FROM'", "'WHERE'", "'AND'", "'OR'",
		"'IS'", "'NOT'", "'NULL'", "'*'", "','", "'.'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "SELECT", "FROM", "WHERE", "AND", "OR", "IS", "NOT", "NULL",
		"STAR", "COMMA", "DOT", "EQUALS", "IDENTIFIER", "STRING", "WS", "URL",
	}
	staticData.RuleNames = []string{
		"T__0", "SELECT", "FROM", "WHERE", "AND", "OR", "IS", "NOT", "NULL",
		"STAR", "COMMA", "DOT", "EQUALS", "IDENTIFIER", "STRING", "WS", "URL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 151, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 5, 13, 94, 8, 13, 10, 13, 12, 13, 97, 9, 13, 1, 14, 1,
		14, 5, 14, 101, 8, 14, 10, 14, 12, 14, 104, 9, 14, 1, 14, 1, 14, 1, 15,
		4, 15, 109, 8, 15, 11, 15, 12, 15, 110, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 124, 8, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 131, 8, 16, 11, 16, 12, 16, 132, 1,
		16, 1, 16, 4, 16, 137, 8, 16, 11, 16, 12, 16, 138, 3, 16, 141, 8, 16, 1,
		16, 1, 16, 5, 16, 145, 8, 16, 10, 16, 12, 16, 148, 9, 16, 3, 16, 150, 8,
		16, 0, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 7,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3,
		0, 10, 10, 13, 13, 39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 5, 0, 45, 46, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 8, 0, 35, 35, 37, 38, 46, 57,
		61, 61, 63, 63, 65, 90, 95, 95, 97, 122, 159, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		1, 35, 1, 0, 0, 0, 3, 46, 1, 0, 0, 0, 5, 53, 1, 0, 0, 0, 7, 58, 1, 0, 0,
		0, 9, 64, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13, 71, 1, 0, 0, 0, 15, 74, 1,
		0, 0, 0, 17, 78, 1, 0, 0, 0, 19, 83, 1, 0, 0, 0, 21, 85, 1, 0, 0, 0, 23,
		87, 1, 0, 0, 0, 25, 89, 1, 0, 0, 0, 27, 91, 1, 0, 0, 0, 29, 98, 1, 0, 0,
		0, 31, 108, 1, 0, 0, 0, 33, 123, 1, 0, 0, 0, 35, 36, 5, 97, 0, 0, 36, 37,
		5, 116, 0, 0, 37, 38, 5, 116, 0, 0, 38, 39, 5, 114, 0, 0, 39, 40, 5, 105,
		0, 0, 40, 41, 5, 98, 0, 0, 41, 42, 5, 117, 0, 0, 42, 43, 5, 116, 0, 0,
		43, 44, 5, 101, 0, 0, 44, 45, 5, 115, 0, 0, 45, 2, 1, 0, 0, 0, 46, 47,
		5, 83, 0, 0, 47, 48, 5, 69, 0, 0, 48, 49, 5, 76, 0, 0, 49, 50, 5, 69, 0,
		0, 50, 51, 5, 67, 0, 0, 51, 52, 5, 84, 0, 0, 52, 4, 1, 0, 0, 0, 53, 54,
		5, 70, 0, 0, 54, 55, 5, 82, 0, 0, 55, 56, 5, 79, 0, 0, 56, 57, 5, 77, 0,
		0, 57, 6, 1, 0, 0, 0, 58, 59, 5, 87, 0, 0, 59, 60, 5, 72, 0, 0, 60, 61,
		5, 69, 0, 0, 61, 62, 5, 82, 0, 0, 62, 63, 5, 69, 0, 0, 63, 8, 1, 0, 0,
		0, 64, 65, 5, 65, 0, 0, 65, 66, 5, 78, 0, 0, 66, 67, 5, 68, 0, 0, 67, 10,
		1, 0, 0, 0, 68, 69, 5, 79, 0, 0, 69, 70, 5, 82, 0, 0, 70, 12, 1, 0, 0,
		0, 71, 72, 5, 73, 0, 0, 72, 73, 5, 83, 0, 0, 73, 14, 1, 0, 0, 0, 74, 75,
		5, 78, 0, 0, 75, 76, 5, 79, 0, 0, 76, 77, 5, 84, 0, 0, 77, 16, 1, 0, 0,
		0, 78, 79, 5, 78, 0, 0, 79, 80, 5, 85, 0, 0, 80, 81, 5, 76, 0, 0, 81, 82,
		5, 76, 0, 0, 82, 18, 1, 0, 0, 0, 83, 84, 5, 42, 0, 0, 84, 20, 1, 0, 0,
		0, 85, 86, 5, 44, 0, 0, 86, 22, 1, 0, 0, 0, 87, 88, 5, 46, 0, 0, 88, 24,
		1, 0, 0, 0, 89, 90, 5, 61, 0, 0, 90, 26, 1, 0, 0, 0, 91, 95, 7, 0, 0, 0,
		92, 94, 7, 1, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1,
		0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 28, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98,
		102, 5, 39, 0, 0, 99, 101, 8, 2, 0, 0, 100, 99, 1, 0, 0, 0, 101, 104, 1,
		0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 1, 0, 0,
		0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 39, 0, 0, 106, 30, 1, 0, 0, 0, 107,
		109, 7, 3, 0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108,
		1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 6, 15,
		0, 0, 113, 32, 1, 0, 0, 0, 114, 115, 5, 104, 0, 0, 115, 116, 5, 116, 0,
		0, 116, 117, 5, 116, 0, 0, 117, 124, 5, 112, 0, 0, 118, 119, 5, 104, 0,
		0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 116, 0, 0, 121, 122, 5, 112, 0,
		0, 122, 124, 5, 115, 0, 0, 123, 114, 1, 0, 0, 0, 123, 118, 1, 0, 0, 0,
		124, 125, 1, 0, 0, 0, 125, 126, 5, 58, 0, 0, 126, 127, 5, 47, 0, 0, 127,
		128, 5, 47, 0, 0, 128, 130, 1, 0, 0, 0, 129, 131, 7, 4, 0, 0, 130, 129,
		1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 140, 1, 0, 0, 0, 134, 136, 5, 58, 0, 0, 135, 137, 7, 5, 0, 0,
		136, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 134, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 149, 1, 0, 0, 0, 142, 146, 5, 47, 0, 0, 143, 145, 7, 6,
		0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0,
		146, 147, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149,
		142, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 34, 1, 0, 0, 0, 10, 0, 95,
		102, 110, 123, 132, 138, 140, 146, 149, 1, 6, 0, 0,
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

// HtqlLexerInit initializes any static state used to implement HtqlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewHtqlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HtqlLexerInit() {
	staticData := &HtqlLexerLexerStaticData
	staticData.once.Do(htqllexerLexerInit)
}

// NewHtqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewHtqlLexer(input antlr.CharStream) *HtqlLexer {
	HtqlLexerInit()
	l := new(HtqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &HtqlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Htql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HtqlLexer tokens.
const (
	HtqlLexerT__0       = 1
	HtqlLexerSELECT     = 2
	HtqlLexerFROM       = 3
	HtqlLexerWHERE      = 4
	HtqlLexerAND        = 5
	HtqlLexerOR         = 6
	HtqlLexerIS         = 7
	HtqlLexerNOT        = 8
	HtqlLexerNULL       = 9
	HtqlLexerSTAR       = 10
	HtqlLexerCOMMA      = 11
	HtqlLexerDOT        = 12
	HtqlLexerEQUALS     = 13
	HtqlLexerIDENTIFIER = 14
	HtqlLexerSTRING     = 15
	HtqlLexerWS         = 16
	HtqlLexerURL        = 17
)
