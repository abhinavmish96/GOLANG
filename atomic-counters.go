package main

//The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.

import "fmt"

func main() {
	//We’ll use an unsigned integer to represent our (always-positive) counter.

	//To simulate concurrent updates, we’ll start 50 goroutines that each increment the counter about once a millisecond.

	//To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.

	//Wait a bit between increments.

	// Wait a second to allow some ops to accumulate.

	//In order to safely use the counter while it’s still being updated by other goroutines, we extract a copy of the current value into opsFinal via LoadUint64. As above we need to give this function the memory address &ops from which to fetch the value.
}

//Running the program shows that we executed about 40,000 operations.