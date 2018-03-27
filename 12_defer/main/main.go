package main

import "fmt"

func world(){
	fmt.Printf("world")
}
func hello(){
	//world()
	//hello()
	defer world()
	fmt.Printf("\nhello")
}

func main()  {
	hello()

}