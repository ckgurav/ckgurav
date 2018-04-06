package main

import "fmt"

func main()  {
	a:=13
	b:=13.999

	fmt.Println(float64(a)+b)
	fmt.Println(a+int(b))

	var r rune = 'a'
	var i1 int = 'b'

	fmt.Println(int(r)+i1)
	fmt.Println(string(r+rune(i1)))

	fmt.Println(string([]byte{'c','h','a','n','d','a','n'}))

	fmt.Println([]byte(string("Chandan")))

}
