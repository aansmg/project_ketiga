package main

import (
	"fmt"
)

// func main() {
// 	greet("Hacktiv8", "Jalan Iskandar Muda")
// function with return
// 	var names = []string("Hacktiv8", "Jakarta")
// 	var printMsg = greet ("Hei", names)

// 	fmt.Println(printMsg)
// }
// func greet(msg string, names []string) string) {
// 	// fmt.Printf("Hello There! Myname is %s and i'm %d years old", name age)
// 	// fmt.Println("Hello", name)
// 	// fmt.Println("I live", address)
// 	var joinStr = string.Join(names,"")
// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)
// 	fmt.Println(printMsg)
// }
// func main() {
// 	var diameter float64 = 15
// 	var area, circumference = calculate(diameter)

// 	fmt.Println("Area:", area)
// 	fmt.Println("circumference:", circumference)
// }

// func calculate(d float64) (float64, float64) {
// 	var area float64 = math.Pi * math.Pow(d/2, 2)

// 	var circumference = math.Pi * d
// 	return area, circumference
// }

// func calculate(d float64) (area float64, circumference float64) {
// 	area = math.Pi * math.Pow(d/2, 2)

// 	circumference = math.Pi * d
// 	return
// }

// func main() {
// 	studentLists := print("Hacktiv8", "Hacktivkidz", "Kode")
// 	fmt.Printf("%v", studentLists)
// }

// func print(names ...string) []map[string]string {
// 	var result []map[string]string

// 	for i, v := range names {
// 		key := fmt.Sprintf("student%d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)

// 	}
// 	return result
// }

func main() {
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)
	fmt.Println("result:", result)
}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	profile("A", "B", "C", "D", "E")
}
func profile(name string, FavFoods ...string) {
	mergeFavFoods := string.join(FavFoods, ",")

	fmt.Println("Hello")
}
