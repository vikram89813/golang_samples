package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// fmt.Println("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "test@test.com",
			zipCode: 94300,
		},
	}

	// test := person{}

	// test.contact.email = "test1@test.com"
	// test.contact.zipCode = 560037
	// test.firstName = "t1"
	// test.lastName = "t2"

	// fmt.Println(test)

	// jptr := &jim
	// jptr.updateName("jimmy")

	jim.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
