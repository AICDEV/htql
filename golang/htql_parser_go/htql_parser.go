// Code generated from /Users/aicdev/coding/htql/kotlin/app/src/main/antlr/Htql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package htql // Htql
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

type HtqlParser struct {
	*antlr.BaseParser
}

var HtqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func htqlParserInit() {
	staticData := &HtqlParserStaticData
	staticData.LiteralNames = []string{
		"", "'attributes'", "'SELECT'", "'FROM'", "'WHERE'", "'AND'", "'OR'",
		"'IS'", "'NOT'", "'NULL'", "'*'", "','", "'.'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "SELECT", "FROM", "WHERE", "AND", "OR", "IS", "NOT", "NULL",
		"STAR", "COMMA", "DOT", "EQUALS", "IDENTIFIER", "STRING", "WS", "URL",
	}
	staticData.RuleNames = []string{
		"query", "selectStmt", "elementList", "fromStmt", "whereStmt", "conditionExpr",
		"condition", "attributeExpr", "attributeNullCheck", "logicalOp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 75, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 3, 0, 24, 8, 0, 1, 1, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1, 2, 1, 2,
		1, 2, 5, 2, 34, 8, 2, 10, 2, 12, 2, 37, 9, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 49, 8, 5, 10, 5, 12, 5, 52, 9, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 58, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 3, 8, 69, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 0, 0,
		10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 2, 2, 0, 14, 14, 17, 17, 1, 0,
		5, 6, 71, 0, 20, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 30, 1, 0, 0, 0, 6, 38,
		1, 0, 0, 0, 8, 41, 1, 0, 0, 0, 10, 44, 1, 0, 0, 0, 12, 57, 1, 0, 0, 0,
		14, 59, 1, 0, 0, 0, 16, 65, 1, 0, 0, 0, 18, 72, 1, 0, 0, 0, 20, 21, 3,
		2, 1, 0, 21, 23, 3, 6, 3, 0, 22, 24, 3, 8, 4, 0, 23, 22, 1, 0, 0, 0, 23,
		24, 1, 0, 0, 0, 24, 1, 1, 0, 0, 0, 25, 28, 5, 2, 0, 0, 26, 29, 5, 10, 0,
		0, 27, 29, 3, 4, 2, 0, 28, 26, 1, 0, 0, 0, 28, 27, 1, 0, 0, 0, 29, 3, 1,
		0, 0, 0, 30, 35, 5, 14, 0, 0, 31, 32, 5, 11, 0, 0, 32, 34, 5, 14, 0, 0,
		33, 31, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1,
		0, 0, 0, 36, 5, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 3, 0, 0, 39,
		40, 7, 0, 0, 0, 40, 7, 1, 0, 0, 0, 41, 42, 5, 4, 0, 0, 42, 43, 3, 10, 5,
		0, 43, 9, 1, 0, 0, 0, 44, 50, 3, 12, 6, 0, 45, 46, 3, 18, 9, 0, 46, 47,
		3, 12, 6, 0, 47, 49, 1, 0, 0, 0, 48, 45, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0,
		50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 11, 1, 0, 0, 0, 52, 50, 1,
		0, 0, 0, 53, 58, 3, 14, 7, 0, 54, 58, 3, 16, 8, 0, 55, 56, 5, 8, 0, 0,
		56, 58, 3, 14, 7, 0, 57, 53, 1, 0, 0, 0, 57, 54, 1, 0, 0, 0, 57, 55, 1,
		0, 0, 0, 58, 13, 1, 0, 0, 0, 59, 60, 5, 1, 0, 0, 60, 61, 5, 12, 0, 0, 61,
		62, 5, 14, 0, 0, 62, 63, 5, 13, 0, 0, 63, 64, 5, 15, 0, 0, 64, 15, 1, 0,
		0, 0, 65, 66, 5, 1, 0, 0, 66, 68, 5, 7, 0, 0, 67, 69, 5, 8, 0, 0, 68, 67,
		1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 5, 9, 0, 0,
		71, 17, 1, 0, 0, 0, 72, 73, 7, 1, 0, 0, 73, 19, 1, 0, 0, 0, 6, 23, 28,
		35, 50, 57, 68,
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

// HtqlParserInit initializes any static state used to implement HtqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHtqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HtqlParserInit() {
	staticData := &HtqlParserStaticData
	staticData.once.Do(htqlParserInit)
}

// NewHtqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHtqlParser(input antlr.TokenStream) *HtqlParser {
	HtqlParserInit()
	this := new(HtqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &HtqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Htql.g4"

	return this
}

// HtqlParser tokens.
const (
	HtqlParserEOF        = antlr.TokenEOF
	HtqlParserT__0       = 1
	HtqlParserSELECT     = 2
	HtqlParserFROM       = 3
	HtqlParserWHERE      = 4
	HtqlParserAND        = 5
	HtqlParserOR         = 6
	HtqlParserIS         = 7
	HtqlParserNOT        = 8
	HtqlParserNULL       = 9
	HtqlParserSTAR       = 10
	HtqlParserCOMMA      = 11
	HtqlParserDOT        = 12
	HtqlParserEQUALS     = 13
	HtqlParserIDENTIFIER = 14
	HtqlParserSTRING     = 15
	HtqlParserWS         = 16
	HtqlParserURL        = 17
)

// HtqlParser rules.
const (
	HtqlParserRULE_query              = 0
	HtqlParserRULE_selectStmt         = 1
	HtqlParserRULE_elementList        = 2
	HtqlParserRULE_fromStmt           = 3
	HtqlParserRULE_whereStmt          = 4
	HtqlParserRULE_conditionExpr      = 5
	HtqlParserRULE_condition          = 6
	HtqlParserRULE_attributeExpr      = 7
	HtqlParserRULE_attributeNullCheck = 8
	HtqlParserRULE_logicalOp          = 9
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SelectStmt() ISelectStmtContext
	FromStmt() IFromStmtContext
	WhereStmt() IWhereStmtContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) SelectStmt() ISelectStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStmtContext)
}

func (s *QueryContext) FromStmt() IFromStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromStmtContext)
}

func (s *QueryContext) WhereStmt() IWhereStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereStmtContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HtqlParserRULE_query)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.SelectStmt()
	}
	{
		p.SetState(21)
		p.FromStmt()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HtqlParserWHERE {
		{
			p.SetState(22)
			p.WhereStmt()
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

// ISelectStmtContext is an interface to support dynamic dispatch.
type ISelectStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	STAR() antlr.TerminalNode
	ElementList() IElementListContext

	// IsSelectStmtContext differentiates from other interfaces.
	IsSelectStmtContext()
}

type SelectStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStmtContext() *SelectStmtContext {
	var p = new(SelectStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_selectStmt
	return p
}

func InitEmptySelectStmtContext(p *SelectStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_selectStmt
}

func (*SelectStmtContext) IsSelectStmtContext() {}

func NewSelectStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStmtContext {
	var p = new(SelectStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_selectStmt

	return p
}

func (s *SelectStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStmtContext) SELECT() antlr.TerminalNode {
	return s.GetToken(HtqlParserSELECT, 0)
}

func (s *SelectStmtContext) STAR() antlr.TerminalNode {
	return s.GetToken(HtqlParserSTAR, 0)
}

func (s *SelectStmtContext) ElementList() IElementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementListContext)
}

func (s *SelectStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterSelectStmt(s)
	}
}

func (s *SelectStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitSelectStmt(s)
	}
}

func (s *SelectStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitSelectStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) SelectStmt() (localctx ISelectStmtContext) {
	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HtqlParserRULE_selectStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(HtqlParserSELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HtqlParserSTAR:
		{
			p.SetState(26)
			p.Match(HtqlParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case HtqlParserIDENTIFIER:
		{
			p.SetState(27)
			p.ElementList()
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

// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_elementList
	return p
}

func InitEmptyElementListContext(p *ElementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_elementList
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_elementList

	return p
}

func (s *ElementListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(HtqlParserIDENTIFIER)
}

func (s *ElementListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(HtqlParserIDENTIFIER, i)
}

func (s *ElementListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(HtqlParserCOMMA)
}

func (s *ElementListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(HtqlParserCOMMA, i)
}

func (s *ElementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterElementList(s)
	}
}

func (s *ElementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitElementList(s)
	}
}

func (s *ElementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitElementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) ElementList() (localctx IElementListContext) {
	localctx = NewElementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HtqlParserRULE_elementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(HtqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HtqlParserCOMMA {
		{
			p.SetState(31)
			p.Match(HtqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.Match(HtqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IFromStmtContext is an interface to support dynamic dispatch.
type IFromStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	URL() antlr.TerminalNode

	// IsFromStmtContext differentiates from other interfaces.
	IsFromStmtContext()
}

type FromStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromStmtContext() *FromStmtContext {
	var p = new(FromStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_fromStmt
	return p
}

func InitEmptyFromStmtContext(p *FromStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_fromStmt
}

func (*FromStmtContext) IsFromStmtContext() {}

func NewFromStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromStmtContext {
	var p = new(FromStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_fromStmt

	return p
}

func (s *FromStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FromStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(HtqlParserFROM, 0)
}

func (s *FromStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(HtqlParserIDENTIFIER, 0)
}

func (s *FromStmtContext) URL() antlr.TerminalNode {
	return s.GetToken(HtqlParserURL, 0)
}

func (s *FromStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterFromStmt(s)
	}
}

func (s *FromStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitFromStmt(s)
	}
}

func (s *FromStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitFromStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) FromStmt() (localctx IFromStmtContext) {
	localctx = NewFromStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HtqlParserRULE_fromStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(HtqlParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		_la = p.GetTokenStream().LA(1)

		if !(_la == HtqlParserIDENTIFIER || _la == HtqlParserURL) {
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

// IWhereStmtContext is an interface to support dynamic dispatch.
type IWhereStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	ConditionExpr() IConditionExprContext

	// IsWhereStmtContext differentiates from other interfaces.
	IsWhereStmtContext()
}

type WhereStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereStmtContext() *WhereStmtContext {
	var p = new(WhereStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_whereStmt
	return p
}

func InitEmptyWhereStmtContext(p *WhereStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_whereStmt
}

func (*WhereStmtContext) IsWhereStmtContext() {}

func NewWhereStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereStmtContext {
	var p = new(WhereStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_whereStmt

	return p
}

func (s *WhereStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereStmtContext) WHERE() antlr.TerminalNode {
	return s.GetToken(HtqlParserWHERE, 0)
}

func (s *WhereStmtContext) ConditionExpr() IConditionExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionExprContext)
}

func (s *WhereStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterWhereStmt(s)
	}
}

func (s *WhereStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitWhereStmt(s)
	}
}

func (s *WhereStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitWhereStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) WhereStmt() (localctx IWhereStmtContext) {
	localctx = NewWhereStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HtqlParserRULE_whereStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(HtqlParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.ConditionExpr()
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

// IConditionExprContext is an interface to support dynamic dispatch.
type IConditionExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCondition() []IConditionContext
	Condition(i int) IConditionContext
	AllLogicalOp() []ILogicalOpContext
	LogicalOp(i int) ILogicalOpContext

	// IsConditionExprContext differentiates from other interfaces.
	IsConditionExprContext()
}

type ConditionExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionExprContext() *ConditionExprContext {
	var p = new(ConditionExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_conditionExpr
	return p
}

func InitEmptyConditionExprContext(p *ConditionExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_conditionExpr
}

func (*ConditionExprContext) IsConditionExprContext() {}

func NewConditionExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionExprContext {
	var p = new(ConditionExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_conditionExpr

	return p
}

func (s *ConditionExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionExprContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionExprContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
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

	return t.(IConditionContext)
}

func (s *ConditionExprContext) AllLogicalOp() []ILogicalOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalOpContext); ok {
			len++
		}
	}

	tst := make([]ILogicalOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalOpContext); ok {
			tst[i] = t.(ILogicalOpContext)
			i++
		}
	}

	return tst
}

func (s *ConditionExprContext) LogicalOp(i int) ILogicalOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOpContext); ok {
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

	return t.(ILogicalOpContext)
}

func (s *ConditionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterConditionExpr(s)
	}
}

func (s *ConditionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitConditionExpr(s)
	}
}

func (s *ConditionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitConditionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) ConditionExpr() (localctx IConditionExprContext) {
	localctx = NewConditionExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HtqlParserRULE_conditionExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Condition()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HtqlParserAND || _la == HtqlParserOR {
		{
			p.SetState(45)
			p.LogicalOp()
		}
		{
			p.SetState(46)
			p.Condition()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttributeExpr() IAttributeExprContext
	AttributeNullCheck() IAttributeNullCheckContext
	NOT() antlr.TerminalNode

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AttributeExpr() IAttributeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeExprContext)
}

func (s *ConditionContext) AttributeNullCheck() IAttributeNullCheckContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeNullCheckContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeNullCheckContext)
}

func (s *ConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(HtqlParserNOT, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HtqlParserRULE_condition)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.AttributeExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.AttributeNullCheck()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Match(HtqlParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.AttributeExpr()
		}

	case antlr.ATNInvalidAltNumber:
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

// IAttributeExprContext is an interface to support dynamic dispatch.
type IAttributeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsAttributeExprContext differentiates from other interfaces.
	IsAttributeExprContext()
}

type AttributeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeExprContext() *AttributeExprContext {
	var p = new(AttributeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_attributeExpr
	return p
}

func InitEmptyAttributeExprContext(p *AttributeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_attributeExpr
}

func (*AttributeExprContext) IsAttributeExprContext() {}

func NewAttributeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeExprContext {
	var p = new(AttributeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_attributeExpr

	return p
}

func (s *AttributeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(HtqlParserDOT, 0)
}

func (s *AttributeExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(HtqlParserIDENTIFIER, 0)
}

func (s *AttributeExprContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(HtqlParserEQUALS, 0)
}

func (s *AttributeExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(HtqlParserSTRING, 0)
}

func (s *AttributeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterAttributeExpr(s)
	}
}

func (s *AttributeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitAttributeExpr(s)
	}
}

func (s *AttributeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitAttributeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) AttributeExpr() (localctx IAttributeExprContext) {
	localctx = NewAttributeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HtqlParserRULE_attributeExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(HtqlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(HtqlParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(HtqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Match(HtqlParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(HtqlParserSTRING)
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

// IAttributeNullCheckContext is an interface to support dynamic dispatch.
type IAttributeNullCheckContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IS() antlr.TerminalNode
	NULL() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsAttributeNullCheckContext differentiates from other interfaces.
	IsAttributeNullCheckContext()
}

type AttributeNullCheckContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeNullCheckContext() *AttributeNullCheckContext {
	var p = new(AttributeNullCheckContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_attributeNullCheck
	return p
}

func InitEmptyAttributeNullCheckContext(p *AttributeNullCheckContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_attributeNullCheck
}

func (*AttributeNullCheckContext) IsAttributeNullCheckContext() {}

func NewAttributeNullCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeNullCheckContext {
	var p = new(AttributeNullCheckContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_attributeNullCheck

	return p
}

func (s *AttributeNullCheckContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeNullCheckContext) IS() antlr.TerminalNode {
	return s.GetToken(HtqlParserIS, 0)
}

func (s *AttributeNullCheckContext) NULL() antlr.TerminalNode {
	return s.GetToken(HtqlParserNULL, 0)
}

func (s *AttributeNullCheckContext) NOT() antlr.TerminalNode {
	return s.GetToken(HtqlParserNOT, 0)
}

func (s *AttributeNullCheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeNullCheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeNullCheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterAttributeNullCheck(s)
	}
}

func (s *AttributeNullCheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitAttributeNullCheck(s)
	}
}

func (s *AttributeNullCheckContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitAttributeNullCheck(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) AttributeNullCheck() (localctx IAttributeNullCheckContext) {
	localctx = NewAttributeNullCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HtqlParserRULE_attributeNullCheck)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(HtqlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(HtqlParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HtqlParserNOT {
		{
			p.SetState(67)
			p.Match(HtqlParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(70)
		p.Match(HtqlParserNULL)
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

// ILogicalOpContext is an interface to support dynamic dispatch.
type ILogicalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsLogicalOpContext differentiates from other interfaces.
	IsLogicalOpContext()
}

type LogicalOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOpContext() *LogicalOpContext {
	var p = new(LogicalOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_logicalOp
	return p
}

func InitEmptyLogicalOpContext(p *LogicalOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HtqlParserRULE_logicalOp
}

func (*LogicalOpContext) IsLogicalOpContext() {}

func NewLogicalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOpContext {
	var p = new(LogicalOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HtqlParserRULE_logicalOp

	return p
}

func (s *LogicalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOpContext) AND() antlr.TerminalNode {
	return s.GetToken(HtqlParserAND, 0)
}

func (s *LogicalOpContext) OR() antlr.TerminalNode {
	return s.GetToken(HtqlParserOR, 0)
}

func (s *LogicalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.EnterLogicalOp(s)
	}
}

func (s *LogicalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HtqlListener); ok {
		listenerT.ExitLogicalOp(s)
	}
}

func (s *LogicalOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HtqlVisitor:
		return t.VisitLogicalOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HtqlParser) LogicalOp() (localctx ILogicalOpContext) {
	localctx = NewLogicalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HtqlParserRULE_logicalOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		_la = p.GetTokenStream().LA(1)

		if !(_la == HtqlParserAND || _la == HtqlParserOR) {
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
