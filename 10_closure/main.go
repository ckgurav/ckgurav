package main

import "fmt"

func wrapper() func() int{
	x:=0

	return func() int{
		x++	
		return x
	}
}

func main(){
	innerFn :=wrapper()
	fmt.Printf("\n%d",innerFn())
	fmt.Printf("\n%d",innerFn())
	fmt.Printf("\n%d",wrapper()())
	fmt.Printf("\n%d",wrapper()())
}