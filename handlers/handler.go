package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	//h := &Handler{}
	c.R.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	c.R.GET("/metrics", gin.WrapH(promhttp.Handler()))
	c.R.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, gin.H{"liveness": "true"})
	})

	eth := c.R.Group("/eth")
	eth.GET("/balance/:address", func(c *gin.Context) {
		c.JSON(200, gin.H{"balance": "0x" + c.Param("address")})
	})

}
