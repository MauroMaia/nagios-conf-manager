package utils

import (
	"io/ioutil"
	"os"

	"nagios-conf-manager/src/utils/exceptions"
)

func IsDir(dirName string) bool {
	fi, err := os.Stat(dirName)
	if err != nil {
		Log.Fatalln(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}

func IsFile(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		Log.Fatalln(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return false
	case mode.IsRegular():
		return true
	}
	return false
}

func ReadFileOrPanic(fileFullPath string) string {
	Log.Println("Reading file: " + fileFullPath)
	dat, err := ioutil.ReadFile(fileFullPath)
	exceptions.CheckAndPanic(err)
	return string(dat)
}
