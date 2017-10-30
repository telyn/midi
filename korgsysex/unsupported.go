package korgsysex

/*

type UnsupportedMessage struct {
	Manufacturer byte
	Bytes        []byte
}

func (msg UnsupportedMessage) String() string {
	return fmt.Sprintf("Unsupported message from manufacturer %x: %x", msg.Manufacturer, msg.Bytes)
}

func readUnsupportedSysEx(mfr byte, buf *bytes.Buffer) (msg Message, err error) {
	msg.Manufacturer = mfr
	bs, err := buf.ReadBytes(SysExEndOfMessage)
	if err != nil {
		return
	}
	msg.Bytes = bs
}
*/
