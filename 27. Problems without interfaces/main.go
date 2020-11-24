package main

import "fmt"

type englisthBot struct{}
type spanishBot struct{}


func main(){
	eb := englisthBot{}
	sb := spanishBot{}

	printGreeting(eb)
	//printGreeting(sb)

}


func printGreeting(eb englisthBot){
	fmt.Println(eb.getGreeting())
}

//func printGreeting(sb spanishBot){
//	fmt.Println(sb.getGreeting())
//}

func (englisthBot) getGreeting() string {
	//Very costum logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}