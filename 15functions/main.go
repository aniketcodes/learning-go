package main

import "fmt"

func main()  {
	res:=adder(2,3)
	fmt.Println("Simple sum",res)
	proRes,message := proAdder(1,2, 3, 4, 5, 6, 7)
	fmt.Println(message,proRes)

}

func adder(a int , b int) int {
	return a + b
}

func proAdder(vals...int) (int,string) {
	var message string = "This is a pro adder function"
	total:=0
	for _, val := range vals {
		total += val
	}
	return total,message
}