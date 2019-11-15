package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintCommandList(commands []*model.Command) {
	fmt.Printf("%30s\t%15s\n", "CommandName", "CommandLine")
	for _, command := range commands {
		fmt.Printf("%30s\t%s\n", command.CommandName, command.CommandLine)
	}
}