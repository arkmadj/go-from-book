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
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%s %vs \n", s, time.Since(start).Nanoseconds())
}

func echo3() {
	start := time.Now()
	fmt.Printf("%s %.vs \n", strings.Join(os.Args[1:], " "), time.Since(start).Nanoseconds())
}

func main() {
	echo1()
	echo3()
}
