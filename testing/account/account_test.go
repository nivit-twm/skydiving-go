package account

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type creditServiceMock struct {
	mock.Mock
}

func (m *creditServiceMock) CheckCredit(firstName string) bool {
	fmt.Printf("input: %s\n", firstName)
	args := m.Called(firstName)

	return args.Bool(0)
}

func TestAccount_CanDepositIfAccountPassBackgroundCheck(t *testing.T) {
	creditService := new(creditServiceMock)
	creditService.On("CheckCredit", "John").Return(true)

	accA, _ := NewAccount(creditService, "John", "Doe", 100)
	accA.Deposit(200.0)

	want := 300.0
	got := accA.Balance
	if want != got {
		t.Errorf("Cannot deposit want: %f, but got: %f", want, got)
	}

	creditService.AssertExpectations(t)
}
