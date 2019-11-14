package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintHostList(hostList []*model.Host) {
	fmt.Printf("%15s\t%15s\tTempalte\tHostName\n", "Address", "Name")
	for _, host := range hostList {
		fmt.Printf("%15s\t%15s\t%t\t\t%s\n", host.Address, host.Name, !host.Register, "host.HostName")
	}
}
