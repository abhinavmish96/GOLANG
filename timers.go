package main

//We often want to execute Go code at some point in the future, or repeatedly at some interval. Go’s built-in timer and ticker features make both of these tasks easy.

import "fmt"
import "time"

func main() {
	//Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.

	//The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.

	//If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it expires. Here’s an example of that.
}