package api

// ThermostatInfo contains the full info values from the thermostat api
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
	AvaliableModes  int    `json:"avaliablemodes"`
}
