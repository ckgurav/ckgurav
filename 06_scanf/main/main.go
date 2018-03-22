package main

import (
	"fmt"
	"math"
)

const pi float64 = 4.14

func main(){
	fmt.Println("Enter redius: ")
	var r float64
	fmt.Scan(r)

	var surface float64 = pi*(math.Pow(r,2))

	fmt.Printf("Surface is:%f",surface)

}