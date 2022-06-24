package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

// Goodbye returns a goodbye for the named person
func Goodbye(name string) string {
   // Returns a goodbye that embeds the the name in a message.
   message := fmt.Sprintf("Goodbye %v!", name)

   return message
}

