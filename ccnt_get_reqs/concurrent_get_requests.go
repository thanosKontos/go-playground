package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func makeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
	urls := [2]string{"http://www.google.com", "https://golang.org"}

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go makeRequest(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
