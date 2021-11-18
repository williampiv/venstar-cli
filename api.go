package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ThermostatInfo struct {
	Name            string `json:"name"`
	Mode            int    `json:"mode"`
	State           int    `json:"state"`
	Fan             int    `json:"fan"`
	FanState        int    `json:"fanstate"`
	TempUnits       int    `json:"tempunits"`
	Schedule        int    `json:"schedule"`
	SchedulePart    int    `json:"schedulepart"`
	Away            int    `json:"away"`
	Holiday         int    `json:"holiday"`
	Override        int    `json:"override"`
	OverrideTime    int    `json:"overridetime"`
	ForceUnoccupied int    `json:"forceunocc"`
	SpaceTemp       int    `json:"spacetemp"`
	HeatTemp        int    `json:"heattemp"`
	CoolTemp        int    `json:"cooltemp"`
	CoolTempMin     int    `json:"cooltempmin"`
	CoolTempMax     int    `json:"cooltempmax"`
	HeatTempMin     int    `json:"heattempmin"`
	HeatTempMax     int    `json:"heattempmax"`
	SetPointDelta   int    `json:"setpointdelta"`
	Humidity        int    `json:"hum"`
	AvaliableModes  []int  `json:"avaliablemodes"`
}

func ConvertThermostatMode(mode string) int {
	switch mode {
	case "off":
		return 0
	case "heat":
		return 1
	case "cool":
		return 2
	case "auto":
		return 3
	default:
		return 0
	}
}

func GetThermostatInfo(ipaddress string) ThermostatInfo {
	resp, err := http.Get(fmt.Sprintf("http://%s/query/info", ipaddress))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	resp_bytes, _ := ioutil.ReadAll(resp.Body)
	var thermostat_response ThermostatInfo
	json.Unmarshal(resp_bytes, &thermostat_response)
	return thermostat_response
}

func SetThermostatMode(ipaddress string, mode int) bool {
	data := make(map[string]int, 3)
	currentInfo := GetThermostatInfo(ipaddress)
	data["heattemp"] = currentInfo.HeatTemp
	data["cooltemp"] = currentInfo.CoolTemp
	data["mode"] = mode
	jsonData, err := json.Marshal(data)
	if err != nil {
		return false
	}
	_, reqErr := http.Post(fmt.Sprintf("http://%s/control", ipaddress), "application/json", bytes.NewBuffer(jsonData))
	return reqErr == nil
}
