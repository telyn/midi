package test

import (
	"reflect"
	"testing"

	"github.com/telyn/midi"
	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/korg/korgsysex"
	"github.com/telyn/midi/korg/korgsysex/format4"
	"github.com/telyn/midi/korg/korgsysex/search"
	"github.com/telyn/midi/msgs"
	"github.com/telyn/midi/sysex"
)

// this is just a single test. I want a vague idea whether the whole system works together
func TestDeepDispatch(t *testing.T) {
	noteOff := []byte{}
	nanoKontrol := []byte{}
	searchResponseMajor := uint16(0)

	dispatch := midi.Dispatcher{
		Handlers: map[msgs.Kind]msgs.Handler{
			msgs.SystemExclusive: sysex.Handler{
				sysex.Korg: korgsysex.MultiFormatHandler{
					Format4: format4.MultiDeviceHandler{
						Handlers: map[uint8]map[korgdevices.Device]format4.SingleDeviceHandler{
							format4.AllChannels: map[korgdevices.Device]format4.SingleDeviceHandler{
								korgdevices.NanoKONTROL2: format4.SingleDeviceHandlerFunc(func(data []byte) error {
									nanoKontrol = data
									return nil
								}),
							},
						},
					},
					Search: search.Handler{
						ResponseHandler: func(res search.Response) error {
							searchResponseMajor = res.Major
							return nil
						},
					},
				},
			},
			msgs.NoteOff: msgs.HandlerFunc(func(msg msgs.Message) error {
				noteOff = msg.Data
				return nil
			}),
		},
	}
	err := dispatch.HandleMessage(msgs.Message{
		Kind:    msgs.NoteOff,
		Channel: 7,
		Data:    []byte{0x03, 0x46},
	})
	if err != nil {
		t.Errorf("Error running NoteOff test: %v", err)
	}
	if !reflect.DeepEqual([]byte{0x03, 0x46}, noteOff) {
		t.Errorf("noteOff didn't get parsed correctly or smth, got %x", noteOff)
	}
	err = dispatch.HandleMessage(msgs.Message{
		Kind: msgs.SystemExclusive,
		Data: []byte{0x42, 0x43, 0x00, 0x01, 0x13, 0x00, 0x56, 0x34},
	})
	if err != nil {
		t.Errorf("Error running NanoKontrol test: %v", err)
	}
	if !reflect.DeepEqual([]byte{0x56, 0x34}, nanoKontrol) {
		t.Errorf("nanoKontrol didn't get parsed correctly or smth, got %x", nanoKontrol)
	}
	err = dispatch.HandleMessage(msgs.Message{
		Kind: msgs.SystemExclusive,
		Data: []byte{0x42, 0x50, 0x01, 0x07, 0x22, 0x13, 0x01, 0x00, 0x00, 0x11, 0x22, 0x33, 0x44},
	})
	if err != nil {
		t.Errorf("Error running SearchResponse test: %v", err)
	}
	if searchResponseMajor != 0x4433 {
		t.Errorf("searchResponseMajor didn't get parsed correctly or smth, got %x", searchResponseMajor)
	}
}
