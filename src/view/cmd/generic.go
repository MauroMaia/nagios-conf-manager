package cmd

import (
	"errors"
	"fmt"
	"os"

	"nagios-conf-manager/src/utils"
)

/*func init(){
	f, err := os.OpenFile("/home/mauro.maia/go/src/nagios-conf-manager/log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetPrefix("cmd/generic")
	log.Println("Init generic")
}*/

func PrintError(err error, errorCode int8) {
	utils.Log.Printf("Printing error request with:  %s, exiting with error code: %i", err.Error(), errorCode)
	_ = fmt.Errorf("Found error: %s, exiting with error code: %i.", err.Error(), errorCode)
	os.Exit(int(errorCode))
}

func PrintErrorString(err string, errorCode int8) {
	PrintError(errors.New(err), errorCode)
}
