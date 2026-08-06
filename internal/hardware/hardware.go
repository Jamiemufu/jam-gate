package hardware

import (
	"jam-gate/internal/display"
	"jam-gate/internal/gate"
	"jam-gate/internal/keypad"
	"jam-gate/internal/led"
	"jam-gate/internal/status"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type Hardware struct {
	Keypad      *keypad.Keypad
	Gate        *gate.Gate
	StatusLight *status.Status
	OLED        *display.Display
}

func Init() (*Hardware, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}

	redLED, err := led.New("GPIO27")
	if err != nil {
		return nil, err
	}

	greenLED, err := led.New("GPIO17")
	if err != nil {
		return nil, err
	}

	rows := [4]gpio.PinIO{
		gpioreg.ByName("GPIO26"),
		gpioreg.ByName("GPIO19"),
		gpioreg.ByName("GPIO13"),
		gpioreg.ByName("GPIO6"),
	}

	cols := [4]gpio.PinIO{
		gpioreg.ByName("GPIO21"),
		gpioreg.ByName("GPIO20"),
		gpioreg.ByName("GPIO16"),
		gpioreg.ByName("GPIO12"),
	}

	pad, err := keypad.New(rows, cols)
	if err != nil {
		return nil, err
	}

	statusLight := status.New(redLED, greenLED)

	display, err := display.New()
	if err != nil {
		return nil, err
	}

	display.Show("Ready...")

	return &Hardware{
		Keypad:      pad,
		Gate:        gate.NewSimulator(statusLight),
		StatusLight: statusLight,
		OLED:        display,
	}, nil
}
