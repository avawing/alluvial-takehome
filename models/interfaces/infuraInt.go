package interfaces

import "context"

// InfuraServiceInterface defines methods the handler layer expects
// any service it interacts with to implement
type InfuraServiceInterface interface {
	GetBalance(c context.Context, id string) (string, error)
}

// InfuraRepositoryInterface defines methods the service layer expects
// any repository it interacts with to implement
type InfuraRepositoryInterface interface {
	GetBalanceByID(c context.Context, id string) (string, error)
}
