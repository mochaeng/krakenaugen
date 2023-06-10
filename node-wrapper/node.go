package node_wrapper

import (
	"fmt"

	"golang.org/x/net/html"
)

type KrakenNode struct {
	*html.Node
}

func NewKrakenNode(node *html.Node) *KrakenNode {
	return &KrakenNode{Node: node}
}

func (k *KrakenNode) PrintData() {
	fmt.Println(k.Data + " FROM")
}

func (k *KrakenNode) InnerText() string {
	for child := k.FirstChild; child != nil; child = child.NextSibling {
		if k.Type == html.ElementNode && k.isTextElement() {
			if child.Type == html.TextNode {
				return child.Data
			}
		}
	}
	return ""
}

func (k *KrakenNode) isTextElement() bool {
	return k.Data == "li" || k.Data == "p" || k.Data == "span"
}
