package main

import (
	"fmt"
	"time"
	"sync"
)
var w sync.WaitGroup

func main() {
	w.Add(2)
	go foo ("Hello 1")
	go bar ("hello 2")
	w.Wait()

}

func foo(s string)  {

	for i:=0;i<10;i++ {
		fmt.Println(s,i)
		time.Sleep(3*time.Millisecond)
	}
	w.Done()
}

func bar(s string)  {

	for i:=0;i<10;i++ {
		fmt.Println(s,i)
		time.Sleep(3*time.Millisecond)
	}
	w.Done()
}