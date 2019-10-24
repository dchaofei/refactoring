package movie

import (
	"testing"
)

func TestCustomer_Statement(t *testing.T) {
	customer := Customer{
		name: "dingcf",
		rentals: []*Rental{
			NewRental(NewMovie("战狼", Regular), 1),
			NewRental(NewMovie("喜洋洋", ChildRens), 1),
			NewRental(NewMovie("攀登者", NewRelease), 3),
		},
	}
	want := "Rental Record for dingcf\n"
	want += "\t战狼\t2.00\n"
	want += "\t喜洋洋\t1.50\n"
	want += "\t攀登者\t9.00\n"
	want += "Amount owed is 12.50\n"
	want += "You earned 4 frequent renter points"
	got := customer.Statement()
	if got != want {
		t.Errorf("customer.Statement got %s want %s", got, want)
	}
}
