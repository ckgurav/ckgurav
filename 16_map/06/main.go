package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main()  {
	result,err:=http.Get("https://raw.githubusercontent.com/first20hours/google-10000-english/master/google-10000-english-no-swears.txt")

	if err!=nil{
		log.Fatalln(err)
	}

	words,_:= ioutil.ReadAll(result.Body)

	//str:=string(words)

	fmt.Println(string(words))
	fmt.Println(len(words),cap(words))

}
