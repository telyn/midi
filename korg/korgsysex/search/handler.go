package search

import (
	"fmt"
)

const _searchRequest byte = 0x00
const _searchResponse byte = 0x01

type Handler struct {
	RequestHandler  func(Request) error
	ResponseHandler func(Response) error
}

func (h Handler) Handle(data []byte) error {
	switch data[0] {
	case _searchRequest:
		fmt.Errorf("requests aren't supported yet") //TODO
		req, err := ParseRequest(data)
		if err != nil {
			return err
		}
		if h.RequestHandler != nil {
			h.RequestHandler(req)
		}
	case _searchResponse:
		res, err := ParseResponse(data[1:])
		if err != nil {
			return err
		}
		if h.ResponseHandler != nil {
			return h.ResponseHandler(res)
		}
		return nil
	}
	return fmt.Errorf("search.Handler.Handle was called with %x - neither a request nor a response", data[0])
}
