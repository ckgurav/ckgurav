package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

type student struct {
	Name string
	Address string
	Age int
}

func main() {
	//s1:=student{"Chandan","Kagal",26}
	var s1 student
	//json.NewDecoder(os.Stdin).Decode(&s1)
	json.NewDecoder(strings.NewReader(`{"Name":"Chandan","Address":"Kagal","Age":26}")`)).Decode(&s1)
	fmt.Println(s1)
}
//{"Name":"Chandan","Address":"Kagal","Age":26}