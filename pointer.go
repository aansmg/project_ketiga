package main

import "fmt"

// func main() {
// 	var firstNumber *int
// 	var secondNumber *int

// 	_, _ = firstNumber, secondNumber
// }

// func main() {
// 	var firstNumber = 4
// 	var secondNumber = *int == &firstNumber

// 	fmt.Println("firstNumber (value)", firstNumber)
// 	fmt.Println("firstNumber (memori address):", &firstNumber)

// 	fmt.Println("secondtNumber (value)", *secondNumber)
// 	fmt.Println("secondNumber (memori address):", secondNumber)
// }

// func main (){
// 	var firstPerson string = "John"
// 	var secondPerson *string = &firstPerson

// 	fmt.Println("firstPerson (value) :", firstPerson)
// 	fmt.Println("firstPerson (memory address) :", &firstPerson)

// 	fmt.Println("secondPerson (value) :", *secondPerson)
// 	fmt.Println("secondPerson (memory address) :", secondtPerson)

// fmt.Println("\n", string.Repeat("#"30)"\n")

// *secondPerson = "Doe"
// fmt.Println("firstPerson (value) :", firstPerson)
// 	fmt.Println("firstPerson (memory address) :", &firstPerson)

// 	fmt.Println("secondPerson (value) :", *secondPerson)
// 	fmt.Println("secondPerson (memory address) :", secondtPerson)

// }

func main() {
	var a int = 10
	fmt.Println("Before:", a)
	changeValue(&a)
	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
