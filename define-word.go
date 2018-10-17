package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"golang.org/x/net/html"
	"net/http"
	//	"os"
	//	"strings"
)

func getContentExplanation(t html.Token) bool {
	for _, a := range t.Attr {
		if a.Key == "class" && a.Val == "content-explanation ej" {
			pp.Print(t)
			pp.Print(t.String())
			return true
		}
	}
	return false
}

func crawl(url string) (definition string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	b := resp.Body
	defer b.Close()

	tokenizer := html.NewTokenizer(b)
	foundDefinition := false

	for {
		token := tokenizer.Next()

		switch {
		case token == html.ErrorToken:
			return "definition not found"
		case token == html.StartTagToken:
			t := tokenizer.Token()

			isTableData := t.Data == "td"
			if !isTableData {
				continue
			}
			foundDefinition = getContentExplanation(t)

		case token == html.TextToken && foundDefinition:
			t := tokenizer.Token()
			return t.Data
		}
	}
}

func main() {
	url := "https://ejje.weblio.jp/content/play"

	// Channels
	// chUrls := make(chan string)
	// chFinished := make(chan bool)

	definition := crawl(url)

	fmt.Printf("\nFound %s", definition)
}
