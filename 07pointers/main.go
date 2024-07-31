package main

import "fmt"


func main()  {
	num:=23
	ptr := &num

	fmt.Printf("Value of num is %v , the value of pointer is %v and data stored in pointer is %v \n",num,ptr,*ptr);
	fmt.Printf("Type of pointer is %t",ptr);
	
}