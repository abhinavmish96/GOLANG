package main

import "fmt"

func main(){
	//var a
	 var a [2][2] int
	 a[0][1] = 9
	 fmt.Println(a[0][1])
	 fmt.Println(a)
	 for j:=0;j<2;j++ {
		for i:=0;i<2;i++{
			a[j][i]=i+j
		}
	 }
	 fmt.Println(a)
	}