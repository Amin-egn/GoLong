package main

import (
	"fmt"
)

func for_statement() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func if_statement() {
	name := "Amin"

	if name == "Amin" {
		fmt.Println("Your name is Amin")
	} else if name == "mohammad" {
		fmt.Println("Your name is Mohammad")
	} else {
		fmt.Println("I don't know your name")
	}

}

func switch_statement() {
	i := 1
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}
}

func main() {
	for_statement()
	if_statement()
	switch_statement()
}
