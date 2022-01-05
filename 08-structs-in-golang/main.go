package main

import "fmt"

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func main(){
	fmt.Println("structs in golang")
	sandeep := User{"sandeep","sandeep@gmail.com",true,23}
	fmt.Println(sandeep)
	fmt.Println(sandeep.Name)

}