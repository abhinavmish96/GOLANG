package main

import "fmt"

//By convention, errors are the last return value and have type error, a built-in interface.

//errors.New constructs a basic error value with the given error message.

//A nil value in the error position indicates that there was no error.

// It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

func main() {

	//In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.

	//The two loops below test out each of our error-returning functions. Note that the use of an inline error check on the if line is a common idiom in Go code.

	//If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion.
}