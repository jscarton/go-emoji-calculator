/*
Package calculator defines a CustomSyntaxError handler, a CustomErrorListener for ANTLR4 parser and lexer,a listener for the parser and
also exposes the Calc function  which start the expression parsing process
*/
package calculator

import (
	"fmt"
	"go-emoji-calculator/parser"
	"math"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// implements the basic Listener functionality for parsing expressions
type emojiCalculatorListener struct {
	*parser.BaseEmojiCalculatorListener

	stack []float64
}

// implements a ccustom error handler
type CustomSyntaxError struct {
	line, column int
	msg          string
}

func (c *CustomSyntaxError) Error() string {
	return fmt.Sprintf("%s at (%d,%d)", c.msg, c.line, c.column)
}

// implements a custom error listener for the parser and lexer
type CustomErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []CustomSyntaxError
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (l *emojiCalculatorListener) push(i float64) {
	l.stack = append(l.stack, i)
}

func (l *emojiCalculatorListener) pop() float64 {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Pop the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

// ExitMulDiv is called when exiting the MulDiv production.
func (l *emojiCalculatorListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	if math.IsNaN(left) || math.IsNaN(right) {
		l.push(math.NaN())
		return
	}

	switch c.GetOp().GetTokenType() {
	case parser.EmojiCalculatorParserMULTIPLICATION:
		l.push(left * right)
	case parser.EmojiCalculatorParserDIVISION:
		if right != 0 {
			l.push(left / right)
		} else {
			l.push(math.NaN())
			fmt.Printf("error: can not divide by 0\n")
		}
	default:
		panic(fmt.Sprintf("unexpected operation: %s", c.GetOp().GetText()))
	}
}

// ExitAddSub is called when exiting the AddSub production.
func (l *emojiCalculatorListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	if math.IsNaN(left) || math.IsNaN(right) {
		l.push(math.NaN())
		return
	}

	switch c.GetOp().GetTokenType() {
	case parser.EmojiCalculatorParserADDITION:
		l.push(left + right)
	case parser.EmojiCalculatorParserSUBSTRACTION:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected operation: %s", c.GetOp().GetText()))
	}
}

// ExitNumber is called when exiting the Number production.
func (l *emojiCalculatorListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		l.push(math.NaN())
		//panic(err.Error())
	}

	l.push(float64(i))
}

// calc takes a string expression and returns the evaluated result.
func Calc(input string) float64 {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexerErrors := &CustomErrorListener{}
	lexer := parser.NewEmojiCalculatorLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	parserErrors := &CustomErrorListener{}
	p := parser.NewEmojiCalculatorParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	// Finally parse the expression (by walking the tree)
	var listener emojiCalculatorListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	if len(lexerErrors.Errors) > 0 {
		fmt.Printf(" %d Lexer errors found\n", len(lexerErrors.Errors))
		for _, e := range lexerErrors.Errors {
			fmt.Println(e.Error())
		}
	}

	if len(parserErrors.Errors) > 0 {
		fmt.Printf("%d Parser errors found\n", len(parserErrors.Errors))
		for _, e := range parserErrors.Errors {
			fmt.Println(e.Error())
		}
	}

	return listener.pop()
}
