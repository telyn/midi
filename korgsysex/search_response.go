package korgsysex

import "encoding/binary"

type SearchResponse struct {
	Identity   Identity
	EchoBackID uint8
}

// KorgSysEx encodes the SearchResponse, ignoring the channel as there's one in the Identity
func (sr SearchResponse) KorgSysEx(channel uint8) {
	major := make([]byte, 2)
	minor := make([]byte, 2)
	family := make([]byte, 2)
	member := make([]byte, 2)
	binary.LittleEndian.PutUint16(major, sr.Identity.MajorVer)
	binary.LittleEndian.PutUint16(minor, sr.Identity.MinorVer)
	binary.LittleEndian.PutUint16(member, sr.Identity.MemberID)
	binary.LittleEndian.PutUint16(family, sr.Identity.FamilyID)
	return korgsysex.Message{
		Format: 5,
		Data: []byte{
			0x01,
			sr.Identity.Channel & 0xF,
			sr.EchoBackID,
			family[0], family[1],
			member[0], member[1],
			minor[0], minor[1],
			major[0], major[1],
		},
	}
}

func (sr *SearchResponse) Parse(b []byte) error {
	sr.Identity.Channel = b[0]
	sr.EchoBackID = b[1]
	sr.Identity.FamilyID = binary.LittleEndian.Uint16(b[2:4])
	sr.Identity.MemberID = binary.LittleEndian.Uint16(b[4:6])
	sr.Identity.MinorVer = binary.LittleEndian.Uint16(b[6:8])
	sr.Identity.MajorVer = binary.LittleEndian.Uint16(b[8:10])
}
