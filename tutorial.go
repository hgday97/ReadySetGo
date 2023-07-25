package main

import "fmt"

func main() {

	// strings ==================================

	// var nameOne string = "Hello" // Explicitly declare a string
	// var nameTwo string = "World" // Automatically inferring the type
	// var nameThree string // Setting up a string-only variable for future use
	// nameThree = "and"
	// nameFour := "Goodbye" // Shorthand for nameOne or nameTwo (infers the type), only used first time it's assigned (can't be used outside of a function)

	// fmt.Println(nameOne + " " + nameTwo + " " + nameThree + " " + nameFour)

	// integers ==================================
	// var ageOne int = 30 // Explicitly declare an integer
	// var ageTwo int = 40 // Automatically inferring the type
	// ageThree := 50      // Shorthand assignment

	// fmt.Println(ageOne, ageTwo, ageThree)

	// Bits and Memory ===========================

	// var bitNum int8 = 25 // numbers within 8-bit range -127 to 127
	// var bitNum2 uint8 = 25 // numbers within 8-bit range 0 to 255

	// Floats ====================================

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 124087950177644710644764722.7 // More precision, mostly used, not much of a memory hit
	// NOTE: float division is not supported by the usual methods

	// Print =====================================

	fmt.Print("hello, ")           // No new line
	fmt.Print("world!\n")          // This will be on the same line (with a new line after)
	fmt.Println("How's it going?") // This will be on a new line
	age := 26
	name := "Harry"
	fmt.Printf("My name is %v and I'm %v years old\n", name, age) //formatted strings, %v (variable) is replaced with the value of the variable
	// We also have %q for quotes, %T gives the type, %f for floats, %e for exponents, %x for hex, %b for binary
	fmt.Printf("Age is of type %T", age)
	// var str = fmt.Sprintf("My name is %v and I'm %v years old\n", name, age) // Saves a formatted string to a variable
}
