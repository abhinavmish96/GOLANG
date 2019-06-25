package main

//Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

import "fmt"

func main(){
	//For our example we’ll select across two channels.

	//Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.

	//We’ll use select to await both of these values simultaneously, printing each one as it arrives.
}