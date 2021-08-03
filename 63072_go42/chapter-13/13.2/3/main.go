package main

import (
	"flag"
	"fmt"
	"log"

	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	go compute()
	time.Sleep(10 * time.Second)
}
func compute() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(Factorial(uint64(40)))
		}()
		time.Sleep(time.Millisecond * 1)
	}
}
