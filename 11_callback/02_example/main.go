package main

import (

	"fmt"
)

func greaterThan(arr []int,callback func(int)bool) []int{
	var rArr []int
	for _,no:=range(arr){
		if callback(no){
			rArr=append(rArr,no)
		}
	}
	return rArr
}

func main()  {
	fn:= greaterThan([]int{12,23,32,0,12,1},func(no int) bool{
		if no<=1 {
			return true
		}
		return false
	})

	fmt.Printf("/n%v",fn)
}