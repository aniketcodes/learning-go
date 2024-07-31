package main

import "fmt"

func main()  {
	 arrOne:= [4]string{"a", "b", "c", "d"}

	 var arrTwo [4] string

	 arrTwo[0] = "Aniket"
	 arrTwo[1] = "niket"
	 arrTwo[3] = "ket"

	 fmt.Printf("arrOne %v has len %v \n", arrOne , len(arrOne))
	 fmt.Printf("arrTwo %v has len %v\n", arrTwo,len(arrTwo))
}