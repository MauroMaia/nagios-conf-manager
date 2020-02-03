package fs

import (
	"errors"
	"io/ioutil"
	"strings"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

type TimePeriodConcurrentFile struct {
	FileName string
	mutex    sync.RWMutex
}

func NewTimePeriodConcurrentFile(fileName string) (*TimePeriodConcurrentFile, error) {
	if ! utils.IsFile(fileName) {
		return nil, errors.New("file is not a file")
	}

	return &TimePeriodConcurrentFile{
		FileName: fileName,
	}, nil
}

func (timePeriodConcurrentFile *TimePeriodConcurrentFile) ReadFile() ([]string, error) {
	timePeriodConcurrentFile.mutex.RLock()
	defer timePeriodConcurrentFile.mutex.RUnlock()

	utils.Log.Println("Reading file: " + timePeriodConcurrentFile.FileName)
	fileRawData, err := ioutil.ReadFile(timePeriodConcurrentFile.FileName)
	if err != nil {
		utils.Log.Printf("Cant't read file %s error %s\n", timePeriodConcurrentFile.FileName, err)
		return nil, err
	}
	linesOfText := strings.Split(string(fileRawData), string('\n'))
	return linesOfText, nil
}

func (timePeriodConcurrentFile *TimePeriodConcurrentFile) AddNewToFile(timePeriod *model.TimePeriods) {
	timePeriodConcurrentFile.mutex.Lock()
	defer timePeriodConcurrentFile.mutex.Unlock()

}

func (timePeriodConcurrentFile *TimePeriodConcurrentFile) ReplaceNewToFile(timePeriod *model.TimePeriods) {
	timePeriodConcurrentFile.mutex.Lock()
	defer timePeriodConcurrentFile.mutex.Unlock()

}

func (timePeriodConcurrentFile *TimePeriodConcurrentFile) DeleteToFile(timePeriod *model.TimePeriods) {
	timePeriodConcurrentFile.mutex.Lock()
	defer timePeriodConcurrentFile.mutex.Unlock()

}
