package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func readAllContactGroups(nagiosConfigDir string) (chan *model.ContactGroup, error) {

	configFiles, err := dal.GetConfigurationFies(nagiosConfigDir)
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

	return channelOutput, nil
}

func ListAllContactGroups(nagiosConfigDir string) ([]*model.ContactGroup, error) {

	channelOutput, _ := readAllContactGroups(nagiosConfigDir)

	var contactGroups []*model.ContactGroup

	for contactGroup := range channelOutput {
		contactGroups = append(contactGroups, contactGroup)
	}
	return contactGroups, nil
}

func FindContactGroupByName(nagiosConfigDir string, name string) (*model.ContactGroup, error) {

	channelOutput, _ := readAllContactGroups(nagiosConfigDir)

	for contactGroup := range channelOutput {
		if contactGroup.ContactGroupName == name {
			return contactGroup, nil
		}
	}
	return nil, nil
}
