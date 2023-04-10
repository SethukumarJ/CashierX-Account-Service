package repository

import (
	"context"

	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/domain"
	interfaces "github.com/SethukumarJ/CashierX-Auth-Service/pkg/repository/interface"
	"gorm.io/gorm"
)

type accountDatabase struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) interfaces.AccountRepository {
	return &accountDatabase{DB}
}

func (a *accountDatabase) CreateAccount(ctx context.Context, account domain.Accounts) (domain.Accounts, error) {
	// Create Account
	err := a.DB.Create(&account).Error
	if err != nil {
		return domain.Accounts{}, err
	}
	return account, nil
}
