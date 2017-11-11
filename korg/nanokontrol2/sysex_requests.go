package nanokontrol

import (
	"fmt"

	"github.com/telyn/nanokontrol/korgsysex"
)

const (
	SetModeRequestID  byte = 0x00
	DataDumpRequestID byte = 0x1F
)

const (
	ModeRequestFunctionID byte = 0x12
)

type DataDumpRequest struct {
	FunctionID byte
}

func (ddr *DataDumpRequest) Parse(b []byte) error {
	ddr.FunctionID = b[0]
	return nil
}

func (ddr DataDumpRequest) KorgSysEx(channel uint8) korgsysex.Message {
	return NewKorgSysEx(channel, []byte{
		DataDumpRequestID,
		ddr.FunctionID,
		0x00,
	})
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

type GetModeRequest struct{}

func (gmr GetModeRequest) KorgSysEx(channel uint8) korgsysex.Message {
	return DataDumpRequest{
		FunctionID: ModeRequestFunctionID,
	}.KorgSysEx(channel)
}

func (msg GetModeRequest) String() string {
	return "Request to get mode"
}

type SetModeRequest struct {
	Mode bool
}

func (smr SetModeRequest) KorgSysEx(channel uint8) korgsysex.Message {
	mode := byte(0)
	if smr.Mode {
		mode = 0x01
	}
	return NewKorgSysEx(channel, []byte{
		SetModeRequestID,
		0x00,
		mode,
	})
}

func (smr *SetModeRequest) Parse(b []byte) error {
	smr.Mode = b[1] == 0x01
	return nil
}

func (smr SetModeRequest) String() string {
	if smr.Mode {
		return "Request to set NanoKONTROL2 to Native Mode"
	}
	return "Request to set NanoKONTROL2 mode to MIDI mode"
}

type SearchRequest struct {
	EchoBackID byte
}

func (sr SearchRequest) String() string {
	return fmt.Sprintf("NanoKONTROL2 search request with echo ID %x", sr.EchoBackID)
}

// KorgSysEx converts the SearchRequest into a korgsysex.Message
// channel is unused in search request messages
func (sr SearchRequest) KorgSysEx(channel uint8) korgsysex.Message {
	return korgsysex.Message{
		Format: 5,
		Data: []byte{
			0x00, // 0 for request, 1 for response
			sr.EchoBackID,
		},
	}
}

func (sr *SearchRequest) Parse(b []byte) error {
	sr.EchoBackID = b[1]
}
