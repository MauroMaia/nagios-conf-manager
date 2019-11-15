package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintContactGroupList(contactGroups []*model.ContactGroup) {
	fmt.Printf("%20s\t%20s\t%s\n", "ContactGroupName", "Display Name", "Members")
	for _, group := range contactGroups {
		fmt.Printf("%20s\t%20s\t%s\n", group.ContactGroupName, group.Alias, group.Members)
	}
}
