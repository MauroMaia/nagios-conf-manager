package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func readAllCommands(nagiosConfigDir string) (chan *model.Command, error) {
	configFiles, err := dal.GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.Command, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosCommandFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	return channelOutput, nil
}

func ListAllCommands(nagiosConfigDir string) ([]*model.Command, error) {

	channelOutput, _ := readAllCommands(nagiosConfigDir)

	var commands []*model.Command

	for command := range channelOutput {
		commands = append(commands, command)
	}
	return commands, nil
}

func FindCommandByName(nagiosConfigDir string, name string) (*model.Command, error) {

	channelOutput, _ := readAllCommands(nagiosConfigDir)

	for command := range channelOutput {
		if command.CommandName == name {
			return command, nil
		}
	}
	return nil, nil
}
