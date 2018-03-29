package main

import "fmt"

func main()  {
	myMap := map[int]string{
		0:"chandan",
		1:"Gurav",
		2:"Sap",
		3:"abc",
	}

	fmt.Println(myMap)
	fmt.Println(len(myMap))
	delete(myMap,1)
	fmt.Println(myMap)
	fmt.Println(len(myMap))

}
