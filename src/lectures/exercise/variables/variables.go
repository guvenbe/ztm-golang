//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor = "Blue"
	fmt.Println("My favorite color is", favoriteColor)

	birthYear, age := 1966, 56
	fmt.Println("bitheYear is", birthYear, "and age is", age)

	var (
		firstInitial = `J`
		lastInitial  = `G`
	)

	fmt.Println("Initials=", firstInitial, lastInitial)

	fmt.Println(5 + 5)
	fmt.Println(5 - 3)
	fmt.Println(25 / 5)
	fmt.Println(5 * 5)

	var ageInDays int
	ageInDays = 365 * age
	fmt.Println("age in day:", ageInDays)

}
