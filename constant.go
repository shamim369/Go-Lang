package main
import "fmt"

/*
	Constant Rules
		Constant names follow the same naming rules as variables
		Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
		Constants can be declared both inside and outside of a function

	There are two types of constants:
		Typed constants
		Untyped constants
*/

func main() {

	const PI = 3.1416 // Typed constants
	const SECRET_KEY = "5976@54632" // Typed constants

	// Multiple Constants Declaration
	const (
		A int = 1
		B = 3.14
		C = "Hi!"
	)

	fmt.Println(A)
}