package controller

import (
	"fmt"
	"os"

	"nagios-conf-manager/src/view/cmd"
)

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
	case "contacts":
		parseContactsActions()
		break
	case "contactgroups":
		parseContactGroupsActions()
		break
	case "services":
		parseServicesActions()
		break
	case "timeperiods":
		parseTimePeriodActions()
		break
	default:
		cmd.PrintErrorString(fmt.Sprintf("Cli expected %v subcommands", []string{
			"host",
			"hostgroup",
			"commands",
			"contacts",
			"contactgroups",
			"services",
			"timeperiods",
		}), 1)
	}
}

func parseTimePeriodActions() {
	switch os.Args[3] {
	case "list":
		timePeriods, err := ListAllTimePeriods(os.Getenv("NAGIOS_BASE_PATH"))
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintTimePeriodsList(timePeriods)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{"list"}), 1)
	}
}

func parseServicesActions() {
	switch os.Args[3] {
	case "list":
		hosts, err := ListAllService(os.Getenv("NAGIOS_BASE_PATH"))
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintServiceList(hosts)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{"list"}), 1)
	}
}

func parseHostActions() {
	switch os.Args[3] {
	case "list":
		hosts, err := ListAllHosts(os.Getenv("NAGIOS_BASE_PATH"))
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
		hostGroups, err := ListAllHostGroups(os.Getenv("NAGIOS_BASE_PATH"))
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintHostGroupList(hostGroups)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{"list"}), 1)
	}
}

func parseCommandsActions() {
	switch os.Args[3] {
	case "list":
		commands, err := ListAllCommands(os.Getenv("NAGIOS_BASE_PATH"))
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintCommandList(commands)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{"list"}), 1)
	}
}

func parseContactGroupsActions() {
	switch os.Args[3] {
	case "list":
		contactGroups, err := ListAllContactGroups(os.Getenv("NAGIOS_BASE_PATH"))
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintContactGroupList(contactGroups)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{
			"list",
		}), 1)
	}
}

func parseContactsActions() {
	switch os.Args[3] {
	case "list":
		contacts, err := ListAllContacts(os.Getenv("NAGIOS_BASE_PATH"))
		if err != nil {
			cmd.PrintError(err, 2)
		} else {
			cmd.PrintContactList(contacts)
		}
	default:
		cmd.PrintErrorString(fmt.Sprintf("'cli host' expected %v subcommands", []string{
			"list",
		}), 1)
	}
}
