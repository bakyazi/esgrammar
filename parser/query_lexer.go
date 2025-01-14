// Code generated from parser/Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

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

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var QueryLexerLexerStaticData struct {
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

func querylexerLexerInit() {
	staticData := &QueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'SEARCH'", "'COUNT'", "'FROM'", "'WHERE'", "'ORDER'", "'BY'", "'LIMIT'",
		"'OFFSET'", "'AND'", "'OR'", "'NOT'", "'LIKE'", "'IN'", "", "'ASC'",
		"'DESC'", "", "", "", "", "'('", "')'", "'.'", "','", "'='", "'!='",
		"'>'", "'<'", "'>='", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "SEARCH", "COUNT", "FROM", "WHERE", "ORDER", "BY", "LIMIT", "OFFSET",
		"AND", "OR", "NOT", "LIKE", "IN", "ISNULL", "ASC", "DESC", "IDENTIFIER",
		"NUMBER", "STRING", "WS", "LPAR", "RPAR", "DOT", "COMMA", "EQUAL", "NOT_EQUAL",
		"GREATER", "LESSER", "GREATER_EQ", "LESSER_EQ",
	}
	staticData.RuleNames = []string{
		"SEARCH", "COUNT", "FROM", "WHERE", "ORDER", "BY", "LIMIT", "OFFSET",
		"AND", "OR", "NOT", "LIKE", "IN", "ISNULL", "ASC", "DESC", "IDENTIFIER",
		"NUMBER", "STRING", "WS", "LPAR", "RPAR", "DOT", "COMMA", "EQUAL", "NOT_EQUAL",
		"GREATER", "LESSER", "GREATER_EQ", "LESSER_EQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 196, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 5, 16, 146, 8, 16, 10, 16, 12, 16, 149, 9, 16, 1, 17, 4,
		17, 152, 8, 17, 11, 17, 12, 17, 153, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18,
		160, 8, 18, 10, 18, 12, 18, 163, 9, 18, 1, 18, 1, 18, 1, 19, 4, 19, 168,
		8, 19, 11, 19, 12, 19, 169, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 0, 0, 30, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41,
		21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59,
		30, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13,
		32, 32, 200, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1,
		61, 1, 0, 0, 0, 3, 68, 1, 0, 0, 0, 5, 74, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0,
		9, 85, 1, 0, 0, 0, 11, 91, 1, 0, 0, 0, 13, 94, 1, 0, 0, 0, 15, 100, 1,
		0, 0, 0, 17, 107, 1, 0, 0, 0, 19, 111, 1, 0, 0, 0, 21, 114, 1, 0, 0, 0,
		23, 118, 1, 0, 0, 0, 25, 123, 1, 0, 0, 0, 27, 126, 1, 0, 0, 0, 29, 134,
		1, 0, 0, 0, 31, 138, 1, 0, 0, 0, 33, 143, 1, 0, 0, 0, 35, 151, 1, 0, 0,
		0, 37, 155, 1, 0, 0, 0, 39, 167, 1, 0, 0, 0, 41, 173, 1, 0, 0, 0, 43, 175,
		1, 0, 0, 0, 45, 177, 1, 0, 0, 0, 47, 179, 1, 0, 0, 0, 49, 181, 1, 0, 0,
		0, 51, 183, 1, 0, 0, 0, 53, 186, 1, 0, 0, 0, 55, 188, 1, 0, 0, 0, 57, 190,
		1, 0, 0, 0, 59, 193, 1, 0, 0, 0, 61, 62, 5, 83, 0, 0, 62, 63, 5, 69, 0,
		0, 63, 64, 5, 65, 0, 0, 64, 65, 5, 82, 0, 0, 65, 66, 5, 67, 0, 0, 66, 67,
		5, 72, 0, 0, 67, 2, 1, 0, 0, 0, 68, 69, 5, 67, 0, 0, 69, 70, 5, 79, 0,
		0, 70, 71, 5, 85, 0, 0, 71, 72, 5, 78, 0, 0, 72, 73, 5, 84, 0, 0, 73, 4,
		1, 0, 0, 0, 74, 75, 5, 70, 0, 0, 75, 76, 5, 82, 0, 0, 76, 77, 5, 79, 0,
		0, 77, 78, 5, 77, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 87, 0, 0, 80, 81,
		5, 72, 0, 0, 81, 82, 5, 69, 0, 0, 82, 83, 5, 82, 0, 0, 83, 84, 5, 69, 0,
		0, 84, 8, 1, 0, 0, 0, 85, 86, 5, 79, 0, 0, 86, 87, 5, 82, 0, 0, 87, 88,
		5, 68, 0, 0, 88, 89, 5, 69, 0, 0, 89, 90, 5, 82, 0, 0, 90, 10, 1, 0, 0,
		0, 91, 92, 5, 66, 0, 0, 92, 93, 5, 89, 0, 0, 93, 12, 1, 0, 0, 0, 94, 95,
		5, 76, 0, 0, 95, 96, 5, 73, 0, 0, 96, 97, 5, 77, 0, 0, 97, 98, 5, 73, 0,
		0, 98, 99, 5, 84, 0, 0, 99, 14, 1, 0, 0, 0, 100, 101, 5, 79, 0, 0, 101,
		102, 5, 70, 0, 0, 102, 103, 5, 70, 0, 0, 103, 104, 5, 83, 0, 0, 104, 105,
		5, 69, 0, 0, 105, 106, 5, 84, 0, 0, 106, 16, 1, 0, 0, 0, 107, 108, 5, 65,
		0, 0, 108, 109, 5, 78, 0, 0, 109, 110, 5, 68, 0, 0, 110, 18, 1, 0, 0, 0,
		111, 112, 5, 79, 0, 0, 112, 113, 5, 82, 0, 0, 113, 20, 1, 0, 0, 0, 114,
		115, 5, 78, 0, 0, 115, 116, 5, 79, 0, 0, 116, 117, 5, 84, 0, 0, 117, 22,
		1, 0, 0, 0, 118, 119, 5, 76, 0, 0, 119, 120, 5, 73, 0, 0, 120, 121, 5,
		75, 0, 0, 121, 122, 5, 69, 0, 0, 122, 24, 1, 0, 0, 0, 123, 124, 5, 73,
		0, 0, 124, 125, 5, 78, 0, 0, 125, 26, 1, 0, 0, 0, 126, 127, 5, 73, 0, 0,
		127, 128, 5, 83, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 5, 78, 0, 0, 130,
		131, 5, 85, 0, 0, 131, 132, 5, 76, 0, 0, 132, 133, 5, 76, 0, 0, 133, 28,
		1, 0, 0, 0, 134, 135, 5, 65, 0, 0, 135, 136, 5, 83, 0, 0, 136, 137, 5,
		67, 0, 0, 137, 30, 1, 0, 0, 0, 138, 139, 5, 68, 0, 0, 139, 140, 5, 69,
		0, 0, 140, 141, 5, 83, 0, 0, 141, 142, 5, 67, 0, 0, 142, 32, 1, 0, 0, 0,
		143, 147, 7, 0, 0, 0, 144, 146, 7, 1, 0, 0, 145, 144, 1, 0, 0, 0, 146,
		149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 34, 1,
		0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 152, 7, 2, 0, 0, 151, 150, 1, 0, 0,
		0, 152, 153, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		36, 1, 0, 0, 0, 155, 161, 5, 34, 0, 0, 156, 160, 8, 3, 0, 0, 157, 158,
		5, 34, 0, 0, 158, 160, 5, 34, 0, 0, 159, 156, 1, 0, 0, 0, 159, 157, 1,
		0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165, 5, 34, 0, 0, 165,
		38, 1, 0, 0, 0, 166, 168, 7, 4, 0, 0, 167, 166, 1, 0, 0, 0, 168, 169, 1,
		0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 1, 0, 0,
		0, 171, 172, 6, 19, 0, 0, 172, 40, 1, 0, 0, 0, 173, 174, 5, 40, 0, 0, 174,
		42, 1, 0, 0, 0, 175, 176, 5, 41, 0, 0, 176, 44, 1, 0, 0, 0, 177, 178, 5,
		46, 0, 0, 178, 46, 1, 0, 0, 0, 179, 180, 5, 44, 0, 0, 180, 48, 1, 0, 0,
		0, 181, 182, 5, 61, 0, 0, 182, 50, 1, 0, 0, 0, 183, 184, 5, 33, 0, 0, 184,
		185, 5, 61, 0, 0, 185, 52, 1, 0, 0, 0, 186, 187, 5, 62, 0, 0, 187, 54,
		1, 0, 0, 0, 188, 189, 5, 60, 0, 0, 189, 56, 1, 0, 0, 0, 190, 191, 5, 62,
		0, 0, 191, 192, 5, 61, 0, 0, 192, 58, 1, 0, 0, 0, 193, 194, 5, 60, 0, 0,
		194, 195, 5, 61, 0, 0, 195, 60, 1, 0, 0, 0, 6, 0, 147, 153, 159, 161, 169,
		1, 6, 0, 0,
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

// QueryLexerInit initializes any static state used to implement QueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLexerInit() {
	staticData := &QueryLexerLexerStaticData
	staticData.once.Do(querylexerLexerInit)
}

// NewQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQueryLexer(input antlr.CharStream) *QueryLexer {
	QueryLexerInit()
	l := new(QueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &QueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerSEARCH     = 1
	QueryLexerCOUNT      = 2
	QueryLexerFROM       = 3
	QueryLexerWHERE      = 4
	QueryLexerORDER      = 5
	QueryLexerBY         = 6
	QueryLexerLIMIT      = 7
	QueryLexerOFFSET     = 8
	QueryLexerAND        = 9
	QueryLexerOR         = 10
	QueryLexerNOT        = 11
	QueryLexerLIKE       = 12
	QueryLexerIN         = 13
	QueryLexerISNULL     = 14
	QueryLexerASC        = 15
	QueryLexerDESC       = 16
	QueryLexerIDENTIFIER = 17
	QueryLexerNUMBER     = 18
	QueryLexerSTRING     = 19
	QueryLexerWS         = 20
	QueryLexerLPAR       = 21
	QueryLexerRPAR       = 22
	QueryLexerDOT        = 23
	QueryLexerCOMMA      = 24
	QueryLexerEQUAL      = 25
	QueryLexerNOT_EQUAL  = 26
	QueryLexerGREATER    = 27
	QueryLexerLESSER     = 28
	QueryLexerGREATER_EQ = 29
	QueryLexerLESSER_EQ  = 30
)
