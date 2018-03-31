package main

import "fmt"

type  foo int

type student struct {
	name string
	age int
	std int
}

type school struct {
	student
	name string

}

func main()  {
	s1:= school{
		student{
			"Chandan",
			25,
			10,
		},
		"Shahu High",

	}
	s7:=school{student{"lmn",11,10},"Jaysnighrao"}

	fmt.Println(s1.name)
	fmt.Println(s1)
	fmt.Println(s1.student)
	fmt.Println(s1.student.name)
	s1.age = 33
	var s2 foo
	s2 = 100
	fmt.Println(s1.age)
	fmt.Println(s1.std)
	s3:=0
	fmt.Printf("S2 Type:%T and Value %v, Type of S1:%T",s2,s2,s3)

	s4:=s2+foo(s3)

	fmt.Printf("\n S4 type:%T and value:%v\n",s4,s4)

	var s5 school
	s5.age=10

	fmt.Println("S5",s5.age," ",s5)

	s6:=school{}

	s6.student.name="abc"
	fmt.Println(s7)


}