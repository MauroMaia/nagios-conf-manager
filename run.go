package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"nagios-conf-manager/src/controller"
	"nagios-conf-manager/src/view/api"
	"nagios-conf-manager/src/view/cmd"
)

/*func init(){
	f, err := os.OpenFile("/home/mauro.maia/go/src/nagios-conf-manager/log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetPrefix("main/run")
	log.Println("Init generic")
}*/

func startWebservice() {
	// http.HandleFunc("/", webGetHostList)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	r := gin.Default()
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
	r.GET("/hostgroups", api.GetHostGroupList)
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
