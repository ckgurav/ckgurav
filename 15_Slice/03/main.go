package main

import "fmt"

func main()  {
	var arr []int
	i:=0
	fmt.Printf("\n Enter values for slice, Press '0' for return\n ")
	fmt.Scan(&i)
	arr = append(arr,i)

	for arr[len(arr)-1] != 0{
		fmt.Scan(&i)

		arr = append(arr,i)
	}

	fmt.Println(arr)

}
