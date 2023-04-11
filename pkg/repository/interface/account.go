package interfaces

import (
	"context"

	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/domain"
)

type AccountRepository interface {
	// CreateAccount creates a new account
	CreateAccount(ctx context.Context, account domain.Accounts) (domain.Accounts, error)
	UpdateAccount(ctx context.Context, account domain.Accounts, id int64) (domain.Accounts, error)
	FindByID(ctx context.Context, id uint) (domain.Accounts, error)
}
