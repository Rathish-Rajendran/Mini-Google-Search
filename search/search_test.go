package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch( t *testing.T ) {
	t.Run("Successful Run", func(t *testing.T) {
		res, _ := SearchGoogle( "xmldata" )
		assert.Equal(t, res, true, "Expected true")
	})
}