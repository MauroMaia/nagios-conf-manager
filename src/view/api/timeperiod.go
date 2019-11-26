package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetTimePeriods(c *gin.Context) {
	timePeriods, err := controller.ListAllTimePeriods("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if len(timePeriods) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, timePeriods)
}

func GetTimePeriodByName(c *gin.Context) {
	requestParameterName := c.Param("name")
	timePeriod, err := controller.FindTimePeriodByName("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles", requestParameterName)
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if timePeriod != nil {
		c.JSON(http.StatusOK, timePeriod)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}
