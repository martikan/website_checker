package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	sites := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range sites {
		go checkSite(link, c)
	}

	for s := range c {
		go func(site string) {
			time.Sleep(5 * time.Second)
			checkSite(site, c)
		}(s)
	}

}

func checkSite(l string, c chan string) {

	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "is down")
		c <- l
		return
	}

	fmt.Println(l, "is up")
	c <- l
}
