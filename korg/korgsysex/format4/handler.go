package format4

type Handler interface {
	Handle(Message) error
}
