package main

import "fmt"

func main(){
	arr:=[]int{7,3,0,45,3,957,35325,866}
	fmt.Printf("\n%d",largeNo(arr...))
}

func largeNo(arr ...int) int  {
	largestNo:=0
	largestNo=arr[0]
	for _,v:=range (arr){

		if v>largestNo{
			largestNo=v
		}
	}
	return  largestNo
}
