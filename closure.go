package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var evenNumbers = func(numbers ...int) []int {
// 		var result []int

// 		for _, v := range numbers {
// 			if v%2 == 0 {
// 				result = append(result, v)
// 			}
// 		}
// 		return result
// 	}
// 	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
// 	fmt.Println(evenNumbers(numbers...))

// }

// func main() {
// 	var isPalindrome = func(str string) bool {
// 		var temp string = ""

// 		for i := len(str) - 1; i >= 0; i-- {
// 			temp += string(byte(str[i]))
// 		}
// 		return temp == str
// 	}("katak")
// 	fmt.Println(isPalindrome)
// }


// func main(){
// 	var studentLists = []string{"Ani", "Abang", "Adi", "Nani"}
// 	var find = findStudent(studentLists)
// 	fmt.Println(find("Ani"))
// }
// func findStudent(students []string) func (string)string{
// 	return func(s string)string {
// 		var student string
// 		var position int

// 		for i, v := range students{
// 			if strings.ToLower(v)==strings.ToLower(s)
// 			student = v
// 			position = i
// 			break
// 		}
// 	}
// 	if student == ""{
// 		return fmt.Sprintf("%s doesnt exist", s)
// 	}
// 	return fmt.Sprintf("We found %s at position %d", s, position+1)
// }

func main(){
	var numbers := []int{2,5,8,10,3,99,23}

	var findOddNumbers(numbers, func(number int))bool{
		return number%2 !=0
	})
	fmt.Println("Total Odd Numbers", find)
}

func findOddNumbers(numbers []int, callback func(int))bool{
	var totalOddNumbers int

	for _, v:= range numbers {
		if callback (v){
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}