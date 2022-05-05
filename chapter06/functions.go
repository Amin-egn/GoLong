/*
	use 'func' keyword to define a function
	functions are stacking called
	and look like this:

	func function name( function parameters: param1 type, param2 type ) (return value: reutrn1 type) {
		implementation
	}
*/

package main

import (
	"fmt"
)

func not_implement() {
	panic("NOT IMPLEMENT!!!")
}

func args_function(args ...int) int {
	var total int
	for _, n := range args {
		total += n
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first_function() {
	fmt.Println("first one")
}

func second_function() {
	fmt.Println("second one")
}

func panic_recover() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC!")
}

func zero(xptr *int){
	*xptr = 0
}


func main() {
	x := new(int)
	zero(x)
	fmt.Println(x)
}
