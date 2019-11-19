package model

import (
	"encoding/json"
	"regexp"
	"strings"

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

func NewNagiosTimePeriods(defineString string) *TimePeriods {
	name := utils.FindFirstStringOrDefault(reTimePeriodsName, defineString, "")
	timePeriodName := utils.FindFirstStringOrDefault(reTimePeriodsTimePeriodName, defineString, "")
	alias := utils.FindFirstStringOrDefault(reTimePeriodsAlias, defineString, "")

	resultArray := reLine.FindAllStringSubmatch(defineString, -1)

	var sunday, monday, tuesday, wednesday, thursday, friday, saturday string
	var other map[string]string = make(map[string]string)
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

func (host *TimePeriods) String() string {
	b, err := json.Marshal(host)
	if err != nil {
		utils.Log.Println(err)
		utils.Log.Println("Returning empty line")
		return ""
	}
	return string(b)
}

func (host *TimePeriods) FormatToNagiosCFG() (string, error) {
	return "", exceptions.NewErrorGeneratingNagiosObjDefinition("Function not implementedF")
}
