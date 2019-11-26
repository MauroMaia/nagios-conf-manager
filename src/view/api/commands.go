package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetCommands(c *gin.Context) {
	commands, err := controller.ListAllCommands("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if len(commands) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, commands)
}

func GetCommandByName(c *gin.Context) {
	requestParameterName := c.Param("name")
	command, err := controller.FindCommandByName("/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles", requestParameterName)
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if command != nil {
		c.JSON(http.StatusOK, command)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}
