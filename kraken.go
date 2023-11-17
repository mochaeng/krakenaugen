package krakenaugen

import (
	"log"
)

type Kraken struct {
	AllowedDomains []string
	StartURL       string
	HtmlCallbacks  []*HTMLCallbackContainer

	krakenClient KrakenClient
}

type HTMLCallbackContainer struct {
	Selector string
	Function HTMLCallback
}

type HTMLCallback func(*HTMLNode)

func CreateNewKraken(startURL string) *Kraken {
	return &Kraken{
		StartURL:     startURL,
		krakenClient: KrakenClient{},
	}
}

func (k *Kraken) AddProxy(proxyURL string) error {
	err := k.krakenClient.SetTransport(proxyURL)
	if err != nil {
		log.Fatalf("not possible to create proxy dialer: %s", err)
	}
	return nil
}

func (k *Kraken) OnHTML(goqueryselector string, f HTMLCallback) {
	k.HtmlCallbacks = append(k.HtmlCallbacks, &HTMLCallbackContainer{
		Selector: goqueryselector,
		Function: f,
	})
}

func (k *Kraken) Start(url string) {
	for _, callBackContainer := range k.HtmlCallbacks {
		k.fetch(callBackContainer.Selector, callBackContainer.Function)
	}
}
