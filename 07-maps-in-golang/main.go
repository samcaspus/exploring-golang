package main

import "fmt"


func main(){

	fmt.Println("maps in go lang")

	languages := make(map[string]string)
	languages["sandeep"] = "guptan"
	languages["a"] = "b"
	languages["c"] = "d"
	
	fmt.Println(languages)
	fmt.Println(languages["c"]=="")

	for key, value := range languages{
		fmt.Println(key, value)
	}
}