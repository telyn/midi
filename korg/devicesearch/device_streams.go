package devicesearch

import (
	"github.com/rakyll/portmidi"
	"github.com/telyn/midi"
	"github.com/telyn/midi/stream"
)

type deviceStream struct {
	stream     *portmidi.Stream
	info       *portmidi.DeviceInfo
	msgStream  stream.Stream
	dispatcher midi.Dispatcher
}

type deviceStreams []deviceStream

// Closes all the streams except notMe
func (d deviceStreams) CloseAllButNot(notMe *portmidi.Stream) error {
	errs := []error{}
	for _, device := range d {
		if device.stream == notMe {
			continue
		}
		err := device.stream.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 1 {
		return errs[0]
	}
	return nil
}
