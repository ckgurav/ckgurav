package main

import "fmt"

func pritnNumbers (arr []int, callback func(int)){

	for _,no:=range(arr){
		callback(no)

	}
}

func main(){
	//
	//printFn:=func(no int){
	//	fmt.Printf("No:%d\n",no)
	//}

	pritnNumbers([]int{10,20,30,40},func(no int){
		fmt.Printf("No:%d\n",no)
	})
}
