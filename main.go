package main

import (
	"fmt"

	search "github.com/Rathish-Rajendran/mini-search/search"
)

func main() {
	baseSearchUrl := "https://www.google.com/search?q="
	res, err := search.Search(baseSearchUrl,"getstream.io")
	if err != nil {
		fmt.Println("ERROR...!!!")
	}
	fmt.Println("getstream.io present in google search:", res)
}