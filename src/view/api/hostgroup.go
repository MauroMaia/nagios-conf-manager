package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetHostGroupList(c *gin.Context) {
	groups, err := controller.ListAllHostGroups(os.Getenv("NAGIOS_BASE_PATH"))
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if len(groups) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, groups)
}

func GetHostGroupByName(c *gin.Context) {
	requestParameterName := c.Param("name")
	hostGroup, err := controller.FindHostGroupByName(os.Getenv("NAGIOS_BASE_PATH"), requestParameterName)
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if hostGroup != nil {
		c.JSON(http.StatusOK, hostGroup)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}
