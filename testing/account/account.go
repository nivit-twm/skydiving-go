package account

import (
	"bank-testing/credit"
	"errors"
	"fmt"
)

type Account struct {
	creditService credit.CreditService
	FirstName     string
	LastName      string
	Balance       float64
}

func NewAccount(creditService credit.CreditService, firstName, lastName string, balance float64) (*Account, error) {
	if balance < 0 {
		return nil, errors.New("cannot create account w/ negative balance")
	}
	return &Account{creditService, firstName, lastName, balance}, nil
}

func (a *Account) ChangeName(firstName, lastName string) {
	a.FirstName = firstName
	a.LastName = lastName
}

func (a Account) String() string {
	return fmt.Sprintf("Name: %s %s\nBalance: %.2f\n", a.FirstName, a.LastName, a.Balance)
}

func (a *Account) Deposit(amount float64) (float64, error) {
	credit := a.creditService.CheckCredit(a.FirstName)
	if !credit {
		return 0.0, errors.New("this account is in the blacklist")
	}
	if amount > 0.0 {
		a.Balance += amount
		return a.Balance, nil
	}
	return 0.0, errors.New("invalid amount")
}

func (a *Account) Withdraw(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= a.Balance {
			a.Balance -= amount
			return a.Balance, nil
		}
		return 0.0, errors.New("you can't withdraw more than the account has")
	}
	return 0.0, errors.New("you can't remove negative numbers")
}

// lowercase function/type will not be export from the package
func (a *Account) loan(amount float64) (float64, error) {
	return 0.0, errors.New("not implement yet")
}

func (a *Account) CheckBalance() float64 {
	return a.Balance
}

func init() {
	fmt.Println("Will be called whenever package is imported.")
}
