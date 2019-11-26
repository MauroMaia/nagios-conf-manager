package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetHostList(c *gin.Context) {
	hosts, err := controller.ListAllHosts("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if len(hosts) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, hosts)
}


func GetHostByName(c *gin.Context) {
	requestParameterName := c.Param("name")
	host, err := controller.FindHostByName("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles", requestParameterName)
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if host != nil {
		c.JSON(http.StatusOK, host)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}