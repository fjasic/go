package main

import (
	"fmt"
	"net/http"
)

type listWebsite []string

func main() {
	l := listWebsite{"https://www.google.com", "https://www.facebook.com",
		"https://www.stackoverflow.com", "https://www.golang.com", "https://www.amazon.com"}
	c := make(chan string)
	for _, w := range l {
		go checkLink(w, c)
	}
	for link := range c {
		go checkLink(link, c)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "is down!")

		c <- l
		return
	}
	fmt.Println(l, "is up!")
	c <- l
}
