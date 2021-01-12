package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var host string

	fmt.Print("Full URL To Check: ")
	fmt.Scan(&host)

	c := make(chan string)
	go check(host, c)

	for {
		go func(link string) {
			time.Sleep(5 * time.Second)
			check(link, c)
		}(<-c)
	}
}

func check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding")
		c <- link
		return
	}

	fmt.Println(link, "is active and responding")
	c <- link
}
