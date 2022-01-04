package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("hello");

	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"));

	createdtime := time.Date(2020,time.August, 10,23,23,0,0,time.UTC)
	updatedtime := createdtime.Format("01-02-2006")
	fmt.Println(createdtime)
	fmt.Println(updatedtime)
}