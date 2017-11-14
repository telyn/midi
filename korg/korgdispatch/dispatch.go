package korgdispatch

import (
	"github.com/telyn/midi"
	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/korg/korgsysex"
	"github.com/telyn/midi/korg/korgsysex/format4"
	"github.com/telyn/midi/korg/korgsysex/search"
	"github.com/telyn/midi/msgs"
	"github.com/telyn/midi/sysex"
)

type EasyDispatch struct {
	Device                korgdevices.Device
	DeviceHandler         format4.SingleDeviceHandlerFunc
	MIDIMessageHandlers   map[msgs.Kind]msgs.Handler
	SearchResponseHandler func(search.Response) error
}

func (ed EasyDispatch) Dispatcher() midi.Dispatcher {
	midiHandlers := ed.MIDIMessageHandlers
	if midiHandlers == nil {
		midiHandlers = map[msgs.Kind]msgs.Handler{}
	}
	midiHandlers[msgs.SystemExclusive] = sysex.Handler{
		sysex.Korg: korgsysex.MultiFormatHandler{
			Format4: format4.MultiDeviceHandler{
				Handlers: map[uint8]map[korgdevices.Device]format4.SingleDeviceHandler{
					format4.AllChannels: map[korgdevices.Device]format4.SingleDeviceHandler{
						ed.Device: ed.DeviceHandler,
					},
				},
			},
			Search: search.Handler{
				ResponseHandler: ed.SearchResponseHandler,
			},
		},
	}
	return midi.Dispatcher{
		Handlers: midiHandlers,
	}
}
