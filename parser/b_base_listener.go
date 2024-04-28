// Code generated from b.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // b

import "github.com/antlr4-go/antlr/v4"

// BasebListener is a complete listener for a parse tree produced by bParser.
type BasebListener struct{}

var _ bListener = &BasebListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasebListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasebListener) ExitProg(ctx *ProgContext) {}

// EnterLine is called when production line is entered.
func (s *BasebListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasebListener) ExitLine(ctx *LineContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BasebListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BasebListener) ExitInstruction(ctx *InstructionContext) {}

// EnterRepeat is called when production repeat is entered.
func (s *BasebListener) EnterRepeat(ctx *RepeatContext) {}

// ExitRepeat is called when production repeat is exited.
func (s *BasebListener) ExitRepeat(ctx *RepeatContext) {}

// EnterLbl is called when production lbl is entered.
func (s *BasebListener) EnterLbl(ctx *LblContext) {}

// ExitLbl is called when production lbl is exited.
func (s *BasebListener) ExitLbl(ctx *LblContext) {}

// EnterArgumentlist is called when production argumentlist is entered.
func (s *BasebListener) EnterArgumentlist(ctx *ArgumentlistContext) {}

// ExitArgumentlist is called when production argumentlist is exited.
func (s *BasebListener) ExitArgumentlist(ctx *ArgumentlistContext) {}

// EnterLabel is called when production label is entered.
func (s *BasebListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BasebListener) ExitLabel(ctx *LabelContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasebListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasebListener) ExitArgument(ctx *ArgumentContext) {}

// EnterPrefix_ is called when production prefix_ is entered.
func (s *BasebListener) EnterPrefix_(ctx *Prefix_Context) {}

// ExitPrefix_ is called when production prefix_ is exited.
func (s *BasebListener) ExitPrefix_(ctx *Prefix_Context) {}

// EnterString_ is called when production string_ is entered.
func (s *BasebListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BasebListener) ExitString_(ctx *String_Context) {}

// EnterName is called when production name is entered.
func (s *BasebListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasebListener) ExitName(ctx *NameContext) {}

// EnterNumber is called when production number is entered.
func (s *BasebListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasebListener) ExitNumber(ctx *NumberContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BasebListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BasebListener) ExitOpcode(ctx *OpcodeContext) {}
