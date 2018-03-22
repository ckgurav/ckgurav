package main

import "fmt"

const(
	_=iota
	a = 1<<(iota*10)
	b = 1<<(iota*10)
	c = 1<<(iota*10)
)

func main(){
	fmt.Printf("%d \t %d \t %d\n",a,b,c)
	fmt.Printf("%b \t %b \t %b\n",a,b,c)


}