package main

import (
	"fmt"
	"jam-gate/internal/access"
	"jam-gate/internal/hardware"
	"log"
	"time"
)

func main() {
	d, err := hardware.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err := d.StatusLight.Waiting(); err != nil {
		log.Fatal(err)
	}

	pinControl := access.New("0000")

	for {
		// Read the key from the keypad
		key, err := d.Keypad.ReadKey()
		if err != nil {
			log.Fatal(err)
		}
		// Print the key if it is not empty
		if key != "" {
			switch key {
			case "*":
				fmt.Println("Resetting input")
				pinControl.Reset()
				if err := d.StatusLight.KeyPress(); err != nil {
					log.Fatal(err)
				}
			default:
				// Pass the key to the pin controller
				result := pinControl.PinController(key)
				// Handle the result of the pin controller
				if err := handleAccessResult(result, d); err != nil {
					log.Fatal(err)
				}

			}
		}
		// Check if the button has been inactive for 5 minutes
		if d.Button.InactiveFor(5 * time.Minute) {
			if err := d.StatusLight.Sleep(); err != nil {
				log.Fatal(err)
			}
		}
		// Check if the button has been pressed
		if d.Button.Pressed() {
			if err := d.Gate.Open(); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
}

// handleAccessResult handles the result of the access control and updates the hardware accordingly.
func handleAccessResult(result access.Result,
	devices *hardware.Hardware) error {
	switch result {
	case access.Granted:
		fmt.Println("Access granted")
		return devices.Gate.Open()
	case access.Denied:
		fmt.Println("Access denied")
		return devices.StatusLight.AccessDenied()
	case access.Pending:
		return devices.StatusLight.KeyPress()
	default:
		fmt.Println("Unknown access result")
		return devices.StatusLight.Error()
	}
}
