package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

type Handler struct {
}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}
	c.R.GET("/", h.HelloWorld)

	c.R.GET("/metrics", gin.WrapH(promhttp.Handler()))
	c.R.GET("/health", h.Healthz)
	c.R.GET("/healthz", h.Healthz)

	eth := c.R.Group("/eth")
	eth.GET("/balance/:address", func(c *gin.Context) {
		c.JSON(200, gin.H{"balance": "0x" + c.Param("address")})
	})

}

func (h *Handler) HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}

func (h *Handler) Healthz(c *gin.Context) {
	started := time.Now()
	duration := time.Now().Sub(started)
	if duration.Seconds() > 10 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": duration.Seconds()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}
