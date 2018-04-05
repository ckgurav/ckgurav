package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
)

func main() {
	var w io.Writer

	w = os.Stdout

	fmt.Printf("\n%T",w)

	w= new(bytes.Buffer)

	fmt.Printf("\n%T",w)

	w=nil
	fmt.Printf("\n%T",w)

}

