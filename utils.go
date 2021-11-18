package main

import (
	"fmt"
	"reflect"
)

func printThermostatInfo(info thermostatInfo) {
	values := reflect.ValueOf(info)
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(values.Type().Field(i).Name, "\t", values.Field(i))
	}
}
