package services

import (
	"alluvial/models/interfaces"

	"context"
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
	return a.ChainstackRepository.GetBalanceByIDCS(c, id)
}
