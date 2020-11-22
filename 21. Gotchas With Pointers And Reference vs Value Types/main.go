package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func updateSlice(s []string){
	s[0] = "Bye"
}

//Value Types : int , float, string, bool, structs -> Use pointers to change these things in a function
//Reference Types : slices, maps, channels, pointers, functions -> Don't worry about pointers with these