package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jonny-mark/package-demo/prometheus/pkg/metric"
	"net/http"
	"time"
)

var (
	labels = []string{"status", "endpoint", "method", "service"}

	// QPS
	reqCount = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: "gin",
		Name:      "http_request_count_total",
		Help:      "Total number of HTTP requests made.",
		Labels:    labels,
	})
	// 当前正在处理请求的QPS
	curReqCount = metric.NewGaugeVec(
		&metric.GaugeVecOpts{
			Namespace: "gin",
			Name:      "http_request_in_flight",
			Help:      "Current number of http requests in flight.",
			Labels:    labels,
		})

	// 接口响应时间
	reqDuration = metric.NewHistogramVec(
		&metric.HistogramVecOpts{
			Namespace: "gin",
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
			Labels:    labels,
		})

	// 请求大小
	reqSizeBytes = metric.NewHistogramVec(
		&metric.HistogramVecOpts{
			Namespace: "gin",
			Name:      "http_request_size_bytes",
			Help:      "HTTP request sizes in bytes.",
			Labels:    labels,
		})

	// 响应大小
	respSizeBytes = metric.NewHistogramVec(
		&metric.HistogramVecOpts{
			Namespace: "gin",
			Name:      "http_response_size_bytes",
			Help:      "HTTP request sizes in bytes.",
			Labels:    labels,
		})
)

func Metrics(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		endpoint := c.Request.URL.Path
		method := c.Request.Method

		labels := []string{status, endpoint, method, serviceName}

		curReqCount.Inc(labels...)
		defer curReqCount.Dec(labels...)

		reqCount.Inc(labels...)
		reqDuration.Observe(float64(time.Since(start).Seconds()), labels...)
		reqSizeBytes.Observe(calcRequestSize(c.Request), labels...)

		// no response content will return -1
		respSize := c.Writer.Size()
		if respSize < 0 {
			respSize = 0
		}
		respSizeBytes.Observe(float64(respSize), labels...)
	}
}

// 计算请求的字节长度
func calcRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size += len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)
	size += len(r.Host)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}

	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}
