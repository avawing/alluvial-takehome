package utils

import (
	"alluvial/models/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
)

func MakeRequests(c *gin.Context, s interfaces.InfuraServiceInterface, addr string) (string, error) {
	hexstring, err := s.GetBalance(c, addr)
	log.Printf("Hex String: %s", hexstring)
	if err != nil {
		return "", fmt.Errorf("MakeRequests: %v+", err)
	}

	// Convert hex string to big integer
	weiValue := new(big.Int)
	weiValue, success := weiValue.SetString(hexstring[2:], 16) // Remove '0x' and convert from base 16
	if !success {
		return "", fmt.Errorf("MakeRequests: Error converting hex to string")
	}

	return weiValue.String(), nil
}
