package korgsysex

import "fmt"

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
