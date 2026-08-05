// internal/led/led.go
package led

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

type LED struct {
	pin  gpio.PinIO
	isOn bool
}

// New creates a new LED instance associated with the specified GPIO pin.
func New(pinName string) (*LED, error) {
	pin := gpioreg.ByName(pinName)

	if pin == nil {
		return nil, fmt.Errorf("failed to find pin %s", pinName)
	}

	if err := pin.Out(gpio.Low); err != nil {
		return nil, fmt.Errorf("failed to set pin %s to low: %v", pinName, err)
	}

	return &LED{pin: pin, isOn: false}, nil
}

// On turns the LED on by setting the GPIO pin high.
func (l *LED) On() error {
	if err := l.pin.Out(gpio.High); err != nil {
		return fmt.Errorf("failed to set pin %s to high: %v", l.pin.Name(), err)
	}
	l.isOn = true
	return nil
}

// Off turns the LED off by setting the GPIO pin low.
func (l *LED) Off() error {
	if err := l.pin.Out(gpio.Low); err != nil {
		return fmt.Errorf("failed to set pin %s to low: %v", l.pin.Name(), err)
	}
	l.isOn = false
	return nil
}

// Toggle toggles the state of the LED. If the LED is currently on, it will be turned off, and vice versa.
func (l *LED) Toggle() error {
	if l.isOn {
		return l.Off()
	}
	return l.On()
}

// Blink makes the LED blink for a specified duration and interval. The LED will toggle its state at the specified interval until the total duration has elapsed.
func (l *LED) Blink(duration time.Duration, interval time.Duration) error {
	end := time.Now().Add(duration)
	for time.Now().Before(end) {
		if err := l.Toggle(); err != nil {
			return err
		}
		time.Sleep(interval)
	}
	return nil
}

// BlinkForever makes the LED blink indefinitely at the specified interval. The LED will toggle its state at the specified interval until the program is terminated.
func (l *LED) BlinkForever(interval time.Duration) error {
	for {
		if err := l.Toggle(); err != nil {
			return err
		}
		time.Sleep(interval)
	}
}
