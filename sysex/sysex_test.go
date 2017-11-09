package sysex_test

import (
	"reflect"
	"testing"

	"github.com/telyn/midi/sysex"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Input  []byte
		Vendor sysex.VendorID
		Data   []byte
	}{
		{
			Input:  []byte{0x42, 0x13, 0x15},
			Vendor: sysex.Korg,
			Data:   []byte{0x13, 0x15},
		}, {
			Input:  []byte{0x00, 0x20, 0x20, 0x43, 0x40, 0x55},
			Vendor: sysex.DoepferMusikelektronik,
			Data:   []byte{0x43, 0x40, 0x55},
		},
	}
	for i, test := range tests {
		msg, err := sysex.Parse(test.Input)
		if err != nil {
			t.Errorf("%d error: %v", i, err)
		}
		if test.Vendor != msg.Vendor {
			t.Errorf("%d Vendor %d was not the expected %s (%d)", i, msg.Vendor, test.Vendor.String(), test.Vendor)
		}
		if !reflect.DeepEqual(test.Data, msg.Data) {
			t.Errorf("%d Wrong data\nExpected %x\nReceived %x", i, test.Data, msg.Data)
		}
	}
}
