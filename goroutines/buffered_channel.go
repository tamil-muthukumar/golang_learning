package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Buffered channels are asynchronous. They can be used to send and receive messages between goroutines.

	// --- SCENARIO 1 --- Without any iterations
	// Channels are FIFO queue
	bufferedChannel := make(chan string, 2) // size 2
	//go func() {
	bufferedChannel <- "First message"
	bufferedChannel <- "Second message" // If the buffered channel size is 1, this will lead to a deadlock as the channel can only used to put 1 item as the buffer size is 1
	//}()
	fmt.Println(<-bufferedChannel) // Extracts first message in FIFO Queue
	fmt.Println(<-bufferedChannel) // Extracts next item and Prints second message

	// --- SCENARIO 2 --- With Iterations
	bufferedChannel1 := make(chan string) // size 0
	numRounds := 3

	// With buffer size 0, we will throw one ninja star at a time, get the results back from the channel before throwing the next one
	// When a message is put into the channel, the goroutine will wait until the message is read from the channel before we can put another message into the channel
	go throwingNinjaStar(bufferedChannel1, numRounds)

	for i := 0; i < numRounds; i++ {
		fmt.Println("i:", i, ": ", <-bufferedChannel1) // runs sequentially as you have to get the message out first before the goroutine can put another message in
	}

	// -- SCENARIO 3 --- With buffer size 2
	fmt.Println("SCENARIO 3 - iterating through buffered channel with size 2")
	bufferedChannel2 := make(chan string) // size 2????
	go throwingNinjaStar2(bufferedChannel2)
	// Throw all the ninja stars at once and get the results back from the channel

	// This is a simplified version of the for loop below
	for message := range bufferedChannel2 {
		fmt.Println(message)
	}

	//// Alternate and more elaborate version of above for loop
	//for {
	//	message, open := <-bufferedChannel2
	//	if !open {
	//		break
	//	}
	//	fmt.Println(message)
	//}
}

func throwingNinjaStar(channel chan string, rounds int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("Ninja Star Thrown, score: %d", score)
	}
}

func throwingNinjaStar2(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	rounds := 3
	for i := 0; i < rounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("Ninja Star Thrown, score: %d", score)
	}
	// once all the messages are sent out in the channel, close it, so that the receiver knows that it is the end of the messages
	close(channel)
}
