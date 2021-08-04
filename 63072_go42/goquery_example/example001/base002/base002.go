package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("---------")
	fmt.Println("label选择器")
	test1()
	fmt.Println("---------")
	fmt.Printf("id选择器\n")
	idSelection()

	fmt.Println("----------------------")
	fmt.Println("element 选择器")
	elementSelection()

	fmt.Println("----------------------")
	fmt.Println("class选择器")
	classSelection()

	fmt.Println("----------------------")
	fmt.Println("属性选择器")
	attrSelection()

	fmt.Println("----------------------")
	fmt.Println("父子选择器")
	parentChildSelection()

	fmt.Println("----------------------")
	fmt.Println("相邻选择器")
	preNextSelection()

	fmt.Println("----------------------")
	fmt.Println("兄弟选择器")
	preNextSelection2()

	fmt.Println("----------------------")
	fmt.Println("内容选择器")
	contentFilterSelection()

	fmt.Println("----------------------")
	fmt.Println("first_child择器")
	first_child_selection()

	fmt.Println("----------------------")
	fmt.Println("first_type择器")
	first_type_selection()
}

func test1() {
	html := `<body>

				<div>DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
		// fmt.Printf("%v\n", selection)
	})
}

func idSelection() {
	html := `<body>

	<div id="div1">DIV1</div>
	<div>DIV2</div>
	<span>SPAN</span>

</body>
`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func elementSelection() {
	html := `<body>

	<div id="div1">DIV1</div>
	<div>DIV2</div>
	<span>SPAN</span>
	<span id="div1">span hello</span>

</body>
`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("span#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func classSelection() {
	html := `<body>

				<div id="div1">DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".name").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func attrSelection() {
	html := `<body>

				<div>DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

				<div class="age">30</div>
				<div class="age">35</div>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div[class]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	dom.Find("div[age=30]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	dom.Find("div[age!=30]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func parentChildSelection() {
	html := `<body>

				<div lang="ZH">DIV1</div>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("body>div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	fmt.Println()
	dom.Find("body div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func preNextSelection() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div[lang=zh]+p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

// 兄弟选择器
func preNextSelection2() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div[lang=zh]~p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func contentFilterSelection() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div:contains(DIV2)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	fmt.Println("------------------")
	dom.Find("span:has(div)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func first_child_selection() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div style="display:none;">DIV4</div>
					<div>DIV5</div>
				</span>
				<p>P2</p>
				<div></div>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div:first-child").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}

func first_type_selection() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<p>P2</p>
					<div>DIV5</div>
				</span>
				<div></div>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div:first-of-type").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}
