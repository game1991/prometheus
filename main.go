package main

import (
	"fmt"
	"prometheus/router"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/eapp"
)

func main() {

	fmt.Println(eapp.AppMode())
	fmt.Println(eapp.Name())
	fmt.Println(eapp.AppVersion())
	fmt.Println(eapp.AppRegion())
	fmt.Println(eapp.AppInstance())
	fmt.Println(eapp.BuildUser())
	fmt.Println(eapp.BuildHost())

	server := gin.Default()

	router.API(server)
	router.SYS(server)
	router.Prometheus(server)

	if err := server.Run(":8181"); err != nil {
		panic(err)
	}
}
