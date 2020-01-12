package model

import (
	"encoding/json"
	"strings"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type ContactGroup struct {
	//      # REQUIRED
	ContactGroupName string   `json:"contactgroup_name"`
	Members          []string `json:"members"`
	Alias            string   `json:"alias"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#contact
}

func NewNagiosContactGroup(defineStringMap map[string]string, parent *ContactGroup) *ContactGroup {
	var groupMembers []string
	var contactNameString, aliasString string

	if members, ok := defineStringMap["members"]; ok == true {
		groupMembers = []string{members}

		if strings.Contains(members, ",") {
			groupMembers = strings.Split(members, ",")
		}
	}
	if alias, ok := defineStringMap["alias"]; ok == true {
		aliasString = alias
	}
	if contactName, ok := defineStringMap["contactGroupName"]; ok == true {
		contactNameString = contactName
	}

	return &ContactGroup{
		contactNameString,
		groupMembers,
		aliasString,
	}
}

func (command *ContactGroup) String() string {
	b, err := json.Marshal(command)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (command *ContactGroup) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implemented")
}
