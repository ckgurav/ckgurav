package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64{
	return s.side*s.side
}

func (c circle)area() float64  {
	return c.radius*c.radius*(3.14)
}

type shape interface {
	area() float64

}

//type trangle struct {
//	side float64
//}
func info (z shape)  {
	fmt.Println(z)
	fmt.Println(z.area())
}

func  main()  {
	s:=square{10}
	c:=circle{10}
	//t:=trangle{10}

	info(s)
	info(c)
	//info(t)
	//fmt.Println(info(s))
}