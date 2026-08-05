package access

import (
	"fmt"
	"strings"
)

type Access struct {
	pin    string
	keyset []string
}

type Result int

const (
	Pending Result = iota
	Granted
	Denied
)

func New(pin string) *Access {
	return &Access{
		pin:    pin,
		keyset: []string{},
	}
}

// PinController takes a key input and checks if the accumulated keyset matches the stored pin.
func (a *Access) PinController(key string) Result {
	a.keyset = append(a.keyset, key)

	if len(a.keyset) < len(a.pin) {
		fmt.Printf("Key pressed: %v ", key)
		fmt.Println("Current keyset:", a.keyset)
		return Pending
	}

	fmt.Printf("Current keyset: %v ", a.keyset)
	matches := a.Check()

	if matches {
		return Granted
	}

	return Denied
}

// Check compares the accumulated keyset with the stored pin and resets the keyset afterward.
func (a *Access) Check() bool {
	matches := strings.Join(a.keyset, "") == a.pin
	a.Reset()
	return matches
}

// Reset clears the accumulated keyset.
func (a *Access) Reset() {
	a.keyset = a.keyset[:0]
	fmt.Println("Keyset reset")
}
