package sysex

import (
	"reflect"
	"testing"
)

func TestVendorIDBytes(t *testing.T) {
	tests := []struct {
		vendor VendorID
		bytes  []byte
	}{
		{
			vendor: 0x001234,
			bytes:  []byte{0x00, 0x12, 0x34},
		}, {
			vendor: 0x7D0000,
			bytes:  []byte{0x7D},
		},
	}
	for _, test := range tests {
		if !reflect.DeepEqual(test.bytes, test.vendor.Bytes()) {
			t.Errorf("%x != %x", test.bytes, test.vendor.Bytes())
		}
	}
}
