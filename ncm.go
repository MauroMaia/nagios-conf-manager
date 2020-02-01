package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
	"nagios-conf-manager/src/view/api"
	"nagios-conf-manager/src/view/cmd"
)

func startWebservice() {

	r := gin.Default()

	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"*"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           30 * time.Second,
	}))

	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	r.GET("/hosts", api.GetHostList)
	r.GET("/hosts/:name", api.GetHostByName)

	r.GET("/hostgroups", api.GetHostGroupList)
	r.GET("/hostgroups/:name", api.GetHostGroupByName)

	r.GET("/services", api.GetServices)
	r.GET("/services/:name", api.GetServiceByName)

	r.GET("/commands", api.GetCommands)
	r.GET("/commands/:name", api.GetCommandByName)

	r.GET("/contacts", api.GetContacts)
	r.GET("/contacts/:name", api.GetContactByName)

	r.GET("/contactgroups", api.GetContactGroups)
	r.GET("/contactgroups/:name", api.GetContactGroupByName)

	r.GET("/timeperiods", api.GetTimePeriods)
	r.PUT("/timeperiods", api.PutTimePeriod)
	r.GET("/timeperiods/:name", api.GetTimePeriodByName)

	/*r.POST("/todo", handlers.AddTodoHandler)
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	r.PUT("/todo", handlers.CompleteTodoHandler)*/

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		cmd.PrintErrorString(fmt.Sprintf("Cli expected %v subcommands", []string{"cli", "webserver"}), 1)
	}

	NBP := os.Getenv("NAGIOS_BASE_PATH")
	if NBP == "" {
		err := os.Setenv("NAGIOS_BASE_PATH", "/home/mauro.maia/go/src/nagios-conf-manager/nagiosFiles")
		if err != nil {
			cmd.PrintError(err, 1)
		}

	}
	switch os.Args[1] {
	case "cli":
		controller.CliParseDomain()
		break
	case "web":
		startWebservice()
		break
	default:
		cmd.PrintErrorString(fmt.Sprintf("Cli expected %v subcommands", []string{"cli", "webserver"}), 1)
	}
}
