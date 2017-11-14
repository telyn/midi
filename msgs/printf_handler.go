package msgs

import (
	"fmt"
	"io"
)

type PrintfHandler struct {
	Writer io.Writer
}

func (ph PrintfHandler) Handle(msg Message) error {
	fmt.Fprintf(ph.Writer, "%v: %x\n", msg.Kind, msg.Data)
	return nil
}
