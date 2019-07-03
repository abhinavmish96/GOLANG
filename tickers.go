package main

// Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.

import "fmt"

func main(){
	//Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll use the range builtin on the channel to iterate over the values as they arrive every 500ms.

	//Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
}

//When we run this program the ticker should tick 3 times before we stop it.