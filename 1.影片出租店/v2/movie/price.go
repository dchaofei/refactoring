package movie

type Price interface {
	GetPriceCode() int
	GetCharge(daysRented int) float64
	GetFrequentRenterPoints(daysRented int) int
}

type commonPoints struct{}

func (commonPoints) GetFrequentRenterPoints(daysRented int) int {
	return 1
}

type ChildrensPrice struct{ commonPoints }

func (ChildrensPrice) GetPriceCode() int {
	return ChildRens
}

func (ChildrensPrice) GetCharge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += (float64(daysRented) - 3) * 1.5
	}
	return result
}

type NewReleasePrice struct{ commonPoints }

func (NewReleasePrice) GetPriceCode() int {
	return NewRelease
}

func (NewReleasePrice) GetCharge(daysRented int) float64 {
	return float64(daysRented) * 3
}

func (n NewReleasePrice) GetFrequentRenterPoints(daysRented int) int {
	if daysRented > 1 {
		return 2
	}
	return n.commonPoints.GetFrequentRenterPoints(daysRented)
}

type RegularPrice struct{ commonPoints }

func (RegularPrice) GetPriceCode() int {
	return Regular
}

func (RegularPrice) GetCharge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += (float64(daysRented) - 2) * 1.5
	}
	return result
}
