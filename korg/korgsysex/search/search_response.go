package search

import (
	"bytes"
	"encoding/binary"
)

type Response struct {
	Channel    uint8
	EchoBackID uint8
	Family     uint16
	Member     uint16
	Major      uint16
	Minor      uint16
}

func ParseResponse(b []byte) (sr Response, err error) {
	buf := bytes.NewBuffer(b)
	sr.Channel, err = buf.ReadByte()
	if err != nil {
		return
	}
	sr.EchoBackID, err = buf.ReadByte()
	if err != nil {
		return
	}
	err = binary.Read(buf, binary.LittleEndian, &sr.Family)
	if err != nil {
		return
	}
	err = binary.Read(buf, binary.LittleEndian, &sr.Member)
	if err != nil {
		return
	}
	err = binary.Read(buf, binary.LittleEndian, &sr.Minor)
	if err != nil {
		return
	}
	err = binary.Read(buf, binary.LittleEndian, &sr.Major)
	return
}
