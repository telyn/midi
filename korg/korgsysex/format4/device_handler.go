package format4

import (
	"fmt"

	"github.com/telyn/midi/korg/korgdevices"
)

type SingleDeviceHandler interface {
	Handle(data []byte) error
}

type SingleDeviceHandlerFunc func(data []byte) error

func (handler SingleDeviceHandlerFunc) Handle(data []byte) error {
	return handler(data)
}

// The uint8 to use with MultiDeviceHandlers.Handlers if you want your handler to listen on all channels
const AllChannels uint8 = 0xFF

// MultiDeviceHandler dispatches messages to other handlers depending on what device they are for
type MultiDeviceHandler struct {
	// Channel is a number from 0 to F to determine which
	Default  SingleDeviceHandler
	Handlers map[uint8]map[korgdevices.Device]SingleDeviceHandler
}

// Handle handles messages for whatever devices are on the channel it's attached to, throwing away messages for devices it doesn't.
func (dh MultiDeviceHandler) Handle(msg Message) error {

	if handlers, ok := dh.Handlers[msg.Channel]; ok {
		if h, ok := handlers[msg.Device]; ok {
			return h.Handle(msg.Data)
		}
	}
	if handlers, ok := dh.Handlers[AllChannels]; ok {
		if h, ok := handlers[msg.Device]; ok {
			return h.Handle(msg.Data)
		}
	}
	if dh.Default != nil {
		return dh.Default.Handle(msg.Data)
	}
	return fmt.Errorf("No handler could be found for device %v", msg.Device)
}
