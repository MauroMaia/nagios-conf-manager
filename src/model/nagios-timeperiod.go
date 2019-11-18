package model

import (
	"encoding/json"
	"regexp"

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

	// FOR MOR INFORMATION CHECK: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html#timeperiod
}

var reTimePeriodsName = regexp.MustCompile(`.*name *(.+).*`)
var reTimePeriodsTimePeriodName = regexp.MustCompile(`.*timeperiod_name *(.+).*`)
var reTimePeriodsAlias = regexp.MustCompile(`.*alias *(.+).*`)

var reTimePeriodsSunday = regexp.MustCompile(`.*sunday *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)
var reTimePeriodsMonday = regexp.MustCompile(`.*monday *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)
var reTimePeriodsTuesday = regexp.MustCompile(`.*tuesday *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)
var reTimePeriodsWednesday = regexp.MustCompile(`.*wednesday *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)
var reTimePeriodsThursday = regexp.MustCompile(`.*thursday *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)
var reTimePeriodsFriday = regexp.MustCompile(`.*friday *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)
var reTimePeriodsSaturday = regexp.MustCompile(`.*saturday *([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}).*`)

func NewNagiosTimePeriods(defineString string) *TimePeriods {
	name := utils.FindFirstStringOrDefault(reTimePeriodsName, defineString, "")
	timePeriodName := utils.FindFirstStringOrDefault(reTimePeriodsTimePeriodName, defineString, "")
	alias := utils.FindFirstStringOrDefault(reTimePeriodsAlias, defineString, "")
	sunday := utils.FindFirstStringOrDefault(reTimePeriodsSunday, defineString, "")
	monday := utils.FindFirstStringOrDefault(reTimePeriodsMonday, defineString, "")
	tuesday := utils.FindFirstStringOrDefault(reTimePeriodsTuesday, defineString, "")
	wednesday := utils.FindFirstStringOrDefault(reTimePeriodsWednesday, defineString, "")
	thursday := utils.FindFirstStringOrDefault(reTimePeriodsThursday, defineString, "")
	friday := utils.FindFirstStringOrDefault(reTimePeriodsFriday, defineString, "")
	saturday := utils.FindFirstStringOrDefault(reTimePeriodsSaturday, defineString, "")

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
