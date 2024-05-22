package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Constants for query keys
const (
	QueryMust     = "must"
	QueryMustNot  = "must_not"
	QueryShould   = "should"
	QueryFilter   = "filter"
	QueryBool     = "bool"
	QueryTerm     = "term"
	QueryRange    = "range"
	QueryExists   = "exists"
	QueryWildcard = "wildcard"
	QueryTerms    = "terms"
)

// EsQuery represents the Elasticsearch query
type EsQuery struct {
	Index   string
	Command string
	Query   map[string]any
}

func (eq EsQuery) String() string {
	var sb strings.Builder
	path := "_search"
	if eq.Command == "count" {
		path = "_count"
	}
	sb.WriteString(fmt.Sprintf("GET /%s/%s\n", eq.Index, path))
	q, _ := json.MarshalIndent(eq.Query, "", " ")
	sb.WriteString(string(q))
	return sb.String()
}

// Convert the parsed query to Elasticsearch query
func convertToElasticsearch(parsedQuery ParsedQuery) (EsQuery, error) {
	esQuery := EsQuery{
		Index:   parsedQuery.Index,
		Command: parsedQuery.Type,
		Query:   map[string]any{},
	}

	esQuery.Query = map[string]any{
		"query": map[string]any{
			QueryBool: buildBoolQuery(parsedQuery.Condition),
		},
	}

	if parsedQuery.Type == "search" {
		esQuery.Query["size"] = parsedQuery.Limit
		esQuery.Query["from"] = parsedQuery.Offset
		if len(parsedQuery.OrderBy) > 0 {
			esQuery.Query["sort"] = buildSortQuery(parsedQuery.OrderBy)
		}
	}
	return esQuery, nil
}

func buildBoolQuery(condition Condition) map[string][]any {
	boolQuery := map[string][]any{
		QueryMust:    {},
		QueryMustNot: {},
		QueryShould:  {},
		QueryFilter:  {},
	}

	for _, conditionPart := range condition.ConditionParts {
		conditionPartQuery := buildConditionPartQuery(conditionPart)
		addConditionToBoolQuery(boolQuery, condition, conditionPart, conditionPartQuery)
	}

	// Remove empty keys
	for k, v := range boolQuery {
		if len(v) == 0 {
			delete(boolQuery, k)
		}
	}
	return boolQuery
}

func addConditionToBoolQuery(boolQuery map[string][]any, condition Condition, conditionPart ConditionPart, conditionPartQuery map[string]any) {
	if condition.Operator == "and" {
		if conditionPart.Negate {
			boolQuery[QueryMustNot] = append(boolQuery[QueryMustNot], conditionPartQuery)
		} else {
			boolQuery[QueryMust] = append(boolQuery[QueryMust], conditionPartQuery)
		}
	} else {
		if conditionPart.Negate {
			encapsulated := map[string]any{
				QueryBool: map[string]any{
					QueryMustNot: []any{conditionPartQuery},
				},
			}
			boolQuery[QueryShould] = append(boolQuery[QueryShould], encapsulated)
		} else {
			boolQuery[QueryShould] = append(boolQuery[QueryShould], conditionPartQuery)
		}
	}
}

func buildConditionPartQuery(conditionPart ConditionPart) map[string]any {
	switch conditionPart.Type {
	case "ComparisonCondition":
		return buildComparisonQuery(conditionPart)
	case "GroupCondition":
		return map[string]any{QueryBool: buildBoolQuery(*conditionPart.Condition)}
	case "LikeCondition":
		return buildLikeQuery(conditionPart)
	case "IsNullCondition":
		return buildIsNullQuery(conditionPart)
	case "InCondition":
		return buildInQuery(conditionPart)
	default:
		return nil
	}
}

func buildComparisonQuery(conditionPart ConditionPart) map[string]any {
	if conditionPart.Operator == "=" || conditionPart.Operator == "!=" {
		return map[string]any{
			QueryTerm: map[string]any{
				conditionPart.Field: conditionPart.Value,
			},
		}
	}
	return map[string]any{
		QueryRange: map[string]any{
			conditionPart.Field: map[string]any{
				rangeOperator(conditionPart.Operator): conditionPart.Value,
			},
		},
	}
}

func rangeOperator(operator string) string {
	switch operator {
	case ">":
		return "gt"
	case ">=":
		return "gte"
	case "<":
		return "lt"
	case "<=":
		return "lte"
	default:
		return ""
	}
}

func buildLikeQuery(conditionPart ConditionPart) map[string]any {
	value := strings.Replace(conditionPart.Value.(string), "%", "*", -1)
	return map[string]any{
		QueryWildcard: map[string]any{
			conditionPart.Field: value,
		},
	}
}

func buildIsNullQuery(conditionPart ConditionPart) map[string]any {
	return map[string]any{
		QueryExists: map[string]any{
			"field": conditionPart.Field,
		},
	}
}

func buildInQuery(conditionPart ConditionPart) map[string]any {
	return map[string]any{
		QueryTerms: map[string]any{
			conditionPart.Field: conditionPart.Value,
		},
	}
}

func buildSortQuery(orderBy []Order) []any {
	var sortQuery []any
	for _, order := range orderBy {
		sortQuery = append(sortQuery, map[string]any{
			order.Field: map[string]any{
				"order": strings.ToLower(order.Direction),
			},
		})
	}
	return sortQuery
}
