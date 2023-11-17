package krakenaugen

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type HTMLNode struct {
	Name       string
	Text       string
	attributes []html.Attribute
	DOM        *goquery.Selection
}

func NewHTMLElementFromSelection(s *goquery.Selection, node *html.Node) *HTMLNode {
	return &HTMLNode{
		Name:       node.Data,
		DOM:        s,
		attributes: node.Attr,
		Text:       goquery.NewDocumentFromNode(node).Text(),
	}
}

func (n *HTMLNode) GetAttr(key string) string {
	for _, attr := range n.attributes {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
