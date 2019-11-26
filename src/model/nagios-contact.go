package model

import (
	"encoding/json"
	"strconv"

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
	/*  Missing fields

		host_notifications_enabled	[0/1]
	service_notifications_enabled	[0/1]

	host_notification_period	timeperiod_name
	service_notification_period	timeperiod_name

	host_notification_options	[d,u,r,f,s,n]
	service_notification_options	[w,u,c,r,f,s,n]

	host_notification_commands	command_name
	service_notification_commands	command_name

	*/

	// This information only valid to this application
	IsTemplate bool

	// templateContact *Contact `json:"-"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#contact
}

func NewNagiosContact(defineStringMap map[string]string, parent *Contact) *Contact {
	contactObjectNameString := defineStringMap["name"]

	var contactNameString, useString, aliasString, emailString string
	var isTemplate bool

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
