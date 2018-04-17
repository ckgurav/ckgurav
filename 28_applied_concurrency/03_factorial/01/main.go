package main

import "fmt"

func main() {
	f:=factorial(4)
	d:=factResult(f)
	fmt.Printf("=%d",<-d)
}

func factorial(no int) chan int {
	out:=make(chan int)
	go func() {
		for i:=no;i>0;i--{
			out<-i
			fmt.Printf("%d",i," * ")
		}
		close(out)
	}()
	return out
}

func factResult(fct chan int ) chan int{
	total:=1
	result:=make(chan int)
	go func() {
		for n:= range fct{
			total*=n
		}
		result<-total
		close(result)
	}()
	return result

}


