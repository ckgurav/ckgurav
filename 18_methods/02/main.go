package main

import (
	//"encoding/xml"
	"fmt"
)

type str1 struct {
	name string
	age int
}

type str2 struct {
	a1 str1
	name string
	age int
}

func (s1 str1) printing(){
	fmt.Println("Inside struct Str1",s1.name)
}

func (s2 str2) printing(){
	fmt.Println("Inside struct Str2",s2.name)
}

func main()  {
	strct:=str2{str1{"Inside",25},"outside",26}

	strct.printing()
	strct.a1.printing()

}