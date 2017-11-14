package search

import (
	"fmt"
	"io"
)

func PrintfRequestHandler(wr io.Writer) func(r Request) error {
	return func(r Request) error {
		fmt.Fprintf(wr, "Received search request with echo back ID %v\n", r.EchoBackID)
		return nil
	}
}

func PrintfResponseHandler(wr io.Writer) func(r Response) error {
	return func(r Response) error {
		fmt.Fprintf(wr, "Received search response from device %v with channel %v, echo back ID %x, firmware %v.%v\n", r.Device(), r.Channel, r.EchoBackID, r.Major, r.Minor)
		return nil
	}
}

type PrintfHandler struct {
	Writer io.Writer
}

func (ph PrintfHandler) Handle(data []byte) error {
	return Handler{
		RequestHandler:  PrintfRequestHandler(ph.Writer),
		ResponseHandler: PrintfResponseHandler(ph.Writer),
	}.Handle(data)
}
