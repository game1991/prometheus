package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

// API ...
func API(engine *gin.Engine) {
	engine.Use()
}

// SYS ...
func SYS(engine *gin.Engine) {
	prometheus.MustRegister(pingCounter)
	sysGroup := engine.Group("/sys")
	sysGroup.GET("ping", func(ctx *gin.Context) {
		pingCounter.Inc()
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg":  "pong",
			"code": http.StatusOK,
		})
	})
}

// Prometheus ...
func Prometheus(engine *gin.Engine) {
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
