package v1

import "testing"

func TestGetDistanceTravelled(t *testing.T) {
	cases := []struct {
		primaryForce   float64
		mass           float64
		delay          float64
		secondaryForce float64
		want           float64
	}{
		{1.0, 1.0, 1.0, 1.0, 2.5},
		{3.0, 1.0, 7.0, 2.0, 6.0},
	}

	for _, c := range cases {
		primaryForce = c.primaryForce
		mass = c.mass
		delay = c.delay
		secondaryForce = c.secondaryForce
		got := GetDistanceTravelled(2)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
