package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func ListAllHostsGroups(nagiosConfigDir string) ([]*model.HostGroup, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}
	var hostGroupsList []*model.HostGroup

	chOut := make(chan *model.HostGroup, 20)
	go func() {

		var waitG sync.WaitGroup

		for _, item := range configFiles {

			waitG.Add(1)
			go dal.ReadNagiosHostGroupFromFileTask(item, chOut, &waitG)

			utils.Log.Printf("created a task to process the file %s", item)
		}

		// Wait for all threads/gorutines to stop
		waitG.Wait()
		close(chOut)
	}()

	for host := range chOut {

		hostGroupsList = append(hostGroupsList, host)
	}
	return hostGroupsList, nil
}
