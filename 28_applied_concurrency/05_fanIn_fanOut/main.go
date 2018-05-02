package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {

	c:= fanIn(boring("Chandan"), boring("Abhi"))
	for i:=0;i<10;i++{
		fmt.Println(<-c)
	}
	fmt.Println("I am bored not, Quiting")
}

func boring(name string) <-chan string {

	c := make(chan string)

	go func() {
		for i:=0;;i++{
			c<-fmt.Sprintf("%s %d", name,i)
			time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
		}
	}()
	return c
}

func fanIn(name1,name2 <-chan string) <-chan string {

	c:=make(chan  string)

	go func() {
		for{
			c<- <-name1
		}
	}()
	go func() {
		for{
			c<- <-name2
		}
	}()

	return c

}
