package controller

import (
	"errors"
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func readAllTimePeriods(nagiosConfigDir string) (chan *model.TimePeriods, error) {

	configFiles, err := dal.GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.TimePeriods, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosTimePeriodsFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	return channelOutput, nil
}

func ListAllTimePeriods(nagiosConfigDir string) ([]*model.TimePeriods, error) {

	channelOutput, _ := readAllTimePeriods(nagiosConfigDir)

	var timePeriods []*model.TimePeriods

	for timePeriod := range channelOutput {
		timePeriods = append(timePeriods, timePeriod)
	}
	return timePeriods, nil
}

func FindTimePeriodByName(nagiosConfigDir string, name string) (*model.TimePeriods, error) {

	channelOutput, _ := readAllTimePeriods(nagiosConfigDir)

	// TODO - change/optimise search algorithm
	for timePeriod := range channelOutput {
		if timePeriod.Name == name {
			return timePeriod, nil
		}
	}
	return nil, nil
}

func CreateNewTimePeriod(nagiosConfigDir string, newTimePeriod model.TimePeriods) (*model.TimePeriods, error) {
	channelOutput, err := readAllTimePeriods(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	for timePeriod := range channelOutput {
		if timePeriod.Name == newTimePeriod.Name {
			utils.Log.Printf("Timeperiod with name/id: %s, already exist.", timePeriod.Name)
			return nil, errors.New("duplicated")
		}
	}

	if !model.Validate(&newTimePeriod) {
		return nil, errors.New("invalid object")
	}

	return &newTimePeriod, nil
}
