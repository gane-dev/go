package main

import (
	"fmt"
	"time"
)

func namedFunction() {
	time.Sleep(10000 * time.Microsecond)
	fmt.Println("Printing from namedFunction!")
}

func main5() {
	fmt.Println("Chapter 09 - Goroutines")
	go namedFunction()

	go func() {
		fmt.Println("Anonymous function")
	}()
	time.Sleep(10000 * time.Microsecond)
	fmt.Println("Exiting...")
}
