package cmd

import (
	"errors"
	"fmt"
	"os"

	"nagios-conf-manager/src/utils"
)

func PrintError(err error, errorCode int8) {
	utils.Log.Printf("Printing error request with:  %s, exiting with error code: %d", err.Error(), errorCode)
	_ = fmt.Errorf("Found error: %s, exiting with error code: %d.", err.Error(), errorCode)
	os.Exit(int(errorCode))
}

func PrintErrorString(err string, errorCode int8) {
	PrintError(errors.New(err), errorCode)
}
