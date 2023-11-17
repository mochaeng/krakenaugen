# krakenaugen

Krakenaguen (kraken + eyes), is a hobby project created for learning purposes. It serves as a simple web scraping library implemented using [goquery](https://github.com/PuerkitoBio/goquery) with an API similar to [Colly](http://go-colly.org/), including proxy support.

### Desired Features

- [X] Handle CSS Selectors
- [X] Proxy support
- [ ] Parallelism

## Usage

```go
const link = "https://www.factretriever.com/rhino-facts"

func main() {
	kraken := krakenaugen.CreateNewKraken(link)

	kraken.OnHTML(".factsList li", func(e *krakenaugen.HTMLNode) {
		fmt.Println(e)
		// e.GetAttr()
        // do something with the elment
	})

	kraken.Start(link)
}
```