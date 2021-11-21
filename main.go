package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func additionOfTwoNumbers(){
	/// will also try if else / conditonal statements here
	reader := bufio.NewReader(os.Stdin);

	number1, _ := reader.ReadString('\n');
	num, err := strconv.ParseFloat(strings.TrimSpace(number1), 64);
	if err != nil {
		fmt.Println(err);

	}
	fmt.Println(num + 1);
}

func main(){
	variables();
	inputOutput();
	additionOfTwoNumbers();
}