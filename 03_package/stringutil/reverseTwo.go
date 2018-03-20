package stringutil

import "fmt"

func reverseTwo (s string)string{
	r:= []rune(s)
	for i, j :=0,len(r)-1;i<len(r)/2; i,j=i+1,j-1{

		r[i],r[j]=r[j],r[i]
		fmt.Println(i,j+1,"%u\t%u\n",r[i],r[j])

	}
	return string(r)
}
