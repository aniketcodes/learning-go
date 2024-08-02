package main

import "fmt"

func main()  {
	defer fmt.Println("Last");
	defer fmt.Println("Second last");
	fmt.Println("Hello")
	myDefer()
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}