package search

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiniSearch( t *testing.T ) {
	mockServer := httptest.NewServer(
		http.HandlerFunc( func( w http.ResponseWriter, r *http.Request ) {
			w.WriteHeader(http.StatusOK)
			if ( strings.Contains(r.URL.String(), "stream") ) {
				w.Write([]byte(`This is a sample response that contains getstream.io in it`))
				return
			}
			w.Write([]byte(`This does not contain the desired response`))
			w.WriteHeader(http.StatusOK)
			return
		} ),
	)
	defer mockServer.Close()

	var baseSearchUrl = mockServer.URL + "/search?q="
	t.Run("Testing when the serach is made with keyword stream in it", func(t *testing.T) {
		res, _ := Search( baseSearchUrl,"stream" )
		assert.Equal(t, res, true, "Expected true")
	})

	t.Run("Testing when the serach is made without keyword stream in it", func(t *testing.T) {
		res, _ := Search(baseSearchUrl, "steam-engine" )
		assert.Equal(t, res, false, "Expected false")
	})
}