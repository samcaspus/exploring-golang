package main

import (
	"fmt"
	"sort"
)

func main(){

	fmt.Println("slices in array")

	var fruitList = []string{"apple","mango","banana"}
	fruitList = append(fruitList[0:1],"nope","nope")
	fmt.Println(fruitList)

	number := make([]int, 4)
	number[0]=0
	number[1]=1
	number[2]=2
	number[3]=3
	
	number = append(number,23,11,23)

	fmt.Println(number)
    sort.Ints(number)
	sort.Sort(sort.Reverse(sort.IntSlice(number)))
	
	fmt.Println(number)


	sliceToBeRemoved := 2
	fmt.Printf("slice which is gettin removed is %d\n",sliceToBeRemoved)
	fmt.Println(number)
	number = append(number[0:sliceToBeRemoved],number[sliceToBeRemoved+1:]...)
	fmt.Println(number)

}