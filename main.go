package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("venstar-cli", "Access Venstar Thermostat")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "An IP (or hostname) is required to actually access a thermostat"})
	set_mode := parser.Selector("", "set-mode", []string{"off", "heat", "cool", "auto"}, &argparse.Options{Required: false, Help: "Set the thermostat mode"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	if *set_mode != "" {
		SetThermostatMode(*ip, ConvertThermostatMode(*set_mode))
	}
	thermostat_data := GetThermostatInfo(*ip)
	fmt.Println("Current Thermostat Mode: ", thermostat_data.State)
}
