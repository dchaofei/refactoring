package v1

import "testing"

func TestGetPrice(t *testing.T) {
	cases := []struct {
		quantity  float64
		itemPrice float64
		want      float64
	}{
		{1, 1, 0.98},
		{1, 10000, 9500},
	}

	for _, c := range cases {
		quantity = c.quantity
		itemPrice = c.itemPrice
		got := GetPrice()
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
