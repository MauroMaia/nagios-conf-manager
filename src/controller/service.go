package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func readAllService(nagiosConfigDir string) (chan *model.Service, error) {

	configFiles, err := dal.GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.Service, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosServiceFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("Created a task to process the file: %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	return channelOutput, nil
}

func ListAllService(nagiosConfigDir string) ([]*model.Service, error) {

	var services []*model.Service

	channelOutput,_ := readAllService(nagiosConfigDir)

	for service := range channelOutput {
		services = append(services, service)
	}
	return services, nil
}

func FindServiceByName(nagiosConfigDir string, name string) (*model.Service, error) {

	channelOutput,_ := readAllService(nagiosConfigDir)

	for service := range channelOutput {
		if service.ServiceDescription == name {
			return service, nil
		}
	}
	return nil, nil
}
