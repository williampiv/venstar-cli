package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/williampiv/venstar-cli/api"
	"github.com/williampiv/venstar-cli/server"
	"os"
)

var commitVersion string

func main() {
	parser := argparse.NewParser("venstar-cli", "Access Venstar Thermostat")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "An IP (or hostname) is required to actually access a thermostat"})
	setMode := parser.Selector("", "set-mode", []string{"off", "heat", "cool", "auto"}, &argparse.Options{Required: false, Help: "Set the thermostat mode"})
	coolTemp := parser.Int("", "set-cool-temp", &argparse.Options{Required: false, Help: "Set the cool-to temperature"})
	fanMode := parser.Selector("", "set-fan-mode", []string{"auto", "on"}, &argparse.Options{Required: false, Help: "Set current fan mode to on or auto"})
	serverFlag := parser.Flag("", "server", &argparse.Options{Required: false, Help: "run the server"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	fmt.Println("Version:", commitVersion)

	if *serverFlag {
		server.Entry(*ip)
		os.Exit(0)
	} else {

		thermostatData := api.GetThermostatInfo(*ip)

		if *setMode != "" {
			api.SetThermostatMode(*ip, api.ConvertThermostatMode(*setMode), thermostatData)
		}
		if *coolTemp != 0 {
			api.SetCoolTemp(*ip, *coolTemp, thermostatData)
		}
		if *fanMode != "" {
			api.SetFanMode(*ip, api.ConvertFanMode(*fanMode), thermostatData)
		}
		printThermostatInfo(thermostatData)
	}

}
