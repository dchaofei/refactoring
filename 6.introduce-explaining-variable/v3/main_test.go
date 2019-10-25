package v1

import "testing"

func TestPrice(t *testing.T) {
	cases := []struct {
		quantity  float64
		itemPrice float64
		want      float64
	}{
		{1, 1, 1.1},
		{1, 10000, 10100},
	}

	for _, c := range cases {
		quantity = c.quantity
		itemPrice = c.itemPrice
		got := Price()
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}