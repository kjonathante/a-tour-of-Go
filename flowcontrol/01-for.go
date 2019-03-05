package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
/*
The basic for loop has three components separated by semicolons.

The init statement will often be a short variable declaration, 
and the variables declared there are visible only in the scope of the for statement.

There are no parentheses surrounding the three components of the for statement 
and the braces { } are always required.

The init and post statements are optional.
You can drop the semicolons.

	sum := 1
	for sum < 1000 {
		sum += sum
	}
*/