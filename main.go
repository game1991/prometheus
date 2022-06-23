package main

import (
	"fmt"
	"os"
	"prometheus/router"
	"prometheus/load"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/eapp"
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "version" || args[1] == "-v") {
		fmt.Println(eapp.Name())
		fmt.Println(eapp.AppVersion())
		fmt.Println(eapp.AppZone())
		fmt.Println(eapp.AppRegion())
		fmt.Println(eapp.AppInstance())
		fmt.Println(eapp.BuildUser())
		fmt.Println(eapp.BuildHost())
		fmt.Println(eapp.BuildTime())
		return
	}

	load.Redis()

	server := gin.Default()

	router.API(server)
	router.SYS(server)
	router.Prometheus(server)

	if err := server.Run(":8181"); err != nil {
		panic(err)
	}
}
