package main

import (
	"strconv"
	"fmt"
)

func main(){
	a,_:=strconv.ParseBool("true")
	fmt.Println(a)
	fmt.Printf("\n%T",a)

	b,_:=strconv.ParseFloat("1.23455",32)
	fmt.Printf("\n%v",b)
	fmt.Printf("\n%T",b)

	c,_:=strconv.ParseInt("-12",10,64)

	fmt.Printf("\n%v",c)
	fmt.Printf("\n%T",c )

}
