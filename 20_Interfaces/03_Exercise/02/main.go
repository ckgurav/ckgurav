package main

import (
	"sort"
	"fmt"
)

type people []string

func (p people) Len()int{
	return len(p)
}
func (p people)Swap(i,j int)  {
	p[i],p[j]=p[j],p[i]

}
func (p people) Less(i,j int)bool  {
	return p[i]<p[j]
}

func main() {

	p:=people{"Chandan","Amit","Nikita","Punith", "Robert"}

	sort.Sort(sort.Reverse(sort.StringSlice(p)))

	fmt.Println(p)
}