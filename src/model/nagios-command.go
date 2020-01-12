package model

import (
	"encoding/json"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type Command struct {
	//      # REQUIRED
	Name        string `json:"name"`
	CommandName string `json:"command_name"`
	CommandLine string `json:"command_line"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#command
}

func NewNagiosCommand(defineStringMap map[string]string) *Command {

	var nameString, commandLineString, commandNameString string

	if contactName, ok := defineStringMap["name"]; ok == true {
		nameString = contactName
	}
	if contactName, ok := defineStringMap["commandName"]; ok == true {
		commandNameString = contactName
	}
	if use, ok := defineStringMap["commandLine"]; ok == true {
		commandLineString = use
	}

	return &Command{
		nameString,
		commandNameString,
		commandLineString,
	}
}

func (command *Command) String() string {
	b, err := json.Marshal(command)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (command *Command) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implemented")
}
