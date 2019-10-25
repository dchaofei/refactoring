package v1

type AccountType struct {
	Type int
}

func (a AccountType) IsPremium() bool {
	return a.Type == 1
}

type Account struct {
	t             AccountType
	daysOverdrawn int
}

func (a Account) overdraftCharge() int {
	if a.t.IsPremium() {
		result := 10
		if a.daysOverdrawn > 7 {
			result += (a.daysOverdrawn - 7) * 8
		}
		return result
	}
	return a.daysOverdrawn * 7
}

func (a Account) BankCharge() int {
	result := 4
	if a.daysOverdrawn > 0 {
		result += a.overdraftCharge()
	}
	return result
}
