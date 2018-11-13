package types

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/golib/assert"
)

type (
	zero struct {
		v interface{}
	}

	izero interface {
		f()
	}
)

var (
	v      zero
	pv     *zero
	iv     = zero{}
	npv    = &v
	npiv   = &iv
	av     [2]interface{}
	azv    []zero
	apzv   []*zero
	aiface []izero
	sv     []interface{}
	szv    []zero
	spzv   []*zero
	siface []izero
	mv     map[interface{}]interface{}
	mzv    map[interface{}]zero
	mpzv   map[interface{}]*zero
	miface map[interface{}]izero
	iface  izero
	piface *izero
	chanv  chan interface{}
	funcv  func(interface{}) interface{}
	unpv   unsafe.Pointer

	zeros = []interface{}{
		nil,
		false,
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		float32(0),
		float64(0),
		complex64(0),
		complex128(0),
		"",
		v, pv, iv,
		av, azv, apzv, aiface,
		mv, mzv, mpzv, miface,
		iface, piface,
		chanv,
		funcv,
		unpv,
	}
	nonzeros = []interface{}{
		npv, npiv,
	}
)

func Test_IsZero(t *testing.T) {
	for _, v := range zeros {
		assert.True(t, IsZero(v), fmt.Sprintf("Expected %#v to be zero.", v))
	}

	for _, v := range nonzeros {
		assert.False(t, IsZero(v), fmt.Sprintf("Expected %#v not to be zero.", v))
	}
}
