package usecase

import (
	"context"

	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/domain"
	interfaces "github.com/SethukumarJ/CashierX-Auth-Service/pkg/repository/interface"
	services "github.com/SethukumarJ/CashierX-Auth-Service/pkg/usecase/interface"
)

type accountUsecase struct {
	accountRepo interfaces.AccountRepository
}

// UpdateAccount implements interfaces.AccountUsecase
func (c *accountUsecase) UpdateAccount(ctx context.Context, account domain.Accounts, id int64) (domain.Accounts, error) {
	// Update Account
	account, err := c.accountRepo.UpdateAccount(ctx, account, id)
	if err != nil {
		return domain.Accounts{}, err
	}
	return account, nil
}

// FindByID implements interfaces.AccountUsecase
func (c *accountUsecase) FindByID(ctx context.Context, id uint) (domain.Accounts, error) {
	user, err := c.accountRepo.FindByID(ctx, id)
	return user, err
}

func NewAccountUsecase(repo interfaces.AccountRepository) services.AccountUsecase {
	return &accountUsecase{
		accountRepo: repo,
	}
}
func (a *accountUsecase) CreateAccount(ctx context.Context, account domain.Accounts) (domain.Accounts, error) {
	// Create Account
	account, err := a.accountRepo.CreateAccount(ctx, account)
	if err != nil {
		return domain.Accounts{}, err
	}
	return account, nil
}
