package controller

import (
	"sync"

	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
	"nagios-conf-manager/src/utils"
)

func readAllContacts(nagiosConfigDir string) (chan *model.Contact, error) {

	configFiles, err := GetConfigurationFies(nagiosConfigDir)
	if err != nil {
		return nil, err
	}

	channelOutput := make(chan *model.Contact, 20)

	go func() {
		var waitGroup sync.WaitGroup

		for _, file := range configFiles {
			waitGroup.Add(1)
			go dal.ReadNagiosContactFromFileTask(file, channelOutput, &waitGroup)

			utils.Log.Printf("created a task to process the file %s", file)
		}

		// Wait for all threads/goroutines to stop
		waitGroup.Wait()
		close(channelOutput)
	}()

	return channelOutput, nil
}

func ListAllContacts(nagiosConfigDir string) ([]*model.Contact, error) {

	channelOutput, _ := readAllContacts(nagiosConfigDir)

	var contacts []*model.Contact

	for contact := range channelOutput {
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func FindContactByName(nagiosConfigDir string, name string) (*model.Contact, error) {

	channelOutput, _ := readAllContacts(nagiosConfigDir)

	for contact := range channelOutput {
		if contact.ContactName == name {
			return contact, nil
		}
	}
	return nil, nil
}
