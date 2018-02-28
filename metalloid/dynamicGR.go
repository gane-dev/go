package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main8() {
	if len(os.Args) != 2 {
		fmt.Printf("suage: % integer \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	var numGR, _ = strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Printf("going to create%d gorroutines .\n", numGR)
	var waitGroup sync.WaitGroup

	var i int64
	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)

		go func(x int64) {
			defer waitGroup.Done()
			fmt.Printf(" %d ", x)
		}(i)
	}
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
