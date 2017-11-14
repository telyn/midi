package devicesearch

import (
	"github.com/rakyll/portmidi"
	"github.com/telyn/midi"
	"github.com/telyn/midi/portbidi"
	"github.com/telyn/midi/stream"
)

type SearchResult struct {
	In      *portmidi.Stream
	Out     *portmidi.Stream
	Stream  stream.Stream
	Channel byte
}

func (res SearchResult) Processor(dispatch midi.Dispatcher) (p *midi.Processor) {
	return midi.NewProcessorWithStream(&portbidi.Stream{
		In:  res.In,
		Out: res.Out,
	}, dispatch, res.Stream)
}
