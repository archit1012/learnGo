// Go's select let us you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go

package main

import (
	"fmt"
	"time"
)


func main(){
	// For our example we will select across two channels
	c1 := make(chan string)
	c2 := make(chan string)


// Each channel will receive a value after some amount of 
	// time, to simulate e.g. blocking RPC operations executing in concurrent goroutines 
	
	// Anonymous function
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func(){
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

// We’ll use select to await both of these values simultaneously, 
// printing each one as it arrives.

// ToDo : understand why if i = 3 it throws error : "fatal error: all goroutines are asleep - deadlock!"

	for i := 0 ; i < 2 ; i++ {
		fmt.Println("before case")
		select{
			case msg1 := <-c1:
				fmt.Println("Received from c1", msg1)
			case msg2 := <-c2:
				fmt.Println("Received from c2", msg2)
		}
		fmt.Println("after case")
	}
}

//Output : 
// Received from c1 one
// Received from c2 two
// Received from c2 two


// We receive the values "one" and then "two" as expected.
// Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently