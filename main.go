package main

import (
	"fmt"
	hw "jam-gate/internal/hardware"
	"log"
	"time"
)

func main() {
	devices, err := hw.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err := devices.StatusLight.Waiting(); err != nil {
		log.Fatal(err)
	}

	for {
		// Read the key from the keypad
		key, err := devices.Keypad.ReadKey()
		if err != nil {
			log.Fatal(err)
		}
		// Print the key if it is not empty
		if key != "" {
			switch key {
			case "A":
				fmt.Println("Unlocking the gate")
				if err := devices.StatusLight.Unlock(); err != nil {
					log.Fatal(err)
				}
			case "D":
				fmt.Println("Locking the gate")
				if err := devices.StatusLight.Lock(); err != nil {
					log.Fatal(err)
				}
			default:
				if err := devices.StatusLight.KeyPress(); err != nil {
					log.Fatal(err)
				}
				fmt.Println(key)
			}
		}
		// Check if the button has been inactive for 5 minutes
		if devices.Button.InactiveFor(5 * time.Minute) {
			if err := devices.StatusLight.Sleep(); err != nil {
				log.Fatal(err)
			}
		}
		// Check if the button has been pressed
		if devices.Button.Pressed() {
			if err := devices.StatusLight.Start(); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
}
