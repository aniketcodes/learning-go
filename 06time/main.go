package main

import (
	"fmt"
	"time"
)

func main()  {
	//currnet tine
	var now = time.Now().Format("01-02-2006 15:04:05 Monday")
	fmt.Printf("now %v \n",now)

	//created data
	createdDatae:= time.Date(1999,time.April,13,5,0,0,0,time.UTC)
	fmt.Printf("createdDatae %v \n", createdDatae)

}