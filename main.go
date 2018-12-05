package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Print(link, "might be down")
		return false
	}

	fmt.Println(link, "is up")
	return true
}
