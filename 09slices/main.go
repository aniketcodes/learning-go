package main

import (
	"fmt"
	"sort"
)

func main() {
	mySlice := make([]int, 4)
	fmt.Printf("Values in slice are %v \n", mySlice)

	mySlice[1] = -2

	//append to slice
	mySlice = append(mySlice, 5)
	fmt.Printf("Values in slice are %v \n", mySlice)

	//append multiple values
	mySlice = append(mySlice, 6, 7, 8)
	fmt.Printf("Values in slice are %v \n", mySlice)

	//append a slice
	mySlice = append(mySlice, []int{9, 10, 11}...)
	fmt.Printf("Values in slice are %v \n", mySlice)

	fmt.Println("Is myslice sorted", sort.IntsAreSorted(mySlice))

	//sorting a slice
	sort.Ints(mySlice)

	fmt.Printf("Sored values %v \n", mySlice)
	fmt.Println("Is myslice sorted", sort.IntsAreSorted(mySlice))

}
