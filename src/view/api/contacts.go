package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
)

func GetContacts(c *gin.Context) {
	contacts, err := controller.ListAllContacts(os.Getenv("NAGIOS_BASE_PATH"))
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if len(contacts) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, contacts)
}

func GetContactByName(c *gin.Context) {
	requestParameterName := c.Param("name")
	contact, err := controller.FindContactByName(os.Getenv("NAGIOS_BASE_PATH"), requestParameterName)
	if err != nil {
		// http.Error(response, "já foste: "+err.Error(), http.StatusInternalServerError)
		_ = c.Error(err)
		return
	} else if contact != nil {
		c.JSON(http.StatusOK, contact)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}
