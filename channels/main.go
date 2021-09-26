package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)
	pauseSeconds := 20
	links := []string{
		"http://amazon.com",
		"http://golang.org",
		"http://google.com",
		"http://iamreliq.com",
	}

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Duration(pauseSeconds) * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("`%v` might be down!\n", link)
		c <- link
		return
	}

	fmt.Printf("`%v` is OK!\n", link)
	c <- link
}
