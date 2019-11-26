package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetServices(c *gin.Context) {
	services, err := controller.ListAllService("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if len(services) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, services)
}

func GetServiceByName(c *gin.Context) {
	requestParameterName := c.Param("name")
	service, err := controller.FindServiceByName("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles", requestParameterName)
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if service != nil {
		c.JSON(http.StatusOK, service)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}
