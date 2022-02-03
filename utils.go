package main

import (
	"fmt"
	"github.com/williampiv/venstar-cli/api"
	"reflect"
)

func printThermostatInfo(info api.ThermostatInfo) {
	values := reflect.ValueOf(info)
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(values.Type().Field(i).Name, "\t", values.Field(i))
	}
}
