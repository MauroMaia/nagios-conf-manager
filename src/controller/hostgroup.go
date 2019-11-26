package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func ListAllHostGroups(nagiosConfigDir string) ([]*model.HostGroup, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	var hostGroupsList []*model.HostGroup

	channelOutput := make(chan *model.HostGroup, 20)
	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosHostGroupFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/gorutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	for hostGroup := range channelOutput {
		hostGroupsList = append(hostGroupsList, hostGroup)
	}
	return hostGroupsList, nil
}

func FindHostGroupByName(nagiosConfigDir string, name string) (*model.HostGroup, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.HostGroup, 20)
	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosHostGroupFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/gorutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	for hostGroup := range channelOutput {
		if hostGroup.HostGroupName == name {
			return hostGroup, nil
		}
	}
	return nil, nil
}
