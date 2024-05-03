package main

import (
	"corgcompl/parser"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type opcode struct {
	name string
	code int
}

var opcodes = []opcode{
	{"BRA", 0x00},
	{"BNE", 0x01},
	{"BEQ", 0x02},
	{"POP", 0x03},
	{"PSH", 0x04},
	{"INC", 0x05},
	{"DEC", 0x06},
	{"LSL", 0x07},
	{"LSR", 0x08},
	{"ASR", 0x09},
	{"CSL", 0x0A},
	{"CSR", 0x0B},
	{"AND", 0x0C},
	{"ORR", 0x0D},
	{"NOT", 0x0E},
	{"XOR", 0x0F},
	{"NAND", 0x10},
	{"MOVH", 0x11},
	{"LDR", 0x12},
	{"STR", 0x13},
	{"MOVL", 0x14},
	{"ADD", 0x15},
	{"ADC", 0x16},
	{"SUB", 0x17},
	{"MOVS", 0x18},
	{"ADDS", 0x19},
	{"SUBS", 0x1A},
	{"ANDS", 0x1B},
	{"ORRS", 0x1C},
	{"XORS", 0x1D},
	{"BX", 0x1E},
	{"BL", 0x1F},
	{"LDRIM", 0x20},
	{"STRIM", 0x21},
	{"NOP", 0x39},
}

func strToHex(s string) int {
	var result int

	switch s {
	case "R0":
		result = 0x00
	case "R1":
		result = 0x01
	case "R2":
		result = 0x02
	case "R3":
		result = 0x03
	}

	for _, v := range opcodes {
		if v.name == s {
			result = v.code
			break
		}
	}

	// weird tricks instead of rewrite
	switch s {
	case "PC":
		result = 0x00 - 4
	case "SP":
		result = 0x10 - 4
	case "AR":
		result = 0x11 - 4
	}

	if labels[s] != 0 {
		result = labels[s]
	}

	fmt.Printf("- map: %s 0x%x \n", s, result)

	return result
}

type bListener struct {
	*parser.BasebListener

	stack []int
}

func (l *bListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *bListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *bListener) ExitRepeat(c *parser.RepeatContext) {
	opcode := strToHex(c.Opcode().GetText())

	repeat := c.TIMES().GetText()

	repeatInt, err := strconv.Atoi(repeat[1:])

	if err != nil {
		panic(err)
	}

	for i := 0; i < repeatInt/2; i++ {
		fmt.Printf("%d %d, %04x \n", line, line+1, opcode|opcode<<10)

		output += fmt.Sprintf("%02x\n", opcode)
		output += fmt.Sprintf("%02x\n", opcode)
		line += 2
	}

}

func (l *bListener) EnterArgument(c *parser.ArgumentContext) {
	s := c.GetText()

	if s[0] == '0' && s[1] == 'x' {
		// hex
		if stoi, err := strconv.ParseInt(s[2:], 16, 64); err != nil {
			panic(err)
		} else {
			l.push(int(stoi))
		}
	} else {
		l.push(strToHex(s))
	}
}

var labels = make(map[string]int)

func (l *bListener) EnterLbl(c *parser.LblContext) {
	labels[c.Label().GetText()] = line
}

var line = 0
var output = ""

func (l *bListener) ExitInstruction(c *parser.InstructionContext) {
	a := strToHex(c.Opcode().GetText())

	if len(l.stack) == 1 {
		// like bra
		a = a << 10
		// rsel 0
		a = a | l.pop()
	} else if len(l.stack) == 2 {
		// like movh r1, 0x00
		a = a << 10
		a = a | l.pop()        // adr
		a = a | (l.pop() << 8) // rsel
	} else if len(l.stack) == 3 {
		// like add r2, r2, r3
		a = a << 10
		// s 0

		a = a | (l.pop() + 4)
		a = a | ((l.pop() + 4) << 3)
		a = a | ((l.pop() + 4) << 6)
	}

	for _, v := range l.stack {
		v = v << 4
		a = a | v
	}

	fmt.Println(c.GetText())
	fmt.Printf("%d %d, %04x \n", line, line+1, a)
	line += 2

	output += fmt.Sprintf("%02x\n", a&0xff00>>8)
	output += fmt.Sprintf("%02x\n", a&0x00ff)
}

// func (f *bListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Printf("EnterEveryRule %v\n", ctx.GetText())
// }

func main() {
	// Setup the input

	// NOP is special empty instruction
	// (NOP x27) means 27 NOP instructions
	is := antlr.NewInputStream(`
	BRA 0x28 # This instruction is written to the memory address 0x00,
	(NOP x39) # The first instruction must be written to address 0x28, this creates 39 NOP instructions(it's my addition to make it easier)
	MOVH R1, 0x00
	MOVL R1, 0x0A # This first instruction is written to the address 0x28,
	# R1 is used for the number of iterations
	MOVH R2, 0x00
	MOVL R2, 0x00 # R2 is used to store total
	MOVH R3, 0x00
	MOVL R3, 0xB0
	MOVS AR, R3 # AR is used to track data address: starts from 0xA0
	LABEL:
	LDR R3 # R3 ← M[AR] (AR = 0xB0 to 0xBA)
	ADD R2, R2, R3 # R2 ← R2 + R3 (Total = Total + M[AR])
	INC AR, AR # AR ← AR + 1 (Next Data)
	DEC R1, R1 # R1 ← R1 – 1 (Decrement Iteration Counter)
	BNE LABEL # Go back to LABEL if Z=0 (Iteration Counter > 0)
	INC AR, AR # AR ← AR + 1 (Total will be written to 0xBB)
	STR R2 # M[AR] ← R2 (Store Total at 0xBB)
	`)

	// Create the Lexer
	lexer := parser.NewbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// for {
	// 	t := lexer.NextToken()
	// 	if t.GetTokenType() == antlr.TokenEOF {
	// 		fmt.Println("EOF")
	// 		break
	// 	}
	//
	// 	fmt.Printf("%s (%q)\n",
	// 		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	// }
	// lexer.Reset()

	// Create the Parser
	p := parser.NewbParser(stream)

	// Parse the input
	var listener bListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Prog())

	fmt.Println(labels)
	fmt.Println(listener.stack)

	for range 256 - line {
		output += "00\n"
	}

	fmt.Println("output:")
	fmt.Println(output)
}
