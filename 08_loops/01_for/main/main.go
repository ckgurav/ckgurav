package main

import "fmt"

func main(){
	for i:=0;i<100;i++{
		if i%2==1 {
			fmt.Printf("%d is odd number\n", i)
		}else{
			fmt.Printf("%d is even number\n", i)
		}

	}
}
