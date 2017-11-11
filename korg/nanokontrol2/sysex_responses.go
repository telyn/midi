package nanokontrol

import (
	"fmt"

	"github.com/telyn/nanokontrol/korgsysex"
)

const (
	SetModeResponseID         byte = 0x40
	DataDumpTwoByteResponseID byte = 0x5F
	DataDumpResponseID        byte = 0x7F
)

const (
	ModeResponseFunctionID byte = 0x42
)

type DataDumpTwoByteResponse struct {
	FunctionID byte
	Data       byte
}

func (ddr *DataDumpTwoByteResponse) Parse(b []byte) error {
	ddr.FunctionID = b[0]
	ddr.Data = b[1]
	return nil
}

func (ddr DataDumpTwoByteResponse) KorgSysEx(channel uint8) korgsysex.Message {
	return NewKorgSysEx(channel, []byte{
		DataDumpResponseID,
		ddr.FunctionID,
		ddr.Data,
	})
}
func (ddr DataDumpTwoByteResponse) String() string {
	return fmt.Sprintf("Data response: Function %v: %v", ddr.FunctionID, ddr.Data)
}

// Message returns the most-canonical form of this data-dump response.
// Basically, if we have a better type for the request (e.g. GetModeResponse), it uses that instead.
func (ddr DataDumpTwoByteResponse) Message() Message {
	switch ddr.FunctionID {
	case ModeResponseFunctionID:
		return GetModeResponse{
			Mode: (ddr.Data == 0x03),
		}
	}
	return ddr
}

type GetModeResponse struct {
	Mode bool
}

func (gmr GetModeResponse) KorgSysEx(channel uint8) korgsysex.Message {
	data := byte(0x2)
	if gmr.Mode {
		data = 0x3
	}
	return DataDumpTwoByteResponse{
		FunctionID: ModeResponseFunctionID,
		Data:       data,
	}.KorgSysEx(channel)
}

func (gmr GetModeResponse) String() string {
	if gmr.Mode {
		return "NanoKONTROL2 mode status: Native Mode"
	}
	return "NanoKONTROL2 mode status: MIDI mode"
}

type SetModeResponse struct {
	Mode bool
}

func (gmr SetModeResponse) KorgSysEx(channel uint8) korgsysex.Message {
	data := byte(0x2)
	if gmr.Mode {
		data = 0x3
	}
	return NewKorgSysEx(channel, []byte{
		SetModeResponseID,
		data,
	})
}

func (gmr *SetModeResponse) Parse(bytes []byte) error {
	gmr.Mode = (bytes[0] == 0x03)
	return nil
}

func (gmr SetModeResponse) String() string {
	if gmr.Mode {
		return "NanoKONTROL2 mode set to Native Mode"
	}
	return "NanoKONTROL2 mode set to MIDI mode"
}
