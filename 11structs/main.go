package main

import "fmt"

func main()  {
	type User struct {
		Name string
		Email string
		Age int
		IsEmailVerified bool
	}

	var aniket User = User{
		Name: "Aniket",
		Email: "aniket@go.dev",
		Age: 25,
		IsEmailVerified: true,
	}

	fmt.Printf("Aniket Details are: %v\n", aniket)
	fmt.Printf("Aniket Details are: %+v\n", aniket)
	fmt.Printf("Name is %v and email is %v",aniket.Name, aniket.Email)
}