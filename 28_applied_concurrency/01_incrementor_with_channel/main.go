package main

import "fmt"

func main()  {

	c1:=incrementor("First func")
	c2:=incrementor("Second func")
	c3:=puller(c1)
	c4:=puller(c2)

	fmt.Println("Total number of clounts:",<-c3+<-c4)
}

func incrementor(s string)chan int  {
	out:=make(chan int)

	go func() {
		for i:=0;i<10 ;i++{
			out<-1
			fmt.Println(s,i)
		}
		close(out)
	}()
	return out
}

func puller (in chan int) chan int{

	sum:=0
	out:=make(chan int)
	go func() {
		for n:=range in{
			sum+=n
		}
		out<-sum
		close(out)
	}()
	return out
}
