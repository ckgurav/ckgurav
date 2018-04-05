package main

import "fmt"

type animal struct {
	voice string
}
type cat struct {
	animal
	canFly bool
}
type parrot struct {
	animal
	canFly bool
}

func printing(a interface{}){
	fmt.Println(a)
}

func main(){
	Cat:=cat{animal{""},false}
	Parrot:=parrot{animal{"hahahahhaa"},true}

	printing(Cat)
	printing(Parrot)
	printing(111)
	printing("String")


}


