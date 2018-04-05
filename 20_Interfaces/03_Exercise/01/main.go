package main

import (
	"sort"
	"fmt"
)

type people []string
type Age []int

func (p1 people) Len() int{
	return len(p1)
}

func (p1 people) Swap(i,j int){
	p1[i],p1[j]=p1[j],p1[i]
}
func (p1 people) Less(i,j int)bool  {
	return p1[i]<p1[j]

}

func (p1 Age) Len() int{
	return len(p1)
}

func (p1 Age) Swap(i,j int){
	p1[i],p1[j]=p1[j],p1[i]
}
func (p1 Age) Less(i,j int)bool  {
	return p1[i]<p1[j]

}

func  main()  {
	p1:=people{"Chandan","Amit","Nikita","Punith", "Robert"}

	b:=sort.StringsAreSorted(p1)
	fmt.Println(b)
//Method 1
	sort.Slice(p1, func(i, j int) bool {
		return p1[i]>p1[j]
	})
	fmt.Println(p1)
	b=sort.StringsAreSorted(p1)
	fmt.Println(b)

	//sort.StringSlice(p1).Sort()
//Method 2
	sort.Sort(p1)
	b=sort.StringsAreSorted(p1)
	fmt.Println(b)
	fmt.Println(p1)

	//Method 3
	p2:=people{"Chandan","Amit","Nikita","Punith", "Robert"}
	sort.StringSlice(p2).Sort()
	b=sort.StringsAreSorted(p2)
	fmt.Println(b)
	fmt.Println(p2)
//Method 4
	p3:=people{"Chandan","Amit","Nikita","Punith", "Robert"}
	sort.Sort(sort.StringSlice(p3))
	sort.StringSlice(p3).Sort()
	b=sort.StringsAreSorted(p3)
	fmt.Println(b)
	fmt.Println(p3)

//Method 5
	p4:=people{"Chandan","Amit","Nikita","Punith", "Robert"}
	sort.Strings(p4)
	sort.StringSlice(p4).Sort()
	b=sort.StringsAreSorted(p4)
	fmt.Println(b)
	fmt.Println(p4)


	A1:=Age{3,21,1,5,2,445,1,332,54,22}

	sort.Sort(A1)
	fmt.Println(A1)

	
	
}
