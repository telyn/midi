package sysex

import (
	"fmt"

	"github.com/rakyll/portmidi"
)

// SysExers can convert themselves into a SysEx
type SysExer interface {
	SysEx() SysEx
}

type SysExBytesWriter interface {
	WriteSysExBytes(when portmidi.Timestamp, msg []byte) error
}

func Write(stream SysExBytesWriter, time portmidi.Timestamp, sysexer SysExer) error {
	sysex := sysexer.SysEx()
	bytes := []byte{0xF0}
	bytes = append(bytes, sysex.Vendor.Bytes()...)
	bytes = append(bytes, sysex.Data...)
	bytes = append(bytes, 0xF7)
	fmt.Printf("stream.WriteSysExBytes(%v,%x) len(bytes)=%d\n", time, bytes, len(bytes))
	return stream.WriteSysExBytes(time, bytes)

}
