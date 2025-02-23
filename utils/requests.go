package utils

import (
	"alluvial/models/interfaces"
	"github.com/gin-gonic/gin"
)

func MakeRequests(c *gin.Context, s interfaces.InfuraServiceInterface, addr string) (string, error) {

	return s.GetBalance(c, addr)
	// implement Round Robin
}
