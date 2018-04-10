package main

import "fmt"

func main() {
	c:=make(chan int) // Wrong way-- We dont have any buffer defined for channel C,
	// system will wait immediately after line 13 for some process to take out value of c
	// to slove this issue, we can add buffer (limit) to the channel c
	//c:=make(chan int,30)
	done:=make(chan bool)

	go func() {
		for i:=0;i<10;i++{
			c<-i
		}
		done<-true
	}()

	go func() {
		for i:=20;i<30;i++{
			c<-i
		}
		done<-true
	}()

	<-done
	<-done
	close(c)

	for n:=range c{
		fmt.Println(n)
	}
}
