package v1

type Account struct{}

func (Account) delta() int {
	return 1
}

func (a Account) Gamma(inputVal, quantity, yearToDate int) int {
	importantValue1 := inputVal*quantity + a.delta()
	importantValue2 := inputVal*yearToDate + 100
	if yearToDate-importantValue1 > 100 {
		importantValue2 -= 20
	}
	importantValue3 := importantValue2 * 7
	return importantValue3 - 2*importantValue1
}
