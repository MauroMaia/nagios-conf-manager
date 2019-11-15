package model

import (
	"encoding/json"
	"regexp"
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

var reContactGroupName = regexp.MustCompile(`.*contactgroup_name *(.+).*`)
var reContactGroupMembers = regexp.MustCompile(`.*members *(.+).*`)
var reContactGroupAlias = regexp.MustCompile(`.*alias *(.+).*`)

func NewNagiosContactGroup(defineString string) *ContactGroup {
	contactNameString := utils.FindFirstStringOrDefault(reContactGroupName, defineString, "")
	groupMembersString := utils.FindFirstStringOrDefault(reContactGroupMembers, defineString, "")
	aliasString := utils.FindFirstStringOrDefault(reContactGroupAlias, defineString, "")

	var hostGroupMembers = []string{groupMembersString}

	if strings.Contains(groupMembersString, ","){
		hostGroupMembers = strings.Split(groupMembersString, ",")
	}

	return &ContactGroup{
		contactNameString,
		hostGroupMembers,
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
