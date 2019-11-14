package dal

import (
	"fmt"
	"strings"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func ReadNagiosHostFromFileTask(file string, outputChannel chan *model.Host, waitG *sync.WaitGroup) {
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
			if reEndDefineStatment.MatchString(line) && strings.Compare(define, "") > 0 {
				define += "\n"
				define += line
				outputChannel <- model.NewNagiosHost(define)
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
