package search

import (
	"bytes"
	"fmt"

	"github.com/telyn/midi/sysex"
)

type Request struct {
	EchoBackID byte
}

func (sr Request) SysEx() sysex.SysEx {
	return sysex.SysEx{
		Vendor: sysex.Korg,
		Data: []byte{
			0x50,
			0x00,
			sr.EchoBackID,
		},
	}
}

func ParseRequest(b []byte) (sr Request, err error) {
	buf := bytes.NewBuffer(b)
	sr.EchoBackID, err = buf.ReadByte()
	return
}

func (sr Request) String() string {
	return fmt.Sprintf("Korg Search request with echo ID %x", sr.EchoBackID)
}
