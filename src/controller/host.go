package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func ListAllHosts(nagiosConfigDir string) ([]*model.Host, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	var hostList []*model.Host

	channelOutput := make(chan *model.Host, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, item := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosHostFromFileTask(item, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", item)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	for host := range channelOutput {
		hostList = append(hostList, host)
	}
	return hostList, nil
}

func FindHostByName(nagiosConfigDir string, name string) (*model.Host, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.Host, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, item := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosHostFromFileTask(item, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", item)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	for host := range channelOutput {
		if host.Name == name {
			return host, nil
		}
	}
	return nil, nil
}
