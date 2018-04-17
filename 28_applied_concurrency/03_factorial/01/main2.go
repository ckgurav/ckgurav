package main

import "fmt"
var no int = 10

func main() {

	f1:=factorial2()
	f2:=factorial2()
	f3:=puller(f1)
	f4:=puller(f2)
	n1,n2:=0,0

	go func() {
		for n1=range f4{
			n1=n1
			fmt.Println(n1)
		}
		for n2=range f3{
			n2=n2
			fmt.Println(n2)
		}
	}()
	if no==0{
		//fmt.Println("final",n1*n2)
	}

}

func factorial2() chan int {
	chan1 := make(chan int)
	go func() {
		if no==0{
			close(chan1)
		}
		tmp := no
		no--
		tmp *= no
		no--
		fmt.Println(tmp,no)
		chan1 <- tmp
	}()
	return chan1
}

func puller(f1 chan int)chan int {
	result:=make(chan int)
	total:=1
	go func() {
		for n:=range f1{
			total*=n
			fmt.Println(total)
		}

		result<-total
		close(result)
	}()
	return result
}