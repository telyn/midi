package msgs

type Message struct {
	Kind    Kind
	Channel byte
	Data    []byte
}

// Status reassmbles the status byte for this message
func (m Message) Status() byte {
	if m.Kind.HasChannel() {
		return m.Kind.Byte() | m.Channel
	}
	return m.Kind.Byte()
}
