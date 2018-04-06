package main

import (
	"strconv"
	"fmt"
)

func main(){
	x:="12"
	y,err:=strconv.Atoi(string(x))

	fmt.Println(y)
	fmt.Println(err)
	fmt.Printf("\n%T",x)
}
