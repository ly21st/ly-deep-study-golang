package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
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
		fmt.Print(err)
	}

	dom.Find("div:contains(DIV2)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	// // dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(rsp.Body())))
	// dom, err = goquery.NewDocumentFromReader(strings.NewReader(str1))
	// if err != nil {
	// 	fmt.Print(err)
	// 	return

	// }
	// fmt.Printf("%v\n", dom)

	// dom.Find("div[class=inner header-inner]")

	// // dom.Find("div[class='inner header-inner'] select[name='ctl00$ddlAirport']#ddlAirport[class='form-select-cmn header-airport-select'] option").Each(func(i int, selection *goquery.Selection) {
	// // 	attr, _ := selection.Attr("value")
	// // 	fmt.Printf("%s:%s\n", attr, selection.Text())})

}
