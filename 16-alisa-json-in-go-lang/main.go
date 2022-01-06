package main

import (
	"encoding/json"
	"fmt"
	"strings"
)


type Course struct{

	Name string `json:"courseName"`
	Price int `json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main(){
	fmt.Println("json in golang")
	EncodingJson()
}


func EncodingJson(){

	courses := []Course{
		{"one",10,"one","one",[]string{"web","android"}},
		{"two",10,"two","two",[]string{"android"}},
		{"three",10,"three","three",nil},	
	}

	fmt.Println(courses)
	var responseReader strings.Builder
	finalJson, err := json.MarshalIndent(courses,"","\t")
	if err!=nil{
		panic(err)
	}
	responseReader.Write(finalJson)
	fmt.Println(responseReader.String())
}