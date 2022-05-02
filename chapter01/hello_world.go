// every Go file must has a package name
package main

// this is for import other packages include standard
import (
	"fmt" // a standard package and it's shorthand of 'format'
)

/*
	func stands for function to declare functions
	main stands for function's name
	and at last in { } we do what we want
*/
func main() {
	fmt.Println("Hello, World!")
}

/*
	in terminal .../GoLong/chapter01/go_01.go> go run ./go_01.go
	ouput: Hello, World!
*/
