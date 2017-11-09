package msgs

type Handler interface {
	Handle(Message) error
}

type HandlerFunc func(Message) error

func (hf HandlerFunc) Handle(msg Message) error {
	return hf(msg)
}
