package main

import (

	"fmt"
)

type person struct {
	fname string
	lname	string
	age int
}

func main()  {

	p1:=person{"Chandan","Gurav",27}

	fmt.Println(printName(p1))
}

func printName(p1 person) string{

	return  p1.fname+" "+p1.lname
}
