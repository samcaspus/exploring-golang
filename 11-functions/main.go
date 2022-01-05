package main

import "fmt"

func main(){

	fmt.Println("functions in golang")
	greeter()

	sandeep := User{"sandeep",23,true}
	fmt.Println(sandeep)
	sandeep.getStatus()
	sandeep.getAge()
}

func greeter(){
	fmt.Println("this is greeter")
}


type User struct{
	Name string
	Age int
	Status bool
}


func (user User) getStatus(){
	fmt.Println(user.Status)
}


func (user User) getAge(){
	fmt.Println(user.Age)
}