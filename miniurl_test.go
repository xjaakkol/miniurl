package miniurl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xjaakkol/miniurl"
)

func TestHashLength(t *testing.T) {
	const (
		input          = "https://github.com/xjaakkol/miniurl"
		expectedLength = 32
	)

	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}

func TestHashIsDeterministic(t *testing.T) {
	const (
		input = "https://github.com/xjaakkol/miniurl"
	)

	output1 := miniurl.Hash(input)
	output2 := miniurl.Hash(input)
	assert.Equal(t, output1, output2)
}
