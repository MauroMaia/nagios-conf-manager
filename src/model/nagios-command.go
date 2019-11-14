package model

import (
	"encoding/json"
	"regexp"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type Command struct {
	//      # REQUIRED
	CommandName string `json:"command_name"`
	CommandLine string `json:"command_line"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#command
}

var reCommandName = regexp.MustCompile(`.*command_name *(.+).*`)
var reCommandLine = regexp.MustCompile(`.*command_line *(.+).*`)

func NewNagiosCommand(defineString string) *Command {
	hostNameString := utils.FindFirstStringOrDefault(reCommandName, defineString, "")
	nameString := utils.FindFirstStringOrDefault(reCommandLine, defineString, "")

	return &Command{
		hostNameString,
		nameString,
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
