package services

import (
	"alluvial/models/interfaces"

	"context"
	"fmt"
)

// ChainstackService acts as a struct for injecting an implementation of ChainstackRepository
// for use in service methods
type ChainstackService struct {
	ChainstackRepository interfaces.ChainStackRepositoryInterface
}

func NewChainstackService(c *ClientConfig) *ChainstackService {
	return &ChainstackService{
		ChainstackRepository: c.ChainstackRepository,
	}
}

func (a *ChainstackService) GetBalance(c context.Context, id string) (string, error) {
	if a.ChainstackRepository == nil {
		return "", fmt.Errorf("GetBalance: chainstackRespository is nil")
	}
	return a.ChainstackRepository.GetBalanceByIDCS(c, id)
}
