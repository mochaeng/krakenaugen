package krakenaugen

import (
	"fmt"
	"os"

	node_wrapper "github.com/mochaeng/krakenaugen/node-wrapper"
	"github.com/mochaeng/krakenaugen/selectors"
	"github.com/mochaeng/krakenaugen/traversal"
)

type Kraken struct {
	AllowedDomains []string
	StartURL       string
}

func CreateNewAugen(startURL string) *Kraken {
	return &Kraken{StartURL: startURL}
}

func (k *Kraken) OnHTML(rawSelector string, onEach func(node *node_wrapper.KrakenNode)) {
	// Visit(k.StartURL, tag, onEach)
	// traversal.Visit()
	// selector := selectors.BuildSelector(cssSelector)

	matcher := selectors.BuildMatcher(rawSelector)

	// fmt.Println(matcher)

	err := traversal.Visit(k.StartURL, matcher, onEach)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}

}
