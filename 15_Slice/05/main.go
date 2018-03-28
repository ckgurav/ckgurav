package main

import "fmt"

func main()  {
	tranx:=make([][]int,0)

	for i:=0;i<5;i++{
		trans:=make([]int,0)
		for j:=100;j>=95;j--{
			trans=append(trans,j)
		}
		tranx=append(tranx,trans)
	}

	fmt.Println(tranx)
}
