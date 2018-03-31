package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Fname string
	Lname string `json:"-"`
	Age	int `json:"Testing"`
}

func main() {
	p1:=Person{"Chandan","Gurav",77}

	bs,_:=json.Marshal(p1)

	fmt.Println(bs)
	fmt.Println(string(bs))

}

