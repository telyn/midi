package nanokontrol

import "fmt"

type Message interface {
	fmt.Stringer
}

type EnableLED struct {
	LED ControlID
}

func (msg EnableLED) Bytes() []byte {
	return []byte{
		0xBF, byte(msg.LED), 0x7F, 0x00,
	}
}

func (msg EnableLED) String() string {
	return fmt.Sprintf("Enable LED for %s", msg.LED.String())
}

type DisableLED struct {
	LED ControlID
}

func (msg DisableLED) Bytes() []byte {
	return []byte{
		0xBF, byte(msg.LED), 0x00, 0x00,
	}
}

func (msg DisableLED) String() string {
	return fmt.Sprintf("Disable LED for %s", msg.LED.String())
}

type ButtonOn struct {
	Button ControlID
}

func (msg ButtonOn) String() string {
	return fmt.Sprintf("%s on", msg.Button.String())
}

type ButtonOff struct {
	Button ControlID
}

func (msg ButtonOff) String() string {
	return fmt.Sprintf("%s off", msg.Button.String())
}

type ValueChanged struct {
	Control ControlID
	Value   uint8
}

func (msg ValueChanged) String() string {
	return fmt.Sprintf("%s: %d", msg.Control.String, msg.Value)
}
