package main

import "fmt"

func main()  {
	myObj:=make(map[string]string)

	myObj["js"] = "Javascript"
	myObj["rb"] = "Ruby"
	myObj["py"] = "Python"
	myObj["go"] = "Golang"
	myObj["ts"] = "Typescript"

	//keys are sorted alphabetically
	fmt.Printf("Object data %v \n",myObj);

	delete(myObj,"rb")

	fmt.Printf("Object data %v \n",myObj);

	//looping over map
	for key, value := range myObj {
		fmt.Printf("Key: %s, Value: %s \n", key, value)
	}

	//check if a key exists
	_, ok := myObj["jvm"]
	fmt.Printf("Key exists: %v \n", ok)

}