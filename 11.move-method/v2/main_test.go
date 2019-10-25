package v1

import "testing"

func TestAccount_BankCharge(t *testing.T) {
	want := 14
	got := Account{
		t:             AccountType{1},
		daysOverdrawn: 1,
	}.BankCharge()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
