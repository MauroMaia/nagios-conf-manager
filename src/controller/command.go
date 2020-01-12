package controller

import (
	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
)

func ListAllCommands(nagiosConfigDir string) ([]*model.Command, error) {

	channelOutput, _ := dal.ConcurrentCommandRead(nagiosConfigDir)

	var commands []*model.Command

	for command := range channelOutput {
		commands = append(commands, command)
	}
	return commands, nil
}

func FindCommandByName(nagiosConfigDir string, name string) (*model.Command, error) {

	channelOutput, _ := dal.ConcurrentCommandRead(nagiosConfigDir)

	for command := range channelOutput {
		if command.CommandName == name {
			return command, nil
		}
	}
	return nil, nil
}
