package main

import "fmt"

func factorial(no uint64) uint64  {

	if no==0{
		return 1
	}
	return uint64(no*factorial(no-1))

}

func main()  {

	fmt.Printf("\nFactorial:%v",factorial(60))
}
