package main

import (
	"log"
	"math/rand"

	"github.com/gocolly/colly"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	c := colly.NewCollector()
	visited := false

	// extensions.RandomUserAgent(c)
	// extensions.Referer(c)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(string(r.Body))
		if !visited {
			visited = true
			r.Request.Visit("/get?q=2")
		}
	})

	c.Visit("http://httpbin.org/get")
}
