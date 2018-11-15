package types

import (
	"testing"
	"time"
)

var (
	nonemptys = []interface{}{
		true,
		int(1),
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		uint(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),
		float32(1),
		float64(1),
		complex64(1),
		complex128(1),
		"a",
		zero{"a"},
		&zero{"a"},
		[2]interface{}{0},
		[2]zero{zero{"a"}},
		[2]*zero{&zero{"a"}},
		[]interface{}{0},
		[]zero{zero{}},
		[]*zero{&zero{}},
		time.Now(),
	}
)

func Test_IsEmpty(t *testing.T) {
	for _, v := range zeros {
		if !IsEmpty(v) {
			t.Errorf("Expected %#v to be empty.", v)
		}
	}

	for _, v := range nonzeros {
		if !IsEmpty(v) {
			t.Errorf("Expected %#v to be empty.", v)
		}
	}
}

func Test_IsEmptyWithNonemptys(t *testing.T) {
	for _, v := range nonemptys {
		if IsEmpty(v) {
			t.Errorf("Expected %#v not to be empty.", v)
		}
	}
}
