package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	//"os"
	"bufio"
	"strings"
)

func main()  {
	response,err:=http.Get("http://humanstxt.org/humans.txt")
	//response,err:=http.Get("file:///C:/Users/I335763/Desktop/sample.txt")

	if err!=nil {
		log.Fatalln(err)
	}

	inside,_:=ioutil.ReadAll(response.Body)
	bs:=bufio.NewScanner(strings.NewReader(string(inside)))
	bs.Split(bufio.ScanWords)
	i:=0
	for bs.Scan(){
		i++
	}

	fmt.Println(i)
}
