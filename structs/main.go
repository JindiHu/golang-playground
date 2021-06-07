package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.updateContactEmail("jimmy@gmail.com")
	jim.print()
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *Person) updateContactEmail(email string) {
	(*p).contact.email = email
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
