# GoLang Cheat Sheet
	
	strings ==================================

	var nameOne string = "Hello" // Explicitly declare a string
	var nameTwo string = "World" // Automatically inferring the type
	var nameThree string // Setting up a string-only variable for future use
	nameThree = "and"
	nameFour := "Goodbye" // Shorthand for nameOne or nameTwo (infers the type), only used first time it's assigned (can't be used outside of a function)

	fmt.Println(nameOne + " " + nameTwo + " " + nameThree + " " + nameFour)

	integers ==================================
	var ageOne int = 30 // Explicitly declare an integer
	var ageTwo int = 40 // Automatically inferring the type
	ageThree := 50      // Shorthand assignment

	fmt.Println(ageOne, ageTwo, ageThree)

	Bits and Memory ===========================

	var bitNum int8 = 25 // numbers within 8-bit range -127 to 127
	var bitNum2 uint8 = 25 // numbers within 8-bit range 0 to 255

	Floats ====================================

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 124087950177644710644764722.7 // More precision, mostly used, not much of a memory hit
	NOTE: float division is not supported by the usual methods

	Print =====================================

	fmt.Print("hello, ")           // No new line
	fmt.Print("world!\n")          // This will be on the same line (with a new line after)
	fmt.Println("How's it going?") // This will be on a new line
	age := 26
	name := "Harry"
	fmt.Printf("My name is %v and I'm %v years old\n", name, age) //formatted strings, %v (variable) is replaced with the value of the variable
	// We also have %q for quotes, %T gives the type, %f for floats, %e for exponents, %x for hex, %b for binary
	fmt.Printf("Age is of type %T", age)
	// var str = fmt.Sprintf("My name is %v and I'm %v years old\n", name, age) // Saves a formatted string to a variable

	Arrays and Slices ===========================

	var ages [3]int = [3]int{10, 20, 30} // [3]int declares an array of 3 integers, right hand side must include [3]int
	var shortHandAges = [3]int{10, 20, 30} // Shorthand for above
	names := []string{"Harry", "Ron", "Hermione"} // It's a slice (an array of undefined length)
	You can edit the elements of both arrays/slices, but you can only append to slices
	names = append(names, "Neville")
	sliceRange := names[1:3] // You can take slices of arrays/slices

	Loops ========================================

	names := []string{"Harry", "Ron", "Hermione"}
	for index, name := range names { // Classic for loop
		fmt.Printf("Name %d is %s\n", index, name)
	}
	index := 0
	for index < len(names) { // While loop alternative
		fmt.Printf("Name %d is %s\n", index, names[index])
		index++
	}
	You also have the classic
	for i := 0; i < len(names); i++ {
	fmt.Printf("Name %d is %s\n", i, names[i])
	}
	You can also have an infinite loop with some exit condition

	If/Else =======================================

	num := 8
	if num > 10 {
	    fmt.Println("num is greater than 10")
	} else {
	    fmt.Println("num is less than or equal to 10")
	}

	Switch Statements =============================

	city := "London"

	switch city {
		case "New York":
			fmt.Println("New York")
	    case "London":
			fmt.Println("London")
		case "Paris":
			fmt.Println("Paris")
	    default:
			fmt.Println("Invalid city")
	}

	Logical operators =============================

	and is &&
	or is ||
	not is !
	xor is ^
	xnor is ~
	nand is &
	nor is |
	xnor is ^

	User Input =====================================

	fmt.Println("Enter a number: ")
	var userInput int
	fmt.Scan(&userInput) // This will take a user input and store it in variable userInput

	Scope Levels ===================================

	Local Scope (declaration within a function or within a block - ie for loop)
	Package Scope (Multiple functions can access these variables)
	Global Scope (All functions can access these variables, using capital variable name)

	Maps ===========================================

	Maps unique keys to values (kind of like a dictionary in Python)
	Cannot mix data types in a map
	var userData = make(map[string]string) // Creating an empty map
	userData["name"] = "Harry"
	strconv.FormatUint((uint64(26)), 10) // Formats a uint64 into a string, as we cannot mix data types in a map
	userData["city"] = "Bath"

	Then you can append this to other variables etc.

	Structs ========================================

	A struct is a collection of variables and methods (it's like maps but can have multiple mixed data types)

	type UserData struct { // This is a custom data type (a class in Python)
		firstName string
	    lastName  string
	    age       int
	}

