package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func variables(){
	fmt.Println("hello how are you");
	fmt.Println("variables can be used as follows");
	var number int = 10;
	fmt.Println(number)

	var name string= "sandeep";
	fmt.Println(name)
}

func inputOutput(){
	reader := bufio.NewReader(os.Stdin);
	input, _ := reader.ReadString('\n');
	value, _ := strconv.Atoi(input)
	fmt.Println(value);

}

func main(){
	variables();
	inputOutput();
}