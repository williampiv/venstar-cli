package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("venstar-cli", "Access Venstar Thermostat")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "An IP (or hostname) is required to actually access a thermostat"})
	setMode := parser.Selector("", "set-mode", []string{"off", "heat", "cool", "auto"}, &argparse.Options{Required: false, Help: "Set the thermostat mode"})
	coolTemp := parser.Int("", "set-cool-temp", &argparse.Options{Required: false, Help: "Set the cool-to temperature"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	thermostatData := getThermostatInfo(*ip)

	if *setMode != "" {
		setThermostatMode(*ip, convertThermostatMode(*setMode), thermostatData)
	}
	if *coolTemp != 0 {
		setCoolTemp(*ip, *coolTemp, thermostatData)
	}
	printThermostatInfo(thermostatData)
}
