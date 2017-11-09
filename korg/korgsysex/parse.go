package korgsysex

/*
const NativeModeMessage byte = 0xBF

func (c Conn) parseMessage(buf *bytes.Buffer) (Message, error) {
	switch buf.ReadByte() {
	case NativeModeMessage:
		switch device {
		case NanoKONTROL2:
			nanokontrol.Parse(buf)
		}
	case SysExMessage:
		mfr, err := buf.ReadByte()
		if err != nil {
			return nil, err
		}
		switch mfr {
		case Korg:
			return parseKorgMessage(buf)
		case Universal:
		default:
			return readUnsupportedSysEx(mfr, buf)
		}
	}
}

func (c Conn) parsePacket(b []byte) ([]Message, error) {
	buf := bytes.NewBuffer(b)
	for buf.Len() > 0 {
		ParseMessage(buf)
	}
}
*/
