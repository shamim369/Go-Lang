package main
import "fmt"

/*
	In Go, there are different types of variables, for example:
		int- stores integers (whole numbers), such as 123 or -123
		float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
		string - stores text, such as "Hello World". String values are surrounded by double quotes
		bool- stores values with two states: true or false

		1. With the var keyword:
		var variablename type = value

		2. With the := sign:
		variablename := value

	Go variable naming rules:
		A variable name must start with a letter or an underscore character (_)
		A variable name cannot start with a digit
		A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		Variable names are case-sensitive (age, Age and AGE are three different variables)
		There is no limit on the length of the variable name
		A variable name cannot contain spaces
		The variable name cannot be any Go keywords

*/

func main() {
	// Variable Declaration With Initial Value
	var name string 	= "Md. Shamim Hossain" 	// type is string
	var age int 		= 35					// type is integer
	var salary float32 	= 19.35 				// type is floating number
	var active bool 	= true 					// type is boolean
	var student2 		= "Jane" 				//type is inferred
  	x := 2 										//type is inferred
	fmt.Println(name, age, salary, active, student2, x)

	// Variable Declaration Without Initial Value
	var a string
	var b int
	var c float32
	var d bool
	fmt.Println(a, b, c, d)

}