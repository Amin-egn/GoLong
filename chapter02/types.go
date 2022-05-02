/*
	the types we have are

	Numbers {
		Integers,
		Float
	}
	Strings
	Booleans
*/
/*
	Integers {
		these are unsigned integers and used for positive integers
		uint8,
		uint16,
		uint32,
		uint64,

		these are signed integers
		int8,
		int16,
		int32,
		int64
	}

	Float {
		the size of float numbers
		float32,
		float64
	}
*/
/*
	the Go operators are

	+	addition

	-	subtraction

	*	multipliction

	/	division

	%	reminder

*/
/*
	the Go logical operators

	&&	and

	||	or

	!	not
*/

package main

import (
	"fmt"
)

func number() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)
}

func strings() {
	fmt.Println(len("Hello, World!")) // show the length of the string
	fmt.Println("Hello, World!"[1])   // access the particular character in the string (byte)
	fmt.Println("Hello, " + "World!") // concatenates two strings together
}

func bools() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(false && true)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(false || true)
	fmt.Println(!true)
	fmt.Println(!false)
}

func main() {
	number()
	strings()
	bools()
}
