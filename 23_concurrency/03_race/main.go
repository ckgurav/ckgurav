package main

import (
	"sync"
	"fmt"
	"time"
)

var counter int=0
var wg sync.WaitGroup
func main() {

	wg.Add(2)

	go increment("Foo:")
	go increment("Bar:")
	wg.Wait()
	fmt.Println("Final counter:",counter)
}

func increment(s string) {

	for i:=0;i<1000;i++{
		x:=counter
		x++
		time.Sleep(3*time.Millisecond)
		counter=x
		fmt.Println(s,counter)

	}
	wg.Done()
}
