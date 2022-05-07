/*
	structs are a type that contains name fields
	for define a struct we follow this way:
	type Name struct {}		use type keyword, a name and struct keyword and define fields
	type Circle struct {
		x, y, r float64
	}


	interfaces are a way to combine similarities
	for define an interface we follow this way:
	type Name interface {}		use type keyword, a name and interface keyword but instead of
								define fields, we define method set
	type Shape interface {
		area() float64
	}

*/

package main

import (
	"fmt"
	"math"
)

/* ***************** STRUCT ***************** */

type Circle struct {
	x float64
	y float64
	r float64
	// or x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func print_values() {
	c := Circle{0, 0, 5}
	fmt.Println(c, c.x, c.y, c.r)
}

// Circle area
func circle_area(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// also Circle area but method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// rectangle has area method
func (r *Rectangle) area() float64 {
	a1 := r.x1 * r.y1
	return a1
}

type Person struct {
	name string
}

func (p *Person) introduce() {
	fmt.Println("Hi, my name is", p.name)
}

// Teacher has nested struct
type Teacher struct {
	person Person // teacher has a Person
	lesson string
}

// Student has promoted field that directly access to Person method
type Student struct {
	Person // student is a Person
	age    int
}

func teacher_struct() {
	t := Teacher{
		person: Person{"Richard"},
		lesson: "math",
	}
	fmt.Println(t, t.person.name)
	t.person.introduce()
}

func student_struct() {
	s := Student{
		Person: Person{"Adrian"}, // promote field
		age:    18,
	}
	fmt.Println(s, s.name)
	// direct access
	s.introduce()
}

/* ***************** INTERFACE ***************** */

type Shape interface {
	area() float64
}

func total_area(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{1, 2, 3}
	r := Rectangle{2, 3, 2, 3}
	fmt.Println(total_area(&c, &r))
}
