package sysex

type SysEx struct {
	Vendor VendorID
	Data   []byte
}

func Parse(b []byte) (sysex SysEx) {
	if b[0] == 0x00 {
		sysex.Vendor = vendorFrom3Bytes(b)
		b = b[3:]
	} else {
		sysex.Vendor = vendorFrom1Byte(b)
		b = b[1:]
	}
	sysex.Data = b
	return
}
