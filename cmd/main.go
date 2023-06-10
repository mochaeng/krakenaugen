package main

import (
	"fmt"
	"net/http"

	krakenaugen "github.com/mochaeng/krakenaugen/core"
	node_wrapper "github.com/mochaeng/krakenaugen/node-wrapper"
	"golang.org/x/net/html"
)

func TraverseHTML(htmlDoc *html.Node) {
	if htmlDoc.Type == html.ElementNode && htmlDoc.Data == "img" {
		for _, attr := range htmlDoc.Attr {
			if attr.Key == "src" {
				fmt.Println(attr.Val)
			}
		}
	}
	for child := htmlDoc.FirstChild; child != nil; child = child.NextSibling {
		TraverseHTML(child)
	}
}

func ScrapeAndParseHTML(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got http status code %d at %s", response.StatusCode, url)
	}

	htmlDoc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error: %v while parsing the url: %s", err, url)
	}

	return htmlDoc, nil
}

// const link = "https://www.reddit.com/r/tzuyu/comments/13j02bt/181128_tzuyu/"
const link = "https://www.factretriever.com/rhino-facts"

type Question struct {
	question string
}

var allQuestions []string

func main() {
	// links := []string{link}

	// for _, url := range links {
	// 	htmlDoc, err := ScrapeAndParseHTML(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error while scrape into html: %v\n", err)
	// 	}
	// 	TraverseHTML(htmlDoc)
	// }

	kraken := krakenaugen.CreateNewAugen(link)

	kraken.OnHTML(".factsList li", func(node *node_wrapper.KrakenNode) {
        text := node.InnerText()
        if text == "" {
            return
        }
        fmt.Println(node.InnerText())
	})

	// fmt.Println(kraken)
}
