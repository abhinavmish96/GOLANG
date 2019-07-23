package main

//In the previous example we saw how to manage simple counter state using atomic operations. For more complex state we can use a mutex to safely access data across multiple goroutines.

import "fmt"

func main(){

	//For our example the state will be a map.

	// This mutex will synchronize access to state.

	//We’ll keep track of how many read and write operations we do.

	//Here we start 100 goroutines to execute repeated reads against the state, once per millisecond in each goroutine.

	//For each read we pick a key to access, Lock() the mutex to ensure exclusive access to the state, read the value at the chosen key, Unlock() the mutex, and increment the readOps count.

	//Wait a bit between reads.

	//We’ll also start 10 goroutines to simulate writes, using the same pattern we did for reads.

	//Let the 10 goroutines work on the state and mutex for a second.

	//Take and report final operation counts.

	//With a final lock of state, show how it ended up.

	//Running the program shows that we executed about 90,000 total operations against our mutex-synchronized state.

	
}