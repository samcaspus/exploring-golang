package main

import "fmt"

func main(){

	fmt.Println("if and else statement")

	var result string;
	loginCount := 11

	if loginCount < 10{
		result = "regularUser";
	
	}else if loginCount == 11{
		result = "number 11 here"
	
	}else{
		result = "not a regular user"
	}

	fmt.Println(result)


}