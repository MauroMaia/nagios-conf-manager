package model

import (
	"encoding/json"
	"regexp"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type Service struct {
	//      # REQUIRED
	HostName             string `json:"host_name"`
	ServiceDescription   string `json:"service_description"`
	CheckCommand         string `json:"check_command"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
	CheckInterval        int    `json:"check_interval"`
	RetryInterval        int    `json:"retry_interval"`

	// use - Name of host template to use
	Use string `json:"use"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#host
}

var reServiceHostName = regexp.MustCompile(`.*host_name *(.+).*`)
var reServiceDescription = regexp.MustCompile(`.*service_description *(.+).*`)
var reCheckCommand = regexp.MustCompile(`.*check_command *(.+).*`)
var reNotificationsEnabled = regexp.MustCompile(`.*notifications_enabled *(.+).*`)
var reCheckInterval = regexp.MustCompile(`.*check_interval *(.+).*`)
var reRetryInterval = regexp.MustCompile(`.*retry_interval *(.+).*`)
var reServiceUse = regexp.MustCompile(`.*use *(.+).*`)

func NewNagiosService(defineString string) *Service {
	hostName := utils.FindFirstStringOrDefault(reServiceHostName, defineString, "")
	serviceDescription := utils.FindFirstStringOrDefault(reServiceDescription, defineString, "")
	checkCommand := utils.FindFirstStringOrDefault(reCheckCommand, defineString, "")
	notificationsEnabled, _ := utils.FindFirstBooleanOrDefault(reNotificationsEnabled, defineString, true)
	checkInterval, _ := utils.FindFirstIntOrDefault(reCheckInterval, defineString, 1)
	retryInterval, _ := utils.FindFirstIntOrDefault(reRetryInterval, defineString, 60)
	use := utils.FindFirstStringOrDefault(reServiceUse, defineString, "")

	return &Service{
		hostName,
		serviceDescription,
		checkCommand,
		notificationsEnabled,
		checkInterval,
		retryInterval,
		use,
	}
}

func (host *Service) String() string {
	b, err := json.Marshal(host)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (host *Service) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implementedF")
}
