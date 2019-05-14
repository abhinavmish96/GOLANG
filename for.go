package main

//import "fmt"

func main() {

	i := 1
	for i<=3 {
		//fmt.Println(i)
		println(i)
		i +=1
	}

	for j:=1; j<=3; j++ {
		println(j)
	} 

	for {
		println("loop")
		break
	}

	for i:=0; i<=5; i++ {
		if i%2!=0{
			println(i)
		}
	}
}