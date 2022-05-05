/*
	array is a numbered sequence of elements of a single type with a fixed lenth
	array starts with 0 and indexable
	to define array:

	var x [n]int					# array of n element with int type
	var s [n]float64 or float32 	# array of n element with float type
	var y [n]string					# array of n element with string type
	var z [n]bool					# array of n element with bool type

	default values of int and float64 array with n element is [0 0 ... 0]
	default values of string array with n element is ["" "" ... ""]
	default values of bool array with n element is [false false ... false]

*/
/*
	slices are same as arrays but unlike arrays not have length and length may has change
	var x []int
	var s []float64 or float32
	var y []string
	var z []bool

	default values of all types are []
*/
/*
	maps are unordered collection of key-value pairs
	maps are also called hash tables or dictionaries

	var x map[string]int
*/

package main

import (
	"fmt"
)

func structure_array() {
	var x [5]int
	var s [5]float64
	var y [5]string
	var z [5]bool
	fmt.Println(x, s, y, z)
}

func structure_slice() {
	var x []int
	var s []float64
	var y []string
	var z []bool
	fmt.Println(x, s, y, z, len(x), len(y))
}

func structure_test() {
	arr := [5]int{1, 2, 3, 4, 5}
	x := arr[1:3]
	fmt.Println(x)
}

func slice_append() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, "has changed to", slice2)
}

func slice_copy() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func structure_map() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x, x["key"])
}

func map_delete() {
	numbers := make(map[string]int)
	numbers["Zero"] = 0
	numbers["One"] = 1
	numbers["Two"] = 2
	numbers["Three"] = 3
	numbers["Four"] = 4
	numbers["Five"] = 5
	numbers["Six"] = 6
	numbers["Seven"] = 7
	numbers["Eight"] = 8
	numbers["Nine"] = 9
	delete(numbers, "Zero")
	fmt.Println(numbers, numbers["Five"])
}

func map_lookup() {
	// shorter way
	numbers := map[string]int{
		"Zero":  0,
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Six":   6,
		"Seven": 7,
		"Eight": 8,
		"Nine":  9,
	}
	// show default value for non existing key
	fmt.Println(numbers["Nine"], numbers["Ten"])
	num, ok := numbers["Ten"]
	fmt.Println(num, ok)
	/*	better way
		if key value exist 
	*/
	if num, ok := numbers["Ten"]; ok {
		fmt.Println(num, ok)
	}
}

func map_nested() {
	room_temp := map[string]map[string]int{
		"room_101": map[string]int{
			"c": 25,
			"f": 78,
		},
		"room_202": map[string]int{
			"c": 26,
			"f": 83,
		},
		"room_303": map[string]int{
			"c": 23,
			"f": 74,
		},
	}
	fmt.Println(room_temp)
	if room, ok := room_temp["room_404"]; ok{
		fmt.Println(room["c"], room["f"])
	}
}

func main() {
	map_nested()
}
