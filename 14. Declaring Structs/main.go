package main

import "fmt"

type person struct {
	firstname string
	lastname string
}

func  main()  {
	alex := person{firstname: "Alex", lastname: "Anderson"}
	fmt.Println(alex)
}
