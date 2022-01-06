package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main(){


	response , _ := http.Post("https://lco.dev","application/json",nil)

	result, _ := ioutil.ReadAll(response.Body)

	var responseReader strings.Builder
	responseReader.Write(result)
	fmt.Println(responseReader.String())


}