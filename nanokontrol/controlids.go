package nanokontrol

//go:generate stringer -type=ControlID

type ControlID uint8

func (c ControlID) HasLED() bool {
	if c < 0x20 {
		return false
	}

	if c >= 0x3A && c < 0x40 {
		return false
	}
	return true
}

const (
	Slider1 ControlID = 0x00 + iota
	Slider2
	Slider3
	Slider4
	Slider5
	Slider6
	Slider7
	Slider8
)
const (
	Knob1 ControlID = 0x10 + iota
	Knob2
	Knob3
	Knob4
	Knob5
	Knob6
	Knob7
	Knob8
)
const (
	SoloButton1 ControlID = 0x20 + iota
	SoloButton2
	SoloButton3
	SoloButton4
	SoloButton5
	SoloButton6
	SoloButton7
	SoloButton8
)
const (
	PlayButton ControlID = 0x29 + iota
	StopButton
	Rewind
	FastForward
	Record
	Cycle
)
const (
	MuteButton1 ControlID = 0x30 + iota
	MuteButton2
	MuteButton3
	MuteButton4
	MuteButton5
	MuteButton6
	MuteButton7
	MuteButton8
)
const (
	TrackLeft ControlID = 0x3A + iota
	TrackRight
	MarkerSet
	MarkerLeft
	MarkerRight
)
const (
	RecordButton1 ControlID = 0x40 + iota
	RecordButton2
	RecordButton3
	RecordButton4
	RecordButton5
	RecordButton6
	RecordButton7
	RecordButton8
)
