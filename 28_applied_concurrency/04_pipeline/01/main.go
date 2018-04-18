package main

import "fmt"

func main() {
	a1:=sqAdd(1,2,3)
	a2:=seResult(a1)

	for n:=range a2{
		fmt.Println(n)
	}

}

func sqAdd(number ...int)chan int {

	out:=make(chan int)

	go func() {
		for _,i:=range number{
			out<-i
		}
		close(out)
	}()
	return out
}

func seResult(in chan int) chan int {

	out:=make(chan  int)

	go func() {
		for n:=range in{
			out<- n*n
		}
		close(out)
	}()
	return out
}
