package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.golang.org",
		"https://www.linkedin.com",
	}

	// Create a channel through which go routines can communicate with each other
	// make a channel through which go routines can pass string values
	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	// an infinite loop
	for link := range ch {
		go checkLink(link, ch)
	}
}

// check if the weblink is responding to the get request
func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be Down")
		ch <- link
		return
	}

	fmt.Println(link, "is Up")
	ch <- link
}
