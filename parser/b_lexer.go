// Code generated from b.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type bLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BLexerLexerStaticData struct {
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

func blexerLexerInit() {
	staticData := &BLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "':'", "','", "'*'", "'+'", "'-'", "'#'", "", "'ADD'",
		"'AND'", "'BNE'", "'BRA'", "'DEC'", "'INC'", "'MOVH'", "'MOVL'", "'MOVS'",
		"'LDR'", "'STR'", "'NOP'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "TIMES", "ADD", "AND", "BNE", "BRA",
		"DEC", "INC", "MOVH", "MOVL", "MOVS", "LDR", "STR", "NOP", "NAME", "NUMBER",
		"HEX", "COMMENT", "STRING", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "TIMES",
		"ADD", "AND", "BNE", "BRA", "DEC", "INC", "MOVH", "MOVL", "MOVS", "LDR",
		"STR", "NOP", "NAME", "NUMBER", "HEX", "COMMENT", "STRING", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 180, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 4, 8, 76,
		8, 8, 11, 8, 12, 8, 77, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 5, 21, 133, 8, 21, 10, 21, 12, 21, 136, 9, 21,
		1, 22, 3, 22, 139, 8, 22, 1, 22, 4, 22, 142, 8, 22, 11, 22, 12, 22, 143,
		1, 23, 1, 23, 1, 23, 1, 23, 4, 23, 150, 8, 23, 11, 23, 12, 23, 151, 1,
		24, 1, 24, 5, 24, 156, 8, 24, 10, 24, 12, 24, 159, 9, 24, 1, 24, 1, 24,
		1, 25, 1, 25, 5, 25, 165, 8, 25, 10, 25, 12, 25, 168, 9, 25, 1, 25, 1,
		25, 1, 26, 4, 26, 173, 8, 26, 11, 26, 12, 26, 174, 1, 27, 1, 27, 1, 27,
		1, 27, 0, 0, 28, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 1, 0, 24, 2, 0, 88, 88, 120, 120, 1, 0, 48, 57, 2, 0, 65, 65,
		97, 97, 2, 0, 68, 68, 100, 100, 2, 0, 78, 78, 110, 110, 2, 0, 66, 66, 98,
		98, 2, 0, 69, 69, 101, 101, 2, 0, 82, 82, 114, 114, 2, 0, 67, 67, 99, 99,
		2, 0, 73, 73, 105, 105, 2, 0, 77, 77, 109, 109, 2, 0, 79, 79, 111, 111,
		2, 0, 86, 86, 118, 118, 2, 0, 72, 72, 104, 104, 2, 0, 76, 76, 108, 108,
		2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 80, 80, 112, 112,
		2, 0, 65, 90, 97, 122, 5, 0, 34, 34, 46, 46, 48, 57, 65, 90, 97, 122, 3,
		0, 48, 57, 65, 70, 97, 102, 2, 0, 10, 10, 13, 13, 1, 0, 34, 34, 2, 0, 9,
		9, 32, 32, 187, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0,
		0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0,
		0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1,
		0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45,
		1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0,
		53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 1, 57, 1, 0, 0, 0, 3, 59, 1, 0, 0, 0,
		5, 61, 1, 0, 0, 0, 7, 63, 1, 0, 0, 0, 9, 65, 1, 0, 0, 0, 11, 67, 1, 0,
		0, 0, 13, 69, 1, 0, 0, 0, 15, 71, 1, 0, 0, 0, 17, 73, 1, 0, 0, 0, 19, 79,
		1, 0, 0, 0, 21, 83, 1, 0, 0, 0, 23, 87, 1, 0, 0, 0, 25, 91, 1, 0, 0, 0,
		27, 95, 1, 0, 0, 0, 29, 99, 1, 0, 0, 0, 31, 103, 1, 0, 0, 0, 33, 108, 1,
		0, 0, 0, 35, 113, 1, 0, 0, 0, 37, 118, 1, 0, 0, 0, 39, 122, 1, 0, 0, 0,
		41, 126, 1, 0, 0, 0, 43, 130, 1, 0, 0, 0, 45, 138, 1, 0, 0, 0, 47, 145,
		1, 0, 0, 0, 49, 153, 1, 0, 0, 0, 51, 162, 1, 0, 0, 0, 53, 172, 1, 0, 0,
		0, 55, 176, 1, 0, 0, 0, 57, 58, 5, 40, 0, 0, 58, 2, 1, 0, 0, 0, 59, 60,
		5, 41, 0, 0, 60, 4, 1, 0, 0, 0, 61, 62, 5, 58, 0, 0, 62, 6, 1, 0, 0, 0,
		63, 64, 5, 44, 0, 0, 64, 8, 1, 0, 0, 0, 65, 66, 5, 42, 0, 0, 66, 10, 1,
		0, 0, 0, 67, 68, 5, 43, 0, 0, 68, 12, 1, 0, 0, 0, 69, 70, 5, 45, 0, 0,
		70, 14, 1, 0, 0, 0, 71, 72, 5, 35, 0, 0, 72, 16, 1, 0, 0, 0, 73, 75, 7,
		0, 0, 0, 74, 76, 7, 1, 0, 0, 75, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77,
		75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 18, 1, 0, 0, 0, 79, 80, 7, 2, 0,
		0, 80, 81, 7, 3, 0, 0, 81, 82, 7, 3, 0, 0, 82, 20, 1, 0, 0, 0, 83, 84,
		7, 2, 0, 0, 84, 85, 7, 4, 0, 0, 85, 86, 7, 3, 0, 0, 86, 22, 1, 0, 0, 0,
		87, 88, 7, 5, 0, 0, 88, 89, 7, 4, 0, 0, 89, 90, 7, 6, 0, 0, 90, 24, 1,
		0, 0, 0, 91, 92, 7, 5, 0, 0, 92, 93, 7, 7, 0, 0, 93, 94, 7, 2, 0, 0, 94,
		26, 1, 0, 0, 0, 95, 96, 7, 3, 0, 0, 96, 97, 7, 6, 0, 0, 97, 98, 7, 8, 0,
		0, 98, 28, 1, 0, 0, 0, 99, 100, 7, 9, 0, 0, 100, 101, 7, 4, 0, 0, 101,
		102, 7, 8, 0, 0, 102, 30, 1, 0, 0, 0, 103, 104, 7, 10, 0, 0, 104, 105,
		7, 11, 0, 0, 105, 106, 7, 12, 0, 0, 106, 107, 7, 13, 0, 0, 107, 32, 1,
		0, 0, 0, 108, 109, 7, 10, 0, 0, 109, 110, 7, 11, 0, 0, 110, 111, 7, 12,
		0, 0, 111, 112, 7, 14, 0, 0, 112, 34, 1, 0, 0, 0, 113, 114, 7, 10, 0, 0,
		114, 115, 7, 11, 0, 0, 115, 116, 7, 12, 0, 0, 116, 117, 7, 15, 0, 0, 117,
		36, 1, 0, 0, 0, 118, 119, 7, 14, 0, 0, 119, 120, 7, 3, 0, 0, 120, 121,
		7, 7, 0, 0, 121, 38, 1, 0, 0, 0, 122, 123, 7, 15, 0, 0, 123, 124, 7, 16,
		0, 0, 124, 125, 7, 7, 0, 0, 125, 40, 1, 0, 0, 0, 126, 127, 7, 4, 0, 0,
		127, 128, 7, 11, 0, 0, 128, 129, 7, 17, 0, 0, 129, 42, 1, 0, 0, 0, 130,
		134, 7, 18, 0, 0, 131, 133, 7, 19, 0, 0, 132, 131, 1, 0, 0, 0, 133, 136,
		1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 44, 1, 0,
		0, 0, 136, 134, 1, 0, 0, 0, 137, 139, 5, 36, 0, 0, 138, 137, 1, 0, 0, 0,
		138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 142, 7, 20, 0, 0, 141,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144,
		1, 0, 0, 0, 144, 46, 1, 0, 0, 0, 145, 146, 5, 48, 0, 0, 146, 147, 7, 0,
		0, 0, 147, 149, 1, 0, 0, 0, 148, 150, 7, 20, 0, 0, 149, 148, 1, 0, 0, 0,
		150, 151, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152,
		48, 1, 0, 0, 0, 153, 157, 5, 35, 0, 0, 154, 156, 8, 21, 0, 0, 155, 154,
		1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0,
		0, 0, 158, 160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 6, 24, 0, 0,
		161, 50, 1, 0, 0, 0, 162, 166, 5, 34, 0, 0, 163, 165, 8, 22, 0, 0, 164,
		163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167,
		1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 170, 5, 34,
		0, 0, 170, 52, 1, 0, 0, 0, 171, 173, 7, 21, 0, 0, 172, 171, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		54, 1, 0, 0, 0, 176, 177, 7, 23, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179,
		6, 27, 0, 0, 179, 56, 1, 0, 0, 0, 9, 0, 77, 134, 138, 143, 151, 157, 166,
		174, 1, 6, 0, 0,
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

// bLexerInit initializes any static state used to implement bLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewbLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BLexerInit() {
	staticData := &BLexerLexerStaticData
	staticData.once.Do(blexerLexerInit)
}

// NewbLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewbLexer(input antlr.CharStream) *bLexer {
	BLexerInit()
	l := new(bLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "b.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// bLexer tokens.
const (
	bLexerT__0    = 1
	bLexerT__1    = 2
	bLexerT__2    = 3
	bLexerT__3    = 4
	bLexerT__4    = 5
	bLexerT__5    = 6
	bLexerT__6    = 7
	bLexerT__7    = 8
	bLexerTIMES   = 9
	bLexerADD     = 10
	bLexerAND     = 11
	bLexerBNE     = 12
	bLexerBRA     = 13
	bLexerDEC     = 14
	bLexerINC     = 15
	bLexerMOVH    = 16
	bLexerMOVL    = 17
	bLexerMOVS    = 18
	bLexerLDR     = 19
	bLexerSTR     = 20
	bLexerNOP     = 21
	bLexerNAME    = 22
	bLexerNUMBER  = 23
	bLexerHEX     = 24
	bLexerCOMMENT = 25
	bLexerSTRING  = 26
	bLexerEOL     = 27
	bLexerWS      = 28
)
