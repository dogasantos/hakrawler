package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)
func main() {

	c := colly.NewCollector(
        colly.AllowedDomains("telefonica.com","apis.telefonica.com"),
    )

    c.OnHTML("body", func(e *colly.HTMLElement) {
		dom := e.DOM
		inputs := dom.Find("input")
		for oneinput := range inputs.Nodes {
			input := inputs.Eq(oneinput)
			val,_ := input.Attr("type")
			if val == "password" {
				fmt.Println(input.Attr("id"))
				
			}
		}
	})

    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        RandomDelay: 1 * time.Second,
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    c.Visit("https://apis.telefonica.com/login")

}