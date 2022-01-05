package main

import "fmt"

func main(){


	number := 2

	switch number{
	case 1: fmt.Println("do something1");
	case 2: fmt.Println("do something else2")
	fallthrough
	case 3: fmt.Println("this is case 3")
	default:
		fmt.Println("this is default")
	}
}