package services

import (
	"alluvial/models/interfaces"

	"context"
)

// AlchemyService acts as a struct for injecting an implementation of AlchemyRepository
// for use in service methods
type AlchemyService struct {
	AlchemyRepository interfaces.AlchemyRepositoryInterface
}

func NewAlchemyService(c *ClientConfig) *AlchemyService {
	return &AlchemyService{
		AlchemyRepository: c.AlchemyRepository,
	}
}

func (a *AlchemyService) GetBalance(c context.Context, id string) (string, error) {
	return a.AlchemyRepository.GetBalanceByIDAlchemy(c, id)
}
