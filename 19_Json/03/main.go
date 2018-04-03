package main

import (
	"fmt"
	"encoding/json"
)

type student struct {
	Name string
	Address string
	Age int
}

func main() {
	var s1 student

	fmt.Println(s1)

	bs:=[]byte(`{"Name":"Chandan","Address":"Kagal","Age":26}`)

	json.Unmarshal(bs,&s1)

	fmt.Println(s1)
}
