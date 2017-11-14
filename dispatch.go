package midi

import (
	"fmt"

	"github.com/telyn/midi/msgs"
)

type Dispatcher struct {
	DefaultHandler msgs.Handler
	Handlers       map[msgs.Kind]msgs.Handler
}

// HandleMessage handles the given message. Note that it is NOT thread-safe!
func (d *Dispatcher) HandleMessage(msg msgs.Message) error {
	h := d.Handlers[msg.Kind]
	if h == nil {
		if d.DefaultHandler != nil {
			d.DefaultHandler.Handle(msg)
		}
		return fmt.Errorf("There is no handler for %v", msg.Kind)
	}
	return h.Handle(msg)
}
