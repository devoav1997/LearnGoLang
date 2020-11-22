package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstname string
	lastname string
	contact contactInfo
}

func  main()  {
	jim := person{
		firstname: "jim",
		lastname: "party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v",jim)
}
