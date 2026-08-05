package keypad

import (
	"time"

	"periph.io/x/conn/v3/gpio"
)

type Keypad struct {
	rows [4]gpio.PinIO
	cols [4]gpio.PinIO
}

// New initializes a new Keypad with the given row and column GPIO pins.
func New(rows [4]gpio.PinIO, cols [4]gpio.PinIO) (*Keypad, error) {
	// initialize rows as output and set them low
	for _, row := range rows {
		if err := row.Out(gpio.Low); err != nil {
			return nil, err
		}
	}
	// initialize columns as input with pull-down resistors
	for _, col := range cols {
		if err := col.In(gpio.PullDown, gpio.NoEdge); err != nil {
			return nil, err
		}
	}

	return &Keypad{
		rows: rows,
		cols: cols,
	}, nil
}

// Scan checks the keypad for a pressed key and returns the corresponding key value.
func (k *Keypad) Scan() (string, error) {
	for rowIndex, row := range k.rows {
		// set the current row high
		if err := row.Out(gpio.High); err != nil {
			return "", err
		}

		// check each column for a high signal
		for colIndex, col := range k.cols {
			if col.Read() == gpio.High {
				// reset the current row to low before returning
				if err := row.Out(gpio.Low); err != nil {
					return "", err
				}
				return k.getKey(rowIndex, colIndex), nil
			}
		}

		// reset the current row to low before moving to the next row
		if err := row.Out(gpio.Low); err != nil {
			return "", err
		}
	}

	return "", nil // no key pressed
}

// getKey returns the key value based on the row and column indices.
func (k *Keypad) getKey(row, col int) string {
	keys := [4][4]string{
		{"1", "2", "3", "A"},
		{"4", "5", "6", "B"},
		{"7", "8", "9", "C"},
		{"*", "0", "#", "D"},
	}

	// check for out-of-bounds access
	if row < 0 || row >= len(keys) || col < 0 || col >= len(keys[row]) {
		return ""
	}

	return keys[row][col]
}

// ReadKey waits for a key press and returns the key value when released.
func (k *Keypad) ReadKey() (string, error) {
	key, err := k.Scan()
	if err != nil {
		return "", err
	}

	for {
		releasedKey, err := k.Scan()
		if err != nil {
			return "", err
		}

		if releasedKey == "" {
			return key, nil
		}

		time.Sleep(20 * time.Millisecond)
	}
}
