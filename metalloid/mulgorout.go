package main

import (
	"fmt"
	"time"
)

func main6() {
	fmt.Println("Chapter 09 - goroutines")

	for i := 0; 1 < 10; i++ {
		go func(x int) {
			time.Sleep(100)
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(100)
	fmt.Println("Exiting ...")
}
