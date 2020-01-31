package model

import (
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"nagios-conf-manager/src/utils"
	"nagios-conf-manager/src/utils/exceptions"
)

type TimePeriods struct {
	//      # REQUIRED
	Name           string `json:"name"`
	TimePeriodName string `json:"timeperiod_name"`
	Alias          string `json:"alias"`

	Sunday    string `json:"sunday"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`

	Other map[string]string

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#timeperiod
}

var reTimePeriodsName = regexp.MustCompile(`.*name *(.+).*`)
var reTimePeriodsTimePeriodName = regexp.MustCompile(`.*timeperiod_name *(.+).*`)
var reTimePeriodsAlias = regexp.MustCompile(`.*alias *(.+).*`)

var reLine = regexp.MustCompile(` *(.+) *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)
var reTime = regexp.MustCompile(`([0-9]{2}:[0-9]{2})-([0-9]{2}:[0-9]{2}).*`)

func NewNagiosTimePeriods(defineString string) *TimePeriods {
	name := utils.FindFirstStringOrDefault(reTimePeriodsName, defineString, "")
	timePeriodName := utils.FindFirstStringOrDefault(reTimePeriodsTimePeriodName, defineString, "")
	alias := utils.FindFirstStringOrDefault(reTimePeriodsAlias, defineString, "")

	resultArray := reLine.FindAllStringSubmatch(defineString, -1)

	var sunday, monday, tuesday, wednesday, thursday, friday, saturday string
	var other = make(map[string]string)
	for _, matchElement := range resultArray {
		matchElement[1] = strings.TrimSpace(matchElement[1])
		switch matchElement[1] {
		case "sunday":
			sunday = matchElement[2]
			break
		case "monday":
			monday = matchElement[2]
			break
		case "tuesday":
			tuesday = matchElement[2]
			break
		case "wednesday":
			wednesday = matchElement[2]
			break
		case "thursday":
			thursday = matchElement[2]
			break
		case "friday":
			friday = matchElement[2]
			break
		case "saturday":
			saturday = matchElement[2]
			break
		default:
			other[matchElement[1]] = matchElement[2]
		}
	}

	return &TimePeriods{
		name,
		timePeriodName,
		alias,
		sunday,
		monday,
		tuesday,
		wednesday,
		thursday,
		friday,
		saturday,
		other,
	}
}

func Validate(timePeriod *TimePeriods) bool {
	// TODO - validate the other

	if timePeriod.Name == "" || timePeriod.Alias == "" || strings.Contains(timePeriod.Name, string(' ')) {
		return false
	}

	var toCheck = []string{
		timePeriod.Monday,
		timePeriod.Tuesday,
		timePeriod.Wednesday,
		timePeriod.Thursday,
		timePeriod.Friday,
		timePeriod.Saturday,
		timePeriod.Sunday,
	}

	for _, element := range toCheck {
		if ! reTime.MatchString(element) {
			// TODO: comment here
			return false
		}

		resultArray := reTime.FindAllStringSubmatch(element, -1)
		if len(resultArray[0]) != 3 {
			// TODO: comment here
			return false
		}

		startTime, _ := time.Parse("15:04", resultArray[0][1])
		endTime, _ := time.Parse("15:04", resultArray[0][2])
		if endTime.Before(startTime) || endTime.Equal(startTime) {
			// TODO: comment here
			return false
		}
	}
	return true
}

func (timePeriod *TimePeriods) String() string {
	b, err := json.Marshal(timePeriod)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (timePeriod *TimePeriods) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implementedF")
}
