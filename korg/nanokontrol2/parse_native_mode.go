package nanokontrol

import (
	"bytes"
	"fmt"
)

// DecodeNative assumes that the 0xBF was already ripped off the message
func DecodeNative(buf *bytes.Buffer) (Message, error) {
	cid, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}
	val, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}
	nul, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}
	if nul != 0x00 {
		return nil, fmt.Errorf("Expected a nil but didn't get one sadface")
	}

	switch {
	case 0x00 <= cid && cid < 0x20:
		return ValueChanged{
			Control: ControlID(cid),
			Value:   val,
		}, nil
	case 0x20 <= cid && cid < 0x50:
		if val == 0 {
			return ButtonOff{
				Button: ControlID(cid),
			}, nil
		} else if val == 127 {
			return ButtonOn{
				Button: ControlID(cid),
			}, nil
		}
		return nil, err
	default:
		return nil, err
	}
}
