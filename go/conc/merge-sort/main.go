// https://www.educative.io/courses/mastering-concurrency-in-go/N7A80vPyLk2

package main

import (
	"fmt"
	"log"
	"time"
)

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSortSync(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	left := MergeSortConc(data[:mid])

	right := MergeSortConc(data[mid:])

	return Merge(left, right)
}

func MergeSortConc(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	done1 := make(chan bool)
	done2 := make(chan bool)
	var left []int
	var right []int

	go func() {
		left = MergeSortConc(data[:mid])
		done1 <- true
	}()

	go func() {
		right = MergeSortConc(data[mid:])
		done2 <- true
	}()

	<-done1
	<-done2

	return Merge(left, right)
}

func main() {
	// expected output: [1 2 3 4 5 6 7 8 9 10]
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}

	startSync := time.Now()
	fmt.Printf("%v\n%v\n", data, MergeSortSync(data))
	elapsedSync := time.Since(startSync)
	log.Printf("Conc Took %s", elapsedSync)

	startConc := time.Now()
	fmt.Printf("%v\n%v\n", data, MergeSortConc(data))
	elapsedConc := time.Since(startConc)
	log.Printf("Conc Took %s", elapsedConc)
}
