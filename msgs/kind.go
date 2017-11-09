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
	RealTimeClock
	RealTimeUndefined
	RealTimeStart
	RealTimeContinue
	RealTimeStop
	RealTimeActiveSensing
	RealTimeSystemReset
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
	case b >= 0xF8 && b <= 0xF8:
		return RealTimeClock
	case b >= 0xF9 && b <= 0xF9:
		return RealTimeUndefined
	case b >= 0xFA && b <= 0xFA:
		return RealTimeStart
	case b >= 0xFB && b <= 0xFB:
		return RealTimeContinue
	case b >= 0xFC && b <= 0xFC:
		return RealTimeStop
	case b >= 0xFD && b <= 0xFD:
		return RealTimeUndefined
	case b >= 0xFE && b <= 0xFE:
		return RealTimeActiveSensing
	case b >= 0xFF && b <= 0xFF:
		return RealTimeSystemReset
	}
	return DataByte
}

// ChannelOf returns the MIDI channel for this status byte
// if the status byte does not contain a MIDI channel then 0xFF is returned
func ChannelOf(b byte) byte {
	if KindOf(b).HasChannel() {
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
	case RealTimeClock, RealTimeUndefined, RealTimeStart, RealTimeContinue, RealTimeStop, RealTimeActiveSensing, RealTimeSystemReset:
		return 0

	}
	return -2
}

func (k Kind) IsStatus() bool {
	return k != DataByte
}

func (k Kind) HasChannel() bool {
	switch k {
	case NoteOff, NoteOn, KeyPressure, ControlChange, ProgramChange, ChannelPressure, PitchBend:
		return true
	}
	return false
}

func (k Kind) RealTime() bool {
	switch k {
	case RealTimeClock, RealTimeUndefined, RealTimeStart, RealTimeContinue, RealTimeStop, RealTimeActiveSensing, RealTimeSystemReset:
		return true
	}
	return false
}

// Byte gets the status byte for this Kind.
// If k.HasChannel() then the calling function will need to bitwise-or the result with a channel between 0x00 and 0x0F
func (k Kind) Byte() byte {
	switch k {
	case NoteOff:
		return 0x80
	case NoteOn:
		return 0x90
	case KeyPressure:
		return 0xA0
	case ControlChange:
		return 0xB0
	case ProgramChange:
		return 0xC0
	case ChannelPressure:
		return 0xD0
	case PitchBend:
		return 0xE0
	case SystemExclusive:
		return 0xF0
	case SystemCommonTimeCode:
		return 0xF1
	case SystemCommonSongPositionPointer:
		return 0xF2
	case SystemCommonSongSelect:
		return 0xF3
	case SystemCommonUndefined:
		return 0xF4
	case SystemCommonTuneRequest:
		return 0xF6
	case SystemCommonEOX:
		return 0xF7
	case RealTimeClock:
		return 0xF8
	case RealTimeUndefined:
		return 0xF9
	case RealTimeStart:
		return 0xFA
	case RealTimeContinue:
		return 0xFB
	case RealTimeStop:
		return 0xFC
	case RealTimeActiveSensing:
		return 0xFE
	case RealTimeSystemReset:
		return 0xFF
	}
	return 0x00
}
