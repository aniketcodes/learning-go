package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "File contents for learning go"

	fileName := "./file.txt"

	file, err := os.Create(fileName)

	checkNilError(err)

	writeLen, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println(writeLen)
	readFile(fileName)
	defer file.Close()
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Here are the file contents")
	fmt.Println(string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
