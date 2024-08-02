package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://www.aniketcodes.tech/blogs/golang-tut-101?country=IN&mode=dark&subscribed=true"

func main()  {
	urlData,err := url.Parse(URL)

	if(err!=nil){
		panic(err)
	}

	path:=urlData.Path
	query:=urlData.Query();
	fmt.Println(path);
	fmt.Println(query)

	for _,val := range query {
		fmt.Printf("%T\n",val)
	}


}