package main

import "fmt"

func main(){


	a,b := test("Chandan","Gurav")

	fmt.Printf("%s and %s",a,b)

}

func test(a,b string) (string,string) {
	return fmt.Sprint(a," ",b), fmt.Sprint(b," ",a)

	}