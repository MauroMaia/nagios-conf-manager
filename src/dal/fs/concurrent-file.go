package fs

import (
	"errors"
	"io/ioutil"
	"sync"

	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

type ConcurrentFile struct {
	FileName string
	mutex    sync.RWMutex
}

func NewConcurrentFile(fileName string) (*ConcurrentFile, error) {
	if ! utils.IsFile(fileName) {
		return nil, errors.New("file is not a file")
	}

	return &ConcurrentFile{
		FileName: fileName,
	}, nil
}

func (concurrentFile *ConcurrentFile) ReadFile() (string, error) {
	concurrentFile.mutex.RLock()
	defer concurrentFile.mutex.RUnlock()

	utils.Log.Println("Reading file: " + concurrentFile.FileName)
	dat, err := ioutil.ReadFile(concurrentFile.FileName)
	return string(dat), err
}

func (concurrentFile *ConcurrentFile) AddNewToFile(timePeriod *model.TimePeriods) {
	concurrentFile.mutex.Lock()
	defer concurrentFile.mutex.Unlock()

}

func (concurrentFile *ConcurrentFile) ReplaceNewToFile(timePeriod *model.TimePeriods) {
	concurrentFile.mutex.Lock()
	defer concurrentFile.mutex.Unlock()

}

func (concurrentFile *ConcurrentFile) DeleteToFile(timePeriod *model.TimePeriods) {
	concurrentFile.mutex.Lock()
	defer concurrentFile.mutex.Unlock()

}