// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Sorter

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Sorter struct {
	*antlr.BaseParser
}

var sorterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sorterParserInit() {
	staticData := &sorterParserStaticData
	staticData.literalNames = []string{
		"", "'CASE'", "'WHEN'", "'END'", "", "", "'0'",
	}
	staticData.symbolicNames = []string{
		"", "CASE", "WHEN", "END", "IDENT", "NUMBER1", "ZERO", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"sorter", "whenStmt", "attribute", "prio",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 30, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		1, 0, 5, 0, 12, 8, 0, 10, 0, 12, 0, 15, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 26, 8, 2, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0,
		2, 4, 6, 0, 1, 1, 0, 5, 6, 28, 0, 8, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4,
		25, 1, 0, 0, 0, 6, 27, 1, 0, 0, 0, 8, 9, 5, 1, 0, 0, 9, 13, 5, 4, 0, 0,
		10, 12, 3, 2, 1, 0, 11, 10, 1, 0, 0, 0, 12, 15, 1, 0, 0, 0, 13, 11, 1,
		0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 16, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 16,
		17, 5, 3, 0, 0, 17, 1, 1, 0, 0, 0, 18, 19, 5, 2, 0, 0, 19, 20, 3, 4, 2,
		0, 20, 21, 3, 6, 3, 0, 21, 3, 1, 0, 0, 0, 22, 26, 5, 4, 0, 0, 23, 26, 3,
		6, 3, 0, 24, 26, 5, 7, 0, 0, 25, 22, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25,
		24, 1, 0, 0, 0, 26, 5, 1, 0, 0, 0, 27, 28, 7, 0, 0, 0, 28, 7, 1, 0, 0,
		0, 2, 13, 25,
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

// SorterInit initializes any static state used to implement Sorter. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSorter(). You can call this function if you wish to initialize the static state ahead
// of time.
func SorterInit() {
	staticData := &sorterParserStaticData
	staticData.once.Do(sorterParserInit)
}

// NewSorter produces a new parser instance for the optional input antlr.TokenStream.
func NewSorter(input antlr.TokenStream) *Sorter {
	SorterInit()
	this := new(Sorter)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sorterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// Sorter tokens.
const (
	SorterEOF     = antlr.TokenEOF
	SorterCASE    = 1
	SorterWHEN    = 2
	SorterEND     = 3
	SorterIDENT   = 4
	SorterNUMBER1 = 5
	SorterZERO    = 6
	SorterSTRING  = 7
	SorterWS      = 8
)

// Sorter rules.
const (
	SorterRULE_sorter    = 0
	SorterRULE_whenStmt  = 1
	SorterRULE_attribute = 2
	SorterRULE_prio      = 3
)

// ISorterContext is an interface to support dynamic dispatch.
type ISorterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSorterContext differentiates from other interfaces.
	IsSorterContext()
}

type SorterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySorterContext() *SorterContext {
	var p = new(SorterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SorterRULE_sorter
	return p
}

func (*SorterContext) IsSorterContext() {}

func NewSorterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SorterContext {
	var p = new(SorterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SorterRULE_sorter

	return p
}

func (s *SorterContext) GetParser() antlr.Parser { return s.parser }

func (s *SorterContext) CASE() antlr.TerminalNode {
	return s.GetToken(SorterCASE, 0)
}

func (s *SorterContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SorterIDENT, 0)
}

func (s *SorterContext) END() antlr.TerminalNode {
	return s.GetToken(SorterEND, 0)
}

func (s *SorterContext) AllWhenStmt() []IWhenStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhenStmtContext); ok {
			len++
		}
	}

	tst := make([]IWhenStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhenStmtContext); ok {
			tst[i] = t.(IWhenStmtContext)
			i++
		}
	}

	return tst
}

func (s *SorterContext) WhenStmt(i int) IWhenStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhenStmtContext); ok {
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

	return t.(IWhenStmtContext)
}

func (s *SorterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SorterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SorterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.EnterSorter(s)
	}
}

func (s *SorterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.ExitSorter(s)
	}
}

func (s *SorterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SorterVisitor:
		return t.VisitSorter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sorter) Sorter() (localctx ISorterContext) {
	this := p
	_ = this

	localctx = NewSorterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SorterRULE_sorter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.Match(SorterCASE)
	}
	{
		p.SetState(9)
		p.Match(SorterIDENT)
	}
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SorterWHEN {
		{
			p.SetState(10)
			p.WhenStmt()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(16)
		p.Match(SorterEND)
	}

	return localctx
}

// IWhenStmtContext is an interface to support dynamic dispatch.
type IWhenStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the Name rule contexts.
	GetName() IAttributeContext

	// GetPriority returns the Priority rule contexts.
	GetPriority() IPrioContext

	// SetName sets the Name rule contexts.
	SetName(IAttributeContext)

	// SetPriority sets the Priority rule contexts.
	SetPriority(IPrioContext)

	// IsWhenStmtContext differentiates from other interfaces.
	IsWhenStmtContext()
}

type WhenStmtContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Name     IAttributeContext
	Priority IPrioContext
}

func NewEmptyWhenStmtContext() *WhenStmtContext {
	var p = new(WhenStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SorterRULE_whenStmt
	return p
}

func (*WhenStmtContext) IsWhenStmtContext() {}

func NewWhenStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenStmtContext {
	var p = new(WhenStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SorterRULE_whenStmt

	return p
}

func (s *WhenStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenStmtContext) GetName() IAttributeContext { return s.Name }

func (s *WhenStmtContext) GetPriority() IPrioContext { return s.Priority }

func (s *WhenStmtContext) SetName(v IAttributeContext) { s.Name = v }

func (s *WhenStmtContext) SetPriority(v IPrioContext) { s.Priority = v }

func (s *WhenStmtContext) WHEN() antlr.TerminalNode {
	return s.GetToken(SorterWHEN, 0)
}

func (s *WhenStmtContext) Attribute() IAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *WhenStmtContext) Prio() IPrioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrioContext)
}

func (s *WhenStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.EnterWhenStmt(s)
	}
}

func (s *WhenStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.ExitWhenStmt(s)
	}
}

func (s *WhenStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SorterVisitor:
		return t.VisitWhenStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sorter) WhenStmt() (localctx IWhenStmtContext) {
	this := p
	_ = this

	localctx = NewWhenStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SorterRULE_whenStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(SorterWHEN)
	}
	{
		p.SetState(19)

		var _x = p.Attribute()

		localctx.(*WhenStmtContext).Name = _x
	}
	{
		p.SetState(20)

		var _x = p.Prio()

		localctx.(*WhenStmtContext).Priority = _x
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SorterRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SorterRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SorterIDENT, 0)
}

func (s *AttributeContext) Prio() IPrioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrioContext)
}

func (s *AttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(SorterSTRING, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (s *AttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SorterVisitor:
		return t.VisitAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sorter) Attribute() (localctx IAttributeContext) {
	this := p
	_ = this

	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SorterRULE_attribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(25)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SorterIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(SorterIDENT)
		}

	case SorterNUMBER1, SorterZERO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Prio()
		}

	case SorterSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Match(SorterSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrioContext is an interface to support dynamic dispatch.
type IPrioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrioContext differentiates from other interfaces.
	IsPrioContext()
}

type PrioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrioContext() *PrioContext {
	var p = new(PrioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SorterRULE_prio
	return p
}

func (*PrioContext) IsPrioContext() {}

func NewPrioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrioContext {
	var p = new(PrioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SorterRULE_prio

	return p
}

func (s *PrioContext) GetParser() antlr.Parser { return s.parser }

func (s *PrioContext) ZERO() antlr.TerminalNode {
	return s.GetToken(SorterZERO, 0)
}

func (s *PrioContext) NUMBER1() antlr.TerminalNode {
	return s.GetToken(SorterNUMBER1, 0)
}

func (s *PrioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.EnterPrio(s)
	}
}

func (s *PrioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SorterListener); ok {
		listenerT.ExitPrio(s)
	}
}

func (s *PrioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SorterVisitor:
		return t.VisitPrio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sorter) Prio() (localctx IPrioContext) {
	this := p
	_ = this

	localctx = NewPrioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SorterRULE_prio)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SorterNUMBER1 || _la == SorterZERO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
