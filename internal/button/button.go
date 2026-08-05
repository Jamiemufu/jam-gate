package button

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

type Button struct {
	pin         gpio.PinIO
	isPressed   bool
	wasPressed  bool
	lastPressed time.Time
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
		pin:         pin,
		isPressed:   false,
		wasPressed:  false,
		lastPressed: time.Now(),
	}, nil
}

// Pressed returns true if the button was pressed since the last call to Pressed.
func (b *Button) Pressed() bool {
	b.wasPressed = b.isPressed
	b.isPressed = b.pin.Read() == gpio.Low

	if b.isPressed && !b.wasPressed {
		b.lastPressed = time.Now()
	}

	return b.isPressed && !b.wasPressed
}

// InactiveFor returns true if the button has not been pressed for the specified duration.
func (b *Button) InactiveFor(d time.Duration) bool {
	return time.Since(b.lastPressed) >= d
}
