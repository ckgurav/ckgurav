package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"time"
)

var counter int64=0
var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go increment("Foo:")
	go increment("Bar:")
	wg.Wait()
	fmt.Println("Final Count:",counter)
}

func increment(s string)  {

	for i:=0;i<20;i++{
		//x:=counter
		//x++
		//counter=x
		time.Sleep(3*time.Millisecond)
		atomic.AddInt64(&counter,1)
		fmt.Println(s,counter)

	}
	wg.Done()

}

