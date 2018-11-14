package types

import (
	"reflect"
	"testing"
	"time"
	"unsafe"
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
	v       zero
	iv      = zero{}
	pv      *zero
	ipv     = &zero{}
	av      [2]interface{}
	iav     = [2]interface{}{}
	azv     [2]zero
	iazv    = [2]zero{}
	apzv    [2]*zero
	iapzv   = [2]*zero{}
	aiface  [2]izero
	iaiface = [2]izero{}
	sv      []interface{}
	isv     = []interface{}{}
	szv     []zero
	iszv    = []zero{}
	spzv    []*zero
	ispzv   = []*zero{}
	siface  []izero
	isiface = []izero{}
	mv      map[interface{}]interface{}
	imv     = map[interface{}]interface{}{}
	mzv     map[interface{}]zero
	imzv    = map[interface{}]zero{}
	mpzv    map[interface{}]*zero
	impzv   = map[interface{}]*zero{}
	miface  map[interface{}]izero
	imiface = map[interface{}]izero{}
	chanv   chan interface{}
	ichanv  = make(chan interface{})
	funcv   func(interface{}) interface{}
	ifuncv  = func(in interface{}) (out interface{}, err error) { return }
	timev   time.Time
	itimev  = time.Time{}
	ptimev  *time.Time
	iptimev = &time.Time{}
	iface   izero
	piface  *izero
	unpv    unsafe.Pointer
	npv     = &v
	npiv    = &iv

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
		v,
		iv,
		pv,
		av,
		iav,
		azv,
		iazv,
		apzv,
		iapzv,
		aiface,
		iaiface,
		sv,
		szv,
		spzv,
		siface,
		mv,
		mzv,
		mpzv,
		miface,
		chanv,
		funcv,
		timev,
		itimev,
		ptimev,
		iface,
		piface,
		unpv,
	}
	nonzeros = []interface{}{
		ipv,
		isv,
		iszv,
		ispzv,
		isiface,
		imv,
		imzv,
		impzv,
		imiface,
		ichanv,
		ifuncv,
		iptimev,
		npv,
		npiv,
	}
)

func Test_IsZeroWithZeros(t *testing.T) {
	// it should work with zero of types
	for _, v := range zeros {
		if !IsZero(v) {
			t.Errorf("Expected %#v to be zero.", v)
		}

		if v == nil {
			continue
		}
		if !reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface()) {
			t.Errorf("Expected %#v to be zero.", v)
		}
	}
}

func Test_IsZeroWithNonzeros(t *testing.T) {
	// it should work with nonzero of types
	for _, v := range nonzeros {
		if IsZero(v) {
			t.Errorf("Expected %#v not to be zero.", v)
		}

		if v == nil {
			continue
		}
		if reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface()) {
			t.Errorf("Expected %#v not to be zero.", v)
		}
	}
}
