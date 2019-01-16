package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func makeRequest(url string) {
	defer wg.Done()
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(len(body), url)
}

func main() {
	urls := [2]string{"http://www.google.com", "https://golang.org"}

	for _, url := range urls {
		wg.Add(1)
		go makeRequest(url)
	}

	fmt.Println("The end")
	wg.Wait()
}
