package controller

import (
	"fmt"
	"regexp"
	"strings"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

var cfgDeclarationLine = regexp.MustCompile("\ncfg_file=.*")

func GetConfigurationFies(nagiosConfigDir string) ([]string, error) {
	nagiosConfigFile := nagiosConfigDir + "/nagios.cfg"
	var configFiles []string

	if utils.IsFile(nagiosConfigFile) {
		text := utils.ReadFileOrPanic(nagiosConfigFile)
		listOfConfigurationFiles := cfgDeclarationLine.FindAllString(text, -1)

		for _, file := range listOfConfigurationFiles {
			// TODO - Remove next line
			// Only for local development
			file = strings.Replace(file, "\ncfg_file=/usr/local/nagios/etc/", "/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles/", -1)
			// file = strings.Replace(file, "\ncfg_file=", "", -1)

			// Check ifthe file exist
			if utils.IsFile(file) {
				configFiles = append(configFiles, file)
			} else {
				utils.Log.Println(fmt.Sprintf("Cant't find file %s ", file))
			}

			utils.Log.Printf("Removed '^\\ncfg_file=' from configuration line line: %s", file)
		}

		return configFiles, nil
	}

	return nil, exceptions.NewConfigurationError(
		"Error reading nagios configuration directory",
	)
}
