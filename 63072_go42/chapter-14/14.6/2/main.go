package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h, H bool

	v bool
	q *bool

	D    string
	Conf string
)

func init() {
	flag.BoolVar(&h, "h", false, "帮助信息")
	flag.BoolVar(&h, "H", false, "帮助信息")

	flag.BoolVar(&v, "v", false, "显示版本号")

	//
	flag.StringVar(&D, "D", "deamon", "set descripton ")
	flag.StringVar(&Conf, "Conf", "/dev/conf/cli.conf", "set Conf filename ")

	// 另一种绑定方式
	q = flag.Bool("q", false, "退出程序")

	// 像flag.Xxx函数格式都是一样的，第一个参数表示参数名称，
	// 第二个参数表示默认值，第三个参数表示使用说明和描述。
	// flag.XxxVar这样的函数第一个参数换成了变量地址，
	// 后面的参数和flag.Xxx是一样的。

	// 改变默认的 Usage
	flag.Usage = usage

	flag.Parse()

	var cmd string = flag.Arg(0)

	fmt.Printf("-----------------------\n")
	fmt.Printf("cli non=flags      : %s\n", cmd)

	fmt.Printf("q: %t\n", *q)

	fmt.Printf("descripton:  %s\n", D)
	fmt.Printf("Conf filename : %s\n", Conf)

	fmt.Printf("-----------------------\n")
	fmt.Printf("there are %d non-flag input param\n", flag.NArg())
	for i, param := range flag.Args() {
		fmt.Printf("#%d    :%s\n", i, param)
	}

}

func main() {
	flag.Parse()

	if h || H {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `CLI: 8.0
Usage: Cli [-hvq] [-D descripton] [-Conf filename] 

`)
	flag.PrintDefaults()
}
