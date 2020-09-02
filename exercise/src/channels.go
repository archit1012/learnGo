// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working ...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	// Send a value to notify that we are done
	done <- true
}

func main() {

	// 1
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel.
	// Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	//  When we run the program the "ping" message is successfully passed
	// from one goroutine to another via our channel.

	// By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the "ping"
	// message without having to use any other synchronization.

	// 3  - Channel Synchronization
	// We can use channels to synchronize execution
	// across goroutines. Here's an example of using a
	// blocking receive to wait for a goroutine to finish.
	// When waiting for multiple goroutines to finish,
	// you may prefer to use a WaitGroup - https://gobyexample.com/waitgroups

	//start a worker goroutine , giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	<-done

	// 2 - Buffering

	// By default channels are unbuffered, meaning that they
	// will only accept sends (`chan <-`) if there is a
	// corresponding receive (`<- chan`) ready to receive the
	// sent value. _Buffered channels_ accept a limited
	// number of  values without a corresponding receiver for
	// those values.

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	channel := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	channel <- "buffered"
	channel <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	// 4 - channel-directions
	// When using channels as function parameters, you can
	// specify if a channel is meant to only send or receive
	// values. This specificity increases the type-safety of
	// the program.
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pongs, pings)
	fmt.Println(<-pongs)

}

// this `ping` function only accepts a channel for sending values.
// It would be a compile time error to try to receive on this channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// the `pong` function accepts one channel for receives ('pings') and a second for sends ('pongs')
func pong(pongs chan<- string, pings <-chan string) {
	msg := <-pings
	pongs <- msg

}
