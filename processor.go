package midi

import (
	"fmt"

	"github.com/rakyll/portmidi"
	"github.com/telyn/midi/portbidi"
	"github.com/telyn/midi/stream"
)

type Processor struct {
	bidi     *portbidi.Stream
	stream   stream.Stream
	dispatch Dispatcher
}

func NewProcessor(bidi *portbidi.Stream, dispatch Dispatcher) (p *Processor) {
	return &Processor{
		bidi:     bidi,
		dispatch: dispatch,
	}
}
func NewProcessorWithStream(bidi *portbidi.Stream, dispatch Dispatcher, stream stream.Stream) (p *Processor) {
	return &Processor{
		bidi:     bidi,
		dispatch: dispatch,
		stream:   stream,
	}
}

func (p *Processor) Close() error {
	err := p.bidi.Close()
	p.bidi = nil
	return err
}

// Process runs a process step. This reads a bundle of bytes from the MIDI input channel,
// calling HandleMessage on the Dispatcher whenever relevant
func (p *Processor) Process() error {
	bytes, err := p.bidi.ReadSysExBytes(1024)
	if err != nil {
		return fmt.Errorf("Couldn't read from stream: %s\n", err)
	}
	messages := p.stream.ConsumeBytes(bytes)
	for _, msg := range messages {
		err = p.dispatch.HandleMessage(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// Write implements the usual golang Write interface bc FUCK portmidi's stupid []Events crap
func (p *Processor) Write(data []byte) (err error) {
	return p.WriteSysExBytes(portmidi.Time(), data)
}

// WriteSysExBytes writes the bytes to the midi stream at the given time.
func (p *Processor) WriteSysExBytes(time portmidi.Timestamp, data []byte) error {
	return p.bidi.WriteSysExBytes(time, data)
}
