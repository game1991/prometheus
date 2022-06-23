package controller

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTPDurations http请求耗时指标
	HTTPDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_durations_seconds",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path"},
	)
	// HTTPDurationsHistogram http耗时直方图分析指标
	HTTPDurationsHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_durations_histogram_seconds",
			Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
		},
		[]string{"path"},
	)
	// MetricServerReqCodeTotal http请求数量
	MetricServerReqCodeTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
		},
		[]string{"method", "path"},
	)
	// MetricServerRequestFailCounter http失败请求数量及信息
	MetricServerRequestFailCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_fail",
		},
		[]string{"time", "method", "path"},
	)
)
// MetricsHandler 请求的监控中间件
func MetricsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("start monitor request...")
		purl, _ := url.Parse(c.Request.RequestURI)
		defer func() {
			MetricServerReqCodeTotal.With(prometheus.Labels{
				"method": c.Request.Method,
				"path":   purl.Path,
			}).Inc() //请求数量累加
			if c.Writer.Status() != http.StatusOK {
				MetricServerRequestFailCounter.With(prometheus.Labels{
					"time":   time.Now().Format("2006-01-02 15:04:05"),
					"method": c.Request.Method,
					"path":   purl.Path,
				})
			}
			HTTPDurations.With(prometheus.Labels{"path": purl.Path}).Observe(float64(rand.Intn(30)))          //监测请求耗时
			HTTPDurationsHistogram.With(prometheus.Labels{"path": purl.Path}).Observe(float64(rand.Intn(30))) //统计分析请求
			log.Println("monitor request success...")
		}()
		//请求处理
		c.Next()
	}
}
