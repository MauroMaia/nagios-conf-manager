package controller

import (
	"fmt"
	"os"

	"nagios-conf-manager/src/view/cmd"
)

func parseHostActions() {
	switch os.Args[3] {
	case "list":
		hosts, err := ListAllHosts("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintHostList(hosts)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{"list"}), 1)
	}
}

func parseHostGroupActions() {
	switch os.Args[3] {
	case "list":
		hostGroups, err := ListAllHostsGroups("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintHostGroupList(hostGroups)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{"list"}), 1)
	}
}

func CliParseDomain() {
	switch os.Args[2] {
	case "host":
		parseHostActions()
		break
	case "hostgroup":
		parseHostGroupActions()
		break
	case "commands":
		parseCommandsActions()
		break
	default:
		cmd.PrintErrorString(fmt.Sprintf("Cli expected %v subcommands", []string{
			"host",
			"hostgroup",
			"command",
		}), 1)
	}
}

func parseCommandsActions() {
	switch os.Args[3] {
	case "list":
		commands, err := ListAllCommands("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintCommandList(commands)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{"list"}), 1)
	}
}