package midi

import (
	"fmt"

	"github.com/telyn/midi/msgs"
)

type ChannelSplitHandler map[uint8]msgs.Handler

func (csh ChannelSplitHandler) Handle(msg msgs.Message) error {
	if !msg.Kind.HasChannel() {
		return fmt.Errorf("%v messages don't have channels - a ChannelSplitHandler is a mistake", msg.Kind)
	}
	if h, ok := csh[msg.Channel]; ok {
		return h.Handle(msg)
	}
	return fmt.Errorf("No handler available for channel %d", msg.Channel)
}
