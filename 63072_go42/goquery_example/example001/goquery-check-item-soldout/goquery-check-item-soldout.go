package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var str1 = `
<div class="product-cv">
    <a id="ContentPlaceHolder1_BtnAddCart" class="btn-cv--lg product-cv__btn js-btn-cv ItemCART_ show-loader" onMouseOver="JavaScript:this.style.cursor=&#39;pointer&#39;;" href="javascript:__doPostBack(&#39;ctl00$ContentPlaceHolder1$BtnAddCart&#39;,&#39;&#39;)">
        <svg class="btn-cv__ico--lg">
            <use xlink:href="#ico-cart"/>
        </svg>
        <span class="btn-cv__txt">加入预约购物车</span>
    </a>
    <div class="product-sns grid sns a2a_kit a2a_kit_size_32 a2a_default_style">
        <a href="" class="product-sns__item a2a_button_line" target="_blank">
            <svg class="product-sns__ico">
                <use xlink:href="#ico-line"/>
            </svg>
        </a>
        <a href="" class="product-sns__item a2a_button_twitter" target="_blank">
            <svg class="product-sns__ico">
                <use xlink:href="#ico-twitter"/>
            </svg>
        </a>
        <a href="" class="product-sns__item a2a_button_facebook" target="_blank">
            <svg class="product-sns__ico">
                <use xlink:href="#ico-facebook"/>
            </svg>
        </a>
    </div>
</div>
`

var str2 = `
<div class="product-cv">
    <button class="btn-cv--lg btn-cv--disabled product-cv__btn" disabled>
        <svg class="btn-cv__ico--lg">
            <use xlink:href="#ico-cart-sold"/>
        </svg>
        <span class="btn-cv__txt">SOLD OUT</span>
    </button>
    <div class="product-sns grid sns a2a_kit a2a_kit_size_32 a2a_default_style">
        <a href="" class="product-sns__item a2a_button_line" target="_blank">
            <svg class="product-sns__ico">
                <use xlink:href="#ico-line"/>
            </svg>
        </a>
        <a href="" class="product-sns__item a2a_button_twitter" target="_blank">
            <svg class="product-sns__ico">
                <use xlink:href="#ico-twitter"/>
            </svg>
        </a>
        <a href="" class="product-sns__item a2a_button_facebook" target="_blank">
            <svg class="product-sns__ico">
                <use xlink:href="#ico-facebook"/>
            </svg>
        </a>
    </div>
</div>`

func main() {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(str1))
	if err != nil {
		fmt.Print(err)
		return
	}

	v := dom.Find("div[class='product-cv'] a#ContentPlaceHolder1_BtnAddCart span[class='btn-cv__txt']").Eq(0).Text()
	v = strings.Trim(v, " ")
	fmt.Printf("v=%s", v)

	dom, err = goquery.NewDocumentFromReader(strings.NewReader(str2))
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("\n")
	v = dom.Find("div[class='product-cv'] button[class='btn-cv--lg btn-cv--disabled product-cv__btn'] span[class='btn-cv__txt']").Eq(0).Text()
	v = strings.Trim(v, " ")
	fmt.Printf("v=%s", v)

}
