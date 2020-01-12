package model

import (
	"encoding/json"
	"strconv"
	"strings"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type Contact struct {
	//      # REQUIRED
	Name        string `json:"name"`
	ContactName string `json:"contact_name"`
	Use         string `json:"use"`
	Alias       string `json:"alias"`
	Email       string `json:"email"`

	HostNotificationsEnabled    bool `json:"host_notifications_enabled"`
	ServiceNotificationsEnabled bool `json:"service_notifications_enabled"`

	HostNotificationPeriod    string `json:"host_notification_period"`
	ServiceNotificationPeriod string `json:"service_notification_period"`

	HostNotificationCommand    string `json:"host_notification_commands"`
	ServiceNotificationCommand string `json:"service_notification_commands"`

	HostNotificationOptions    []string `json:"host_notification_options"`
	ServiceNotificationOptions []string `json:"service_notification_options"`

	// This information only valid to this application
	IsTemplate bool

	// templateContact *Contact `json:"-"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#contact
}

func NewNagiosContact(defineStringMap map[string]string, parent *Contact) *Contact {
	contactObjectNameString := defineStringMap["name"]

	var contactNameString, useString, aliasString, emailString, hostNotificationPeriodString, serviceNotificationPeriodString, hostNotificationCommandString, serviceNotificationCommandString string
	var isTemplate, hostNotificationsEnabledBool, serviceNotificationsEnabledBool bool
	var hostNotificationOptionsStringArray, serviceNotificationOptionsStringArray []string

	if contactName, ok := defineStringMap["contactName"]; ok == true {
		contactNameString = contactName
	} else if parent != nil && parent.Name != "" {
		contactNameString = parent.Name
	}

	if use, ok := defineStringMap["use"]; ok == true {
		useString = use
	} else if parent != nil && parent.Name != "" {
		useString = parent.Name
	}

	if use, ok := defineStringMap["email"]; ok == true {
		emailString = use
	} else if parent != nil && parent.Email != "" {
		emailString = parent.Name
	}
	if hostNotificationPeriod, ok := defineStringMap["hostNotificationPeriod"]; ok == true {
		hostNotificationPeriodString = hostNotificationPeriod
	} else if parent != nil && parent.HostNotificationPeriod != "" {
		hostNotificationPeriodString = parent.HostNotificationPeriod
	}

	if serviceNotificationPeriod, ok := defineStringMap["serviceNotificationPeriod"]; ok == true {
		serviceNotificationPeriodString = serviceNotificationPeriod
	} else if parent != nil && parent.ServiceNotificationPeriod != "" {
		serviceNotificationPeriodString = parent.ServiceNotificationPeriod
	}

	if hostNotificationCommand, ok := defineStringMap["hostNotificationCommand"]; ok == true {
		hostNotificationCommandString = hostNotificationCommand
	} else if parent != nil && parent.HostNotificationCommand != "" {
		hostNotificationCommandString = parent.HostNotificationCommand
	}

	if serviceNotificationCommand, ok := defineStringMap["serviceNotificationCommand"]; ok == true {
		serviceNotificationCommandString = serviceNotificationCommand
	} else if parent != nil && parent.ServiceNotificationCommand != "" {
		serviceNotificationCommandString = parent.ServiceNotificationCommand
	}

	if hostNotificationsEnabled, ok := defineStringMap["hostNotificationsEnabled"]; ok == true {
		hostNotificationsEnabledBool, _ = strconv.ParseBool(hostNotificationsEnabled)
	} else if parent != nil {
		hostNotificationsEnabledBool = parent.HostNotificationsEnabled
	} else {
		hostNotificationsEnabledBool = true
	}

	if serviceNotificationsEnabled, ok := defineStringMap["serviceNotificationsEnabled"]; ok == true {
		serviceNotificationsEnabledBool, _ = strconv.ParseBool(serviceNotificationsEnabled)
	} else if parent != nil {
		serviceNotificationsEnabledBool = parent.ServiceNotificationsEnabled
	} else {
		serviceNotificationsEnabledBool = true
	}

	if hostNotificationOptions, ok := defineStringMap["hostNotificationOptions"]; ok == true {
		hostNotificationOptionsStringArray = strings.Split(hostNotificationOptions, string(','))
	} else if parent != nil {
		hostNotificationOptionsStringArray = parent.HostNotificationOptions
	}
	if serviceNotificationOptions, ok := defineStringMap["serviceNotificationOptions"]; ok == true {
		serviceNotificationOptionsStringArray = strings.Split(serviceNotificationOptions, string(','))
	} else if parent != nil {
		serviceNotificationOptionsStringArray = parent.ServiceNotificationOptions
	}

	if alias, ok := defineStringMap["alias"]; ok == true {
		aliasString = alias
	}
	if register, ok := defineStringMap["register"]; ok == true {
		isTemplate, _ = strconv.ParseBool(register)
		isTemplate = !isTemplate
	} else {
		isTemplate = false
	}

	return &Contact{
		contactObjectNameString,
		contactNameString,
		useString,
		aliasString,
		emailString,
		hostNotificationsEnabledBool,
		serviceNotificationsEnabledBool,
		hostNotificationPeriodString,
		serviceNotificationPeriodString,
		hostNotificationCommandString,
		serviceNotificationCommandString,
		hostNotificationOptionsStringArray,
		serviceNotificationOptionsStringArray,
		isTemplate,
	}
}

func (command *Contact) String() string {
	b, err := json.Marshal(command)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (command *Contact) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implemented")
}

func FindContactInArray(array []*Contact, use string) *Contact {
	if len(array) > 0 {
		for _, element := range array {
			if element.Name == use {
				return element
			}
		}
	}

	return nil
}
