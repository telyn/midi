package format4

import (
	"bytes"
	"encoding/binary"

	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/sysex"
)

// Message is a type of message which
// is used by (at least) the NanoKONTROL
//
// From what I can tell, it's always this, followed by the actual command to sent to/from the device
//
//     0x4g, <Project> <SubID>
//
// AFAICT <Project> is always either 1-byte or 3-bytes, the same
// as sysex vendors in MIDI.
//
// <SubID> always seems to be 1-byte and can be 00
//
// And I *think* <Project> == Family ID in the search response
// and more uncertainly, <SubID> == Member ID in the search response
//
// I intend to only match project against family id.
type Message struct {
	Channel byte
	Device  korgdevices.Device
	SubID   byte
	Data    []byte
}

// Parse parses the format-4 message described by the channel and data into a Message,
// setting all the fields
func Parse(channel uint8, data []byte) (message Message, err error) {
	buf := bytes.NewBuffer(data)
	message.Channel = channel & 0xF

	firstProjectByte, err := buf.ReadByte()
	if err != nil {
		return
	}
	message.Device = korgdevices.Device(firstProjectByte)

	// decode 3-byte projects
	if message.Device == 0x00 {
		var proj uint16
		err = binary.Read(buf, binary.BigEndian, &proj)
		if err != nil {
			return
		}
		message.Device = korgdevices.Device(proj)
	}

	message.SubID, err = buf.ReadByte()
	message.Data = buf.Bytes()
	return
}

func (m Message) SysEx() sysex.SysEx {
	msg := sysex.SysEx{
		Vendor: sysex.Korg,
		Data: []byte{
			0x40 | m.Channel,
		},
	}
	msg.Data = append(msg.Data, m.Device.Bytes()...)
	msg.Data = append(msg.Data, m.SubID)
	msg.Data = append(msg.Data, m.Data...)
	return msg
}
