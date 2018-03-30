package main

import "fmt"

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

	fmt.Println(s1.name)
	fmt.Println(s1)
	fmt.Println(s1.student)
	fmt.Println(s1.student.name)
	s1.age = 33
	fmt.Println(s1.age)
	fmt.Println(s1.std)


}