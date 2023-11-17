package krakenaugen

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (k *Kraken) fetch(selector string, callback HTMLCallback) error {
	response, err := k.krakenClient.HttpClient.Get(k.StartURL)
	if err != nil {
		return fmt.Errorf("error while making a get to: %s", k.StartURL)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code %d on %s", response.StatusCode, selector)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return fmt.Errorf("could not create document from response body: %s", err)
	}

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			element := NewHTMLElementFromSelection(s, n)
			callback(element)
		}
	})

	return nil
}
