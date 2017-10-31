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

	button1 *gpio.ButtonDriver
	button2 *gpio.ButtonDriver
)

func Reset() {
	fmt.Println("Reset the rooster.")
	led1.Off()
	led2.Off()
}

func LightsOn() {
	fmt.Println("lights on!")
	led1.On()
	led2.On()
}

func main() {
	master := gobot.NewMaster()

	a := api.NewAPI(master)
	a.Start()

	board := firmata.NewAdaptor(os.Args[1])

	led1 = gpio.NewLedDriver(board, "6")
	led2 = gpio.NewLedDriver(board, "7")

	button1 = gpio.NewButtonDriver(board, "4")
	button2 = gpio.NewButtonDriver(board, "5")

	// digital devices
	work := func() {

		button1.On(gpio.ButtonPush, func(data interface{}) {
			Reset()
		})

		button2.On(gpio.ButtonPush, func(data interface{}) {
			LightsOn()
		})
	}

	robot := gobot.NewRobot("rooster",
		[]gobot.Connection{board},
		[]gobot.Device{led1, led2, button1, button2},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
