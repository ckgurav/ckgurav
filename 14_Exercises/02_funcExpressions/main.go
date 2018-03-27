package main
import "fmt"

//func half(number int) (int,bool){
//
//	ret1:=number/2
//
//	if number%2 == 0{
//		return ret1,true
//	}
//	return number,false
//}

func main(){

	half:= func(number int) (int,bool){

			ret1:=number/2

			if number%2 == 0{
				return ret1,true
			}
			return number,false
	}
	fmt.Printf("\n Func Expressions\n")
	fmt.Println(half(1))
	fmt.Println(half(2))
}
