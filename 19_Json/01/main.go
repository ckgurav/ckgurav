package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//import "go/ast"

type Person struct {
	Name string
	Country	string
	Age int
	notExported bool
}

func main()  {
	JsonPerson:=Person{"Chandan","India",26,false}

	bs,err:=json.Marshal(JsonPerson)
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println(bs)
	fmt.Printf("\n%T\n",bs)
	fmt.Println(string(bs))
}

