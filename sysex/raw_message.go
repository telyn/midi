package sysex

import (
	"bufio"
	"fmt"
	"io"
)

const SysExMessage byte = 0xF0
const SysExStop byte = 0xF7

type RawMessage struct {
	VendorID byte
	Bytes    []byte
}

func (msg *RawMessage) DecodeSysEx(msg2 RawMessage) error {
	msg.VendorID = msg2.VendorID
	msg.Bytes = msg2.Bytes
	return nil
}

func (msg RawMessage) WriteSysEx(wr io.Writer) error {
	_, err := wr.Write(append([]byte{
		msg.VendorID,
	}, msg.Bytes...))
	return err
}

// DecodeMIDI takes a reader and populates the raw message.
// The reader is expected to have only just read the 0xF0 which indicates a SysEx message
func (msg *RawMessage) DecodeMIDI(r io.Reader) (err error) {
	buf := bufio.NewReader(r)
	msg.VendorID, err = buf.ReadByte()
	if err != nil {
		return
	}
	fmt.Println(msg.VendorID)
	msg.Bytes, err = buf.ReadBytes(SysExStop)
	return
}
