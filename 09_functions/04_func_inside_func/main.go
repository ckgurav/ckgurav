package main

import "fmt"

func helloWorld() func() string{

	strng:=func() string{
		return fmt.Sprint("\n Hello World, from inside function")
	}
	return strng
}

func main(){

	fmt.Println((helloWorld())())
}