package main

import (
	"fmt"
	"sync"
)

func printTable(n int, wg *sync.WaitGroup) {
	wg.Add(12)
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
		wg.Done()
	}

}

func main() {
	var wg sync.WaitGroup

	for number := 1; number <= 12; number++ {
		go printTable(number, &wg)
	}

	wg.Wait()
}
