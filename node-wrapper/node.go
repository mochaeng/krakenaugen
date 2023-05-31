package node_wrapper

import (
	"fmt"

	"golang.org/x/net/html"
)

type KrakenNode struct {
	*html.Node
}

func (k *KrakenNode) PrintData() {
	fmt.Println(k.Data + " FROM")
}

func (k *KrakenNode) InnerText() string {
	if k.Type == html.TextNode {
		return k.Data
	}
	return ""
}

func NewKrakenNode(node *html.Node) *KrakenNode {
	return &KrakenNode{Node: node}
}
