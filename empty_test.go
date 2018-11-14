package types

import (
	"fmt"
	"testing"

	"github.com/golib/assert"
)

func Test_IsEmpty(t *testing.T) {
	for _, v := range zeros {
		assert.True(t, IsEmpty(v), fmt.Sprintf("Expected %#v to be empty.", v))
	}

	for _, v := range nonzeros {
		assert.True(t, IsEmpty(v), fmt.Sprintf("Expected %#v to be empty.", v))
	}
}
