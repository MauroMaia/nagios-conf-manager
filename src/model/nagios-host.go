package model

import (
	"encoding/json"
	"regexp"

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

var reHostName = regexp.MustCompile(`.*hostName *(.+).*`)
var reName = regexp.MustCompile(`.*name *(.+).*`)
var reHostAlias = regexp.MustCompile(`.*alias *(.+).*`)
var reAddress = regexp.MustCompile(`.*address *(.+).*`)
var reMaxCheckAttempts = regexp.MustCompile(`.*max_check_attempts *(.+).*`)
var reNotificationInterval = regexp.MustCompile(`.*notification_interval *(.+).*`)
var reNotificationPeriod = regexp.MustCompile(`.*notification_period *(.+).*`)
var reContact = regexp.MustCompile(`.*contact *(.+).*`)
var reContactGroups = regexp.MustCompile(`.*contact_groups *(.+).*`)
var reCheckPeriod = regexp.MustCompile(`.*check_period *(.+).*`)
var reRegister = regexp.MustCompile(`.*register *(.+).*`)

func NewNagiosHost(defineString string) *Host {
	hostNameString := utils.FindFirstStringOrDefault(reHostName, defineString, "")
	nameString := utils.FindFirstStringOrDefault(reName, defineString, "")
	aliasString := utils.FindFirstStringOrDefault(reHostAlias, defineString, "")
	addressString := utils.FindFirstStringOrDefault(reAddress, defineString, "")

	maxCheckAttemptsInt, _ := utils.FindFirstIntOrDefault(reMaxCheckAttempts, defineString, 1)
	notificationInterval, _ := utils.FindFirstIntOrDefault(reNotificationInterval, defineString, 60)

	registerBool, _ := utils.FindFirstBooleanOrDefault(reRegister, defineString, true)

	return &Host{
		HostName:             hostNameString,
		Name:                 nameString,
		Alias:                aliasString,
		Address:              addressString,
		MaxCheckAttempts:     maxCheckAttemptsInt,
		NotificationInterval: notificationInterval,

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
