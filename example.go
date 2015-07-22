package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	e := edison.NewEdisonAdaptor("edison")

	led := gpio.NewLedDriver(e, "led", "13")
	buzzer := gpio.NewBuzzerDriver(e, "buzzer", "4")

	// Blink the LED
	blink_work := func() {
		gobot.Every(10*time.Second, func() {
			led.Toggle()
		})
	}

	// Ring the buzzer
	buzzer_work := func() {
		gobot.Every(4*time.Second, func() {
			buzzer.Tone(gpio.G5, gpio.Eighth)
		})
	}

	blink_bot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{e},
		[]gobot.Device{led},
		blink_work,
	)

	buzz_bot := gobot.NewRobot("buzzBot",
		[]gobot.Connection{e},
		[]gobot.Device{led},
		buzzer_work,
	)

	gbot.AddRobot(blink_bot)
	gbot.AddRobot(buzz_bot)

	a := api.NewAPI(gbot)
	a.Debug()
	a.Start()

	gbot.Start()
}
