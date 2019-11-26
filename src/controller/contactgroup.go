package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func ListAllContactGroups(nagiosConfigDir string) ([]*model.ContactGroup, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}
	var contactGroups []*model.ContactGroup

	channelOutput := make(chan *model.ContactGroup, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosContactGroupFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	for contactGroup := range channelOutput {
		contactGroups = append(contactGroups, contactGroup)
	}
	return contactGroups, nil
}

func FindContactGroupByName(nagiosConfigDir string, name string) (*model.ContactGroup, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.ContactGroup, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosContactGroupFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	for contactGroup := range channelOutput {
		if contactGroup.ContactGroupName == name {
			return contactGroup, nil
		}
	}
	return nil, nil
}
