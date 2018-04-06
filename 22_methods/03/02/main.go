package main

import (
	"strconv"
	"fmt"
)

func main(){
	x:=12
	y:=strconv.Itoa(x)
	fmt.Println(y)

	fmt.Printf("\n%T",y)
}
