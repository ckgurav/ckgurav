package main

import "fmt"

func main() {

	c:=make(chan int)
	done:=make(chan bool)
	fmt.Println("Hii")
	go func() {
		for i:=0;i<100;i++{
			c<-i
		}
		fmt.Println("Hii")
		close(c)

	}()

	go func() {
		for n:=range c{
			fmt.Println("First one priting:",n)
		}
	done<-true
	}()

	go func() {
		for n:=range c{
			fmt.Println("Second one printing:",n)
		}
		done<-true
	}()
		<-done
		<-done
}
