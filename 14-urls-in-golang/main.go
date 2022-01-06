package main

import (
	"fmt"
	"net/url"
)

const baseUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dsadasdasdas"
func main(){

	fmt.Println("welcome to handeling url in golang")
	fmt.Println(baseUrl)

	result, _ := url.Parse(baseUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Query())
	fmt.Println(result.Port())
	fmt.Println(result.ForceQuery)


	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/test",
		RawQuery: "sandeep=guptan&sam=caspus",

	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
	

}