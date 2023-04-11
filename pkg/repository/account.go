package repository

import (
	"context"
	"fmt"

	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/domain"
	interfaces "github.com/SethukumarJ/CashierX-Auth-Service/pkg/repository/interface"
	"gorm.io/gorm"
)

type accountDatabase struct {
	DB *gorm.DB
}

// UpdateAccount implements interfaces.AccountRepository
func (c *accountDatabase) UpdateAccount(ctx context.Context, account domain.Accounts, id int64) (domain.Accounts, error) {
	fmt.Println(account)
	// Update 
	fieldsToUpdate := map[string]interface{}{
		"account_type": account.AccountType,
		"balance":      account.Balance,
	}

	err := c.DB.Model(&domain.Accounts{}).Where("account_id = ?", id).Updates(fieldsToUpdate).Error
	if err != nil {
		return domain.Accounts{}, err
	}

	// Return the updated account
	return account, nil
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

func (c *accountDatabase) FindByID(ctx context.Context, id uint) (domain.Accounts, error) {
	var account domain.Accounts
	err := c.DB.First(&account, id).Error

	return account, err
}
