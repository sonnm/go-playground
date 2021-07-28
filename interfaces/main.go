package main

import "fmt"

type bot interface {
	// getGreeting(string, int) (string, error)
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// eb.printGreeting()
	// sb.printGreeting()

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func (eb englishBot) printGreeting() {
// 	fmt.Println(eb.getGreeting())
// }

// func (sb spanishBot) printGreeting() {
// 	fmt.Println(sb.getGreeting())
// }
