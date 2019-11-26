package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintContactList(contacts []model.Contact) {
	fmt.Printf("%20s\t%20s\t%20s\t%20s\t%20s\t%v\n", "Name","ContactName", "Display Name","Email","Template","IsTemplate")
	for _, contact := range contacts {
		fmt.Printf("%20s\t%20s\t%20s\t%20s\t%20s\t%v\n", contact.Name,contact.ContactName, contact.Alias,contact.Email,contact.Use,contact.IsTemplate)
	}
}