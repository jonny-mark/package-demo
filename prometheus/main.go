package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonny-mark/package-demo/prometheus/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//func main() {
//	http.Handle("/metrics", promhttp.Handler())
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

func main() {
	g := gin.Default()
	g.Use(gin.Recovery())
	g.Use(middleware.Metrics("prometheus-demo"))
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))
	g.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	g.Run(":8080")
}
