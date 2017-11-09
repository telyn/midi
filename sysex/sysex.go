package sysex

import (
	"bytes"
	"encoding/binary"
)

type SysEx struct {
	Vendor VendorID
	Data   []byte
}

func Parse(b []byte) (sysex SysEx, err error) {
	buf := bytes.NewBuffer(b)
	vendor, err := buf.ReadByte()
	if err != nil {
		return
	}
	if vendor == 0x00 {
		bigVendor := uint16(0)
		err = binary.Read(buf, binary.LittleEndian, &bigVendor)
		if err != nil {
			return
		}
		sysex.Vendor = VendorID(bigVendor)
	} else {
		sysex.Vendor = VendorID(uint(vendor) << 16)
	}
	sysex.Data = buf.Bytes()
	return
}
