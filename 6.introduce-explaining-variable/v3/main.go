package v1

import "math"

var (
	quantity  float64
	itemPrice float64
)

func Price() float64 {
	return quantity*itemPrice -
		math.Max(0, quantity-500)*itemPrice*0.05 +
		math.Min(quantity*itemPrice*0.1, 100.0)
}
