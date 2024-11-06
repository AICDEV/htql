package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antlr4-go/antlr/v4"
	htql "htql/htql_parser_go"
	parser2 "htql/parser"
	"log"
	"net/http"
	"os"
	"regexp"
)

var urlPattern = regexp.MustCompile(`(http|https)://[a-zA-Z0-9_.-]+(:[0-9]+)?(/[a-zA-Z0-9_.-/?&%#=]*)?`)

func main() {
	expression := flag.String("query", "", "HTQL query")

	flag.Parse()

	evaluateDocument(*expression, func(document *goquery.Document) {

		input := antlr.NewInputStream(*expression)
		lexer := htql.NewHtqlLexer(input)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		parser := htql.NewHtqlParser(tokens)

		tree := parser.Query()

		visitor := parser2.NewHtqlRuntimeVisitor(document)
		nodes := visitor.Visit(tree)

		fmt.Println(nodes)
	})
}

func evaluateDocument(expression string, block func(document *goquery.Document)) {
	if urlMatch := urlPattern.FindString(expression); urlMatch != "" {

		res, err := http.Get(urlMatch)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		block(doc)
	} else {
		file, err := os.Open("./test.html")
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()

		doc, err := goquery.NewDocumentFromReader(file)
		if err != nil {
			log.Fatalf("Failed to parse HTML file: %v", err)
		}
		block(doc)
	}
}
