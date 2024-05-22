// Code generated from parser/Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Query

import "github.com/antlr4-go/antlr/v4"

// BaseQueryListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryListener struct{}

var _ QueryListener = &BaseQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSearchQuery is called when production SearchQuery is entered.
func (s *BaseQueryListener) EnterSearchQuery(ctx *SearchQueryContext) {}

// ExitSearchQuery is called when production SearchQuery is exited.
func (s *BaseQueryListener) ExitSearchQuery(ctx *SearchQueryContext) {}

// EnterCountQuery is called when production CountQuery is entered.
func (s *BaseQueryListener) EnterCountQuery(ctx *CountQueryContext) {}

// ExitCountQuery is called when production CountQuery is exited.
func (s *BaseQueryListener) ExitCountQuery(ctx *CountQueryContext) {}

// EnterSearchClause is called when production searchClause is entered.
func (s *BaseQueryListener) EnterSearchClause(ctx *SearchClauseContext) {}

// ExitSearchClause is called when production searchClause is exited.
func (s *BaseQueryListener) ExitSearchClause(ctx *SearchClauseContext) {}

// EnterCountClause is called when production countClause is entered.
func (s *BaseQueryListener) EnterCountClause(ctx *CountClauseContext) {}

// ExitCountClause is called when production countClause is exited.
func (s *BaseQueryListener) ExitCountClause(ctx *CountClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseQueryListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseQueryListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseQueryListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseQueryListener) ExitCondition(ctx *ConditionContext) {}

// EnterAndCondition is called when production andCondition is entered.
func (s *BaseQueryListener) EnterAndCondition(ctx *AndConditionContext) {}

// ExitAndCondition is called when production andCondition is exited.
func (s *BaseQueryListener) ExitAndCondition(ctx *AndConditionContext) {}

// EnterOrCondition is called when production orCondition is entered.
func (s *BaseQueryListener) EnterOrCondition(ctx *OrConditionContext) {}

// ExitOrCondition is called when production orCondition is exited.
func (s *BaseQueryListener) ExitOrCondition(ctx *OrConditionContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseQueryListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseQueryListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseQueryListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseQueryListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOffsetClause is called when production offsetClause is entered.
func (s *BaseQueryListener) EnterOffsetClause(ctx *OffsetClauseContext) {}

// ExitOffsetClause is called when production offsetClause is exited.
func (s *BaseQueryListener) ExitOffsetClause(ctx *OffsetClauseContext) {}

// EnterComparisonCondition is called when production ComparisonCondition is entered.
func (s *BaseQueryListener) EnterComparisonCondition(ctx *ComparisonConditionContext) {}

// ExitComparisonCondition is called when production ComparisonCondition is exited.
func (s *BaseQueryListener) ExitComparisonCondition(ctx *ComparisonConditionContext) {}

// EnterGroupCondition is called when production GroupCondition is entered.
func (s *BaseQueryListener) EnterGroupCondition(ctx *GroupConditionContext) {}

// ExitGroupCondition is called when production GroupCondition is exited.
func (s *BaseQueryListener) ExitGroupCondition(ctx *GroupConditionContext) {}

// EnterLikeCondition is called when production LikeCondition is entered.
func (s *BaseQueryListener) EnterLikeCondition(ctx *LikeConditionContext) {}

// ExitLikeCondition is called when production LikeCondition is exited.
func (s *BaseQueryListener) ExitLikeCondition(ctx *LikeConditionContext) {}

// EnterIsNullCondition is called when production IsNullCondition is entered.
func (s *BaseQueryListener) EnterIsNullCondition(ctx *IsNullConditionContext) {}

// ExitIsNullCondition is called when production IsNullCondition is exited.
func (s *BaseQueryListener) ExitIsNullCondition(ctx *IsNullConditionContext) {}

// EnterInCondition is called when production InCondition is entered.
func (s *BaseQueryListener) EnterInCondition(ctx *InConditionContext) {}

// ExitInCondition is called when production InCondition is exited.
func (s *BaseQueryListener) ExitInCondition(ctx *InConditionContext) {}

// EnterOrderCondition is called when production orderCondition is entered.
func (s *BaseQueryListener) EnterOrderCondition(ctx *OrderConditionContext) {}

// ExitOrderCondition is called when production orderCondition is exited.
func (s *BaseQueryListener) ExitOrderCondition(ctx *OrderConditionContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseQueryListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseQueryListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseQueryListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseQueryListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterValue is called when production value is entered.
func (s *BaseQueryListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseQueryListener) ExitValue(ctx *ValueContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseQueryListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseQueryListener) ExitComparator(ctx *ComparatorContext) {}
