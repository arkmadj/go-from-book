package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.Create("webiste-result.txt")
	if err != nil {
		fmt.Printf("while creating file: %v", err)
	}
	defer file.Close()

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read file: %v\n", err)
			continue
		}
		for _, url := range strings.Split(string(data), "\n") {
			go fetch(url, ch)
		}
	}

	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}
	fmt.Fprintf(file, "%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
