package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintServiceList(services []*model.Service) {
	fmt.Printf("%25s\t%25s\t%s\n", "ServiceDescription","HostName", "CheckCommand")
	for _, service := range services {
		fmt.Printf("%25s\t%25s\t%s\n", service.ServiceDescription, service.HostName, service.CheckCommand)
	}
}