package v1

import "testing"

func TestAccount_Gamma(t *testing.T) {
	want := 715
	got := Account{}.Gamma(1, 2, 3)
	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}
