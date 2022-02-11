package main
import "fmt"

func main() {
	// Declares two arrays with defined lengths:
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// Declares two arrays with inferred lengths:
	var arr3 = [...]int{1,2,3}
	arr4 := [...]int{4,5,6,7,8}
	fmt.Println(arr3)
	fmt.Println(arr4)

	// Declares an array of strings:
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  	fmt.Print(cars)

	// Initializes only the second and third elements of the array: 
	arr5 := [5]int{1:10,2:40}
  	fmt.Println(arr5)

	// Find the Length of an Array
	arr6 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr7 := [...]int{1,2,3,4,5,6}
	fmt.Println(len(arr6))
	fmt.Println(len(arr7))

	// Access Elements of an Array
	prices := [3]int{10,20,30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// // Change Elements of an Array
	prices[2] = 50
	fmt.Println(prices)

	// Array Initialization
	arr8 := [5]int{} 			// not initialized
	arr9 := [5]int{1,2} 		// partially initialized
	arr10 := [5]int{1,2,3,4,5} 	// fully initialized

	fmt.Println(arr8)
	fmt.Println(arr9)
	fmt.Println(arr10)
}