package v1

import "math"

var (
	primaryForce   float64
	mass           float64
	delay          float64
	secondaryForce float64
)

func GetDistanceTravelled(time float64) float64 {
	acc := primaryForce / mass
	primaryTime := math.Min(time, delay)
	result := 0.5 * acc * primaryTime * primaryTime
	secondaryTime := time - delay
	if secondaryTime > 0 {
		primaryVel := acc * delay
		acc = (primaryForce + secondaryForce) / mass
		result += primaryVel*secondaryTime + 0.5*acc*secondaryTime*secondaryTime
	}
	return result
}
