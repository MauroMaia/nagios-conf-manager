package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintContactList(contacts []*model.Contact) {
	fmt.Printf("%20s\t%20s\t%20s\t%20s\n", "ContactName", "Display Name","Email","Template")
	for _, contact := range contacts {
		fmt.Printf("%20s\t%20s\t%20s\t%20s\n", contact.ContactName, contact.Alias,contact.Email,contact.Use)
	}
}