package nanokontrol2

import (
	"fmt"

	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/korg/korgsysex/format4"
	"github.com/telyn/midi/sysex"
)

const (
	SetModeRequestID  byte = 0x00
	DataDumpRequestID byte = 0x1F
)

const (
	ModeRequestFunctionID byte = 0x12
)

type DataDumpRequest struct {
	Channel    byte
	FunctionID byte
}

func (ddr *DataDumpRequest) Parse(b []byte) error {
	ddr.FunctionID = b[0]
	return nil
}

func (ddr DataDumpRequest) SysEx() sysex.SysEx {
	return format4.Message{
		Channel: ddr.Channel,
		Device:  korgdevices.NanoKONTROL2,
		SubID:   0x00,
		Data: []byte{
			DataDumpRequestID,
			ddr.FunctionID,
			0x00,
		},
	}.SysEx()
}

func (ddr DataDumpRequest) String() string {
	return fmt.Sprintf("Request for data using function %v", ddr.FunctionID)
}

// Message returns the most-canonical form of this data-dump request.
// Basically, if we have a better type for the request (e.g. GetModeRequest), it uses that instead.
func (ddr DataDumpRequest) Message() Message {
	switch ddr.FunctionID {
	case ModeRequestFunctionID:
		return GetModeRequest{}
	}
	return ddr
}

type GetModeRequest struct {
	Channel uint8
}

func (gmr GetModeRequest) SysEx() sysex.SysEx {
	return DataDumpRequest{
		Channel:    gmr.Channel,
		FunctionID: ModeRequestFunctionID,
	}.SysEx()
}

func (msg GetModeRequest) String() string {
	return "Request to get mode"
}

// SetModeRequest tells the NanoKONTROL to enter/leave KORG Native Mode
type SetModeRequest struct {
	Channel    uint8
	NativeMode bool
}

func (smr SetModeRequest) SysEx() sysex.SysEx {
	mode := byte(0)
	if smr.NativeMode {
		mode = 0x01
	}
	return format4.Message{
		Channel: smr.Channel,
		Device:  korgdevices.NanoKONTROL2,
		SubID:   0x00,
		Data: []byte{
			SetModeRequestID,
			0x00,
			mode,
		},
	}.SysEx()
}

func (smr *SetModeRequest) Parse(b []byte) error {
	smr.NativeMode = b[1] == 0x01
	return nil
}

func (smr SetModeRequest) String() string {
	if smr.NativeMode {
		return "Request to set NanoKONTROL2 to Native Mode"
	}
	return "Request to set NanoKONTROL2 mode to MIDI mode"
}
