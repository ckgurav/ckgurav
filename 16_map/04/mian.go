package main

import "fmt"

func main()  {
	myMap:= map[int]string{
		1: "one",
		2:"Two",
		3:"Three",

	}

	if val,exists:=myMap[1];exists{

		fmt.Println(val)
	}else{
		fmt.Println("1" + "Not available")
	}
	myMap[4]="Four"
	fmt.Println(myMap)
	myMap[3]="Three again"
	fmt.Println(myMap)
	delete(myMap,3)
	fmt.Println(myMap)
	myMap[3]="Three again"

}
