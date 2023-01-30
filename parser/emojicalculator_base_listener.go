// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // EmojiCalculator

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseEmojiCalculatorListener is a complete listener for a parse tree produced by EmojiCalculatorParser.
type BaseEmojiCalculatorListener struct{}

var _ EmojiCalculatorListener = &BaseEmojiCalculatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEmojiCalculatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEmojiCalculatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEmojiCalculatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEmojiCalculatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseEmojiCalculatorListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseEmojiCalculatorListener) ExitStart(ctx *StartContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseEmojiCalculatorListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseEmojiCalculatorListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseEmojiCalculatorListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseEmojiCalculatorListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseEmojiCalculatorListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseEmojiCalculatorListener) ExitAddSub(ctx *AddSubContext) {}
