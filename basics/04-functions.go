package main

import "fmt"

// Notice that the type comes after the variable name.
func add(x int, y int) int {
	return x + y
}

//When two or more consecutive named function parameters share a type, 
//you can omit the type from all but the last.
func sub(x, y int) int {
	return x - y
}

//A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(10,5))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
