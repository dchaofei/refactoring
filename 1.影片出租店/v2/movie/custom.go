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
	result := "Rental Record for " + c.name + "\n"
	// determine amounts for each line
	for _, rental := range c.rentals {
		// show figures for this rental
		result += "\t" + rental.movie.title + "\t" + strconv.FormatFloat(rental.GetCharge(), 'f', 2, 64) + "\n"
	}
	// add footer lines
	result += "Amount owed is " + strconv.FormatFloat(c.GetTotalCharge(), 'f', 2, 64) + "\n"
	result += "You earned " + strconv.Itoa(c.GetTotalFrequentRenterPoints()) + " frequent renter points"
	return result
}

func (c *Customer) GetTotalCharge() float64 {
	var result float64
	for _, rental := range c.rentals {
		result += rental.GetCharge()
	}
	return result
}

func (c *Customer) GetTotalFrequentRenterPoints() int {
	var result int
	for _, rental := range c.rentals {
		result += rental.GetFrequentRenterPoints()
	}
	return result
}

func (c *Customer) HtmlStatement() string {
	result := "<H1>Rentals for <EM>" + c.name + "</EM></H1><P>\n"
	for _, rental := range c.rentals {
		result += rental.movie.title + ": " + strconv.FormatFloat(rental.GetCharge(), 'f', 2, 64) + "<BR>\n"
	}
	result += "<P>You owe <EM>" + strconv.FormatFloat(c.GetTotalCharge(), 'f', 2, 64) + "</EM><P>\n"
	result += "On this rental you earned <EM>" + strconv.Itoa(c.GetTotalFrequentRenterPoints()) + "</EM> frequent renter points<P>"
	return result
}
