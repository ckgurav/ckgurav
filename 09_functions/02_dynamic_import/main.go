package main

import (
	"fmt"
)

func main(){

	n:= avg (1,2,3,4,5,6,7,8,9)

	fmt.Printf("\n %v",n)

}

func avg(arr ...int ) float64{
	var total int

	for _,v :=range(arr){
		total+=v
	}
	return float64(total/(len(arr)))

}