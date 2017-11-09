package stream

import (
	"fmt"

	"github.com/telyn/midi/msgs"
)

type Stream struct {
	cur msgs.Message
	// number of data bytes for the current message
	len int
}

func (s *Stream) SetStatus(b byte) {
	kind := msgs.KindOf(b)

	s.len = kind.Bytes()
	expectedBytes := s.len
	if s.len == -1 {
		expectedBytes = 16
	}
	if s.len == -2 {
		expectedBytes = 0
	}

	s.cur = msgs.Message{
		Kind:    kind,
		Channel: msgs.ChannelOf(b),
		Data:    make([]byte, 0, expectedBytes),
	}
}

func (s *Stream) ConsumeByte(b byte) (message msgs.Message, messageReady bool) {
	k := msgs.KindOf(b)
	fmt.Printf("byte %x kind %v status %x\n", b, k, k.Byte())
	//  process realtime messages first, so that they can interrupt other messages
	if k.RealTime() {
		return msgs.Message{Kind: k}, true
	}

	// if we're assembling a sysex message,
	if s.cur.Kind == msgs.SystemExclusive {
		// append data bytes to it,
		if k == msgs.DataByte {
			s.cur.Data = append(s.cur.Data, b)
			return msgs.Message{Kind: msgs.DataByte}, false
		} else {
			// if a status byte comes in, consider it an EOX and unset status
			message = s.cur
			s.SetStatus(0x00)
			return message, true
		}
	}
	// not assembling a sysex message
	// if a data byte comes in
	if k == msgs.DataByte {
		// if we're not desiring any data bytes then return an invalid message
		if s.len < 1 {
			return msgs.Message{Kind: msgs.DataByte}, false
		}
		fmt.Println("s.len >= 1")
		// we do desire data, so append it to the current message
		s.cur.Data = append(s.cur.Data, b)
		fmt.Printf("appended, s.cur.Data: %x\n", s.cur.Data)
		fmt.Printf("len: %d, len(s.cur.Data): %d\n", s.len, len(s.cur.Data))
		// if we fill the data buffer, return it and start a new message with the current status
		if len(s.cur.Data) == s.len {
			message = s.cur
			s.SetStatus(s.cur.Status())
			return message, true
		}
		return msgs.Message{Kind: msgs.DataByte}, false
	}
	// if we're assembling a non-sysex message and a status byte comes in, drop the message and set status
	s.SetStatus(b)
	return message, false

}
