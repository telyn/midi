package sysex

import (
	"fmt"

	"github.com/telyn/midi/msgs"
)

type MessageHandler interface {
	Handle(SysEx) error
}

type Handler map[VendorID]MessageHandler

func (h Handler) Handle(msg msgs.Message) error {
	if msg.Kind != msgs.SystemExclusive {
		return fmt.Errorf("Message wasn't a SystemExclusive message.")
	}
	sysex, err := Parse(msg.Data)
	if err != nil {
		return err
	}
	if handler, ok := h[sysex.Vendor]; ok {
		return handler.Handle(sysex)
	}
	return fmt.Errorf("No handler for sysex messages from vendor %v", sysex.Vendor)
}
