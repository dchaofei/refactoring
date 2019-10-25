package v2

import "testing"

func TestAccount_InterestForAmountDays(t *testing.T) {
	want := 0
	got := Account{
		t: AccountType{interestRate: 1},
	}.InterestForAmountDays(1, 1)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
