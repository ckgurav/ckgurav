package main

import "fmt"

func main() {
	var test interface{} = "Chandan"

	fmt.Println(test)
	fmt.Printf("\n%T",test)
	fmt.Println(test.(string))
	fmt.Printf("\n%T",test.(string))

}
