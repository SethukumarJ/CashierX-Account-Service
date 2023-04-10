package interfaces

import (
	"context"

	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/domain"
)

type AccountRepository interface {
	// CreateAccount creates a new account
	CreateAccount(ctx context.Context, account domain.Accounts) (domain.Accounts, error)
}
