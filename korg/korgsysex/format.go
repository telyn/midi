package korgsysex

import "github.com/telyn/midi/sysex"

func FormatOf(sysex sysex.SysEx) byte {
	return sysex.Data[0] & 0xF0 >> 4
}
func ChannelOf(sysex sysex.SysEx) byte {
	return sysex.Data[0] & 0x0F
}
