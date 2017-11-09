package korgdevices

type Device uint

func (d Device) Bytes() []byte {
	if d&0xFF0000 != 0 {
		return []byte{byte((d >> 16) & 0xFF)}
	}
	return []byte{
		0x00,
		byte((d >> 8) & 0xFF),
		byte(d & 0xFF),
	}
}

const (
	NanoKONTROL2 Device = 0x000113
)
