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
	}

	c.JSON(http.StatusOK, hosts)
}