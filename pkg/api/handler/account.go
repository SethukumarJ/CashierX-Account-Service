package handler

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/domain"
	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/pb"
	services "github.com/SethukumarJ/CashierX-Auth-Service/pkg/usecase/interface"
	"gorm.io/gorm"
)

type AccountHandler struct {
	accountUsecase services.AccountUsecase
}

func NewAccountHandler(usecase services.AccountUsecase) *AccountHandler {
	return &AccountHandler{
		accountUsecase: usecase,
	}
}

func GenerateRandomNumber(length int) int {
	rand.Seed(time.Now().UnixNano())
	min := int64(1)
	max := int64(1)
	for i := 0; i < length; i++ {
		max *= 10
	}
	return int(rand.Int63n(max-min) + min)
}

// CreateAccount
func (cr *AccountHandler) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	fmt.Println("create account called in service")
	accno := GenerateRandomNumber(11)

	
	account := domain.Accounts{
		AccountType:   domain.AccountType(req.Type),
		AccountNumber: int64(accno),
		AccountHolder: req.AccountHolder,
		Balance:       float64(req.Balance),
		CreatedAt:     time.Now(),
		UserID:        req.UserId,
	}
	fmt.Println("account", account)
	// Create Account
	account, err := cr.accountUsecase.CreateAccount(ctx, account)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return &pb.CreateAccountResponse{
				Status: http.StatusConflict,
				Error:  fmt.Sprint(errors.New("account already exists, try again")),
			}, nil
		}
		return &pb.CreateAccountResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  fmt.Sprint(errors.New("failed to create account")),
		}, nil
	}
	return &pb.CreateAccountResponse{
		Status: http.StatusCreated,
		Error:  "",
		Id:     int64(account.AccountID),
	}, nil
}
func (cr *AccountHandler) FindAccount(ctx context.Context, req *pb.FindAccountRequest) (*pb.FindAccountResponse, error) {

	return &pb.FindAccountResponse{
		Status: http.StatusCreated,
		Error:  "",
		Data:   &pb.FindAccountData{},
	}, nil
}

func (cr *AccountHandler) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {

	return &pb.UpdateAccountResponse{
		Status: http.StatusCreated,
		Error:  "",
		Id:     0,
	}, nil
}

func (cr *AccountHandler) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {

	return &pb.DeleteAccountResponse{
		Status: http.StatusCreated,
		Error:  "",
		Id:     0,
	}, nil
}

// Get Balance
func (cr *AccountHandler) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {

	return &pb.GetBalanceResponse{
		Status:  http.StatusCreated,
		Error:   "",
		Balance: 0,
	}, nil
}

func (cr *AccountHandler) GetAccounts(ctx context.Context, req *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {

	return &pb.GetAccountsResponse{
		Status:   http.StatusCreated,
		Error:    "",
		Accounts: []*pb.AccountData{},
	}, nil
}

func (cr *AccountHandler) GetTransactions(ctx context.Context, req *pb.GetTransactionsRequest) (*pb.GetTransactionsResponse, error) {

	return &pb.GetTransactionsResponse{
		Status:       http.StatusCreated,
		Error:        "",
		Transactions: []*pb.TransactionData{},
	}, nil
}
func (cr *AccountHandler) GetTransferredTransactions(ctx context.Context, req *pb.GetTransferredTransactionsRequest) (*pb.GetTransferredTransactionsResponse, error) {

	return &pb.GetTransferredTransactionsResponse{
		Status:       http.StatusCreated,
		Error:        "",
		Transactions: []*pb.TransactionData{},
	}, nil
}
