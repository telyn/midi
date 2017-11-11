package nanokontrol

import (
	"io"

	"github.com/telyn/nanokontrol/korgsysex"
	"github.com/telyn/nanokontrol/sysex"
)

type SysExMessage interface {
	KorgSysEx(channel uint8) korgsysex.Message
}

func ParseSysEx(in korgsysex.Message) (out Message, err error) {
	msgType := in.Data[0]
	switch msgType {
	case DataDumpRequestID:
		msg := DataDumpRequest{}
		err = msg.Parse(in.Data[1:])
		out = msg.Message()
	case DataDumpTwoByteResponseID:
		msg := DataDumpTwoByteResponse{}
		err = msg.Parse(in.Data[1:])
		out = msg.Message()
	case SetModeRequestID:
		msg := SetModeRequest{}
		err = msg.Parse(in.Data[1:])
		out = msg
	case SetModeResponseID:
		msg := SetModeResponse{}
		err = msg.Parse(in.Data[1:])
		out = msg
	}
	return
}

func WriteSysEx(msg SysExMessage, channel uint8, wr io.Writer) error {
	return sysex.Write(msg.KorgSysEx(channel), wr)
}

func NewKorgSysEx(channel uint8, data []byte) korgsysex.Message {
	return korgsysex.Message{
		Format:  4,
		Channel: channel,
		Device:  DeviceID,
		Data:    data,
	}

}
