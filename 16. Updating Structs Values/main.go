package main

import "fmt"

type person struct {
	firstname string
	lastname string
}

func  main()  {
	var alex person

	alex.firstname = "Alex"
	alex.lastname = "Devo"

	fmt.Println(alex)
	fmt.Printf("%+v",alex)
}
