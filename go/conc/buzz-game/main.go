//https://www.educative.io/courses/mastering-concurrency-in-go/7nGMKVJjo0r

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel1 <- "Player 1 Buzzed"
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel2 <- "Player 2 Buzzed"
	}()

	for i := 0; i < 2; i++ {
		select {
		case ch1 := <-channel1:
			fmt.Println(ch1)
		case ch2 := <-channel2:
			fmt.Println(ch2)
		}
	}
}
