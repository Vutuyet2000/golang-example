package main

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// func main() {
// 	var alex person
// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson"

// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Party",
// 		contact: contactInfo{
// 			email:   "jim@gmail.com",
// 			zipCode: 94000,
// 		},
// 	}

// 	// alex := person{firstName: "Alex", lastName: "Anderson"}
// 	fmt.Println(alex)
// 	fmt.Println(jim)

// 	jimPointer := &jim
// 	jimPointer.updateName("jimmy")

// 	fmt.Println(jimPointer)
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
