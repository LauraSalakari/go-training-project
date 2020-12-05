package main

import "fmt"

// we use interface to declare a method set!
// anything that matches this rule set now becomes part of this type
type bot interface {
	getGreeting() string
}

// because the englishBot and spanishBot have a method called getGreeting() which returns a string,
// these bots are now "honorary" members of the type bot
// and thus can now call the printGreeting() function which takes a variable type bot as arg/param
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating English greeting
	return "Hi There!"
}

// if we are not making use of the receiver variable, the var name for it can be dropped!
func (spanishBot) getGreeting() string {
	// VERY custom logic for generating Spanish greeting
	return "Hola!"
}
