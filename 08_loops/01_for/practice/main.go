package main

import (
	"fmt"
)



func main(){
	const a1=30
	const b1=25
	var arr [a1][b1]int
	var x,t,i,j int = 0,0,0,0
	a := a1
	b:=b1
	for arr[i][t] ==0{

		if arr[i][t] ==0 {
			for j = t; j < b; j++ {
				arr[i][j] = x
				fmt.Printf("1a[%d][%d]=%d\t", i, j, arr[i][j])
				x++
			}
			j--
		}

		if arr[t + 1][j] ==0 {
			for i = t + 1; i < a; i++ {
				arr[i][j] = x
				fmt.Printf("2a[%d][%d]=%d\t", i, j, arr[i][j])
				x++
			}
			i--
		}
		if arr[i][j - 1] ==0 {

			for j = j - 1; j >= t; j-- {
				arr[i][j] = x
				fmt.Printf("3a[%d][%d]=%d\t", i, j, arr[i][j])
				x++
			}
			j++
		}
		if arr[i - 1][j] ==0 {

			for i = i - 1; i > t; i-- {
				arr[i][j] = x
				fmt.Printf("4a[%d][%d]=%d\t", i, j, arr[i][j])
				x++
			}
			i++
		}
		t++
		a--
		b--
		fmt.Printf("a=%d\tb=%d\tt=%d\n",a,b,t)

	}
fmt.Printf("\n")
	for i=0;i<a1;i++{
		for j=0;j<b1;j++{
			fmt.Printf("a[%d][%d]=%d\t",i,j,arr[i][j])
		}
		fmt.Printf("\n")
	}

}
