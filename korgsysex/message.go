package korgsysex

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/telyn/nanokontrol/sysex"
)

const VendorID byte = 0x42

type Message struct {
	// Format can be (at least) 3 or 4.
	// only 4 is supported by korgsysex atm.
	Format  uint8
	Channel uint8
	Device  Device
	Data    []byte
}

func (km *Message) decodeSysEx4(msg sysex.RawMessage) error {
	if len(msg.Bytes) < 6 {
		return fmt.Errorf("message not long enough")
	}
	fmt.Println(hex.Dump(msg.Bytes))
	km.Channel = msg.Bytes[0] & 0xF
	km.Device = Device(binary.BigEndian.Uint32(msg.Bytes[1:5]))
	fmt.Printf("%x\n", km.Device)
	km.Data = msg.Bytes[5:]
	return nil
}

func (km *Message) decodeSysEx5(msg sysex.RawMessage) error {

}

func (km *Message) DecodeSysEx(msg sysex.RawMessage) error {
	km.Format = (msg.Bytes[0] & 0xF0) >> 4
	if km.Format == 4 {
		return km.decodeSysEx4(msg)
	}
	return fmt.Errorf("korgsysex package does not understand Korg SysEx messages with format %x", km.Format)
}

func (km Message) writeSysEx4(wr io.Writer) (err error) {
	_, err = wr.Write([]byte{
		VendorID,
		km.Format<<4 | km.Channel&0xF,
	})
	if err != nil {
		return
	}

	_, err = wr.Write(km.Device.Bytes())
	if err != nil {
		return
	}
	_, err = wr.Write(km.Data)
	return
}

func (km Message) writeSysEx5(wr io.Writer) (err error) {
	_, err = wr.Write(append([]byte{km.Format},
		km.Data...))
	return
}

func (km Message) WriteSysEx(wr io.Writer) error {
	switch km.Format {
	case 4:
		return km.writeSysEx4(wr)
	case 5:
		return km.writeSysEx5(wr)
	}
}
