package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/bakyazi/esgrammar/parser"
	"strconv"
	"strings"
)

func Visit(tree antlr.ParseTree) any {
	if tree == nil {
		return nil
	}

	switch node := tree.(type) {
	case *parser.SearchQueryContext:
		return visitSearchQuery(node)
	case *parser.CountQueryContext:
		return visitCountQuery(node)
	case *parser.WhereClauseContext:
		return visitWhereClause(node)
	case *parser.ConditionContext:
		return visitCondition(node)
	case *parser.AndConditionContext:
		return visitAndCondition(node)
	case *parser.OrConditionContext:
		return visitOrCondition(node)
	case *parser.ComparisonConditionContext:
		return visitComparisonCondition(node)
	case *parser.GroupConditionContext:
		return visitGroupCondition(node)
	case *parser.LikeConditionContext:
		return visitLikeCondition(node)
	case *parser.IsNullConditionContext:
		return visitIsNullCondition(node)
	case *parser.InConditionContext:
		return visitInCondition(node)
	case *parser.OrderByClauseContext:
		return visitOrderByClause(node)
	case *parser.OrderConditionContext:
		return visitOrderCondition(node)
	case *parser.LimitClauseContext:
		return visitLimitClause(node)
	case *parser.OffsetClauseContext:
		return visitOffsetClause(node)
	case *parser.ValueContext:
		return visitValue(node)
	}
	return nil
}

func visitSearchQuery(node *parser.SearchQueryContext) ParsedQuery {
	parsedQuery := ParsedQuery{
		Type:  "search",
		Index: node.SearchClause().IndexName().GetText(),
	}
	if node.WhereClause() != nil {
		parsedQuery.Condition = Visit(node.WhereClause()).(Condition)
	}
	if node.OrderByClause() != nil {
		parsedQuery.OrderBy = Visit(node.OrderByClause()).([]Order)
	}
	if node.LimitClause() != nil {
		parsedQuery.Limit = Visit(node.LimitClause()).(int)
	}
	if node.OffsetClause() != nil {
		parsedQuery.Offset = Visit(node.OffsetClause()).(int)
	}
	return parsedQuery
}

func visitCountQuery(node *parser.CountQueryContext) ParsedQuery {
	parsedQuery := ParsedQuery{
		Type:  "count",
		Index: node.CountClause().IndexName().GetText(),
	}
	if node.WhereClause() != nil {
		parsedQuery.Condition = Visit(node.WhereClause()).(Condition)
	}
	return parsedQuery
}

func visitWhereClause(node *parser.WhereClauseContext) Condition {
	return Visit(node.Condition()).(Condition)
}

func visitCondition(node *parser.ConditionContext) any {
	if node.OrCondition() != nil {
		return Visit(node.OrCondition())
	}
	return Visit(node.AndCondition())
}

func visitAndCondition(node *parser.AndConditionContext) Condition {
	condition := Condition{Operator: "and"}
	for _, cp := range node.AllConditionPart() {
		condition.ConditionParts = append(condition.ConditionParts, Visit(cp).(ConditionPart))
	}
	return condition
}

func visitOrCondition(node *parser.OrConditionContext) Condition {
	condition := Condition{Operator: "or"}
	for _, cp := range node.AllConditionPart() {
		condition.ConditionParts = append(condition.ConditionParts, Visit(cp).(ConditionPart))
	}
	return condition
}

func visitComparisonCondition(node *parser.ComparisonConditionContext) ConditionPart {
	return ConditionPart{
		Type:     "ComparisonCondition",
		Field:    node.FieldName().GetText(),
		Operator: node.Comparator().GetText(),
		Value:    Visit(node.Value()),
		Negate:   node.Comparator().GetText() == "!=",
	}
}

func visitGroupCondition(node *parser.GroupConditionContext) ConditionPart {
	condition := Visit(node.Condition()).(Condition)
	return ConditionPart{
		Type:      "GroupCondition",
		Condition: &condition,
	}
}

func visitLikeCondition(node *parser.LikeConditionContext) ConditionPart {
	return ConditionPart{
		Type:   "LikeCondition",
		Field:  node.FieldName().GetText(),
		Value:  strings.Trim(node.STRING().GetText(), `"`),
		Negate: node.NOT() != nil,
	}
}

func visitIsNullCondition(node *parser.IsNullConditionContext) ConditionPart {
	return ConditionPart{
		Type:   "IsNullCondition",
		Field:  node.FieldName().GetText(),
		Negate: node.NOT() != nil,
	}
}

func visitInCondition(node *parser.InConditionContext) ConditionPart {
	var values []any
	for _, value := range node.AllValue() {
		values = append(values, Visit(value))
	}
	return ConditionPart{
		Type:   "InCondition",
		Field:  node.FieldName().GetText(),
		Value:  values,
		Negate: node.NOT() != nil,
	}
}

func visitOrderByClause(node *parser.OrderByClauseContext) []Order {
	var orders []Order
	for _, orderCondition := range node.AllOrderCondition() {
		orders = append(orders, Visit(orderCondition).(Order))
	}
	return orders
}

func visitOrderCondition(node *parser.OrderConditionContext) Order {
	order := Order{
		Field:     node.FieldName().GetText(),
		Direction: "ASC",
	}
	if node.DESC() != nil {
		order.Direction = "DESC"
	}
	return order
}

func visitLimitClause(node *parser.LimitClauseContext) int {
	return getIntValue(node.NUMBER())
}

func visitOffsetClause(node *parser.OffsetClauseContext) int {
	return getIntValue(node.NUMBER())
}

func visitValue(node *parser.ValueContext) any {
	if node.STRING() != nil {
		return strings.Trim(node.STRING().GetText(), `"`)
	}
	return getIntValue(node.NUMBER())
}

func getIntValue(node antlr.TerminalNode) int {
	value, err := strconv.Atoi(node.GetText())
	if err != nil {
		return 0
	}
	return value
}
