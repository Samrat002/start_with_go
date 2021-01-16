package main

import (
	"fmt"
	"net/http"
)

func checkLink(url string, c chan string){
	_, err := http.Get(url)
	if err != nil {
		c <- "Link Down:- "+url
	} else {
		c <- "Link up:- "+url
	}
}
func main() {
	c:= make(chan string)

	url_list := []string{
		"http://www.google.com",
		"http://www.stackoverflow.com",
		"http://www.facebook.com",
		"http://www.twitter.com",
	}
	for _, url := range url_list {
		go checkLink(url, c)
	}
	for message:= range c{
		fmt.Println(message)
	}
}
