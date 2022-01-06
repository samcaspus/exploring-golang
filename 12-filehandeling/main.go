package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


func main(){
	fmt.Println("file handeling in go lang")

	content := "this is the string am going to write  dasdasda"
	fmt.Println(content)
	file, err := os.Create("./file.txt")
	if err!=nil{
		fmt.Println("something went wrong")
		panic(err)
	}
	file.WriteString(content)
	
	data, err := ioutil.ReadFile(file.Name())
	if err!=nil{
		fmt.Println("something went wrong")
		panic(err)
	}

	fmt.Println(string(data))



}