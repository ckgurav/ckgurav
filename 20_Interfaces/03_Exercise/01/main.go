package main

import (
	"sort"
	"fmt"
)

type people []string

func (p1 people) Len() int{
	return len(p1)
}

func (p1 people) Swap(i,j int){
	p1[i],p1[j]=p1[j],p1[i]
}
func (p1 people) Less(i,j int)bool  {
	return p1[i]<p1[j]

}

func  main()  {
	p1:=people{"Chandan","Amit","Nikita","Punith", "Robert"}
	b:=sort.StringsAreSorted(p1)
	fmt.Println(b)

	sort.Slice(p1, func(i, j int) bool {
		return p1[i]>p1[j]
	})
	fmt.Println(p1)

	sort.Sort(p1)
	fmt.Println(p1)
	
	
}
