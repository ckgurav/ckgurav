package main

import "fmt"

type test int
type test1 string

func (i test)valueReciver(){
	i=i*13
	fmt.Printf("%T\n",i)
	fmt.Println(i)
}

type use interface {
	valueReciver()
}

type applyToAnything interface {

}

func used(s use)  {
	s.valueReciver()
}

func main() {
	i:=test(11)
	used(i)
	used(&i)
	//a:=[]use{i,test(12)}
	b:=[]applyToAnything{i,test1("sadfsd"),test(11),12,"aaa",3.12}


	for key,value:=range (b){
		fmt.Println(key,value)
	}

	for z:=0;z<len(b) ;z++  {
		fmt.Printf("\n %T",b[z])
	}
}
