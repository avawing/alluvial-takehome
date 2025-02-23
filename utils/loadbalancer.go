package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"sync"
)

type Server interface {
	GetBalance(c context.Context, id string) (string, error)
}

type LoadBalancer struct {
	servers []Server
	mu      sync.Mutex
	index   int
}

// NewLoadBalancer creates a new LoadBalancer with the given list of servers.
func NewLoadBalancer(servers []Server) *LoadBalancer {
	return &LoadBalancer{servers: servers}
}

// GetNextServer returns the next server in a round-robin fashion.
func (lb *LoadBalancer) GetNextServer() Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	server := lb.servers[lb.index]
	lb.index = (lb.index + 1) % len(lb.servers)
	return server
}

// MakeRequests uses RoundRobin algorithm to make requests to a new server
func (lb *LoadBalancer) MakeRequests(c *gin.Context, addr string) (string, error) {
	var hexstring string
	var err error

	for attempts := 0; attempts < len(lb.servers); attempts++ {
		s := lb.GetNextServer()
		hexstring, err = s.GetBalance(c, addr)
		if err != nil {
			// If it's the last attempt, return the error to the client
			if attempts == len(lb.servers)-1 {
				return "", fmt.Errorf("MakeRequests: %v+", err)
			}
			// Otherwise, continue retrying with a new server
			continue
		} else {
			break
		}
	}

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
