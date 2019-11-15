package model

import (
	"encoding/json"
	"regexp"
	"strings"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type HostGroup struct {
	//      # REQUIRED
	HostGroupName    string   `json:"hostgroup_name"`
	Alias            string   `json:"alias"`
	Members          []string `json:"members"`           // comma separated list
	HostGroupMembers []string `json:"hostgroup_members"` // comma separated list

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#hostgroup
}

var reHostGroupName = regexp.MustCompile(`.*hostgroup_name *(.+).*`)
var reMembers = regexp.MustCompile(`.*members *(.+).*`)
var reHostGroupAlias = regexp.MustCompile(`.*alias *(.+).*`)
var reHostGroupMembers = regexp.MustCompile(`.*hostgroup_members *(.+).*`)

func NewNagiosHostGroup(defineString string) *HostGroup {
	hostNameString := utils.FindFirstStringOrDefault(reHostGroupName, defineString, "")
	aliasString := utils.FindFirstStringOrDefault(reHostGroupAlias, defineString, "")
	membersString := utils.FindFirstStringOrDefault(reMembers, defineString, "")
	hostGroupMembersString := utils.FindFirstStringOrDefault(reHostGroupMembers, defineString, "")

	var members = []string{membersString}
	var hostGroupMembers = []string{hostGroupMembersString}

	if strings.Contains(membersString, ","){
		members = strings.Split(membersString, ",")
	}
	if strings.Contains(hostGroupMembersString, ","){
		hostGroupMembers = strings.Split(hostGroupMembersString, ",")
	}

	return &HostGroup{
		hostNameString,
		aliasString,
		members,
		hostGroupMembers,
	}
}

func (hostGroup *HostGroup) String() string {
	b, err := json.Marshal(hostGroup)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (hostGroup *HostGroup) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implementedF")
}

