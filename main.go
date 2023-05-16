package main

import (
	"fmt"
	"net/http"
	"os"

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

const link = "https://www.reddit.com/r/tzuyu/comments/13i9o50/200321_tzuyu/"

func main() {
	links := []string{link}

	for _, url := range links {
		htmlDoc, err := ScrapeAndParseHTML(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while scrape into html: %v\n", err)
		}
		TraverseHTML(htmlDoc)
	}
	fmt.Println(links)
}
