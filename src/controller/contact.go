package controller

import (
	"nagios-conf-manager/src/dal"
	"nagios-conf-manager/src/model"
)

func ListAllContacts(nagiosConfigDir string) ([]model.Contact, error) {

	channelOutput, _ := dal.ConcurrentContactRead(nagiosConfigDir)

	var contacts []model.Contact

	for contact := range channelOutput {
		contacts = append(contacts, *contact)
	}
	return contacts, nil
}

func FindContactByName(nagiosConfigDir string, name string) (*model.Contact, error) {

	channelOutput, _ := dal.ConcurrentContactRead(nagiosConfigDir)

	for contact := range channelOutput {
		if contact.ContactName == name {
			return contact, nil
		}
	}
	return nil, nil
}
