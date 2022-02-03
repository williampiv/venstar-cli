package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ConvertThermostatMode converts the "plain text" mode to the numeric value
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

// ConvertFanMode converts the "plain text" mode to the numeric value
func ConvertFanMode(mode string) int {
	switch mode {
	case "auto":
		return 0
	case "on":
		return 1
	default:
		return 0
	}
}

// GetThermostatInfo pulls the full info from the thermostat
func GetThermostatInfo(ipaddress string) ThermostatInfo {
	resp, err := http.Get(fmt.Sprintf("http://%s/query/info", ipaddress))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var thermostatResponse ThermostatInfo
	json.Unmarshal(respBytes, &thermostatResponse)
	return thermostatResponse
}

// SetThermostatMode sets the thermostat mode
func SetThermostatMode(ipaddress string, mode int, currentInfo ThermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", currentInfo.CoolTemp))
	data.Set("mode", fmt.Sprintf("%d", mode))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}

// SetCoolTemp sets the "cool to" temp on the thermostat
func SetCoolTemp(ipaddress string, coolTemp int, currentInfo ThermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", coolTemp))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}

// SetFanMode sets whether the fan is "auto" or "on"
func SetFanMode(ipaddress string, fanMode int, currentInfo ThermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", currentInfo.CoolTemp))
	data.Set("fan", fmt.Sprintf("%d", fanMode))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}
