package cmd

import (
	"fmt"

	"nagios-conf-manager/src/model"
)

func PrintTimePeriodsList(timePeriods []*model.TimePeriods) {
	fmt.Printf("%20s\t%20s\t%12s\t%12s\t%12s\t%12s\t%12s\t%12s\t%12s\t%s\n", "TimePeriodName", "Name", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday","Other")
	for _, period := range timePeriods {
		fmt.Printf("%20s\t%20s\t%12s\t%12s\t%12s\t%12s\t%12s\t%12s\t%12s\t%s\n", period.TimePeriodName, period.Name, period.Sunday, period.Monday, period.Tuesday, period.Wednesday, period.Thursday, period.Friday, period.Saturday,period.Other)
	}
}
