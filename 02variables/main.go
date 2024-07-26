package main

import "fmt"

func main()  {
	var username string = "Aniket"
	var isLoggedIn bool = true
	var age uint8 = 25
	variableDetails(username)
	variableDetails(isLoggedIn)
	variableDetails(age)

	//default values and some aliases
	var anotherUsername string
	var anotherIsLoggedIn bool
	var anotherAge int

	variableDetails(anotherUsername)
	variableDetails(anotherIsLoggedIn)
	variableDetails(anotherAge)

	//implicit type
	var myself = "Aniket"
	variableDetails(myself)

	//no var style
	//it can only be used inside a function
	numberOfUsers := 100
	variableDetails(numberOfUsers)

	//constants
	const PI = 3.14
	variableDetails(PI)


}

func variableDetails (variable interface {}) {
	fmt.Printf("Value of variable is : %v \n", variable)
	// %T is used to print the type of the variable
	// %v is used to print the value of the variable
	fmt.Printf("Type of variable is : %T\n", variable)
}