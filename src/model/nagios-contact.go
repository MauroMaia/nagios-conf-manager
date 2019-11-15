package model

import (
	"encoding/json"
	"regexp"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type Contact struct {
	//      # REQUIRED
	ContactName string `json:"contact_name"`
	Use         string `json:"use"`
	Alias       string `json:"alias"`
	Email       string `json:"email"`

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#contact
}

var reContactName = regexp.MustCompile(`.*contact_name *(.+).*`)
var reContactUse = regexp.MustCompile(`.*use *(.+).*`)
var reContactAlias = regexp.MustCompile(`.*alias *(.+).*`)
var reContactEmail = regexp.MustCompile(`.*email *(.+).*`)

func NewNagiosContact(defineString string) *Contact {
	contactNameString := utils.FindFirstStringOrDefault(reContactName, defineString, "")
	useString := utils.FindFirstStringOrDefault(reContactUse, defineString, "")
	aliasString := utils.FindFirstStringOrDefault(reContactAlias, defineString, "")
	emailString := utils.FindFirstStringOrDefault(reContactEmail, defineString, "")

	return &Contact{
		contactNameString,
		useString,
		aliasString,
		emailString,
	}
}

func (command *Contact) String() string {
	b, err := json.Marshal(command)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (command *Contact) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implemented")
}
