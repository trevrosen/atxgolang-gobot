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

	blink_work := func() {
		gobot.Every(10*time.Second, func() {
			led.Toggle()
		})
	}

	buzz_work := func() {
		gobot.Every(5*time.Second, func() {
			// make it buzz
		})
	}

	blink_bot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{e},
		[]gobot.Device{led},
		blink_work,
	)

	buzz_bot := gobot.NewRobot("buzzBot",
		[]gobot.Connection{e},
		[]gobot.Device{buzzer},
		buzz_work,
	)

	gbot.AddRobot(blink_bot)
	gbot.AddRobot(buzz_bot)

	a := api.NewAPI(gbot)
	a.Debug()
	a.Start()

	gbot.Start()
}
