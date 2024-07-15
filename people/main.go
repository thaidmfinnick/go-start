package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newName string) {
	p.firstName = newName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "fni",
		contact: contactInfo{
			email: "hehe@gmail.com",
			zip:   84,
		},
	}
	jim.updateName("FInnick")
	jim.print()

}
