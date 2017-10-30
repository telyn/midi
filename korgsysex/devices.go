package korgsysex

import "encoding/binary"

type Device uint32

func (d Device) Bytes() []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(d))
	return bytes
}
