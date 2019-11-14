package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
)

var (
	Log *log.Logger
)

func init() {
	// set location of log file
	var logPath = "/home/mauro.maia/go/src/nagios-conf-manager/app.log"

	flag.Parse()

	var file, err1 = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}

	if len(os.Args) == 1 {
		log.Printf(fmt.Sprintf("Cli expected %v subcommands", []string{"cli", "webserver"}))
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cli":
		currentUser, err := user.Current()
		if err != nil {
			log.Fatalln(err)
		}
		Log = log.New(file, currentUser.Username+" ", log.LstdFlags|log.Lshortfile|log.LUTC)
	default:
		Log = log.New(file, "", log.Lshortfile)
	}
}
