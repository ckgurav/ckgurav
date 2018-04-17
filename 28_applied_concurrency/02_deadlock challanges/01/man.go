package main

import "fmt"

//This below code results in deadlock

//func main() {
//
//	c:=make(chan  int)
//
//	c<-1
//	fmt.Println(<-c)
//}

//Solution:

func main()  {
	c:=make(chan int)

	go func() {
		c<-1
	}()
	fmt.Println(<-c)

}