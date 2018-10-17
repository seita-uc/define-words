package main

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/urfave/cli"
	"golang.org/x/net/html"
	"net/http"
	"os"
	//	"strings"
)

func getContentExplanation(t html.Token) bool {
	for _, a := range t.Attr {
		if a.Key == "class" && a.Val == "content-explanation ej" {
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
	app := cli.NewApp()
	app.Name = "shirabete"
	app.Usage = "just set the English words in arguments"
	app.Version = "1.0.0"

	// action
	app.Action = func(c *cli.Context) error {
		url := "https://ejje.weblio.jp/content/play"
		filename := c.Args().Get(0)
		fp, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return err
		}
		definition := crawl(url)

		pp.Print("\nFound %s", definition)
		return nil
	}

	app.Run(os.Args)
}
