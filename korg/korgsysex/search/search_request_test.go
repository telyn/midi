package search_test

import (
	"reflect"
	"testing"

	"github.com/telyn/midi/korg/korgsysex/search"
)

func TestParseResponse(t *testing.T) {
	tests := []struct {
		Data     []byte
		Response search.Response
		Err      error
	}{
		{ // 0 - an actual search response
			Data: []byte{0x07, 0x99,
				0x34, 0x12, 0x78, 0x56,
				0x51, 0x52, 0x53, 0x54,
			},
			Response: search.Response{
				Channel:    7,
				EchoBackID: 0x99,
				Family:     0x1234,
				Member:     0x5678,
				Minor:      0x5251,
				Major:      0x5453,
			},
		}, // TODO add some bad ones
	}
	for i, test := range tests {
		actual, err := search.ParseResponse(test.Data)
		if err != test.Err {
			t.Errorf("%d - got error %v when expecting %v", i, err, test.Err)
		}
		if !reflect.DeepEqual(test.Response, actual) {
			t.Errorf("%d - SearchResponse differed from expected.\nExpected %#v\nReceived %#v", i, test.Response, actual)
		}
	}
}
