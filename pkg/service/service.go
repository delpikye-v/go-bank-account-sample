package service

import (
	"context"

	"github.com/go-kit/kit/bankaccount/pkg/model"
)

// BankaccountService describes the service.
type BankaccountService interface {
	CreateAcc(ctx context.Context, acc model.Account) (string, error)
	GetAcc(ctx context.Context, uuid string) (model.Account, error)
	DepositAmount(ctx context.Context, number int64, ammount float64) (string, error)
	WithDrawAmount(ctx context.Context, number int64, ammount float64) (string, error)
}

type basicBankaccountService struct{}

func (b *basicBankaccountService) CreateAcc(ctx context.Context, acc model.Account) (s0 string, e1 error) {
	e1 = accountStore.CreateAccount(&acc)
	if e1 != nil {
		return "Create error!", e1
	}
	return "Account created!", e1
}
func (b *basicBankaccountService) GetAcc(ctx context.Context, uuid string) (m0 model.Account, e1 error) {
	acc, err := accountStore.GetAccount(uuid)
	return acc, err
}

func (b *basicBankaccountService) DepositAmount(ctx context.Context, number int64, ammount float64) (s0 string, e1 error) {
	acc, err := accountStore.GetAccountByNumber(number)
	if err != nil {
		return "Account not exist!", err
	}
	acc.Balance = acc.Balance + ammount
	errUpdate := accountStore.UpdateAccount(&acc)

	if errUpdate != nil {
		return "Transfer not successfully!", err
	}
	return "Transfer successfully!", nil
}

// NewBasicBankaccountService returns a naive, stateless implementation of BankaccountService.
func NewBasicBankaccountService() BankaccountService {
	return &basicBankaccountService{}
}

// New returns a BankaccountService with all of the expected middleware wired in.
func New(middleware []Middleware) BankaccountService {
	var svc BankaccountService = NewBasicBankaccountService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicBankaccountService) WithDrawAmount(ctx context.Context, number int64, ammount float64) (s0 string, e1 error) {
	acc, err := accountStore.GetAccountByNumber(number)
	if err != nil {
		return "Account not exist!", err
	}

	if acc.Balance < 1000 {
		return "Account balance below 1000, not enough to withdraw money!", err
	}
	if acc.Balance < ammount {
		return "Cannot withdraw the amount beyond the account balance!", err
	}
	if 1000 > acc.Balance-ammount {
		return "cannot withdraw funds beyond the allowed balance of 1000!", err
	}

	acc.Balance = acc.Balance - ammount
	errUpdate := accountStore.UpdateAccount(&acc)

	if errUpdate != nil {
		return "Withdraw not successfully!", err
	}

	return "Withdraw successfully!", nil
}
