package main

import "fmt"

func main()  {

	records:= make([][]string,0)
	student1:=make([]string,0)
	student1=append(student1,"Chandan")
	student1=append(student1,"gurav")
	student1=append(student1,"10")
	student1=append(student1,"20")
	//records=append(records,student1)
	fmt.Println(student1)
	student2:=make([]string,4)
	student2[0]="test"
	student2[1]="abc"
	student2[2]="90"
	student2[3]="10"
	//
	////student2=append(student2,"Chwwwwan")
	//student2=append(student2,"gqqqq")
	//student2=append(student2,"105678")
	//student2=append(student2,"20234567")
	records=append(records,student1)
	records=append(records,student2)

	fmt.Println("records[0][2]",records[0][2])
	fmt.Println("records[0]",records[0])
	fmt.Println("records[1]",records[1])

	fmt.Println(len(records[0]),len(records[1]),len(records))
	fmt.Println()
	for i1,_:=range (records){
		for i,v:=range (records[i1]){
			fmt.Println(i,v)
		}
	}

}
