//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
package main

import (
	"fmt"

)

func main(){
	var i1 uint64
	var i2 uint64
	i1=1
	i2=2
	var addNo uint64
	for i2<=4000000{
		i2,i1 =  i2+i1,i2
		if i1%2==0{
			addNo+=i1
		}

	}
	fmt.Println(addNo)
}