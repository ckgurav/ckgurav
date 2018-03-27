package main

import "fmt"

func main()  {

	func(){
		fmt.Printf("Hello, Anonymous functions\n")
	}()

	noAnonymous:=func(){
		fmt.Printf("Hello, Non-Anonymous functions")
	}

	noAnonymous()
}
