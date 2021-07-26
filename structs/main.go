package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// rocky := person{firstName: "Rocky", lastName: "Nguyen"}
	// var rocky person
	// rocky.firstName = "Rocky"
	// rocky.lastName = "Nguyen"

	rocky := person{
		firstName: "Rocky",
		lastName:  "Nguyen",
		contactInfo: contactInfo{
			email:   "rocky@a.ly",
			zipCode: 100000,
		},
	}

	// fmt.Println(rocky)
	// fmt.Printf("%+v", rocky)

	// &variable get the memory address of the value this variable is pointing at
	// *pointer get the value of this memory address is pointing at
	// rockyPointer := &rocky
	// rockyPointer.updateName("Son", "Nguyen")

	rocky.updateName("Son", "Nguyen")
	rocky.print()
}

// // *[type] type of the receiver is pointer to [type]
// func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
// 	(*pointerToPerson).firstName = newFirstName
// 	(*pointerToPerson).lastName = newLastName
// }

func (p *person) updateName(newFirstName string, newLastName string) {
	p.firstName = newFirstName
	p.lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
