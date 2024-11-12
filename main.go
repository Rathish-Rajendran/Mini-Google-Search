package main

import (
	"fmt"

	search "github.com/Rathish-Rajendran/mini-search/search"
)

func main() {
	res, err := search.SearchGoogle("getstreamio")
	if err != nil {
		fmt.Println("ERROR...!!!")
	}
	fmt.Println("getstream.io present in google search:", res)
}