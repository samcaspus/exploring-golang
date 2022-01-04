package main

import "fmt"

func main(){


	fmt.Println("Pointers")

	var one *int;
	fmt.Println(one)

	var number int = 32;
	var pointer *int = &number
	fmt.Println(pointer)
	fmt.Println(*pointer)
	
}