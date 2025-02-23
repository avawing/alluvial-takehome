package interfaces

import "context"

// ChainstackServiceInterface defines methods the handler layer expects
// any service it interacts with to implement
type ChainstackServiceInterface interface {
	GetBalance(c context.Context, id string) (string, error)
}

// ChainStackRepositoryInterface defines methods the service layer expects
// any repository it interacts with to implement
type ChainStackRepositoryInterface interface {
	GetBalanceByIDCS(c context.Context, id string) (string, error)
}
