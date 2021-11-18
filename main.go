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
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	if *setMode != "" {
		setThermostatMode(*ip, convertThermostatMode(*setMode))
	}
	thermostatData := getThermostatInfo(*ip)
	printThermostatInfo(thermostatData)
}
