package controller

import (
	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
)

func ListAllHosts(nagiosConfigDir string) ([]*model.Host, error) {

	var hostList []*model.Host

	channelOutput, _ := dal.ConcurrentReadHosts(nagiosConfigDir)

	for host := range channelOutput {
		hostList = append(hostList, host)
	}
	return hostList, nil
}

func FindHostByName(nagiosConfigDir string, name string) (*model.Host, error) {

	channelOutput, _ := dal.ConcurrentReadHosts(nagiosConfigDir)

	for host := range channelOutput {
		if host.Name == name {
			return host, nil
		}
	}
	return nil, nil
}
