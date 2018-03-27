package main

import "fmt"

func main(){
	arr:=[]int{7,3,0,45,3,957,35325,866}
	foo(1,2)
	fmt.Println("foo(1,2)")
	foo(1,2,3)
	fmt.Println("foo(1,2,3)")
	foo(arr...)
	fmt.Println("foo(arr)")
	foo()
	fmt.Println("foo()")
}

func foo(arr ...int){
	for _,v:=range (arr){
	fmt.Println("Number passed ",v)
	}

}
