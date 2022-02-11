package main
import "fmt"

func main() {
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
	*/

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