// Code generated from b.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // b

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type bParser struct {
	*antlr.BaseParser
}

var BParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bParserInit() {
	staticData := &BParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "':'", "','", "'*'", "'+'", "'-'", "'#'", "", "'ADD'",
		"'AND'", "'BNE'", "'BRA'", "'DEC'", "'INC'", "'MOVH'", "'MOVL'", "'MOVS'",
		"'LDR'", "'POP'", "'PSH'", "'STR'", "'XOR'", "'SUB'", "'NOP'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "TIMES", "ADD", "AND", "BNE", "BRA",
		"DEC", "INC", "MOVH", "MOVL", "MOVS", "LDR", "POP", "PSH", "STR", "XOR",
		"SUB", "NOP", "NAME", "NUMBER", "HEX", "COMMENT", "STRING", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "line", "instruction", "repeat", "lbl", "argumentlist", "label",
		"argument", "prefix_", "string_", "name", "number", "opcode",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 93, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 1, 1, 1, 1, 2, 3,
		2, 43, 8, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 60, 8, 5, 1, 6, 1, 6, 1, 7, 3,
		7, 65, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 71, 8, 7, 1, 7, 1, 7, 3, 7,
		75, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 81, 8, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 0, 0, 13, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 3, 1, 0, 6, 7, 1, 0, 27, 28, 1,
		0, 10, 25, 92, 0, 29, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0,
		6, 48, 1, 0, 0, 0, 8, 53, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 61, 1, 0,
		0, 0, 14, 80, 1, 0, 0, 0, 16, 82, 1, 0, 0, 0, 18, 84, 1, 0, 0, 0, 20, 86,
		1, 0, 0, 0, 22, 88, 1, 0, 0, 0, 24, 90, 1, 0, 0, 0, 26, 28, 3, 2, 1, 0,
		27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1,
		0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 0, 0, 1, 33,
		1, 1, 0, 0, 0, 34, 38, 3, 4, 2, 0, 35, 38, 3, 8, 4, 0, 36, 38, 3, 6, 3,
		0, 37, 34, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 37, 38,
		1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 5, 31, 0, 0, 40, 3, 1, 0, 0, 0,
		41, 43, 3, 12, 6, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1,
		0, 0, 0, 44, 46, 3, 24, 12, 0, 45, 47, 3, 10, 5, 0, 46, 45, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 5, 1, 0, 0, 0, 48, 49, 5, 1, 0, 0, 49, 50, 3, 24,
		12, 0, 50, 51, 5, 9, 0, 0, 51, 52, 5, 2, 0, 0, 52, 7, 1, 0, 0, 0, 53, 54,
		3, 12, 6, 0, 54, 55, 5, 3, 0, 0, 55, 9, 1, 0, 0, 0, 56, 59, 3, 14, 7, 0,
		57, 58, 5, 4, 0, 0, 58, 60, 3, 10, 5, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1,
		0, 0, 0, 60, 11, 1, 0, 0, 0, 61, 62, 3, 20, 10, 0, 62, 13, 1, 0, 0, 0,
		63, 65, 3, 16, 8, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 70, 1,
		0, 0, 0, 66, 71, 3, 22, 11, 0, 67, 71, 3, 20, 10, 0, 68, 71, 3, 18, 9,
		0, 69, 71, 5, 5, 0, 0, 70, 66, 1, 0, 0, 0, 70, 67, 1, 0, 0, 0, 70, 68,
		1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 73, 7, 0, 0, 0,
		73, 75, 3, 22, 11, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 81, 1,
		0, 0, 0, 76, 77, 5, 1, 0, 0, 77, 78, 3, 14, 7, 0, 78, 79, 5, 2, 0, 0, 79,
		81, 1, 0, 0, 0, 80, 64, 1, 0, 0, 0, 80, 76, 1, 0, 0, 0, 81, 15, 1, 0, 0,
		0, 82, 83, 5, 8, 0, 0, 83, 17, 1, 0, 0, 0, 84, 85, 5, 30, 0, 0, 85, 19,
		1, 0, 0, 0, 86, 87, 5, 26, 0, 0, 87, 21, 1, 0, 0, 0, 88, 89, 7, 1, 0, 0,
		89, 23, 1, 0, 0, 0, 90, 91, 7, 2, 0, 0, 91, 25, 1, 0, 0, 0, 9, 29, 37,
		42, 46, 59, 64, 70, 74, 80,
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

// bParserInit initializes any static state used to implement bParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewbParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BParserInit() {
	staticData := &BParserStaticData
	staticData.once.Do(bParserInit)
}

// NewbParser produces a new parser instance for the optional input antlr.TokenStream.
func NewbParser(input antlr.TokenStream) *bParser {
	BParserInit()
	this := new(bParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "b.g4"

	return this
}

// bParser tokens.
const (
	bParserEOF     = antlr.TokenEOF
	bParserT__0    = 1
	bParserT__1    = 2
	bParserT__2    = 3
	bParserT__3    = 4
	bParserT__4    = 5
	bParserT__5    = 6
	bParserT__6    = 7
	bParserT__7    = 8
	bParserTIMES   = 9
	bParserADD     = 10
	bParserAND     = 11
	bParserBNE     = 12
	bParserBRA     = 13
	bParserDEC     = 14
	bParserINC     = 15
	bParserMOVH    = 16
	bParserMOVL    = 17
	bParserMOVS    = 18
	bParserLDR     = 19
	bParserPOP     = 20
	bParserPSH     = 21
	bParserSTR     = 22
	bParserXOR     = 23
	bParserSUB     = 24
	bParserNOP     = 25
	bParserNAME    = 26
	bParserNUMBER  = 27
	bParserHEX     = 28
	bParserCOMMENT = 29
	bParserSTRING  = 30
	bParserEOL     = 31
	bParserWS      = 32
)

// bParser rules.
const (
	bParserRULE_prog         = 0
	bParserRULE_line         = 1
	bParserRULE_instruction  = 2
	bParserRULE_repeat       = 3
	bParserRULE_lbl          = 4
	bParserRULE_argumentlist = 5
	bParserRULE_label        = 6
	bParserRULE_argument     = 7
	bParserRULE_prefix_      = 8
	bParserRULE_string_      = 9
	bParserRULE_name         = 10
	bParserRULE_number       = 11
	bParserRULE_opcode       = 12
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(bParserEOF, 0)
}

func (s *ProgContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *bParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2281700354) != 0 {
		{
			p.SetState(26)
			p.Line()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(bParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOL() antlr.TerminalNode
	Instruction() IInstructionContext
	Lbl() ILblContext
	Repeat() IRepeatContext

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) EOL() antlr.TerminalNode {
	return s.GetToken(bParserEOL, 0)
}

func (s *LineContext) Instruction() IInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *LineContext) Lbl() ILblContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILblContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILblContext)
}

func (s *LineContext) Repeat() IRepeatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *bParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bParserRULE_line)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(34)
			p.Instruction()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(35)
			p.Lbl()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(36)
			p.Repeat()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(39)
		p.Match(bParserEOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Opcode() IOpcodeContext
	Label() ILabelContext
	Argumentlist() IArgumentlistContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Opcode() IOpcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *InstructionContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *InstructionContext) Argumentlist() IArgumentlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *bParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, bParserRULE_instruction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bParserNAME {
		{
			p.SetState(41)
			p.Label()
		}

	}
	{
		p.SetState(44)
		p.Opcode()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1543504162) != 0 {
		{
			p.SetState(45)
			p.Argumentlist()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeatContext is an interface to support dynamic dispatch.
type IRepeatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Opcode() IOpcodeContext
	TIMES() antlr.TerminalNode

	// IsRepeatContext differentiates from other interfaces.
	IsRepeatContext()
}

type RepeatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatContext() *RepeatContext {
	var p = new(RepeatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_repeat
	return p
}

func InitEmptyRepeatContext(p *RepeatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_repeat
}

func (*RepeatContext) IsRepeatContext() {}

func NewRepeatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatContext {
	var p = new(RepeatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_repeat

	return p
}

func (s *RepeatContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatContext) Opcode() IOpcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *RepeatContext) TIMES() antlr.TerminalNode {
	return s.GetToken(bParserTIMES, 0)
}

func (s *RepeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterRepeat(s)
	}
}

func (s *RepeatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitRepeat(s)
	}
}

func (p *bParser) Repeat() (localctx IRepeatContext) {
	localctx = NewRepeatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bParserRULE_repeat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Opcode()
	}
	{
		p.SetState(50)
		p.Match(bParserTIMES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(bParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILblContext is an interface to support dynamic dispatch.
type ILblContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Label() ILabelContext

	// IsLblContext differentiates from other interfaces.
	IsLblContext()
}

type LblContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLblContext() *LblContext {
	var p = new(LblContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_lbl
	return p
}

func InitEmptyLblContext(p *LblContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_lbl
}

func (*LblContext) IsLblContext() {}

func NewLblContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LblContext {
	var p = new(LblContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_lbl

	return p
}

func (s *LblContext) GetParser() antlr.Parser { return s.parser }

func (s *LblContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LblContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LblContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LblContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterLbl(s)
	}
}

func (s *LblContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitLbl(s)
	}
}

func (p *bParser) Lbl() (localctx ILblContext) {
	localctx = NewLblContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bParserRULE_lbl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Label()
	}
	{
		p.SetState(54)
		p.Match(bParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentlistContext is an interface to support dynamic dispatch.
type IArgumentlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Argument() IArgumentContext
	Argumentlist() IArgumentlistContext

	// IsArgumentlistContext differentiates from other interfaces.
	IsArgumentlistContext()
}

type ArgumentlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentlistContext() *ArgumentlistContext {
	var p = new(ArgumentlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_argumentlist
	return p
}

func InitEmptyArgumentlistContext(p *ArgumentlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_argumentlist
}

func (*ArgumentlistContext) IsArgumentlistContext() {}

func NewArgumentlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentlistContext {
	var p = new(ArgumentlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_argumentlist

	return p
}

func (s *ArgumentlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentlistContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentlistContext) Argumentlist() IArgumentlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *ArgumentlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterArgumentlist(s)
	}
}

func (s *ArgumentlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitArgumentlist(s)
	}
}

func (p *bParser) Argumentlist() (localctx IArgumentlistContext) {
	localctx = NewArgumentlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bParserRULE_argumentlist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Argument()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bParserT__3 {
		{
			p.SetState(57)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Argumentlist()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *bParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bParserRULE_label)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumber() []INumberContext
	Number(i int) INumberContext
	Name() INameContext
	String_() IString_Context
	Prefix_() IPrefix_Context
	Argument() IArgumentContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) AllNumber() []INumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberContext); ok {
			len++
		}
	}

	tst := make([]INumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberContext); ok {
			tst[i] = t.(INumberContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentContext) Number(i int) INumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ArgumentContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ArgumentContext) Prefix_() IPrefix_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefix_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefix_Context)
}

func (s *ArgumentContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *bParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, bParserRULE_argument)
	var _la int

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bParserT__4, bParserT__7, bParserNAME, bParserNUMBER, bParserHEX, bParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bParserT__7 {
			{
				p.SetState(63)
				p.Prefix_()
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case bParserNUMBER, bParserHEX:
			{
				p.SetState(66)
				p.Number()
			}

		case bParserNAME:
			{
				p.SetState(67)
				p.Name()
			}

		case bParserSTRING:
			{
				p.SetState(68)
				p.String_()
			}

		case bParserT__4:
			{
				p.SetState(69)
				p.Match(bParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bParserT__5 || _la == bParserT__6 {
			{
				p.SetState(72)
				_la = p.GetTokenStream().LA(1)

				if !(_la == bParserT__5 || _la == bParserT__6) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(73)
				p.Number()
			}

		}

	case bParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Argument()
		}
		{
			p.SetState(78)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefix_Context is an interface to support dynamic dispatch.
type IPrefix_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrefix_Context differentiates from other interfaces.
	IsPrefix_Context()
}

type Prefix_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_Context() *Prefix_Context {
	var p = new(Prefix_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_prefix_
	return p
}

func InitEmptyPrefix_Context(p *Prefix_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_prefix_
}

func (*Prefix_Context) IsPrefix_Context() {}

func NewPrefix_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_Context {
	var p = new(Prefix_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_prefix_

	return p
}

func (s *Prefix_Context) GetParser() antlr.Parser { return s.parser }
func (s *Prefix_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prefix_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterPrefix_(s)
	}
}

func (s *Prefix_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitPrefix_(s)
	}
}

func (p *bParser) Prefix_() (localctx IPrefix_Context) {
	localctx = NewPrefix_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bParserRULE_prefix_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(bParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_string_
	return p
}

func InitEmptyString_Context(p *String_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_string_
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) STRING() antlr.TerminalNode {
	return s.GetToken(bParserSTRING, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitString_(s)
	}
}

func (p *bParser) String_() (localctx IString_Context) {
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bParserRULE_string_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(bParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(bParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *bParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, bParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(bParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	HEX() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(bParserNUMBER, 0)
}

func (s *NumberContext) HEX() antlr.TerminalNode {
	return s.GetToken(bParserHEX, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *bParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, bParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !(_la == bParserNUMBER || _la == bParserHEX) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpcodeContext is an interface to support dynamic dispatch.
type IOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	AND() antlr.TerminalNode
	BNE() antlr.TerminalNode
	BRA() antlr.TerminalNode
	DEC() antlr.TerminalNode
	INC() antlr.TerminalNode
	MOVH() antlr.TerminalNode
	MOVL() antlr.TerminalNode
	MOVS() antlr.TerminalNode
	LDR() antlr.TerminalNode
	POP() antlr.TerminalNode
	PSH() antlr.TerminalNode
	STR() antlr.TerminalNode
	SUB() antlr.TerminalNode
	XOR() antlr.TerminalNode
	NOP() antlr.TerminalNode

	// IsOpcodeContext differentiates from other interfaces.
	IsOpcodeContext()
}

type OpcodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodeContext() *OpcodeContext {
	var p = new(OpcodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_opcode
	return p
}

func InitEmptyOpcodeContext(p *OpcodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_opcode
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) ADD() antlr.TerminalNode {
	return s.GetToken(bParserADD, 0)
}

func (s *OpcodeContext) AND() antlr.TerminalNode {
	return s.GetToken(bParserAND, 0)
}

func (s *OpcodeContext) BNE() antlr.TerminalNode {
	return s.GetToken(bParserBNE, 0)
}

func (s *OpcodeContext) BRA() antlr.TerminalNode {
	return s.GetToken(bParserBRA, 0)
}

func (s *OpcodeContext) DEC() antlr.TerminalNode {
	return s.GetToken(bParserDEC, 0)
}

func (s *OpcodeContext) INC() antlr.TerminalNode {
	return s.GetToken(bParserINC, 0)
}

func (s *OpcodeContext) MOVH() antlr.TerminalNode {
	return s.GetToken(bParserMOVH, 0)
}

func (s *OpcodeContext) MOVL() antlr.TerminalNode {
	return s.GetToken(bParserMOVL, 0)
}

func (s *OpcodeContext) MOVS() antlr.TerminalNode {
	return s.GetToken(bParserMOVS, 0)
}

func (s *OpcodeContext) LDR() antlr.TerminalNode {
	return s.GetToken(bParserLDR, 0)
}

func (s *OpcodeContext) POP() antlr.TerminalNode {
	return s.GetToken(bParserPOP, 0)
}

func (s *OpcodeContext) PSH() antlr.TerminalNode {
	return s.GetToken(bParserPSH, 0)
}

func (s *OpcodeContext) STR() antlr.TerminalNode {
	return s.GetToken(bParserSTR, 0)
}

func (s *OpcodeContext) SUB() antlr.TerminalNode {
	return s.GetToken(bParserSUB, 0)
}

func (s *OpcodeContext) XOR() antlr.TerminalNode {
	return s.GetToken(bParserXOR, 0)
}

func (s *OpcodeContext) NOP() antlr.TerminalNode {
	return s.GetToken(bParserNOP, 0)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *bParser) Opcode() (localctx IOpcodeContext) {
	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, bParserRULE_opcode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67107840) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
