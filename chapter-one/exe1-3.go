package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	start := time.Now()
	var s, sep string
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%s %.2fs \n", s, time.Since(start).Seconds())
}

func echo3() {
	start := time.Now()
	fmt.Printf("%s %.2fs \n", strings.Join(os.Args[:], " "), time.Since(start).Seconds())
}

func main() {
	go echo1()
	go echo3()
}
