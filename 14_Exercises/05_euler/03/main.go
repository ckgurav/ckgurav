package main

import (
	"math"
	"fmt"
)

func findPrime(number int64) int64{

	var prime int64  = -1

	for number%2==0{
		prime=2
		number/=2
	}
	var i int64
	for i=3;float64(i)<math.Sqrt(float64(number));i+=2{
		for number%i==0  {
			prime=i
			number/=i
		}
	}

	if number>2{
		prime=number
	}

	return prime

}

func main()  {

	fmt.Println(findPrime(600851475143 ))
}
