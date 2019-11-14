package utils

import (
	"regexp"
	"strconv"
)

func FindFirstStringOrDefault(regexp *regexp.Regexp, searchString string, defaultString string) string {
	resultArray := regexp.FindStringSubmatch(searchString)
	if resultArray != nil {
		return resultArray[1]
	}
	return defaultString
}

func FindFirstIntOrDefault(regexp *regexp.Regexp, searchString string, defaultInteger int) (int, error) {
	resultArray := regexp.FindStringSubmatch(searchString)

	if resultArray != nil {
		resultInteger, err := strconv.ParseInt(resultArray[1],10,0)
		return int(resultInteger), err
	}

	return defaultInteger, nil
}

func FindFirstBooleanOrDefault(regexp *regexp.Regexp, searchString string, defaultBoolean bool) (bool, error) {
	resultArray := regexp.FindStringSubmatch(searchString)

	if resultArray != nil {
		resultBool, err := strconv.ParseBool(resultArray[1])
		return resultBool, err
	}

	return defaultBoolean, nil
}
