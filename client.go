package krakenaugen

import (
	"fmt"
	"net/http"

	"golang.org/x/net/proxy"
)

type KrakenClient struct {
	HttpClient http.Client
	Dialer     *proxy.Dialer
}

func (kc *KrakenClient) SetTransport(proxyURL string) error {
	dialer, err := proxy.SOCKS5("tcp", proxyURL, nil, proxy.Direct)
	if err != nil {
		return fmt.Errorf("error while creating proxy dialer: %s", err)
	}
	kc.HttpClient.Transport = &http.Transport{Dial: dialer.Dial}
	return nil
}
