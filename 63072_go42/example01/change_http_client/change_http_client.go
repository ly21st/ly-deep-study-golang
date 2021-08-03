package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	c := colly.NewCollector()

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

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
