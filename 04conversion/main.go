package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("Welocme to my pizza App")
	println("Please rate between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Invalid input")
		panic(err)
	}
	numInput = numInput + 1
	fmt.Printf("You rated %v out of 5 ", numInput)
}
