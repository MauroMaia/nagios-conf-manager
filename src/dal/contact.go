package dal

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#contact

var reContactName = regexp.MustCompile(`.*contact_name *(.+).*`)
var reContactAlias = regexp.MustCompile(`.*alias *(.+).*`)
var reContactHostNotificationsEnabled = regexp.MustCompile(`.*host_notifications_enabled *(.+).*`)
var reContactServiceNotificationsEnabled = regexp.MustCompile(`.*service_notifications_enabled *(.+).*`)
var reContactHostNotificationPeriod = regexp.MustCompile(`.*host_notification_period *(.+).*`)
var reContactServiceNotificationPeriod = regexp.MustCompile(`.*service_notification_period *(.+).*`)
var reContactHostNotificationOptions = regexp.MustCompile(`.*host_notification_options *(.+).*`)
var reContactServiceNotificationOptions = regexp.MustCompile(`.*service_notification_options *(.+).*`)
var reContactHostNotificationCommand = regexp.MustCompile(`.*host_notification_commands *(.+).*`)
var reContactServiceNotificationCommand = regexp.MustCompile(`.*service_notification_commands *(.+).*`)
var reContactEmail = regexp.MustCompile(`.*email *(.+).*`)

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
			go readNagiosContactFromFileTask(file, channelMapOutput, &waitGroup)

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

func readNagiosContactFromFileTask(file string, outputChannel chan map[string]string, waitG *sync.WaitGroup) {
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

	contactMap["name"] = utils.FindFirstStringOrDefault(reGenericName, defineString, "")

	if contactNameString := utils.FindFirstStringOrDefault(reContactName, defineString, ""); contactNameString != "" {
		contactMap["contactName"] = contactNameString
	}
	if useString := utils.FindFirstStringOrDefault(reGenericUse, defineString, ""); useString != "" {
		contactMap["use"] = useString
	}
	if aliasString := utils.FindFirstStringOrDefault(reContactAlias, defineString, ""); aliasString != "" {
		contactMap["alias"] = aliasString
	}
	if emailString := utils.FindFirstStringOrDefault(reContactEmail, defineString, ""); emailString != "" {
		contactMap["email"] = emailString
	}
	if registerString := utils.FindFirstStringOrDefault(reGenericRegister, defineString, ""); registerString != "" {
		contactMap["register"] = registerString
	}
	if hostNotificationsEnabled := utils.FindFirstStringOrDefault(reContactHostNotificationsEnabled, defineString, ""); hostNotificationsEnabled != "" {
		contactMap["hostNotificationsEnabled"] = hostNotificationsEnabled
	}
	if serviceNotificationsEnabled := utils.FindFirstStringOrDefault(reContactServiceNotificationsEnabled, defineString, ""); serviceNotificationsEnabled != "" {
		contactMap["serviceNotificationsEnabled"] = serviceNotificationsEnabled
	}
	if hostNotificationPeriod := utils.FindFirstStringOrDefault(reContactHostNotificationPeriod, defineString, ""); hostNotificationPeriod != "" {
		contactMap["hostNotificationPeriod"] = hostNotificationPeriod
	}
	if serviceNotificationPeriod := utils.FindFirstStringOrDefault(reContactServiceNotificationPeriod, defineString, ""); serviceNotificationPeriod != "" {
		contactMap["serviceNotificationPeriod"] = serviceNotificationPeriod
	}
	if hostNotificationOptions := utils.FindFirstStringOrDefault(reContactHostNotificationOptions, defineString, ""); hostNotificationOptions != "" {
		contactMap["hostNotificationOptions"] = hostNotificationOptions
	}
	if serviceNotificationOptions := utils.FindFirstStringOrDefault(reContactServiceNotificationOptions, defineString, ""); serviceNotificationOptions != "" {
		contactMap["serviceNotificationOptions"] = serviceNotificationOptions
	}
	if hostNotificationCommand := utils.FindFirstStringOrDefault(reContactHostNotificationCommand, defineString, ""); hostNotificationCommand != "" {
		contactMap["hostNotificationCommand"] = hostNotificationCommand
	}
	if serviceNotificationCommand := utils.FindFirstStringOrDefault(reContactServiceNotificationCommand, defineString, ""); serviceNotificationCommand != "" {
		contactMap["serviceNotificationCommand"] = serviceNotificationCommand
	}
	return contactMap
}
