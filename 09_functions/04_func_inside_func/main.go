package main

import "fmt"

func helloWorld() func() string{

	string:=func() string{
		return fmt.Sprint("\n Hello World, from inside function")
	}
	return string
}

func main(){

	fmt.Println((helloWorld())())
}