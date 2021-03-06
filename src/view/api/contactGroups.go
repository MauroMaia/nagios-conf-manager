package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetContactGroups(c *gin.Context) {
	contactGroups, err := controller.ListAllContactGroups(os.Getenv("NAGIOS_BASE_PATH"))
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if len(contactGroups) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, contactGroups)
}

func GetContactGroupByName(c *gin.Context) {
	requestParameterName := c.Param("name")
	contactGroup, err := controller.FindContactGroupByName(os.Getenv("NAGIOS_BASE_PATH"), requestParameterName)
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if contactGroup != nil {
		c.JSON(http.StatusOK, contactGroup)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}
