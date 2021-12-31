package main

import (
	"fmt"
	"strconv"
	"time"
)


func routines () {

	c := make(chan string)
	go count(c)
	fmt.Println(c)
	for {
		msg, open := <- c
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

func count (c chan string) {
	for i := 0; i < 10; i++ {
		c <- "Counting " + strconv.Itoa(i)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}