package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetCommands(c *gin.Context) {
	commands, err := controller.ListAllCommands(os.Getenv("NAGIOS_BASE_PATH"))
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
	command, err := controller.FindCommandByName(os.Getenv("NAGIOS_BASE_PATH"), requestParameterName)
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
