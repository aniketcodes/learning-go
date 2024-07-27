package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Hi, welcome ! Please provide an input"
	println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)

	//comma ok syntax
	input ,_:= reader.ReadString('\n')
	fmt.Println("Thanks for the input", input);
	variableDetails(input);
}

func variableDetails (variable interface {}) {
	fmt.Printf("Value of variable is : %v \n", variable)
	// %T is used to print the type of the variable
	// %v is used to print the value of the variable
	fmt.Printf("Type of variable is : %T\n", variable)
}