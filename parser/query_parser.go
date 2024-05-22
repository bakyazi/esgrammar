// Code generated from parser/Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Query

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

type QueryParser struct {
	*antlr.BaseParser
}

var QueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func queryParserInit() {
	staticData := &QueryParserStaticData
	staticData.LiteralNames = []string{
		"", "'SEARCH'", "'COUNT'", "'FROM'", "'WHERE'", "'ORDER'", "'BY'", "'LIMIT'",
		"'OFFSET'", "'AND'", "'OR'", "'NOT'", "'LIKE'", "'IN'", "", "'ASC'",
		"'DESC'", "", "", "", "", "'('", "')'", "'.'", "','", "'='", "'!='",
		"'>'", "'<'", "'>='", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "SEARCH", "COUNT", "FROM", "WHERE", "ORDER", "BY", "LIMIT", "OFFSET",
		"AND", "OR", "NOT", "LIKE", "IN", "ISNULL", "ASC", "DESC", "IDENTIFIER",
		"NUMBER", "STRING", "WS", "LPAR", "RPAR", "DOT", "COMMA", "EQUAL", "NOT_EQUAL",
		"GREATER", "LESSER", "GREATER_EQ", "LESSER_EQ",
	}
	staticData.RuleNames = []string{
		"query", "searchClause", "countClause", "whereClause", "condition",
		"andCondition", "orCondition", "orderByClause", "limitClause", "offsetClause",
		"conditionPart", "orderCondition", "indexName", "fieldName", "value",
		"comparator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 156, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 3, 0, 35, 8, 0, 1, 0, 3, 0, 38, 8, 0, 1, 0, 3, 0, 41, 8, 0,
		1, 0, 3, 0, 44, 8, 0, 1, 0, 1, 0, 3, 0, 48, 8, 0, 3, 0, 50, 8, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		3, 4, 65, 8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 70, 8, 5, 10, 5, 12, 5, 73, 9,
		5, 1, 6, 1, 6, 1, 6, 5, 6, 78, 8, 6, 10, 6, 12, 6, 81, 9, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 5, 7, 88, 8, 7, 10, 7, 12, 7, 91, 9, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 109, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 116, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 122, 8, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 129, 8, 10, 10, 10, 12, 10, 132, 9,
		10, 1, 10, 1, 10, 3, 10, 136, 8, 10, 1, 11, 1, 11, 3, 11, 140, 8, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 147, 8, 13, 10, 13, 12, 13, 150,
		9, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 2, 71, 79, 0, 16, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 3, 1, 0, 15, 16, 1, 0,
		18, 19, 1, 0, 25, 30, 159, 0, 49, 1, 0, 0, 0, 2, 51, 1, 0, 0, 0, 4, 55,
		1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 66, 1, 0, 0, 0, 12,
		74, 1, 0, 0, 0, 14, 82, 1, 0, 0, 0, 16, 92, 1, 0, 0, 0, 18, 95, 1, 0, 0,
		0, 20, 135, 1, 0, 0, 0, 22, 137, 1, 0, 0, 0, 24, 141, 1, 0, 0, 0, 26, 143,
		1, 0, 0, 0, 28, 151, 1, 0, 0, 0, 30, 153, 1, 0, 0, 0, 32, 34, 3, 2, 1,
		0, 33, 35, 3, 6, 3, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 37,
		1, 0, 0, 0, 36, 38, 3, 14, 7, 0, 37, 36, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0,
		38, 40, 1, 0, 0, 0, 39, 41, 3, 16, 8, 0, 40, 39, 1, 0, 0, 0, 40, 41, 1,
		0, 0, 0, 41, 43, 1, 0, 0, 0, 42, 44, 3, 18, 9, 0, 43, 42, 1, 0, 0, 0, 43,
		44, 1, 0, 0, 0, 44, 50, 1, 0, 0, 0, 45, 47, 3, 4, 2, 0, 46, 48, 3, 6, 3,
		0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1, 0, 0, 0, 49, 32,
		1, 0, 0, 0, 49, 45, 1, 0, 0, 0, 50, 1, 1, 0, 0, 0, 51, 52, 5, 1, 0, 0,
		52, 53, 5, 3, 0, 0, 53, 54, 3, 24, 12, 0, 54, 3, 1, 0, 0, 0, 55, 56, 5,
		2, 0, 0, 56, 57, 5, 3, 0, 0, 57, 58, 3, 24, 12, 0, 58, 5, 1, 0, 0, 0, 59,
		60, 5, 4, 0, 0, 60, 61, 3, 8, 4, 0, 61, 7, 1, 0, 0, 0, 62, 65, 3, 10, 5,
		0, 63, 65, 3, 12, 6, 0, 64, 62, 1, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 9,
		1, 0, 0, 0, 66, 71, 3, 20, 10, 0, 67, 68, 5, 9, 0, 0, 68, 70, 3, 20, 10,
		0, 69, 67, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 71, 69,
		1, 0, 0, 0, 72, 11, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 79, 3, 20, 10,
		0, 75, 76, 5, 10, 0, 0, 76, 78, 3, 20, 10, 0, 77, 75, 1, 0, 0, 0, 78, 81,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 13, 1, 0, 0, 0,
		81, 79, 1, 0, 0, 0, 82, 83, 5, 5, 0, 0, 83, 84, 5, 6, 0, 0, 84, 89, 3,
		22, 11, 0, 85, 86, 5, 24, 0, 0, 86, 88, 3, 22, 11, 0, 87, 85, 1, 0, 0,
		0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 15,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 7, 0, 0, 93, 94, 5, 18, 0, 0,
		94, 17, 1, 0, 0, 0, 95, 96, 5, 8, 0, 0, 96, 97, 5, 18, 0, 0, 97, 19, 1,
		0, 0, 0, 98, 99, 3, 26, 13, 0, 99, 100, 3, 30, 15, 0, 100, 101, 3, 28,
		14, 0, 101, 136, 1, 0, 0, 0, 102, 103, 5, 21, 0, 0, 103, 104, 3, 8, 4,
		0, 104, 105, 5, 22, 0, 0, 105, 136, 1, 0, 0, 0, 106, 108, 3, 26, 13, 0,
		107, 109, 5, 11, 0, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		110, 1, 0, 0, 0, 110, 111, 5, 12, 0, 0, 111, 112, 5, 19, 0, 0, 112, 136,
		1, 0, 0, 0, 113, 115, 3, 26, 13, 0, 114, 116, 5, 11, 0, 0, 115, 114, 1,
		0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 5, 14, 0,
		0, 118, 136, 1, 0, 0, 0, 119, 121, 3, 26, 13, 0, 120, 122, 5, 11, 0, 0,
		121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		124, 5, 13, 0, 0, 124, 125, 5, 21, 0, 0, 125, 130, 3, 28, 14, 0, 126, 127,
		5, 24, 0, 0, 127, 129, 3, 28, 14, 0, 128, 126, 1, 0, 0, 0, 129, 132, 1,
		0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 133, 1, 0, 0,
		0, 132, 130, 1, 0, 0, 0, 133, 134, 5, 22, 0, 0, 134, 136, 1, 0, 0, 0, 135,
		98, 1, 0, 0, 0, 135, 102, 1, 0, 0, 0, 135, 106, 1, 0, 0, 0, 135, 113, 1,
		0, 0, 0, 135, 119, 1, 0, 0, 0, 136, 21, 1, 0, 0, 0, 137, 139, 3, 26, 13,
		0, 138, 140, 7, 0, 0, 0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		23, 1, 0, 0, 0, 141, 142, 5, 17, 0, 0, 142, 25, 1, 0, 0, 0, 143, 148, 5,
		17, 0, 0, 144, 145, 5, 23, 0, 0, 145, 147, 5, 17, 0, 0, 146, 144, 1, 0,
		0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0,
		149, 27, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 7, 1, 0, 0, 152, 29,
		1, 0, 0, 0, 153, 154, 7, 2, 0, 0, 154, 31, 1, 0, 0, 0, 17, 34, 37, 40,
		43, 47, 49, 64, 71, 79, 89, 108, 115, 121, 130, 135, 139, 148,
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

// QueryParserInit initializes any static state used to implement QueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryParserInit() {
	staticData := &QueryParserStaticData
	staticData.once.Do(queryParserInit)
}

// NewQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQueryParser(input antlr.TokenStream) *QueryParser {
	QueryParserInit()
	this := new(QueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &QueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Query.g4"

	return this
}

// QueryParser tokens.
const (
	QueryParserEOF        = antlr.TokenEOF
	QueryParserSEARCH     = 1
	QueryParserCOUNT      = 2
	QueryParserFROM       = 3
	QueryParserWHERE      = 4
	QueryParserORDER      = 5
	QueryParserBY         = 6
	QueryParserLIMIT      = 7
	QueryParserOFFSET     = 8
	QueryParserAND        = 9
	QueryParserOR         = 10
	QueryParserNOT        = 11
	QueryParserLIKE       = 12
	QueryParserIN         = 13
	QueryParserISNULL     = 14
	QueryParserASC        = 15
	QueryParserDESC       = 16
	QueryParserIDENTIFIER = 17
	QueryParserNUMBER     = 18
	QueryParserSTRING     = 19
	QueryParserWS         = 20
	QueryParserLPAR       = 21
	QueryParserRPAR       = 22
	QueryParserDOT        = 23
	QueryParserCOMMA      = 24
	QueryParserEQUAL      = 25
	QueryParserNOT_EQUAL  = 26
	QueryParserGREATER    = 27
	QueryParserLESSER     = 28
	QueryParserGREATER_EQ = 29
	QueryParserLESSER_EQ  = 30
)

// QueryParser rules.
const (
	QueryParserRULE_query          = 0
	QueryParserRULE_searchClause   = 1
	QueryParserRULE_countClause    = 2
	QueryParserRULE_whereClause    = 3
	QueryParserRULE_condition      = 4
	QueryParserRULE_andCondition   = 5
	QueryParserRULE_orCondition    = 6
	QueryParserRULE_orderByClause  = 7
	QueryParserRULE_limitClause    = 8
	QueryParserRULE_offsetClause   = 9
	QueryParserRULE_conditionPart  = 10
	QueryParserRULE_orderCondition = 11
	QueryParserRULE_indexName      = 12
	QueryParserRULE_fieldName      = 13
	QueryParserRULE_value          = 14
	QueryParserRULE_comparator     = 15
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = QueryParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyAll(ctx *QueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SearchQueryContext struct {
	QueryContext
}

func NewSearchQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SearchQueryContext {
	var p = new(SearchQueryContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *SearchQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchQueryContext) SearchClause() ISearchClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchClauseContext)
}

func (s *SearchQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SearchQueryContext) OrderByClause() IOrderByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *SearchQueryContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SearchQueryContext) OffsetClause() IOffsetClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffsetClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffsetClauseContext)
}

func (s *SearchQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterSearchQuery(s)
	}
}

func (s *SearchQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitSearchQuery(s)
	}
}

type CountQueryContext struct {
	QueryContext
}

func NewCountQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountQueryContext {
	var p = new(CountQueryContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *CountQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountQueryContext) CountClause() ICountClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountClauseContext)
}

func (s *CountQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *CountQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterCountQuery(s)
	}
}

func (s *CountQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitCountQuery(s)
	}
}

func (p *QueryParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryParserRULE_query)
	var _la int

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryParserSEARCH:
		localctx = NewSearchQueryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.SearchClause()
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserWHERE {
			{
				p.SetState(33)
				p.WhereClause()
			}

		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserORDER {
			{
				p.SetState(36)
				p.OrderByClause()
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserLIMIT {
			{
				p.SetState(39)
				p.LimitClause()
			}

		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserOFFSET {
			{
				p.SetState(42)
				p.OffsetClause()
			}

		}

	case QueryParserCOUNT:
		localctx = NewCountQueryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.CountClause()
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserWHERE {
			{
				p.SetState(46)
				p.WhereClause()
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

// ISearchClauseContext is an interface to support dynamic dispatch.
type ISearchClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEARCH() antlr.TerminalNode
	FROM() antlr.TerminalNode
	IndexName() IIndexNameContext

	// IsSearchClauseContext differentiates from other interfaces.
	IsSearchClauseContext()
}

type SearchClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchClauseContext() *SearchClauseContext {
	var p = new(SearchClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_searchClause
	return p
}

func InitEmptySearchClauseContext(p *SearchClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_searchClause
}

func (*SearchClauseContext) IsSearchClauseContext() {}

func NewSearchClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchClauseContext {
	var p = new(SearchClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_searchClause

	return p
}

func (s *SearchClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchClauseContext) SEARCH() antlr.TerminalNode {
	return s.GetToken(QueryParserSEARCH, 0)
}

func (s *SearchClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(QueryParserFROM, 0)
}

func (s *SearchClauseContext) IndexName() IIndexNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *SearchClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterSearchClause(s)
	}
}

func (s *SearchClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitSearchClause(s)
	}
}

func (p *QueryParser) SearchClause() (localctx ISearchClauseContext) {
	localctx = NewSearchClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryParserRULE_searchClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(QueryParserSEARCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(QueryParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.IndexName()
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

// ICountClauseContext is an interface to support dynamic dispatch.
type ICountClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUNT() antlr.TerminalNode
	FROM() antlr.TerminalNode
	IndexName() IIndexNameContext

	// IsCountClauseContext differentiates from other interfaces.
	IsCountClauseContext()
}

type CountClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountClauseContext() *CountClauseContext {
	var p = new(CountClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_countClause
	return p
}

func InitEmptyCountClauseContext(p *CountClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_countClause
}

func (*CountClauseContext) IsCountClauseContext() {}

func NewCountClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountClauseContext {
	var p = new(CountClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_countClause

	return p
}

func (s *CountClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CountClauseContext) COUNT() antlr.TerminalNode {
	return s.GetToken(QueryParserCOUNT, 0)
}

func (s *CountClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(QueryParserFROM, 0)
}

func (s *CountClauseContext) IndexName() IIndexNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *CountClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterCountClause(s)
	}
}

func (s *CountClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitCountClause(s)
	}
}

func (p *QueryParser) CountClause() (localctx ICountClauseContext) {
	localctx = NewCountClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryParserRULE_countClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(QueryParserCOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(QueryParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.IndexName()
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

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	Condition() IConditionContext

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_whereClause
	return p
}

func InitEmptyWhereClauseContext(p *WhereClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_whereClause
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(QueryParserWHERE, 0)
}

func (s *WhereClauseContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (p *QueryParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryParserRULE_whereClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(QueryParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Condition()
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
	AndCondition() IAndConditionContext
	OrCondition() IOrConditionContext

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
	p.RuleIndex = QueryParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AndCondition() IAndConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndConditionContext)
}

func (s *ConditionContext) OrCondition() IOrConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrConditionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *QueryParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(62)
			p.AndCondition()
		}

	case 2:
		{
			p.SetState(63)
			p.OrCondition()
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

// IAndConditionContext is an interface to support dynamic dispatch.
type IAndConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConditionPart() []IConditionPartContext
	ConditionPart(i int) IConditionPartContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsAndConditionContext differentiates from other interfaces.
	IsAndConditionContext()
}

type AndConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndConditionContext() *AndConditionContext {
	var p = new(AndConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_andCondition
	return p
}

func InitEmptyAndConditionContext(p *AndConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_andCondition
}

func (*AndConditionContext) IsAndConditionContext() {}

func NewAndConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndConditionContext {
	var p = new(AndConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_andCondition

	return p
}

func (s *AndConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndConditionContext) AllConditionPart() []IConditionPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionPartContext); ok {
			len++
		}
	}

	tst := make([]IConditionPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionPartContext); ok {
			tst[i] = t.(IConditionPartContext)
			i++
		}
	}

	return tst
}

func (s *AndConditionContext) ConditionPart(i int) IConditionPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionPartContext); ok {
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

	return t.(IConditionPartContext)
}

func (s *AndConditionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(QueryParserAND)
}

func (s *AndConditionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(QueryParserAND, i)
}

func (s *AndConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterAndCondition(s)
	}
}

func (s *AndConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitAndCondition(s)
	}
}

func (p *QueryParser) AndCondition() (localctx IAndConditionContext) {
	localctx = NewAndConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QueryParserRULE_andCondition)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.ConditionPart()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(67)
				p.Match(QueryParserAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(68)
				p.ConditionPart()
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
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

// IOrConditionContext is an interface to support dynamic dispatch.
type IOrConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConditionPart() []IConditionPartContext
	ConditionPart(i int) IConditionPartContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsOrConditionContext differentiates from other interfaces.
	IsOrConditionContext()
}

type OrConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrConditionContext() *OrConditionContext {
	var p = new(OrConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_orCondition
	return p
}

func InitEmptyOrConditionContext(p *OrConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_orCondition
}

func (*OrConditionContext) IsOrConditionContext() {}

func NewOrConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrConditionContext {
	var p = new(OrConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_orCondition

	return p
}

func (s *OrConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrConditionContext) AllConditionPart() []IConditionPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionPartContext); ok {
			len++
		}
	}

	tst := make([]IConditionPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionPartContext); ok {
			tst[i] = t.(IConditionPartContext)
			i++
		}
	}

	return tst
}

func (s *OrConditionContext) ConditionPart(i int) IConditionPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionPartContext); ok {
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

	return t.(IConditionPartContext)
}

func (s *OrConditionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(QueryParserOR)
}

func (s *OrConditionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(QueryParserOR, i)
}

func (s *OrConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterOrCondition(s)
	}
}

func (s *OrConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitOrCondition(s)
	}
}

func (p *QueryParser) OrCondition() (localctx IOrConditionContext) {
	localctx = NewOrConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryParserRULE_orCondition)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.ConditionPart()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(75)
				p.Match(QueryParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(76)
				p.ConditionPart()
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
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

// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllOrderCondition() []IOrderConditionContext
	OrderCondition(i int) IOrderConditionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_orderByClause
	return p
}

func InitEmptyOrderByClauseContext(p *OrderByClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_orderByClause
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(QueryParserORDER, 0)
}

func (s *OrderByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(QueryParserBY, 0)
}

func (s *OrderByClauseContext) AllOrderCondition() []IOrderConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderConditionContext); ok {
			len++
		}
	}

	tst := make([]IOrderConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderConditionContext); ok {
			tst[i] = t.(IOrderConditionContext)
			i++
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderCondition(i int) IOrderConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderConditionContext); ok {
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

	return t.(IOrderConditionContext)
}

func (s *OrderByClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QueryParserCOMMA)
}

func (s *OrderByClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QueryParserCOMMA, i)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (p *QueryParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QueryParserRULE_orderByClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(QueryParserORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(QueryParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.OrderCondition()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryParserCOMMA {
		{
			p.SetState(85)
			p.Match(QueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.OrderCondition()
		}

		p.SetState(91)
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

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_limitClause
	return p
}

func InitEmptyLimitClauseContext(p *LimitClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_limitClause
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(QueryParserLIMIT, 0)
}

func (s *LimitClauseContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (p *QueryParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QueryParserRULE_limitClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(QueryParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(QueryParserNUMBER)
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

// IOffsetClauseContext is an interface to support dynamic dispatch.
type IOffsetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsOffsetClauseContext differentiates from other interfaces.
	IsOffsetClauseContext()
}

type OffsetClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetClauseContext() *OffsetClauseContext {
	var p = new(OffsetClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_offsetClause
	return p
}

func InitEmptyOffsetClauseContext(p *OffsetClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_offsetClause
}

func (*OffsetClauseContext) IsOffsetClauseContext() {}

func NewOffsetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetClauseContext {
	var p = new(OffsetClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_offsetClause

	return p
}

func (s *OffsetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(QueryParserOFFSET, 0)
}

func (s *OffsetClauseContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER, 0)
}

func (s *OffsetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterOffsetClause(s)
	}
}

func (s *OffsetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitOffsetClause(s)
	}
}

func (p *QueryParser) OffsetClause() (localctx IOffsetClauseContext) {
	localctx = NewOffsetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QueryParserRULE_offsetClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(QueryParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(QueryParserNUMBER)
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

// IConditionPartContext is an interface to support dynamic dispatch.
type IConditionPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionPartContext differentiates from other interfaces.
	IsConditionPartContext()
}

type ConditionPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionPartContext() *ConditionPartContext {
	var p = new(ConditionPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_conditionPart
	return p
}

func InitEmptyConditionPartContext(p *ConditionPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_conditionPart
}

func (*ConditionPartContext) IsConditionPartContext() {}

func NewConditionPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionPartContext {
	var p = new(ConditionPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_conditionPart

	return p
}

func (s *ConditionPartContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionPartContext) CopyAll(ctx *ConditionPartContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConditionPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ComparisonConditionContext struct {
	ConditionPartContext
}

func NewComparisonConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonConditionContext {
	var p = new(ComparisonConditionContext)

	InitEmptyConditionPartContext(&p.ConditionPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionPartContext))

	return p
}

func (s *ComparisonConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonConditionContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *ComparisonConditionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ComparisonConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ComparisonConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterComparisonCondition(s)
	}
}

func (s *ComparisonConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitComparisonCondition(s)
	}
}

type InConditionContext struct {
	ConditionPartContext
}

func NewInConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InConditionContext {
	var p = new(InConditionContext)

	InitEmptyConditionPartContext(&p.ConditionPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionPartContext))

	return p
}

func (s *InConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InConditionContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *InConditionContext) IN() antlr.TerminalNode {
	return s.GetToken(QueryParserIN, 0)
}

func (s *InConditionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(QueryParserLPAR, 0)
}

func (s *InConditionContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *InConditionContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *InConditionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(QueryParserRPAR, 0)
}

func (s *InConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *InConditionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QueryParserCOMMA)
}

func (s *InConditionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QueryParserCOMMA, i)
}

func (s *InConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterInCondition(s)
	}
}

func (s *InConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitInCondition(s)
	}
}

type IsNullConditionContext struct {
	ConditionPartContext
}

func NewIsNullConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullConditionContext {
	var p = new(IsNullConditionContext)

	InitEmptyConditionPartContext(&p.ConditionPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionPartContext))

	return p
}

func (s *IsNullConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullConditionContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *IsNullConditionContext) ISNULL() antlr.TerminalNode {
	return s.GetToken(QueryParserISNULL, 0)
}

func (s *IsNullConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *IsNullConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterIsNullCondition(s)
	}
}

func (s *IsNullConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitIsNullCondition(s)
	}
}

type GroupConditionContext struct {
	ConditionPartContext
}

func NewGroupConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupConditionContext {
	var p = new(GroupConditionContext)

	InitEmptyConditionPartContext(&p.ConditionPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionPartContext))

	return p
}

func (s *GroupConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupConditionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(QueryParserLPAR, 0)
}

func (s *GroupConditionContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *GroupConditionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(QueryParserRPAR, 0)
}

func (s *GroupConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterGroupCondition(s)
	}
}

func (s *GroupConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitGroupCondition(s)
	}
}

type LikeConditionContext struct {
	ConditionPartContext
}

func NewLikeConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeConditionContext {
	var p = new(LikeConditionContext)

	InitEmptyConditionPartContext(&p.ConditionPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionPartContext))

	return p
}

func (s *LikeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeConditionContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *LikeConditionContext) LIKE() antlr.TerminalNode {
	return s.GetToken(QueryParserLIKE, 0)
}

func (s *LikeConditionContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *LikeConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *LikeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterLikeCondition(s)
	}
}

func (s *LikeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitLikeCondition(s)
	}
}

func (p *QueryParser) ConditionPart() (localctx IConditionPartContext) {
	localctx = NewConditionPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QueryParserRULE_conditionPart)
	var _la int

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.FieldName()
		}
		{
			p.SetState(99)
			p.Comparator()
		}
		{
			p.SetState(100)
			p.Value()
		}

	case 2:
		localctx = NewGroupConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(QueryParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Condition()
		}
		{
			p.SetState(104)
			p.Match(QueryParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewLikeConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.FieldName()
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserNOT {
			{
				p.SetState(107)
				p.Match(QueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(110)
			p.Match(QueryParserLIKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(QueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIsNullConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.FieldName()
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserNOT {
			{
				p.SetState(114)
				p.Match(QueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(117)
			p.Match(QueryParserISNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewInConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(119)
			p.FieldName()
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserNOT {
			{
				p.SetState(120)
				p.Match(QueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(123)
			p.Match(QueryParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(QueryParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Value()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == QueryParserCOMMA {
			{
				p.SetState(126)
				p.Match(QueryParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(127)
				p.Value()
			}

			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(133)
			p.Match(QueryParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IOrderConditionContext is an interface to support dynamic dispatch.
type IOrderConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	ASC() antlr.TerminalNode
	DESC() antlr.TerminalNode

	// IsOrderConditionContext differentiates from other interfaces.
	IsOrderConditionContext()
}

type OrderConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderConditionContext() *OrderConditionContext {
	var p = new(OrderConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_orderCondition
	return p
}

func InitEmptyOrderConditionContext(p *OrderConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_orderCondition
}

func (*OrderConditionContext) IsOrderConditionContext() {}

func NewOrderConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderConditionContext {
	var p = new(OrderConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_orderCondition

	return p
}

func (s *OrderConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderConditionContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OrderConditionContext) ASC() antlr.TerminalNode {
	return s.GetToken(QueryParserASC, 0)
}

func (s *OrderConditionContext) DESC() antlr.TerminalNode {
	return s.GetToken(QueryParserDESC, 0)
}

func (s *OrderConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterOrderCondition(s)
	}
}

func (s *OrderConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitOrderCondition(s)
	}
}

func (p *QueryParser) OrderCondition() (localctx IOrderConditionContext) {
	localctx = NewOrderConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, QueryParserRULE_orderCondition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.FieldName()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryParserASC || _la == QueryParserDESC {
		{
			p.SetState(138)
			_la = p.GetTokenStream().LA(1)

			if !(_la == QueryParserASC || _la == QueryParserDESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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

// IIndexNameContext is an interface to support dynamic dispatch.
type IIndexNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIndexNameContext differentiates from other interfaces.
	IsIndexNameContext()
}

type IndexNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexNameContext() *IndexNameContext {
	var p = new(IndexNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_indexName
	return p
}

func InitEmptyIndexNameContext(p *IndexNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_indexName
}

func (*IndexNameContext) IsIndexNameContext() {}

func NewIndexNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexNameContext {
	var p = new(IndexNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_indexName

	return p
}

func (s *IndexNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QueryParserIDENTIFIER, 0)
}

func (s *IndexNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterIndexName(s)
	}
}

func (s *IndexNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitIndexName(s)
	}
}

func (p *QueryParser) IndexName() (localctx IIndexNameContext) {
	localctx = NewIndexNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, QueryParserRULE_indexName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(QueryParserIDENTIFIER)
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

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(QueryParserIDENTIFIER)
}

func (s *FieldNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(QueryParserIDENTIFIER, i)
}

func (s *FieldNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(QueryParserDOT)
}

func (s *FieldNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(QueryParserDOT, i)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *QueryParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, QueryParserRULE_fieldName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(QueryParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryParserDOT {
		{
			p.SetState(144)
			p.Match(QueryParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(QueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(150)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *QueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, QueryParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryParserNUMBER || _la == QueryParserSTRING) {
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

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	GREATER() antlr.TerminalNode
	GREATER_EQ() antlr.TerminalNode
	LESSER() antlr.TerminalNode
	LESSER_EQ() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryParserEQUAL, 0)
}

func (s *ComparatorContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT_EQUAL, 0)
}

func (s *ComparatorContext) GREATER() antlr.TerminalNode {
	return s.GetToken(QueryParserGREATER, 0)
}

func (s *ComparatorContext) GREATER_EQ() antlr.TerminalNode {
	return s.GetToken(QueryParserGREATER_EQ, 0)
}

func (s *ComparatorContext) LESSER() antlr.TerminalNode {
	return s.GetToken(QueryParserLESSER, 0)
}

func (s *ComparatorContext) LESSER_EQ() antlr.TerminalNode {
	return s.GetToken(QueryParserLESSER_EQ, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *QueryParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, QueryParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2113929216) != 0) {
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
