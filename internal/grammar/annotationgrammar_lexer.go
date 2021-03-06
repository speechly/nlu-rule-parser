// Code generated from internal/grammar/AnnotationGrammar.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 50, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 6, 7, 31, 10, 7, 13, 7, 14, 7, 32, 3, 8, 7, 8, 36,
	10, 8, 12, 8, 14, 8, 39, 11, 8, 3, 8, 6, 8, 42, 10, 8, 13, 8, 14, 8, 43,
	3, 9, 6, 9, 47, 10, 9, 13, 9, 14, 9, 48, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2, 4, 7, 2, 11, 12, 34, 34, 42, 44,
	93, 93, 95, 95, 4, 2, 11, 11, 34, 34, 2, 53, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5,
	21, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11, 27, 3, 2, 2,
	2, 13, 30, 3, 2, 2, 2, 15, 37, 3, 2, 2, 2, 17, 46, 3, 2, 2, 2, 19, 20,
	7, 44, 2, 2, 20, 4, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 6, 3, 2, 2, 2,
	23, 24, 7, 43, 2, 2, 24, 8, 3, 2, 2, 2, 25, 26, 7, 93, 2, 2, 26, 10, 3,
	2, 2, 2, 27, 28, 7, 95, 2, 2, 28, 12, 3, 2, 2, 2, 29, 31, 10, 2, 2, 2,
	30, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3,
	2, 2, 2, 33, 14, 3, 2, 2, 2, 34, 36, 5, 17, 9, 2, 35, 34, 3, 2, 2, 2, 36,
	39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 41, 3, 2, 2,
	2, 39, 37, 3, 2, 2, 2, 40, 42, 7, 12, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 16, 3, 2, 2, 2,
	45, 47, 9, 3, 2, 2, 46, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 46, 3,
	2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 18, 3, 2, 2, 2, 7, 2, 32, 37, 43, 48,
	2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'('", "')'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "INTENT_NAME_START", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_SB", "CLOSE_SB",
	"WORD", "END", "WHITESPACE",
}

var lexerRuleNames = []string{
	"INTENT_NAME_START", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_SB", "CLOSE_SB",
	"WORD", "END", "WHITESPACE",
}

type AnnotationGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAnnotationGrammarLexer(input antlr.CharStream) *AnnotationGrammarLexer {

	l := new(AnnotationGrammarLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "AnnotationGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AnnotationGrammarLexer tokens.
const (
	AnnotationGrammarLexerINTENT_NAME_START = 1
	AnnotationGrammarLexerOPEN_PAREN        = 2
	AnnotationGrammarLexerCLOSE_PAREN       = 3
	AnnotationGrammarLexerOPEN_SB           = 4
	AnnotationGrammarLexerCLOSE_SB          = 5
	AnnotationGrammarLexerWORD              = 6
	AnnotationGrammarLexerEND               = 7
	AnnotationGrammarLexerWHITESPACE        = 8
)
