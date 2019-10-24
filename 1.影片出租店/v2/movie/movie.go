package movie

const (
	// 普通片
	Regular = iota
	// 新片
	NewRelease
	// 儿童片
	ChildRens
)

type Movie struct {
	title string
	price Price // 影片类型
}

func NewMovie(title string, priceCode int) *Movie {
	movie := &Movie{
		title: title,
	}
	movie.setPriceCode(priceCode)
	return movie
}

func (m *Movie) GetCharge(daysRented int) float64 {
	return m.price.GetCharge(daysRented)
}

func (m *Movie) GetFrequentRenterPoints(daysRented int) int {
	return m.price.GetFrequentRenterPoints(daysRented)
}

func (m *Movie) setPriceCode(priceCode int) {
	switch priceCode {
	case Regular:
		m.price = RegularPrice{}
	case ChildRens:
		m.price = ChildrensPrice{}
	case NewRelease:
		m.price = NewReleasePrice{}
	default:
		panic("Incorrect Price Code")
	}
}