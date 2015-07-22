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

	board_led := gpio.NewLedDriver(e, "led", "13")
	red_led := gpio.NewLedDriver(e, "led", "3")
	green_led := gpio.NewLedDriver(e, "led", "2")
	buzzer := gpio.NewBuzzerDriver(e, "buzzer", "4")

	// Blink the Board LED
	board_blink_work := func() {
		gobot.Every(10*time.Second, func() {
			board_led.Toggle()
		})
	}

	// Ring the buzzer
	buzzer_work := func() {
		gobot.Every(4*time.Second, func() {
			buzzer.Tone(gpio.G5, gpio.Eighth)
		})
	}

	board_blink_bot := gobot.NewRobot("Board LED",
		[]gobot.Connection{e},
		[]gobot.Device{board_led},
		board_blink_work,
	)

	buzz_bot := gobot.NewRobot("buzzBot",
		[]gobot.Connection{e},
		[]gobot.Device{buzzer},
		buzzer_work,
	)

	red_blink_bot := gobot.NewRobot("Red LED",
		[]gobot.Connection{e},
		[]gobot.Device{red_led},
	)

	green_blink_bot := gobot.NewRobot("Green LED",
		[]gobot.Connection{e},
		[]gobot.Device{green_led},
	)

	gbot.AddRobot(board_blink_bot)
	gbot.AddRobot(green_blink_bot)
	gbot.AddRobot(red_blink_bot)
	gbot.AddRobot(buzz_bot)

	a := api.NewAPI(gbot)
	a.Debug()
	a.Start()

	gbot.Start()
}
