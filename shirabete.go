package main

import (
	"encoding/csv"
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"sync"
)

func findDefinition(t html.Token) bool {
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

			if t.Data == "td" {
				foundDefinition = findDefinition(t)
			}

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
		url := "https://ejje.weblio.jp/content/"
		filename := c.Args().Get(0)
		fp, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer fp.Close()

		reader := csv.NewReader(fp)
		reader.TrimLeadingSpace = true
		wordChan := make(chan string, 50)

		go func() {
			for {
				record, err := reader.Read()
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Printf("ERROR: Failed to read file: %s", err.Error())
				}

				for _, value := range record {
					if value != "" {
						wordChan <- value
					}
				}
			}

			close(wordChan)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				for word := range wordChan {
					definition := crawl(fmt.Sprintf("%s%s", url, word))
					fmt.Printf("\n%s: %s", word, definition)
				}
				wg.Done()
			}()
		}

		wg.Wait()
		return nil
	}

	app.Run(os.Args)
}
