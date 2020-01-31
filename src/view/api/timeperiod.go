package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
	"nagios-conf-manager/src/model"
)

func GetTimePeriods(c *gin.Context) {
	timePeriods, err := controller.ListAllTimePeriods(os.Getenv("NAGIOS_BASE_PATH"))
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
	timePeriod, err := controller.FindTimePeriodByName(os.Getenv("NAGIOS_BASE_PATH"), requestParameterName)
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

func PutTimePeriod(c *gin.Context) {

	var bodyDecoded model.TimePeriods
	err := c.Bind(&bodyDecoded)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newTp, err := controller.CreateNewTimePeriod(os.Getenv("NAGIOS_BASE_PATH"), bodyDecoded)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if newTp != nil {
		c.JSON(http.StatusCreated, newTp)
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}
