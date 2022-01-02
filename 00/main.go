package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	fmt.Println("welcome to pizza app");
	fmt.Print("please rate the pizza between 1 and 5 = ")

	reader := bufio.NewReader(os.Stdin)
	data , err := reader.ReadString('\n');
	if err != nil{
		
	}else{
		data = strings.TrimSpace(data)
	}
	
	data2 , err := strconv.ParseFloat(data, 64);

	if err != nil{
		fmt.Print("something went wrong");
	}else{
		value := data2+1
		fmt.Println(value)
		
	}
	

	


}