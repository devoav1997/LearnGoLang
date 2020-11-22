package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstname string
	lastname string
	contactInfo
}

func  main()  {
	jim := person{
		firstname: "jim",
		lastname: "party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}
//about access turn  address to value and value to address
// turn address to value with *address
// turn value to address with &value
func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
