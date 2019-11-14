package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintHostGroupList(hostGroupList []*model.HostGroup) {
	fmt.Printf("%15s\t%15s\t%s\t\n", "HostGroupName", "Alias", "HostGroupMembers")
	for _, hostGroup := range hostGroupList {
		fmt.Printf("%15s\t%15s\t%v\n", hostGroup.HostGroupName, hostGroup.Alias, hostGroup.HostGroupMembers)
	}
}
