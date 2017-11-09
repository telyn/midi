package korgsysex

type _korgErr string

const InvalidMessageError = _korgErr("Invalid korg system exclusive message")

func (err _korgErr) Error() string {
	return string(err)
}
