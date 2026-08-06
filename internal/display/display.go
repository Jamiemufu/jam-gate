package display

import (
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/devices/v3/ssd1306/image1bit"
)

type Display struct {
	bus  i2c.Bus
	oled *ssd1306.Dev
}

func New() (*Display, error) {
	// find first available i2c bus
	b, err := i2creg.Open("")
	if err != nil {
		return nil, err
	}

	// create new ssd1306 device
	oled, err := ssd1306.NewI2C(b, &ssd1306.DefaultOpts)
	if err != nil {
		b.Close()
		return nil, err
	}

	img := image1bit.NewVerticalLSB(oled.Bounds())

	// clear the display
	if err := oled.Draw(oled.Bounds(), img, image.Point{}); err != nil {
		b.Close()
		return nil, err
	}

	return &Display{
		bus:  b,
		oled: oled,
	}, nil
}

// Show displays the given text on the OLED display.
func (d *Display) Show(text string) error {
	img := image1bit.NewVerticalLSB(d.oled.Bounds())

	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{image1bit.On},
		Face: basicfont.Face7x13,
		Dot:  fixed.P(0, 13),
	}

	drawer.DrawString(text)

	return d.oled.Draw(d.oled.Bounds(), img, image.Point{})
}
