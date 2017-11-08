package msgs

type Kind byte

const (
	DataByte Kind = iota
	NoteOff
	NoteOn
	KeyPressure
	// for the purpose of making my life easy given that I don't care about Channel Mode messages, they are the same Kind as ControlChange
	ControlChange
	ProgramChange
	ChannelPressure
	PitchBend
	SystemExclusive
	SystemCommonTimeCode
	SystemCommonSongPositionPointer
	SystemCommonSongSelect
	SystemCommonUndefined
	SystemCommonTuneRequest
	SystemCommonEOX
	RealTime
)

func KindOf(b byte) Kind {
	switch {
	case b >= 0x80 && b <= 0x8F:
		return NoteOff
	case b >= 0x90 && b <= 0x9F:
		return NoteOn
	case b >= 0xA0 && b <= 0xAF:
		return KeyPressure
	case b >= 0xB0 && b <= 0xBF:
		return ControlChange
	case b >= 0xC0 && b <= 0xCF:
		return ProgramChange
	case b >= 0xD0 && b <= 0xDF:
		return ChannelPressure
	case b >= 0xE0 && b <= 0xEF:
		return PitchBend
	case b >= 0xF0 && b <= 0xF0:
		return SystemExclusive
	case b >= 0xF1 && b <= 0xF1:
		return SystemCommonTimeCode
	case b >= 0xF2 && b <= 0xF2:
		return SystemCommonSongPositionPointer
	case b >= 0xF3 && b <= 0xF3:
		return SystemCommonSongSelect
	case b >= 0xF4 && b <= 0xF5:
		return SystemCommonUndefined
	case b >= 0xF6 && b <= 0xF6:
		return SystemCommonTuneRequest
	case b >= 0xF7 && b <= 0xF7:
		return SystemCommonEOX
	case b >= 0xF8 && b <= 0xFF:
		return RealTime
	}
	return DataByte
}

// ChannelOf returns the MIDI channel for this status byte
// if the status byte does not contain a MIDI channel then 0xFF is returned
func ChannelOf(b byte) byte {
	switch KindOf(b) {
	case NoteOff, NoteOn, KeyPressure, ControlChange, ProgramChange, ChannelPressure, PitchBend:
		return b & 0xF
	}
	return 0xFF
}

// Bytes returns the number of data bytes to expect
// -1 means 'until 0xF7 is seen'
// -2 means 'this is not a status byte'
func (k Kind) Bytes() int {
	switch k {
	case NoteOff, NoteOn, KeyPressure, ControlChange, PitchBend:
		return 2
	case ProgramChange, ChannelPressure:
		return 1
	case SystemExclusive:
		return -1
	case SystemCommonSongPositionPointer:
		return 2
	case SystemCommonTimeCode, SystemCommonSongSelect:
		return 1
	case SystemCommonUndefined, SystemCommonTuneRequest, SystemCommonEOX:
		return 0
	case RealTime:
		return 0

	}
	return -2
}
