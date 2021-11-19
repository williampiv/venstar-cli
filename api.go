package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type thermostatInfo struct {
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
	AvaliableModes  int    `json:"avaliablemodes"`
}

func convertThermostatMode(mode string) int {
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

func getThermostatInfo(ipaddress string) thermostatInfo {
	resp, err := http.Get(fmt.Sprintf("http://%s/query/info", ipaddress))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var thermostatResponse thermostatInfo
	json.Unmarshal(respBytes, &thermostatResponse)
	return thermostatResponse
}

func setThermostatMode(ipaddress string, mode int, currentInfo thermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", currentInfo.CoolTemp))
	data.Set("mode", fmt.Sprintf("%d", mode))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}

func setCoolTemp(ipaddress string, coolTemp int, currentInfo thermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", coolTemp))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}
