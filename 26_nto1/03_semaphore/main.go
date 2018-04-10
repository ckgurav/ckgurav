package main

import "fmt"

func main() {
	c:=make(chan int)
	s:=make(chan bool)

	go func() {
		for i:=0;i<10;i++{
			c<-i
		}
		s<-true

	}()

	go func() {
		for i:=20;i<30;i++{
			c<-i
		}
		s<-true
	}()

	go func() {
		<-s
		<-s
		close(s)
		close(c)
	}()

	for n:=range c{
		fmt.Println(n)
	}

}
