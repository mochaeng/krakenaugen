package main

import (
	"fmt"

	krakenaugen "github.com/mochaeng/krakenaugen"
)

// const link = "https://www.reddit.com/r/tzuyu/comments/13j02bt/181128_tzuyu/"
const link = "https://www.factretriever.com/rhino-facts"

func main() {
	kraken := krakenaugen.CreateNewKraken(link)

	kraken.OnHTML(".factsList li", func(e *krakenaugen.HTMLNode) {
		fmt.Println(e)
		// e.GetAttr()
	})

	kraken.Start(link)
}
