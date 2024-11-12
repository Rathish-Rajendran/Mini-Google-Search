package search

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SearchGoogle( keyWord string ) ( bool, error ) {
	resp, err := http.Get("https://www.google.com/search?q="+keyWord)
	if err != nil {
		return false, errors.New("HTTP request unsuccessful")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	stringBody := string(body)
	fmt.Println(strings.Contains(stringBody, "getstream.io"))
	return strings.Contains(stringBody, keyWord), nil
}