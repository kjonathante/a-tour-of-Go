package main

import "fmt"

/*
The var statement declares a list of variables, the type is last.

A var statement can be at package or function level.
*/
var c, python, java bool // package level

/*
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; 
the variable will take the type of the initializer.
*/
var x, y int = 1, 2

func main() {
	var i int // function level
	var ruby, scala, swift = true, false, "no!"
	
	fmt.Println(i, c, python, java)
	fmt.Println(x,y,ruby,scala,swift)
}
