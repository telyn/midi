package format4

import (
	"fmt"
	"io"

	"github.com/telyn/midi/korg/korgdevices"
)

type SingleDeviceHandler interface {
	Handle(msg Message) error
}

type SingleDeviceHandlerFunc func(msg Message) error

func (handler SingleDeviceHandlerFunc) Handle(msg Message) error {
	return handler(msg)
}

type PrintfHandler struct {
	Writer io.Writer
}

func (ph PrintfHandler) Handle(msg Message) error {
	fmt.Fprintf(ph.Writer, "%v ch %v subid %v: %x\n", msg.Device, msg.Channel, msg.SubID, msg.Data)
	return nil
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
			return h.Handle(msg)
		}
	}
	if handlers, ok := dh.Handlers[AllChannels]; ok {
		if h, ok := handlers[msg.Device]; ok {
			return h.Handle(msg)
		}
	}
	if dh.Default != nil {
		return dh.Default.Handle(msg)
	}
	return fmt.Errorf("No handler could be found for device %v", msg.Device)
}
