package msgs

type Message struct {
	Kind    Kind
	Channel byte
	Data    []byte
}
