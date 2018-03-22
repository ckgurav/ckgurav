package stringutil

import "fmt"

func reverseTwo (s string)string{
	r:= []rune(s)
	for i, j :=0,len(r)-1;i<len(r)/2; i,j=i+1,j-1{
		fmt.Printf("%v\t%v\t%v\t%v\n",i,j+1,r[i],r[j])
		r[i],r[j]=r[j],r[i]
		fmt.Printf("%v\t%v\t%v\t%v\n",i,j+1,r[i],r[j])

	}
	return string("a")
}
