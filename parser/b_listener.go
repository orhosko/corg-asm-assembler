// Code generated from b.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // b

import "github.com/antlr4-go/antlr/v4"

// bListener is a complete listener for a parse tree produced by bParser.
type bListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterRepeat is called when entering the repeat production.
	EnterRepeat(c *RepeatContext)

	// EnterLbl is called when entering the lbl production.
	EnterLbl(c *LblContext)

	// EnterArgumentlist is called when entering the argumentlist production.
	EnterArgumentlist(c *ArgumentlistContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterPrefix_ is called when entering the prefix_ production.
	EnterPrefix_(c *Prefix_Context)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitRepeat is called when exiting the repeat production.
	ExitRepeat(c *RepeatContext)

	// ExitLbl is called when exiting the lbl production.
	ExitLbl(c *LblContext)

	// ExitArgumentlist is called when exiting the argumentlist production.
	ExitArgumentlist(c *ArgumentlistContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitPrefix_ is called when exiting the prefix_ production.
	ExitPrefix_(c *Prefix_Context)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)
}
