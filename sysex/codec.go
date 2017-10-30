package sysex

import "io"

// Messages can be decoded from RawMessages and can write the SysEx data
// including Vendor ID but excluding SysEx start and stop bytes to a writer.
type Message interface {
	DecodeSysEx(msg RawMessage) error
}

type Writer interface {
	WriteSysEx(wr io.Writer) error
}

// Write writes the
func Write(m Writer, wr io.Writer) (err error) {
	_, err = wr.Write([]byte{0xF0})
	if err != nil {
		return
	}
	err = m.WriteSysEx(wr)
	if err != nil {
		return
	}
	_, err = wr.Write([]byte{0xF7})
	return
}
