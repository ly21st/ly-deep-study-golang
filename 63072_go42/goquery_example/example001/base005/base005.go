package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}

	dhead := doc.Find("head")
	dTitle := dhead.Find("title")
	fmt.Printf("title text:%s\n", dTitle.Text())
	html, _ := dTitle.Html()
	fmt.Printf("title html:%s\n", html)
	metaArr := dhead.Find("meta")
	fmt.Println("111111")
	for i := 0; i < metaArr.Length(); i++ {
		d, _ := metaArr.Eq(i).Attr("name")
		fmt.Println(d)
	}

	fmt.Printf("2222")

	doc.Find("div.wrapper .container .col-lg-9").Each(func(i int, cs *goquery.Selection) {
		d, _ := cs.Attr("class")
		fmt.Println(d)
	})
}

func main() {
	// ExampleScrape()
	// return

	fmt.Println("--------------------")
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(doc.Html())            //.Html()得到html内容
	pTitle := doc.Find("title").Text() //直接提取title的内容
	class := doc.Find("h2").Text()
	fmt.Printf("class:%v\n", class)
	fmt.Printf("title:%v\n", pTitle)
	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		t := contentSelection.Find(".title a")
		log.Printf("the length;%d", t.Length())
		log.Println("第", i+1, "个帖子的标题：", title)
	})
	/*
	   t := doc.Find(".topics .topic")
	   log.Printf("%+v", t)
	   t = doc.Find(".topics")
	   log.Printf("%+v", t)
	   t = doc.Find(".topic")
	   log.Printf("%+v", t)
	   t = doc.Find("div.topic")
	   log.Printf("div.topic:%+v", t)
	*/
	t := doc.Find("div.topic").Find(".title a")
	log.Printf("div.topic.title a:%+v", t)
	for i := 0; i < t.Length(); i++ {
		d, _ := t.Eq(i).Attr("href")
		title, _ := t.Eq(i).Attr("title")
		fmt.Println(d)
		fmt.Println(title)
	}
}
