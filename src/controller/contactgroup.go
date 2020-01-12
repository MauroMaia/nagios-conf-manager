package controller

import (
	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
)

func ListAllContactGroups(nagiosConfigDir string) ([]*model.ContactGroup, error) {

	channelOutput, _ := dal.ConcurrentReadContactGroups(nagiosConfigDir)

	var contactGroups []*model.ContactGroup

	for contactGroup := range channelOutput {
		contactGroups = append(contactGroups, contactGroup)
	}
	return contactGroups, nil
}

func FindContactGroupByName(nagiosConfigDir string, name string) (*model.ContactGroup, error) {

	channelOutput, _ := dal.ConcurrentReadContactGroups(nagiosConfigDir)

	for contactGroup := range channelOutput {
		if contactGroup.ContactGroupName == name {
			return contactGroup, nil
		}
	}
	return nil, nil
}
