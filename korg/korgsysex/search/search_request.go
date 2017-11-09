package search

import (
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

func (sr Request) String() string {
	return fmt.Sprintf("Korg Search request with echo ID %x", sr.EchoBackID)
}
