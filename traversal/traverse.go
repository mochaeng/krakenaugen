package traversal

import (
	"fmt"
	"net/http"

	node_wrapper "github.com/mochaeng/krakenaugen/node-wrapper"
	"golang.org/x/net/html"
)

func forEachNode(node *html.Node, pre, pos func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(child, pre, pos)
	}
	if pos != nil {
		pos(node)
	}
}

func grabNodeFromClass(htmlClass string, node *html.Node) *html.Node {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == htmlClass {
			return node
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		node := grabNodeFromClass(htmlClass, child)
		if node != nil {
			return node
		}
	}

	return nil
}

func Visit(url string, htmlClass string, onEach func(node *node_wrapper.KrakenNode)) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while making a get to: %s", url)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code %d on %s", response.StatusCode, url)
	}
	defer response.Body.Close()

	doc, err := html.Parse(response.Body)
	if err != nil {
		return fmt.Errorf("error: %s while parsing html for url: %s", err, url)
	}

	desiredNode := grabNodeFromClass(htmlClass, doc)

	if desiredNode == nil {
		return fmt.Errorf("No node was founded it with: %s", htmlClass)
	}

	checker := func(node *html.Node) {
		if node.Type == html.TextNode {
			// krakenNode := node_wrapper.KrakenNode{Node: node}
			// onEach(&krakenNode)
			// fmt.Println(node.Attr)
			fmt.Println(node.Data)
		}
	}

	// root := &krakenaugen.KrakenNode{Node: doc}

	forEachNode(desiredNode, checker, nil)

	return nil
}
