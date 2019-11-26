package dal

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

var reContactObjectName = regexp.MustCompile(`.*name *(.+).*`)
var reContactName = regexp.MustCompile(`.*contact_name *(.+).*`)
var reContactUse = regexp.MustCompile(`.*use *(.+).*`)
var reContactAlias = regexp.MustCompile(`.*alias *(.+).*`)
var reContactEmail = regexp.MustCompile(`.*email *(.+).*`)
var reContactRegister = regexp.MustCompile(`.*register *(.+).*`)

func ConcurrentContactRead(nagiosConfigDir string) (chan *model.Contact, error) {
	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelMapOutput := make(chan map[string]string, 20)
	channelOutput := make(chan *model.Contact, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go ReadNagiosContactFromFileTask(file, channelMapOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelMapOutput)
	}()

	go func() {
		var contactsMaps []map[string]string
		var contacts []*model.Contact
		for contactMap := range channelMapOutput {
			if useName, ok := contactMap["use"]; ok == true && useName != "" {
				if parent := model.FindContactInArray(contacts, useName); parent != nil {
					modelContact := model.NewNagiosContact(contactMap, parent)
					channelOutput <- modelContact
					contacts = append(contacts, modelContact)
					continue
				}
			} else if _, ok := contactMap["use"]; ok == false {
				modelContact := model.NewNagiosContact(contactMap, nil)
				channelOutput <- modelContact
				contacts = append(contacts, modelContact)
				continue
			}

			contactsMaps = append(contactsMaps, contactMap)
		}

		for _, contactMap := range contactsMaps {
			modelContact := model.NewNagiosContact(contactMap, model.FindContactInArray(contacts, contactMap["use"]))
			channelOutput <- modelContact
			contacts = append(contacts, modelContact)
		}

		close(channelOutput)
	}()

	return channelOutput, nil
}

func ReadNagiosContactFromFileTask(file string, outputChannel chan map[string]string, waitG *sync.WaitGroup) {
	if utils.IsFile(file) {

		text := utils.ReadFileOrPanic(file)
		linesOfText := strings.Split(text, string('\n'))
		define := ""

		for _, line := range linesOfText {
			// Remove inline comments
			line = RemoveComments(line)

			// Ignore comments and empty lines
			if IgnoreConfigLine(line) {
				continue
			}

			// end of 'define' group
			if reEndDefineStatement.MatchString(line) && strings.Compare(define, "") > 0 {
				define += "\n"
				define += line
				outputChannel <- contactStringToMap(define)
				define = ""
				continue
			}

			if reStartContact.MatchString(line) && !reStartContactGroup.MatchString(line) {
				define = line
			} else if strings.Compare(define, "") > 0 {
				define += "\n"
				define += line
			}
		}
	} else {

		utils.Log.Println(fmt.Sprintf("Cant't find file %s ", file))
		/*exceptions.NewConfigurationError(
			fmt.Sprintf("Cant't find file %s ", file),
		)*/
	}

	waitG.Done()
}

func contactStringToMap(defineString string) map[string]string {

	var contactMap = make(map[string]string)

	contactObjectNameString := utils.FindFirstStringOrDefault(reContactObjectName, defineString, "")
	contactNameString := utils.FindFirstStringOrDefault(reContactName, defineString, "")
	useString := utils.FindFirstStringOrDefault(reContactUse, defineString, "")
	aliasString := utils.FindFirstStringOrDefault(reContactAlias, defineString, "")
	emailString := utils.FindFirstStringOrDefault(reContactEmail, defineString, "")
	registerString := utils.FindFirstStringOrDefault(reContactRegister, defineString, "")

	contactMap["name"] = contactObjectNameString
	if contactNameString != "" {
		contactMap["contactName"] = contactNameString
	}
	if useString != "" {
		contactMap["use"] = useString
	}
	if aliasString != "" {
		contactMap["alias"] = aliasString
	}
	if emailString != "" {
		contactMap["email"] = emailString
	}
	if registerString != "" {
		contactMap["register"] = registerString
	}
	// if _, ok := contactMap["use"]; ok == true {
	return contactMap
}
