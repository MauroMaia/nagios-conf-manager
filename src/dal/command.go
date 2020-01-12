package dal

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

var reCommandName = regexp.MustCompile(`.*command_name *(.+).*`)
var reCommandLine = regexp.MustCompile(`.*command_line *(.+).*`)

func ConcurrentCommandRead(nagiosConfigDir string) (chan *model.Command, error) {
	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.Command, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go readNagiosCommandFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	return channelOutput, nil
}

func readNagiosCommandFromFileTask(file string, outputChannel chan *model.Command, waitG *sync.WaitGroup) {
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
				outputChannel <- model.NewNagiosCommand(commandStringToMap(define))
				define = ""
				continue
			}

			if reStartCommand.MatchString(line) {
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

func commandStringToMap(defineString string) map[string]string {

	var commandMap = make(map[string]string)

	commandMap["name"] = utils.FindFirstStringOrDefault(reGenericName, defineString, "")

	if contactNameString := utils.FindFirstStringOrDefault(reCommandName, defineString, ""); contactNameString != "" {
		commandMap["commandName"] = contactNameString
	}
	if useString := utils.FindFirstStringOrDefault(reGenericUse, defineString, ""); useString != "" {
		commandMap["use"] = useString
	}
	if aliasString := utils.FindFirstStringOrDefault(reCommandLine, defineString, ""); aliasString != "" {
		commandMap["commandLine"] = aliasString
	}
	if registerString := utils.FindFirstStringOrDefault(reGenericRegister, defineString, ""); registerString != "" {
		commandMap["register"] = registerString
	}
	return commandMap
}
