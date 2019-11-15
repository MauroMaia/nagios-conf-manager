package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func ListAllService(nagiosConfigDir string) ([]*model.Service, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}
	var hostList []*model.Service

	chOut := make(chan *model.Service, 20)
	go func() {

		var waitG sync.WaitGroup

		for _, item := range configFiles {

			waitG.Add(1)
			go dal.ReadNagiosServiceFromFileTask(item, chOut, &waitG)

			utils.Log.Printf("created a task to process the file %s", item)
		}

		// Wait for all threads/goroutines to stop
		waitG.Wait()
		close(chOut)
	}()

	for host := range chOut {

		hostList = append(hostList, host)
	}
	return hostList, nil
}
