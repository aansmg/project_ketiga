package main

import (
	"fmt"
	// "strings"
)

// type Employee struck{
// 	name string
// 	age int
// 	division string
// }

// func main (){
// 	var employee1 Employee{}
// employee.name = "Aan"
// employee.age = 30
// employee.division = "Curriculum Developer"

// var employee2 = Employee{name:"Budi", age:31, division:"Finance"}

// fmt.Println(employee.name)
// fmt.Println(employee.age)
// fmt.Println(employee.division)

// fmt.Println("Employee1 : %+v\n", employee1)
// fmt.Println("Employee2 : %+v\n", employee2)

// }

// func main (){
// 	var employee1 = Employee{name:"Budi", age:31, division:"Finance"}
// 	var employee2 *Employee = &employee1

// 	fmt.Println("Employee1 name: ", employee1.name)
// 	fmt.Println("Employee2 name: ", employee2.name)

// 	fmt.Println(strings.Repeat("#"),20)

// 	employee2.name ="ganti"
// 	fmt.Println("Employee1 name: ", employee1.name)
// 	fmt.Println("Employee2 name: ", employee2.name)

// }



// type Employee struct {
// 	division string
// 	person   Person
// }

// func main() {
// 	var employee1 = Employee{}
// 	employee1.person.name = "Aan"
// 	employee1.person.age = "30"
// 	employee1.division = "Curriculum Developer"
// 	fmt.Println("%+v", employee1)
// }


type Person struct {
	name string
	age  int
}


func main(){
	var people =[] Person{
		{name:"Aan", age : 30},
		{name:"Nani", age : 34},
		{name:"Mail", age : 39},


		for_, v:= range people {
			fmt.Printf("%+v\n", v)
		}

	}
}