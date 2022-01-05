package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main(){

	fmt.Println("web requests")
	response, err := http.Get(url)
	if err!=nil{
		panic(err)
	}

	fmt.Println(response)
	defer response.Body.Close()

	content, err :=ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	finalcontent := string(content)
	fmt.Println(finalcontent)

}