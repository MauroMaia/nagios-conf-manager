package dal

import (
	"fmt"
	"strings"
	"sync"

	"nagios-conf-manager/src/dal/fs"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func ReadNagiosTimePeriodsFromFileTask(file string, outputChannel chan *model.TimePeriods, waitG *sync.WaitGroup) {

	defer waitG.Done()

	tpf, err := fs.NewConcurrentFile(file)
	if err != nil {
		utils.Log.Println(fmt.Sprintf("Cant't find file %s ", file))
		return
	}

	text, err := tpf.ReadFile()
	if err != nil {
		utils.Log.Println(fmt.Sprintf("Cant't find file %s error %s", file, err))
		return
	}

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
			defineArray := strings.Split(define, string('\n'))
			defineArray = defineArray[1 : len(defineArray)-1]
			outputChannel <- model.NewNagiosTimePeriods(strings.Join(defineArray, "\n"))
			define = ""
			continue
		}

		if reStartTimePeriod.MatchString(line) {
			define = line
		} else if strings.Compare(define, "") > 0 {
			define += "\n"
			define += line
		}
	}

}
