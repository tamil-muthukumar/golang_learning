package main

import (
	"fmt"
	"time"
)

// / Learning goroutine simple usage of channel
func main() {
	now := time.Now()

	defer func() {
		fmt.Print("time since: ", time.Since(now))
	}()
	channel := make(chan bool)

	evilNinja := "yogi"
	go attack(evilNinja, channel)
	// this could lead to a deadlock as the channel can only used by one function/goroutine at a time.
	// This will lead to a deadlock where the main program and the goroutine are trying to access the channel at the same time
	channel <- false
	// reading the message from the channel. The arrow is on the left of the channel indicating that the data is read from the channel.
	fmt.Println(<-channel)
}

func attack(target string, attacked chan bool) {
	fmt.Println("Attacking", target)
	// writing the message to the channel. The arrow indicates the direction of the data flow.
	// In this case, the data is flowing from the right to the left.
	attacked <- true

}
