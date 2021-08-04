package main

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	c := colly.NewCollector()
	visited := false

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnResponse(func(r *colly.Response) {
		log.Println(string(r.Body))
		if !visited {
			visited = true
			r.Request.Visit("/get?q=2")
		}
	})

	c.Visit("http://httpbin.org/get")
}
