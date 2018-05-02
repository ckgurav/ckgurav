package main

import "fmt"

func main()  {

	for n := range sq(sq(gen(1,2,3))){
		fmt.Println(n)
	}
}

func gen(no ...int) chan int {
	out:= make(chan int)

	go func() {
		for _,n:=range no{
			out<-n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out:=make(chan int)

	go func() {

		for n:=range in{
			out<-n*n
			fmt.Println(n,n*n)
		}
		close(out)
	}()
	return out

}