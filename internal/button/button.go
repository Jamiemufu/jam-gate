package button

import (
	"fmt"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

type Button struct {
	pin        gpio.PinIO
	isPressed  bool
	wasPressed bool
}

func New(pinName string) (*Button, error) {
	pin := gpioreg.ByName(pinName)

	if pin == nil {
		return nil, fmt.Errorf("failed to find pin %s", pinName)
	}

	if err := pin.In(gpio.PullUp, gpio.NoEdge); err != nil {
		return nil, fmt.Errorf("failed to set pin %s as input: %v", pinName, err)
	}

	return &Button{
		pin:        pin,
		isPressed:  false,
		wasPressed: false,
	}, nil
}

// Pressed returns true if the button was pressed since the last call to Pressed.
func (b *Button) Pressed() bool {
	b.wasPressed = b.isPressed
	b.isPressed = b.pin.Read() == gpio.Low

	return b.isPressed && !b.wasPressed
}
