package v1

var (
	quantity  float64
	itemPrice float64
)

func getPrice() float64 {
	basePrice := quantity * itemPrice
	var discountFactor float64
	if basePrice > 1000 {
		discountFactor = 0.95
	} else {
		discountFactor = 0.98
	}
	return basePrice * discountFactor
}
