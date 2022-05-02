/*
	variables are storage locations with a specific type
	some variables aren't have data type, called inferred
	several ways to define a variable like:

	var x string
	var x string = "Hello, World!"
	x := "Hello, World!" (just inside function)

	it can be group
	var i, j, p, k = "amin", 1, 3.14, true

	var (
		i = "amin"
		j = 1
		p = 3.14
		k = true
	)

	variables have default value

	string type is ""
	numeric type is 0
	bool type is false

	constants are same as variables but the difference is they cannot be changed later

	const x string
	const x string = "Hello, World!"

*/
package main

import (
	"fmt"
	"reflect"
)

// global variable
var g = "I'm g"

// declare variables in local
func scope_1() {
	var x string
	var y string = "I'm y"
	q := "I'm q"
	p := 1
	fmt.Println(x, y, q, p, g)
}

// it doesn't access to scope_1
// func scope_2() {
// 	fmt.Println(x, y, q, p)
// }

func scope_3() {
	fmt.Println(g)
}

func reflect_type() {
	x := 1
	y := 1.2
	z := true
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(z))
}

func var_difference() {
	var x string = "Hello"
	fmt.Println(x)
	x = "Hello, World!"
	fmt.Println(x)
	const y float64 = 3.14
	fmt.Println(y)
	// y = 3.15 cause error

}

func main() {
	var_difference()
}
