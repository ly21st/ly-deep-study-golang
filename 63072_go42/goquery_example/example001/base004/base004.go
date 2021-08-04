package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	test1()
}

func test1() {
	html := `<div class="profile-navbar clearfix">
	<a class="item " href="/people/jixin/asks">提问<span class="num">1336</span></a>
	<a class="item " href="/people/jixin/answers">回答<span class="num">785</span></a>
	<a class="item " href="/people/jixin/posts">文章<span class="num">91</span></a>
	<a class="item " href="/people/jixin/collections">收藏<span class="num">44</span></a>
	<a class="item " href="/people/jixin/logs">公共编辑<span class="num">51648</span></a>
	</div>
	`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	spanSelection := doc.Find("div.profile-navbar").Find("span.num")
	log.Printf("spanSelection=%v\n", spanSelection)
	log.Printf("len=%v\n", spanSelection.Length())
	log.Printf("content=%v\n", spanSelection.Eq(1).Text())

}
