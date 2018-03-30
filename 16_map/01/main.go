package main

import "fmt"

func main()  {

	mapExample:=make(map[string]int)

	mapExample["a1"]=11
	mapExample["a2"]=12
fmt.Println(mapExample,len(mapExample))
	for No,V1:=range (mapExample){
		fmt.Println(No,V1)
	}
	//delete(mapExample,"a1")
	i,j:=mapExample["a1"]



	fmt.Println(i,j)
}
