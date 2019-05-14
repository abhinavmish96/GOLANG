package main

import "fmt"

func main() {
	i:=5
	fmt.Print("a " ,i," a ")
	switch i{
	case 1 :println("one")
	case 2 :println("two")
	default :println("nothing")
	}

	kk := func(k int, m float32) {
		println(k, "   " ,m)
	}
	kk(6,6.5)
}