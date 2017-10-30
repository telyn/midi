package portbidi

import (
	"fmt"

	"github.com/rakyll/portmidi"
)

type Err struct {
	In  error
	Out error
}

func (e Err) Error() string {
	return fmt.Sprintf("input err: %s & output err: %s", e.In, e.Out)
}

type Stream struct {
	In  *portmidi.Stream
	Out *portmidi.Stream
}

func (s *Stream) Abort() error {
	err := Err{
		In:  s.In.Abort(),
		Out: s.Out.Abort(),
	}
	if err.In != nil || err.Out != nil {
		return err
	}
	return nil
}
func (s *Stream) Close() error {
	err := Err{
		In:  s.In.Abort(),
		Out: s.Out.Abort(),
	}
	if err.In != nil || err.Out != nil {
		return err
	}
	return nil
}

func (s *Stream) Listen() <-chan portmidi.Event {
	return s.In.Listen()
}
func (s *Stream) Poll() (bool, error) {
	return s.In.Poll()
}
func (s *Stream) Read(max int) (events []portmidi.Event, err error) {
	return s.In.Read(max)
}
func (s *Stream) ReadSysExBytes(max int) ([]byte, error) {
	return s.In.ReadSysExBytes(max)
}
func (s *Stream) SetChannelMask(mask int) error {
	return s.In.SetChannelMask(mask)
}
func (s *Stream) Write(events []portmidi.Event) error {
	return s.Out.Write(events)
}
func (s *Stream) WriteShort(status int64, data1 int64, data2 int64) error {
	return s.Out.WriteShort(status, data1, data2)
}
func (s *Stream) WriteSysEx(when portmidi.Timestamp, msg string) error {
	return s.Out.WriteSysEx(when, msg)
}
func (s *Stream) WriteSysExBytes(when portmidi.Timestamp, msg []byte) error {
	return s.Out.WriteSysExBytes(when, msg)
}

func New(in portmidi.DeviceID, out portmidi.DeviceID, bufferSize int64, latency int64) (stream Stream, err error) {
	bErr := Err{}

	stream.In, bErr.In = portmidi.NewInputStream(in, bufferSize)
	stream.Out, bErr.Out = portmidi.NewOutputStream(out, bufferSize, latency)
	if bErr.In != nil || bErr.Out != nil {
		return stream, bErr
	}
	return
}
