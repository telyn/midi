package devicesearch

import (
	"github.com/telyn/midi/korg/korgdevices"
)

// Search searches for the given device, forever
// it returns a stream to communicate with the device and the device's "global channel" (useful for Korg Native Mode)
// TODO(telyn): add a timeout
func Search(device korgdevices.Device) (res *SearchResult, err error) {
	searcher := initialize(device)

	res, err = searcher.search()
	return
}
