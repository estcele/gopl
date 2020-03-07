package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int, 0)

	go func() {
		for x := 0; ; x++ {
			fmt.Println("counter:x=", x)
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
			fmt.Println("squares:x=", x)
		}
	}()

	for {
		fmt.Println("printer:x=", <-squares)
	}
}
