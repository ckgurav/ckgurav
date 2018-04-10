package main

import (
	"sync"
	"fmt"
)

var wg  = sync.WaitGroup{}
var c=make(chan int)

func main() {
	wg.Add(2)
	go func() {
		//wg.Add(1)
		for i:=0;i<10;i++{
			c<-i
		}
		wg.Done()
	}()
	go func() {
		//wg.Add(1)
		for i:=20;i<30;i++{
			c<-i
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c)
	}()

	for n:=range c{
		fmt.Println(n)
	}



}

