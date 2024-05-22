package main

import (
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/bakyazi/esgrammar/parser"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/convert", convertHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	var req QueryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	is := antlr.NewInputStream(req.Query)
	lexer := parser.NewQueryLexer(is)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewQueryParser(tokenStream)
	tree := p.Query()

	parsedQuery := Visit(tree)

	esQuery, err := convertToElasticsearch(parsedQuery.(ParsedQuery))
	if err != nil {
		fmt.Println("Error converting to Elasticsearch query:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"query": esQuery.String()})
}
