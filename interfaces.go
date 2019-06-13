package main

import "fmt"

//Here’s a basic interface for geometric shapes.

//For our example we’ll implement this interface on rect and circle types.

// To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects.

//The implementation for circles.

// If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.

func main(){

	// The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
}