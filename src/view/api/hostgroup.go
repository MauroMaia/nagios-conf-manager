package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetHostGroupList(c *gin.Context) {
	groups, err := controller.ListAllHostGroups("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
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
	hostGroup, err := controller.FindHostGroupByName("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles", requestParameterName)
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
