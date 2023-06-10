package traversal

import (
	"fmt"
	"net/http"

	node_wrapper "github.com/mochaeng/krakenaugen/node-wrapper"
	"github.com/mochaeng/krakenaugen/selectors"
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

func grabNodes(matcher selectors.Matcher, node *html.Node) []*html.Node {
	nodes := []*html.Node{}

	adder := func(node *html.Node) {
		if matcher.Matches(node) {
			nodes = append(nodes, node)
		}
	}

	var travel func(node *html.Node)
	travel = func(node *html.Node) {
		adder(node)
		for child := node.FirstChild; child != nil; child = node.NextSibling {
			travel(child)
		}
	}

	travel(node)

	return nodes
}

func Visit(url string, matcher selectors.Matcher, onEach func(node *node_wrapper.KrakenNode)) error {
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

    matcherWrapper := func(node *html.Node) {
        if matcher.Matches(node) {
            convertedNode := node_wrapper.NewKrakenNode(node)
            onEach(convertedNode)
        }
    }

	// nodes := grabNodes(matcher, doc)
	// if len(nodes) == 0 {
	// 	return fmt.Errorf("no node was found")
	// }

	// fmt.Println(nodes)

	// nodes := []*html.Node{}
	// forEachNode(doc, checkMatcher, nil)

	// wrapper := func(node *html.Node) {
	// 	krakenNode := node_wrapper.NewKrakenNode(node)
	// 	onEach(krakenNode)
	// }

	forEachNode(doc, matcherWrapper, nil)

	return nil
}
