package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Jiffery",
		contact: ContactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}

	jim.print()
	jim.updateName("Jimbo")
	jim.print()
}

func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
