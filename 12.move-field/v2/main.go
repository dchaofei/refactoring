package v2

type AccountType struct {
	interestRate int
}

func (a AccountType) GetInterestRate() int {
	return a.interestRate
}

type Account struct {
	t AccountType
}

func (a Account) InterestForAmountDays(amount, days int) int {
	return a.t.GetInterestRate() * amount * days / 365
}
