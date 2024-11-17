package search

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func Search( baseSearchUrl, keyWord string ) ( bool, error ) {
	resp, err := http.Get(baseSearchUrl+keyWord)
	if err != nil {
		return false, errors.New("HTTP request unsuccessful")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	stringBody := string(body)
	return strings.Contains(stringBody, "getstream.io"), nil
}