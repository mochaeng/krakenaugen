package krakenaugen

import (
	"fmt"

	node_wrapper "github.com/mochaeng/krakenaugen/node-wrapper"
	"github.com/mochaeng/krakenaugen/traversal"
)

type Kraken struct {
	AllowedDomains []string
	StartURL       string
}

func CreateNewAugen(startURL string) *Kraken {
	return &Kraken{StartURL: startURL}
}

func (k *Kraken) OnHTML(class string, onEach func(node *node_wrapper.KrakenNode)) {
	// Visit(k.StartURL, tag, onEach)
	// traversal.Visit()
	traversal.Visit(k.StartURL, class, onEach)
	fmt.Println("Going to visit: ", k.StartURL)
}
