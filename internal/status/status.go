package status

import (
	"jam-gate/internal/led"
	"time"
)

type Status struct {
	red   *led.LED
	green *led.LED
}

// New creates a new Status instance with the provided red and green LEDs.
func New(red *led.LED, green *led.LED) *Status {
	return &Status{
		red:   red,
		green: green,
	}
}

// Stop turns on the red LED and turns off the green LED, indicating a stopped status.
func (s *Status) Stop() error {
	if err := s.red.On(); err != nil {
		return err
	}
	if err := s.green.Off(); err != nil {
		return err
	}
	return nil
}

// Start turns off the red LED and turns on the green LED, indicating a started status.
func (s *Status) Start() error {

	end := 10 * time.Second
	interval := 500 * time.Millisecond

	if err := s.red.Off(); err != nil {
		return err
	}

	if err := s.green.Blink(end, interval); err != nil {
		return err
	}

	if err := s.Waiting(); err != nil {
		return err
	}

	return nil
}

// Off turns on the red LED and turns off the green LED, indicating an off status.
func (s *Status) Off() error {
	if err := s.red.On(); err != nil {
		return err
	}
	if err := s.green.Off(); err != nil {
		return err
	}
	return nil
}

// On turns off the red LED and turns on the green LED, indicating an on status.
func (s *Status) On() error {
	if err := s.red.Off(); err != nil {
		return err
	}
	if err := s.green.On(); err != nil {
		return err
	}
	return nil
}

// Waiting turns on the red LED to indicate a waiting status.
func (s *Status) Waiting() error {
	if err := s.red.On(); err != nil {
		return err
	}
	return nil
}

// Sleep turns off both the red and green LEDs, indicating a sleep status.
func (s *Status) Sleep() error {
	if err := s.red.Off(); err != nil {
		return err
	}
	if err := s.green.Off(); err != nil {
		return err
	}
	return nil
}

// Error continuously blinks the red LED to indicate an error status.
func (s *Status) Error() error {
	interval := 500 * time.Millisecond
	if err := s.red.BlinkForever(interval); err != nil {
		return err
	}
	return nil
}

// KeyPress toggles the red LED on and off to indicate a key press event.
func (s *Status) KeyPress() error {
	if err := s.red.Toggle(); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	s.red.Toggle()
	return nil
}

// Unlock turns off the red LED and blinks the green LED to indicate an unlock status.
func (s *Status) Unlock() error {
	if err := s.red.Off(); err != nil {
		return err
	}

	if err := s.green.Blink(2*time.Second, 500*time.Millisecond); err != nil {
		return err
	}

	s.Waiting() // Set the status to waiting after unlocking

	return nil
}

// Lock turns on the red LED and turns off the green LED to indicate a lock status.
func (s *Status) Lock() error {
	if err := s.red.Blink(2*time.Second, 500*time.Millisecond); err != nil {
		return err
	}

	s.Waiting() // Set the status to waiting before locking

	return nil
}
