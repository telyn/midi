package devicesearch

import (
	"fmt"
	"time"

	"github.com/rakyll/portmidi"
	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/korg/korgdispatch"
	"github.com/telyn/midi/korg/korgsysex/search"
	"github.com/telyn/midi/sysex"
)

type searcher struct {
	in  deviceStreams
	out deviceStreams

	seeking korgdevices.Device

	result *SearchResult
}

func (s *searcher) searchHandler(inputStream int) func(search.Response) error {
	return func(res search.Response) error {
		fmt.Println(res)
		fmt.Printf("device: %#x, seeking: %#x\n", res.Device(), s.seeking)
		if s.seeking != res.Device() {
			fmt.Println("not equal, carrying on")
			return nil
		}
		s.result = &SearchResult{
			In:      s.in[inputStream].stream,
			Out:     s.out[res.EchoBackID].stream,
			Stream:  s.in[inputStream].msgStream,
			Channel: res.Channel,
		}
		fmt.Printf("%p - %#v\n", s, s.result)
		return nil
	}
}

func (s *searcher) readLoop() error {
	for s.result == nil {
		for _, inStream := range s.in {
			bytes, err := inStream.stream.ReadSysExBytes(32)
			if err != nil {
				fmt.Printf("Couldn't read from stream %d\n")
				continue
			}

			messages := inStream.msgStream.ConsumeBytes(bytes)
			for _, msg := range messages {
				inStream.dispatcher.HandleMessage(msg)
				fmt.Printf("%p - %#v\n", s, s.result)
			}
		}
		time.Sleep(10 * time.Microsecond)
	}
	return nil
}

func (s *searcher) search() (*SearchResult, error) {
	defer func() {
		if s.result == nil {
			s.in.CloseAllButNot(nil)
			s.out.CloseAllButNot(nil)
		} else {
			s.in.CloseAllButNot(s.result.In)
			s.out.CloseAllButNot(s.result.Out)
		}
	}()
	for i, out := range s.out {
		err := sysex.Write(out.stream, portmidi.Time(), search.Request{
			EchoBackID: byte(i),
		})
		if err != nil {
			return s.result, fmt.Errorf("Couldn't write to %s %s - %s", out.info.Interface, out.info.Name, err)
		}
	}

	err := s.readLoop()
	if err != nil {
		return s.result, err
	}
	if s.result == nil {
		return s.result, fmt.Errorf("No result")
	}

	return s.result, nil
}

func initialize(device korgdevices.Device) *searcher {
	s := searcher{
		seeking: device,
	}
	for i := 0; i < portmidi.CountDevices(); i++ {
		device := portmidi.DeviceID(i)
		info := portmidi.Info(device)
		io := ""
		if info.IsInputAvailable {
			io += "i"
		}
		if info.IsOutputAvailable {
			io += "o"
		}
		//fmt.Printf("%d: %s %s: (%s)\n", i, info.Interface, info.Name, io)
		if info.IsInputAvailable {
			input, err := portmidi.NewInputStream(device, 32)
			if err != nil {
				continue
			} else {
				s.in = append(s.in, deviceStream{
					stream: input,
					info:   info,
					dispatcher: korgdispatch.EasyDispatch{
						SearchResponseHandler: s.searchHandler(int(device)),
					}.Dispatcher(),
				})
			}

		}
		if info.IsOutputAvailable {
			output, err := portmidi.NewOutputStream(device, 32, 32)
			if err != nil {
				continue
			} else {
				s.out = append(s.out, deviceStream{stream: output, info: info})
			}
		}

	}
	return &s
}
