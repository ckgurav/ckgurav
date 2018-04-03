package main

import (
	"encoding/json"
	"os"
)

type student struct {
	Name string
	Address string
	Age int
}

func main() {
	s1:=student{"Chandan","Kagal",27}

	json.NewEncoder(os.Stdout).Encode(s1)
}
