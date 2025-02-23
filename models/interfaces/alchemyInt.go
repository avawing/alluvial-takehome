package interfaces

import "context"

// InfuraServiceInterface defines methods the handler layer expects
// any service it interacts with to implement
type AlchemyServiceInterface interface {
	GetBalance(c context.Context, id string) (string, error)
}

// AlchemyRepositoryInterface defines methods the service layer expects
// any repository it interacts with to implement
type AlchemyRepositoryInterface interface {
	GetBalanceByIDAlchemy(c context.Context, id string) (string, error)
}
