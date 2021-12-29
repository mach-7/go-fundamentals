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

	for _, link := range links {
		checkLink(link)
	}
}

// check if the weblink is responding to the get request
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be Down")
		return
	}

	fmt.Println(link, "is Up")
}
