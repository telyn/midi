package korgsysex

import (
	"fmt"

	"github.com/telyn/midi/korg/korgsysex/format4"
	"github.com/telyn/midi/korg/korgsysex/search"
	"github.com/telyn/midi/sysex"
)

// Constants for the various formats korg uses for their sysex messages
const (
	Format4 byte = 0x40
	Search  byte = 0x50
)

// MultiFormatHandler implements sysex.MessageHandler
type MultiFormatHandler struct {
	Format4 format4.Handler
	Search  search.Handler
}

func (h MultiFormatHandler) Handle(msg sysex.SysEx) error {
	format := FormatOf(msg) << 4
	fmt.Printf("FORMAT: %x\n", format)
	switch format {
	case Format4:
		f4m, err := format4.Parse(ChannelOf(msg), msg.Data[1:])
		if err != nil {
			return err
		}
		return h.Format4.Handle(f4m)
	case Search:
		return h.Search.Handle(msg.Data[1:])
	}
	return fmt.Errorf("Unsupported Korg SysEx format %x", format)
}
