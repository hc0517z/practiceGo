package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
		"https://www.youtube.com/",
		"https://forum.dotnetdev.kr/",
		"https://github.com/",
		"https://academy.nomadcoders.co/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://okky.kr/",
	}

	ch := make(chan requestResult)

	for _, url := range urls {
		go hitURL(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		requestResult := <-ch
		results[requestResult.url] = requestResult.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, ch chan requestResult) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- requestResult{url: url, status: status}
}
