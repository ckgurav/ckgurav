package main

import "fmt"

func main()  {
	var mapFn = make( map[string]string,10)
	//mapFn := map[string]string{}
	//var mapFn = map[string]string {}
	var mapFn1 = map[string]string{}
	mapFn ["at"] = "lmp"
	mapFn1 ["at"] = "lmp"

	fmt.Println(mapFn["at"])
	fmt.Println(len(mapFn))

}
