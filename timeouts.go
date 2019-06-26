package main

//Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and select.

import "fmt"

func main(){
	//For our example, suppose we’re executing an external call that returns its result on a channel c1 after 2s

	// Here’s the select implementing a timeout. res := <-c1 awaits the result and <-Time.After awaits a value to be sent after the timeout of 1s. Since select proceeds with the first receive that’s ready, we’ll take the timeout case if the operation takes more than the allowed 1s.

	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.

}

//Running this program shows the first operation timing out and the second succeeding.

//Using this select timeout pattern requires communicating results over channels. This is a good idea in general because other important Go features are based on channels and select. We’ll look at two examples of this next: timers and tickers.