package movie

import "strconv"

type Customer struct {
	name    string
	rentals []*Rental
}

func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (c *Customer) AddRental(r *Rental) {
	c.rentals = append(c.rentals, r)
}

func (c *Customer) Statement() string {
	var (
		totalAmount          float64
		frequentRenterPoints int
		result               = "Rental Record for " + c.name + "\n"
	)

	// determine amounts for each line
	for _, v := range c.rentals {
		var thisAmount float64
		switch v.movie.priceCode {
		case Regular:
			thisAmount += 2
			if v.daysRented > 2 {
				thisAmount += (float64(v.daysRented) - 2) * 1.5
			}
		case NewRelease:
			thisAmount += float64(v.daysRented) * 3
		case ChildRens:
			thisAmount += 1.5
			if v.daysRented > 3 {
				thisAmount += (float64(v.daysRented) - 3) * 1.5
			}
		}

		// add frequent renter points
		frequentRenterPoints++
		// add bonus for a two day new release rental
		if v.movie.priceCode == NewRelease && v.daysRented > 1 {
			frequentRenterPoints++
		}
		// show figures for this rental
		result += "\t" + v.movie.title + "\t" + strconv.FormatFloat(thisAmount, 'f', 2, 64) + "\n"
		totalAmount += thisAmount
	}

	// add footer lines
	result += "Amount owed is " + strconv.FormatFloat(totalAmount, 'f', 2, 64) + "\n"
	result += "You earned " + strconv.Itoa(frequentRenterPoints) + " frequent renter points"
	return result
}
