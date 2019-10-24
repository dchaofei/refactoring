package movie

import (
	"testing"
)

var cases = []struct {
	customer *Customer
	wantStr  string
	wantHtml string
}{
	{
		customer: &Customer{
			name: "dingcf",
			rentals: []*Rental{
				NewRental(NewMovie("战狼", Regular), 1),
				NewRental(NewMovie("喜洋洋", ChildRens), 1),
				NewRental(NewMovie("攀登者", NewRelease), 3),
			},
		},
		wantStr: `Rental Record for dingcf
	战狼	2.00
	喜洋洋	1.50
	攀登者	9.00
Amount owed is 12.50
You earned 4 frequent renter points`,
		wantHtml: `<H1>Rentals for <EM>dingcf</EM></H1><P>
战狼: 2.00<BR>
喜洋洋: 1.50<BR>
攀登者: 9.00<BR>
<P>You owe <EM>12.50</EM><P>
On this rental you earned <EM>4</EM> frequent renter points<P>`,
	},
	{
		customer: &Customer{
			name: "ang",
			rentals: []*Rental{
				NewRental(NewMovie("鸡毛飞上天", Regular), 3),
				NewRental(NewMovie("猪猪侠", ChildRens), 1),
				NewRental(NewMovie("少年的你", NewRelease), 4),
			},
		},
		wantStr: `Rental Record for ang
	鸡毛飞上天	3.50
	猪猪侠	1.50
	少年的你	12.00
Amount owed is 17.00
You earned 4 frequent renter points`,
		wantHtml: `<H1>Rentals for <EM>ang</EM></H1><P>
鸡毛飞上天: 3.50<BR>
猪猪侠: 1.50<BR>
少年的你: 12.00<BR>
<P>You owe <EM>17.00</EM><P>
On this rental you earned <EM>4</EM> frequent renter points<P>`,
	},
	{
		customer: &Customer{
			name: "ang",
			rentals: []*Rental{
				NewRental(NewMovie("鸡毛飞上天", Regular), 3),
				NewRental(NewMovie("猪猪侠", ChildRens), 1),
				NewRental(NewMovie("少年的你", NewRelease), 4),
			},
		},
		wantStr: `Rental Record for ang
	鸡毛飞上天	3.50
	猪猪侠	1.50
	少年的你	12.00
Amount owed is 17.00
You earned 4 frequent renter points`,
		wantHtml: `<H1>Rentals for <EM>ang</EM></H1><P>
鸡毛飞上天: 3.50<BR>
猪猪侠: 1.50<BR>
少年的你: 12.00<BR>
<P>You owe <EM>17.00</EM><P>
On this rental you earned <EM>4</EM> frequent renter points<P>`,
	},
	{
		customer: &Customer{
			name: "ang",
			rentals: []*Rental{
				NewRental(NewMovie("鸡毛飞上天", Regular), 3),
				NewRental(NewMovie("猪猪侠", ChildRens), 1),
				NewRental(NewMovie("少年的你", NewRelease), 1),
			},
		},
		wantStr: `Rental Record for ang
	鸡毛飞上天	3.50
	猪猪侠	1.50
	少年的你	3.00
Amount owed is 8.00
You earned 3 frequent renter points`,
		wantHtml: `<H1>Rentals for <EM>ang</EM></H1><P>
鸡毛飞上天: 3.50<BR>
猪猪侠: 1.50<BR>
少年的你: 3.00<BR>
<P>You owe <EM>8.00</EM><P>
On this rental you earned <EM>3</EM> frequent renter points<P>`,
	},
}

func TestCustomer_Statement(t *testing.T) {
	for _, c := range cases {
		t.Run(c.customer.name, func(t *testing.T) {
			got := c.customer.Statement()
			if got != c.wantStr {
				t.Errorf("got %s want %s", got, c.wantStr)
			}
		})
	}
}

func TestCustomer_HtmlStatement(t *testing.T) {
	for _, c := range cases {
		t.Run(c.customer.name, func(t *testing.T) {
			got := c.customer.HtmlStatement()
			if got != c.wantHtml {
				t.Errorf("got %s want %s", got, c.wantHtml)
			}
		})
	}
}
