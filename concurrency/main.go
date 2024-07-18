package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)
	links := []string{
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://google.com",
	}

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		errString := fmt.Sprint("error from get: ", err)
		c <- link
		panic(errString)
	}

	fmt.Println(link, "is up!")
	c <- link

}
