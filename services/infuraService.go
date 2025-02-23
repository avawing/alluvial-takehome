package services

import (
	"alluvial/models/interfaces"

	"context"
)

// ClientConfig will hold repositories that will eventually be injected into
// the service layer
type ClientConfig struct {
	InfuraRepository     interfaces.InfuraRepositoryInterface
	AlchemyRepository    interfaces.AlchemyRepositoryInterface
	ChainstackRepository interfaces.ChainStackRepositoryInterface
}

// InfuraService acts as a struct for injecting an implementation of InfuraRepository
// for use in service methods
type InfuraService struct {
	InfuraRepository interfaces.InfuraRepositoryInterface
}

func NewInfuraService(c *ClientConfig) *InfuraService {
	return &InfuraService{
		InfuraRepository: c.InfuraRepository,
	}
}

func (i *InfuraService) GetBalance(c context.Context, id string) (string, error) {
	return i.InfuraRepository.GetBalanceByID(c, id)
}
