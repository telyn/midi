package stream

import (
	"reflect"
	"testing"

	"github.com/telyn/midi/msgs"
)

func TestConsumeByte(t *testing.T) {
	invalidMessage := msgs.Message{Kind: msgs.DataByte}
	emptyMessage := msgs.Message{Kind: msgs.DataByte, Data: []byte{}, Channel: 0xFF}
	tests := []struct {
		Before          Stream
		After           Stream
		Byte            byte
		ExpectedMessage msgs.Message
		Ready           bool
	}{
		{ // 0 - starting from an invalid message: does a data byte do nothing?
			Before: Stream{
				cur: invalidMessage,
				len: -2,
			},
			After: Stream{
				cur: invalidMessage,
				len: -2,
			},
			Byte:            0x04,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 1 - starting from an invalid message, does a status byte set a new status?
			Before: Stream{
				cur: invalidMessage,
				len: -2,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{},
				},
				len: 2,
			},
			Byte:            0x87,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 2 - starting from a status byte with no data, does a data byte add to the buffer?
			Before: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{},
				},
				len: 2,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{0x13},
				},
				len: 2,
			},
			Byte:            0x13,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 3 - starting from a NoteOff message with a data byte, does a second finish the message? And does running-status work?
			Before: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{0x13},
				},
				len: 2,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{},
				},
				len: 2,
			},
			Byte: 0x52,
			ExpectedMessage: msgs.Message{
				Kind:    msgs.NoteOff,
				Channel: 7,
				Data:    []byte{0x13, 0x52},
			},
			Ready: true,
		}, { // 4 - starting from a NoteOff message with one data byte, does a new status byte override the first?
			Before: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{0x13},
				},
				len: 2,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.ProgramChange,
					Channel: 4,
					Data:    []byte{},
				},
				len: 1,
			},
			Byte:            0xC4,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 5 - starting from a NoteOff message with one data byte does a real-time byte return immediately?
			Before: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{0x13},
				},
				len: 2,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.NoteOff,
					Channel: 7,
					Data:    []byte{0x13},
				},
				len: 2,
			},
			Byte: 0xFF,
			ExpectedMessage: msgs.Message{
				Kind: msgs.RealTimeSystemReset,
			},
			Ready: true,
		}, { // 6 - starting from an invalid message, does the sysex start byte start a sysex message?
			Before: Stream{
				cur: invalidMessage,
				len: -2,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.SystemExclusive,
					Data:    []byte{},
					Channel: 0xFF,
				},
				len: -1,
			},
			Byte:            0xF0,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 7 - starting from an empty sysex message, does a byte get added?
			Before: Stream{
				cur: msgs.Message{
					Kind:    msgs.SystemExclusive,
					Data:    []byte{},
					Channel: 0xFF,
				},
				len: -1,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.SystemExclusive,
					Data:    []byte{0x4F},
					Channel: 0xFF,
				},
				len: -1,
			},
			Byte:            0x4F,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 8 - starting from a sysex with a byte, does a byte get added?
			Before: Stream{
				cur: msgs.Message{
					Kind:    msgs.SystemExclusive,
					Data:    []byte{0x4F},
					Channel: 0xFF,
				},
				len: -1,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.SystemExclusive,
					Data:    []byte{0x4F, 0x5D},
					Channel: 0xFF,
				},
				len: -1,
			},
			Byte:            0x5D,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 9 - starting from a sysex with two bytes, does a byte get added?
			Before: Stream{
				cur: msgs.Message{
					Kind:    msgs.SystemExclusive,
					Data:    []byte{0x4F, 0x5D},
					Channel: 0xFF,
				},
				len: -1,
			},
			After: Stream{
				cur: msgs.Message{
					Kind:    msgs.SystemExclusive,
					Data:    []byte{0x4F, 0x5D, 0x3B},
					Channel: 0xFF,
				},
				len: -1,
			},
			Byte:            0x3B,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 10 - starting from a sysex with ten bytes, does a byte get added?
			Before: Stream{
				cur: msgs.Message{
					Kind: msgs.SystemExclusive,
					Data: []byte{0x4F, 0x5D, 0x00, 0x23, 0x45,
						0x7D, 0x12, 0x33, 0x00, 0x00},
					Channel: 0xFF,
				},
				len: -1,
			},
			After: Stream{
				cur: msgs.Message{
					Kind: msgs.SystemExclusive,
					Data: []byte{0x4F, 0x5D, 0x00, 0x23, 0x45,
						0x7D, 0x12, 0x33, 0x00, 0x00, 0x5A},
					Channel: 0xFF,
				},
				len: -1,
			},
			Byte:            0x5A,
			ExpectedMessage: invalidMessage,
			Ready:           false,
		}, { // 11 - starting from a sysex with ten bytes, does an EOX return the message and reset the status byte
			Before: Stream{
				cur: msgs.Message{
					Kind: msgs.SystemExclusive,
					Data: []byte{0x4F, 0x5D, 0x00, 0x23, 0x45,
						0x9D, 0x12, 0x33, 0x00, 0x00},
					Channel: 0xFF,
				},
				len: -1,
			},
			After: Stream{
				cur: emptyMessage,
				len: -2,
			},
			Byte: 0xF7,
			ExpectedMessage: msgs.Message{
				Kind: msgs.SystemExclusive,
				Data: []byte{0x4F, 0x5D, 0x00, 0x23, 0x45,
					0x9D, 0x12, 0x33, 0x00, 0x00},
				Channel: 0xFF,
			},
			Ready: true,
		}, { // 12 - starting from a sysex with ten bytes, does a real-time message return without interfering?
			Before: Stream{
				cur: msgs.Message{
					Kind: msgs.SystemExclusive,
					Data: []byte{0x4F, 0x5D, 0x00, 0x23, 0x45,
						0x9D, 0x12, 0x33, 0x00, 0x00},
					Channel: 0xFF,
				},
				len: -1,
			},
			After: Stream{
				cur: msgs.Message{
					Kind: msgs.SystemExclusive,
					Data: []byte{0x4F, 0x5D, 0x00, 0x23, 0x45,
						0x9D, 0x12, 0x33, 0x00, 0x00},
					Channel: 0xFF,
				},
				len: -1,
			},
			Byte: 0xF8,
			ExpectedMessage: msgs.Message{
				Kind:    msgs.RealTimeClock,
				Data:    nil,
				Channel: 0x0,
			},
			Ready: true,
		},
	}

	for i, test := range tests {
		stream := test.Before
		msg, ok := stream.ConsumeByte(test.Byte)
		if !reflect.DeepEqual(test.After, stream) {
			t.Errorf("%d after-consumption stream. Expecting: %#v, got %#v", i, test.After, stream)
		}
		if !reflect.DeepEqual(test.ExpectedMessage, msg) {
			t.Errorf("%d returned wrong message. Expecting: %#v, got %#v", i, test.ExpectedMessage, msg)
		}
		if ok != test.Ready {
			t.Errorf("%d message readiness was wrong. expected %b, got %b", i, test.Ready, ok)
		}

	}
}
