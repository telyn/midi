package search

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/telyn/midi/korg/korgdevices"
)

type Response struct {
	Channel    uint8
	EchoBackID uint8
	Family     uint16
	Member     uint16
	Major      uint16
	Minor      uint16
}

func (sr Response) Device() korgdevices.Device {
	return korgdevices.Device(sr.Family)
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

func (sr Response) String() string {
	return fmt.Sprintf("Search response (%x) from %v on channel %d\nFirmware: %d.%02d", sr.EchoBackID, sr.Device(), sr.Channel, sr.Major, sr.Minor)
}
