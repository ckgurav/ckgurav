package main

import "fmt"

type Vehicels interface {

}

type Vehicel struct {
	Name string
	Age string
}
type car struct {
	Vehicel
	cc int
}
type boat struct {
	Vehicel
	size int
}

type jet struct {
	Vehicel
	speed int
}

func main()  {
	maruti:=car{}
	honda:=car{}
	INS_Virat:=boat{}
	INS_Vikrant:=boat{}
	boing:=jet{}
	airbus:=jet{}

	arr:=[]Vehicels{maruti,honda,INS_Vikrant,INS_Virat,boing,airbus}

	for key,Val:=range(arr){
		fmt.Println(key,"-",Val)
	}

}


