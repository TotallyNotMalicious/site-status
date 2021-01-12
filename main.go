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
		go func(site string) {
			time.Sleep(5 * time.Second)
			check(site, c)
		}(<-c)
	}
}

func check(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is not responding")
		c <- site
		return
	}

	fmt.Println(site, "is active and responding")
	c <- site
}
