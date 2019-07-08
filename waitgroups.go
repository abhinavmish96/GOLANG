package main

//To wait for multiple goroutines to finish, we can use a wait group.

import "fmt"

//This is the function we’ll run in every goroutine. Note that a WaitGroup must be passed to functions by pointer.
func worker(){
	//Sleep to simulate an expensive task.

	// Notify the WaitGroup that this worker is done.
}

func main(){
	//This WaitGroup is used to wait for all the goroutines launched here to finish.

	//Launch several goroutines and increment the WaitGroup counter for each.

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.

}

//The order of workers starting up and finishing is likely to be different for each invocation.