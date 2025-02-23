package handlers

import (
	"alluvial/services"
	"alluvial/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

type Handler struct {
	InfuraService     *services.InfuraService
	AlchemyService    *services.AlchemyService
	ChainstackService *services.ChainstackService
	LoadBalancer      *utils.LoadBalancer
}

type Config struct {
	R                 *gin.Engine
	InfuraService     *services.InfuraService
	AlchemyService    *services.AlchemyService
	ChainstackService *services.ChainstackService
	LoadBalancer      *utils.LoadBalancer
}

func NewHandler(c *Config) {
	h := &Handler{
		InfuraService:     c.InfuraService,
		AlchemyService:    c.AlchemyService,
		ChainstackService: c.ChainstackService,
		LoadBalancer:      c.LoadBalancer,
	}
	c.R.GET("/", h.HelloWorld)

	c.R.GET("/metrics", gin.WrapH(promhttp.Handler()))
	c.R.GET("/healthz", h.Healthz)

	eth := c.R.Group("/eth")
	eth.GET("/balance/:address", h.GetEthBalance)

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

func (h *Handler) GetEthBalance(c *gin.Context) {
	address := c.Param("address")

	if amt, err := h.LoadBalancer.MakeRequests(c, address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"balance": amt})
	}
}
