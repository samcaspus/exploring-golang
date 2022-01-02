package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("hello world");

	var number1 int64 = 10
	var number2 float64 = 10

	fmt.Println(number1+int64(number2))
	rand.Seed(time.Now().UnixMilli())
	
	for i:=0; i<10; i++{
		value := rand.Intn(10000)
		fmt.Println(value)
	}
	
}