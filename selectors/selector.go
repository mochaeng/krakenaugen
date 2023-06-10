package selectors

import (
	"fmt"
	"strings"

	node_wrapper "github.com/mochaeng/krakenaugen/node-wrapper"
	"golang.org/x/net/html"
)

type Matcher interface {
	Matches(node *html.Node) bool
}

type Class struct {
	ClassName  string
	Element    *Element
	IsOneLevel bool
}

type Element struct {
	Name string
}

func (class *Class) Matches(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == class.ClassName {
            if class.Element != nil {
                fmt.Println(class.Element)
                fmt.Println(node.Data)
            }
			return true
		}
	}
	return false
}

func (element *Element) Matches(node *html.Node) bool {
	if node.Type == html.ElementNode && node.Data == element.Name {
		return true
	}
	return false
}

func BuildMatcher(rawSelector string) Matcher {
	terms := strings.Split(rawSelector, " ")

	var matcher Matcher
	firstCharacter := rawSelector[0]
	if firstCharacter == '.' {
		matcher = buildClassSelector(rawSelector)
	} else if firstCharacter == '#' {
		// ids
	} else if len(terms) == 1 {
		matcher = buildElementSelector(rawSelector)
	}

	return matcher
}

func buildElementSelector(rawSelector string) *Element {
	return &Element{Name: rawSelector}
}

func buildClassSelector(rawSelector string) *Class {
	terms := strings.Split(rawSelector, " ")
    className := strings.ReplaceAll(terms[0], ".", "")
	if len(terms) == 1 {
		return &Class{ClassName: className}
	} else if len(terms) == 2 {
		return &Class{ClassName: className, Element: &Element{Name: terms[1]}}
	} else if len(terms) == 3 {
		return &Class{ClassName: className, Element: &Element{Name: terms[2]}, IsOneLevel: true}
	}
	return nil
}

func Text(runner func(krakenNode *node_wrapper.KrakenNode)) func(node *html.Node) {
	return func(node *html.Node) {
		if node.Type == html.TextNode {
			krakenNode := node_wrapper.NewKrakenNode(node)
			runner(krakenNode)
		}
	}
}
