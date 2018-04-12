package main

import (
	"fmt"
)

func main() {
	c:=incrementor()
	cSum:=sumFunc(c)
	fmt.Println("Main printing")

	for n:=range cSum{
		fmt.Println(n)
	}
}

func incrementor()chan int  {
	out:=make(chan int)

	go func() {
		for i:=0;i<10;i++{
			fmt.Println("Incrementor printing",i)
			out<-i
		}
		close(out)
	}()
	fmt.Println("Incrementor returning")
	return out
}

func sumFunc(c chan int) chan int {
	out:=make(chan int)
	sum:=0
	go func() {
		for n:=range c{
			fmt.Println("cSum printing",n)
			sum=sum+n
		}
		out<-sum
		close(out)
	}()
	fmt.Println("cSum returning",sum)
	return out
}
