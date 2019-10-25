package v1

import "fmt"

type Person struct {
	name           string
	officeAreaCode string
	officeNumber   string
}

func (p Person) GetTelephoneNumber() string {
	return fmt.Sprintf("(%s)%s", p.officeAreaCode, p.officeNumber)
}