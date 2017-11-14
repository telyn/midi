package nanokontrol2

import (
	"github.com/telyn/midi/korg/korgsysex/format4"
	"github.com/telyn/midi/sysex"
)

func ParseSysEx(in format4.Message) (out sysex.SysExer, err error) {
	msgType := in.Data[0]
	switch msgType {
	case DataDumpRequestID:
		msg := DataDumpRequest{}
		err = msg.Parse(in.Data[1:])
		out = msg
	case DataDumpTwoByteResponseID:
		msg := DataDumpTwoByteResponse{}
		err = msg.Parse(in.Data[1:])
		out = msg
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
