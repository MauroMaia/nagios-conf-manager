package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetServices(c *gin.Context) {
	services, err := controller.ListAllService(os.Getenv("NAGIOS_BASE_PATH"))
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
	service, err := controller.FindServiceByName(os.Getenv("NAGIOS_BASE_PATH"), requestParameterName)
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
