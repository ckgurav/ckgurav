package main

import "fmt"

func main() {
	c:=make(chan int)

	go func() {
		for i:=0;i<20;i++{
			c<-i
			fmt.Println("inside:",i)

		}
		close(c)
	}()

	for n:=range(c){
		fmt.Println("Outside",n)
	}
}
