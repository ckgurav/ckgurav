package main

import "fmt"

func main()  {
	mySlice:=make([]int,0,5)

	fmt.Println(mySlice,len(mySlice),cap(mySlice))

	for i:=0;i<100;i++{
		mySlice=append(mySlice,i)
		fmt.Println( len(mySlice),cap(mySlice))
	}

}