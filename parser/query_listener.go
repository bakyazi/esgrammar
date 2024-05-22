// Code generated from parser/Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Query

import "github.com/antlr4-go/antlr/v4"

// QueryListener is a complete listener for a parse tree produced by QueryParser.
type QueryListener interface {
	antlr.ParseTreeListener

	// EnterSearchQuery is called when entering the SearchQuery production.
	EnterSearchQuery(c *SearchQueryContext)

	// EnterCountQuery is called when entering the CountQuery production.
	EnterCountQuery(c *CountQueryContext)

	// EnterSearchClause is called when entering the searchClause production.
	EnterSearchClause(c *SearchClauseContext)

	// EnterCountClause is called when entering the countClause production.
	EnterCountClause(c *CountClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterAndCondition is called when entering the andCondition production.
	EnterAndCondition(c *AndConditionContext)

	// EnterOrCondition is called when entering the orCondition production.
	EnterOrCondition(c *OrConditionContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOffsetClause is called when entering the offsetClause production.
	EnterOffsetClause(c *OffsetClauseContext)

	// EnterComparisonCondition is called when entering the ComparisonCondition production.
	EnterComparisonCondition(c *ComparisonConditionContext)

	// EnterGroupCondition is called when entering the GroupCondition production.
	EnterGroupCondition(c *GroupConditionContext)

	// EnterLikeCondition is called when entering the LikeCondition production.
	EnterLikeCondition(c *LikeConditionContext)

	// EnterIsNullCondition is called when entering the IsNullCondition production.
	EnterIsNullCondition(c *IsNullConditionContext)

	// EnterInCondition is called when entering the InCondition production.
	EnterInCondition(c *InConditionContext)

	// EnterOrderCondition is called when entering the orderCondition production.
	EnterOrderCondition(c *OrderConditionContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// ExitSearchQuery is called when exiting the SearchQuery production.
	ExitSearchQuery(c *SearchQueryContext)

	// ExitCountQuery is called when exiting the CountQuery production.
	ExitCountQuery(c *CountQueryContext)

	// ExitSearchClause is called when exiting the searchClause production.
	ExitSearchClause(c *SearchClauseContext)

	// ExitCountClause is called when exiting the countClause production.
	ExitCountClause(c *CountClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitAndCondition is called when exiting the andCondition production.
	ExitAndCondition(c *AndConditionContext)

	// ExitOrCondition is called when exiting the orCondition production.
	ExitOrCondition(c *OrConditionContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOffsetClause is called when exiting the offsetClause production.
	ExitOffsetClause(c *OffsetClauseContext)

	// ExitComparisonCondition is called when exiting the ComparisonCondition production.
	ExitComparisonCondition(c *ComparisonConditionContext)

	// ExitGroupCondition is called when exiting the GroupCondition production.
	ExitGroupCondition(c *GroupConditionContext)

	// ExitLikeCondition is called when exiting the LikeCondition production.
	ExitLikeCondition(c *LikeConditionContext)

	// ExitIsNullCondition is called when exiting the IsNullCondition production.
	ExitIsNullCondition(c *IsNullConditionContext)

	// ExitInCondition is called when exiting the InCondition production.
	ExitInCondition(c *InConditionContext)

	// ExitOrderCondition is called when exiting the orderCondition production.
	ExitOrderCondition(c *OrderConditionContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)
}
