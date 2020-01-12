package dal

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

var reHostName = regexp.MustCompile(`.*hostName *(.+).*`)
var reName = regexp.MustCompile(`.*name *(.+).*`)
var reHostAlias = regexp.MustCompile(`.*alias *(.+).*`)
var reAddress = regexp.MustCompile(`.*address *(.+).*`)
var reMaxCheckAttempts = regexp.MustCompile(`.*max_check_attempts *(.+).*`)
var reNotificationInterval = regexp.MustCompile(`.*notification_interval *(.+).*`)
var reNotificationPeriod = regexp.MustCompile(`.*notification_period *(.+).*`)
var reContact = regexp.MustCompile(`.*contact *(.+).*`)
var reContactGroups = regexp.MustCompile(`.*contact_groups *(.+).*`)
var reCheckPeriod = regexp.MustCompile(`.*check_period *(.+).*`)
var reRegister = regexp.MustCompile(`.*register *(.+).*`)

func ConcurrentReadHosts(nagiosConfigDir string) (chan *model.Host, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.Host, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, item := range configFiles {
			waitGroup.Add(1)
			go readNagiosHostFromFileTask(item, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", item)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	return channelOutput, nil
}

func readNagiosHostFromFileTask(file string, outputChannel chan *model.Host, waitG *sync.WaitGroup) {
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
				outputChannel <- model.NewNagiosHost(hostStringToMap(define))
				define = ""
				continue
			}

			if reStartHost.MatchString(line) && !reStartHostGroup.MatchString(line) {
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

func hostStringToMap(defineString string) map[string]string {

	var contactMap = make(map[string]string)

	contactMap["name"] = utils.FindFirstStringOrDefault(reGenericName, defineString, "")

	if contactNameString := utils.FindFirstStringOrDefault(reHostName, defineString, ""); contactNameString != "" {
		contactMap["hostName"] = contactNameString
	}
	if useString := utils.FindFirstStringOrDefault(reHostAlias, defineString, ""); useString != "" {
		contactMap["alias"] = useString
	}
	if useString := utils.FindFirstStringOrDefault(reAddress, defineString, ""); useString != "" {
		contactMap["address"] = useString
	}
	if useString := utils.FindFirstStringOrDefault(reMaxCheckAttempts, defineString, ""); useString != "" {
		contactMap["max_check_attempts"] = useString
	}
	if useString := utils.FindFirstStringOrDefault(reNotificationInterval, defineString, ""); useString != "" {
		contactMap["notification_interval"] = useString
	}
	if useString := utils.FindFirstStringOrDefault(reRegister, defineString, ""); useString != "" {
		contactMap["register"] = useString
	}

	return contactMap
}

