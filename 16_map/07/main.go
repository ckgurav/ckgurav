package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main()  {
	result,err:=http.Get("https://raw.githubusercontent.com/first20hours/google-10000-english/master/google-10000-english-no-swears.txt")
		if err!=nil{
		log.Fatalln(err)
	}

	words:=map[string]string{}

	sc:=bufio.NewScanner(result.Body)
	sc.Split(bufio.ScanWords)
	i:=0
	for sc.Scan() {
		words[sc.Text()]= fmt.Sprint(i)
		fmt.Println(sc.Text(), " ", i)
		i++

	}
	if err=sc.Err();err!=nil{
		log.Fatalln(err)
	}

	for word,Val:=range words {
		fmt.Println(word, " " , Val)
	}

}
