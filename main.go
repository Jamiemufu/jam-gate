package main

import (
	"jam-gate/internal/button"
	"jam-gate/internal/led"
	"jam-gate/internal/status"
	"log"
	"time"

	"periph.io/x/host/v3"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	redLED, err := led.New("GPIO2")
	if err != nil {
		log.Fatal(err)
	}

	greenLED, err := led.New("GPIO3")
	if err != nil {
		log.Fatal(err)
	}

	button, err := button.New("GPIO4")
	if err != nil {
		log.Fatal(err)
	}

	statusLight := status.New(redLED, greenLED)

	if err := statusLight.Waiting(); err != nil {
		log.Fatal(err)
	}

	for {
		if button.InactiveFor(10 * time.Second) {
			if err := statusLight.Sleep(); err != nil {
				log.Fatal(err)
			}
		}
		if button.Pressed() {
			if err := statusLight.Start(); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
}
