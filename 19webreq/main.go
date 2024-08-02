package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	URL := "https://www.aniketcodes.tech"
	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	data,err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response is")
	fmt.Println(string(data))

	defer response.Body.Close()
}
