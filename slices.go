package main

import "fmt"

func main() {

	s := make([]int, 3)
	fmt.Println(s)

	s[0]=0
	s[1]=1
	s[2]=2
	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s,3)
	fmt.Println(s)

	s = append(s,4,5,6,7)
	fmt.Println(s)

	c := make([]int, len(s))
	copy(c,s)
	fmt.Println(c)

	l := c[:]
	fmt.Println(s)
	fmt.Println(l)

	twoD := make([][]int, 3)
	for i:=0;i<3;i++{
		in:=i+1
		twoD[i]= make([]int, in)
		for j:=0;j<in;j++{
			twoD[i][j]= i+j
		}
	}
fmt.Println(twoD)
}