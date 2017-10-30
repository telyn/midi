package nanokontrol

type Identity struct {
	MIDIChannel uint8
	// FamilyID is expected to be 1301 according to the MIDI implementation doc
	FamilyID uint16
	// MemberID is expected to be 0000 according to the MIDI implementation doc
	MemberID uint16
	// MinorVer is the minor part firmware version for the NanoKONTROL2
	MinorVer uint16
	// MajorVer is the major part of the firmware version for the NanoKONTROL2
	MajorVer uint16
}
