package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var (
	led1 *gpio.LedDriver
	led2 *gpio.LedDriver
)

func Reset() {
	fmt.Println("Reset the rooster.")
	led1.Off()
	led2.Off()
}

func main() {
	master := gobot.NewMaster()

	a := api.NewAPI(master)
	a.Start()

	board := firmata.NewAdaptor(os.Args[1])

	led1 = gpio.NewLedDriver(board, "6")
	led2 = gpio.NewLedDriver(board, "7")

	// digital devices
	work := func() {
		Reset()
	}

	robot := gobot.NewRobot("rooster",
		[]gobot.Connection{board},
		[]gobot.Device{led1, led2},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
