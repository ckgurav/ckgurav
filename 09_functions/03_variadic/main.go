package main

import "fmt"

func main(){
	data := []float64{10,20,30,405,53}

	fmt.Printf("\n average is %v",avg(data...))

}

func avg(slic ...float64) float64{

	total:=0.0
	for _,v:=range(slic){
		total+=v
	}
	return total/float64(len(slic))
}