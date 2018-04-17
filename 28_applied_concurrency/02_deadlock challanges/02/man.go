package main

import "fmt"

//func main() {
//	c:=make(chan int)
//
//	go func() {
//		for i:=0;i<10;i++{
//			c<-i
//		}
//	}()
//
//	fmt.Println(<-c)
//}
//Why this program is just printing 0, what needs to be done to print all values from 0 to9
func main() {
	c:=make(chan int)

	go func() {
		for i:=0;i<10;i++{
			c<-i
		}
		close(c)
	}()
	for n:=range c{
		fmt.Println(n)
	}

}