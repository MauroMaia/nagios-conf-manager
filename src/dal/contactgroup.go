package dal

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

var reContactGroupName = regexp.MustCompile(`.*contactgroup_name *(.+).*`)
var reContactGroupMembers = regexp.MustCompile(`.*members *(.+).*`)
var reContactGroupAlias = regexp.MustCompile(`.*alias *(.+).*`)

func ConcurrentReadContactGroups(nagiosConfigDir string) (chan *model.ContactGroup, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.ContactGroup, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go readNagiosContactGroupFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	return channelOutput, nil
}

func readNagiosContactGroupFromFileTask(file string, outputChannel chan *model.ContactGroup, waitG *sync.WaitGroup) {
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
				outputChannel <- model.NewNagiosContactGroup(contactGroupStringToMap(define), nil)
				define = ""
				continue
			}

			if reStartContactGroup.MatchString(line) {
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

func contactGroupStringToMap(defineString string) map[string]string {

	var contactMap = make(map[string]string)

	if contactGroupNameString := utils.FindFirstStringOrDefault(reContactGroupName, defineString, ""); contactGroupNameString != "" {
		contactMap["contactGroupName"] = contactGroupNameString
	}
	if groupMembers := utils.FindFirstStringOrDefault(reContactGroupMembers, defineString, ""); groupMembers != "" {
		contactMap["members"] = groupMembers
	}
	if aliasString := utils.FindFirstStringOrDefault(reContactGroupAlias, defineString, ""); aliasString != "" {
		contactMap["alias"] = aliasString
	}
	return contactMap
}
