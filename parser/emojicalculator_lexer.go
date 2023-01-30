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

type EmojiCalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var emojicalculatorlexerLexerStaticData struct {
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

func emojicalculatorlexerLexerInit() {
	staticData := &emojicalculatorlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.symbolicNames = []string{
		"", "MULTIPLICATION", "DIVISION", "ADDITION", "SUBSTRACTION", "NUMBER",
		"WHITESPACE",
	}
	staticData.ruleNames = []string{
		"MULTIPLICATION", "DIVISION", "ADDITION", "SUBSTRACTION", "NUMBER",
		"WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 61, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 20, 8, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		33, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 40, 8, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 48, 8, 3, 1, 4, 4, 4, 51, 8, 4, 11, 4, 12, 4,
		52, 1, 5, 4, 5, 56, 8, 5, 11, 5, 12, 5, 57, 1, 5, 1, 5, 0, 0, 6, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 1, 0, 4, 2, 0, 42, 42, 120, 120, 2, 0, 37,
		37, 47, 47, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 66, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3, 32, 1, 0, 0, 0, 5, 39, 1,
		0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 50, 1, 0, 0, 0, 11, 55, 1, 0, 0, 0, 13,
		20, 7, 0, 0, 0, 14, 15, 5, 116, 0, 0, 15, 16, 5, 105, 0, 0, 16, 17, 5,
		109, 0, 0, 17, 18, 5, 101, 0, 0, 18, 20, 5, 115, 0, 0, 19, 13, 1, 0, 0,
		0, 19, 14, 1, 0, 0, 0, 20, 2, 1, 0, 0, 0, 21, 33, 7, 1, 0, 0, 22, 23, 5,
		100, 0, 0, 23, 24, 5, 105, 0, 0, 24, 25, 5, 118, 0, 0, 25, 26, 5, 105,
		0, 0, 26, 27, 5, 100, 0, 0, 27, 28, 5, 101, 0, 0, 28, 29, 5, 100, 0, 0,
		29, 30, 5, 32, 0, 0, 30, 31, 5, 98, 0, 0, 31, 33, 5, 121, 0, 0, 32, 21,
		1, 0, 0, 0, 32, 22, 1, 0, 0, 0, 33, 4, 1, 0, 0, 0, 34, 40, 5, 43, 0, 0,
		35, 36, 5, 112, 0, 0, 36, 37, 5, 108, 0, 0, 37, 38, 5, 117, 0, 0, 38, 40,
		5, 115, 0, 0, 39, 34, 1, 0, 0, 0, 39, 35, 1, 0, 0, 0, 40, 6, 1, 0, 0, 0,
		41, 48, 5, 45, 0, 0, 42, 43, 5, 109, 0, 0, 43, 44, 5, 105, 0, 0, 44, 45,
		5, 110, 0, 0, 45, 46, 5, 117, 0, 0, 46, 48, 5, 115, 0, 0, 47, 41, 1, 0,
		0, 0, 47, 42, 1, 0, 0, 0, 48, 8, 1, 0, 0, 0, 49, 51, 7, 2, 0, 0, 50, 49,
		1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0,
		53, 10, 1, 0, 0, 0, 54, 56, 7, 3, 0, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1,
		0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59,
		60, 6, 5, 0, 0, 60, 12, 1, 0, 0, 0, 7, 0, 19, 32, 39, 47, 52, 57, 1, 6,
		0, 0,
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

// EmojiCalculatorLexerInit initializes any static state used to implement EmojiCalculatorLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEmojiCalculatorLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EmojiCalculatorLexerInit() {
	staticData := &emojicalculatorlexerLexerStaticData
	staticData.once.Do(emojicalculatorlexerLexerInit)
}

// NewEmojiCalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEmojiCalculatorLexer(input antlr.CharStream) *EmojiCalculatorLexer {
	EmojiCalculatorLexerInit()
	l := new(EmojiCalculatorLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &emojicalculatorlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "EmojiCalculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EmojiCalculatorLexer tokens.
const (
	EmojiCalculatorLexerMULTIPLICATION = 1
	EmojiCalculatorLexerDIVISION       = 2
	EmojiCalculatorLexerADDITION       = 3
	EmojiCalculatorLexerSUBSTRACTION   = 4
	EmojiCalculatorLexerNUMBER         = 5
	EmojiCalculatorLexerWHITESPACE     = 6
)
