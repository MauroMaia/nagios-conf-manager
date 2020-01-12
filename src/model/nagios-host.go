package model

import (
	"encoding/json"
	"strconv"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type Host struct {
	//      # REQUIRED
	HostName             string `json:"host_name"`
	Name                 string `json:"name"`
	Alias                string `json:"alias"`
	Address              string `json:"address"`
	MaxCheckAttempts     int    `json:"max_check_attempts"`
	NotificationInterval int    `json:"notification_interval"`
	NotificationPeriod   string `json:"notification_period"` // *NotificationPeriod#timeperiod_name
	Contact              string `json:"contact"`             // *Contact#Name
	ContactGroups        string `json:"contact_groups"`      // *ContactGroups#Name
	CheckPeriod          string `json:"check_period"`        // *TimePeriod#name

	//      # OPTINAL
	// 0 = do NOT register object definition, 1 = register object definition (this is the default).
	Register bool `json:"register"`

	// use - Name of host template to use
	Use string `json:"use"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#host
}

func NewNagiosHost(defineStringMap map[string]string) *Host {
	hostObjectNameString := defineStringMap["name"]

	var hostNameString, aliasString, addressString string
	var maxCheckAttemptsInt, notificationInterval int64
	var registerBool bool

	if contactName, ok := defineStringMap["hostName"]; ok == true {
		hostNameString = contactName
	}
	if alias, ok := defineStringMap["alias"]; ok == true {
		aliasString = alias
	}
	if address, ok := defineStringMap["address"]; ok == true {
		addressString = address
	}
	if checkAtt, ok := defineStringMap["max_check_attempts"]; ok == true {
		maxCheckAttemptsInt, _ = strconv.ParseInt(checkAtt, 10, 0)
	}
	if notificationInt, ok := defineStringMap["notification_interval"]; ok == true {
		notificationInterval, _ = strconv.ParseInt(notificationInt, 10, 0)
	}
	if register, ok := defineStringMap["register"]; ok == true {
		registerBool, _ = strconv.ParseBool(register)
	}

	return &Host{
		Name:                 hostObjectNameString,
		HostName:             hostNameString,
		Alias:                aliasString,
		Address:              addressString,
		MaxCheckAttempts:     int(maxCheckAttemptsInt),
		NotificationInterval: int(notificationInterval),

		Register: registerBool,
		Use:      "linux-server",
	}
}

func (host *Host) String() string {
	b, err := json.Marshal(host)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (host *Host) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implementedF")
}
