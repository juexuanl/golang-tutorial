package main //if the package is main, it tells the GO compiler this code should be compiled into executable program at the end

import "fmt"

func main() {
	// strings
	var name1 string = "hello"
	var name2 = "world"
	var name3 string
	
	fmt.Println(name1, name2, name3)
	name3 = "!"
	fmt.Println(name1, name2, name3)

	name4 := "I love you" //can't be used outside of a function
	fmt.Println(name4)

	// intergers
	var age1 int = 20
	var age2 = 30
	age3 := 40
	fmt.Println(age1, age2, age3)

	// bits $ memory
	var num1 int8 = 123
	var num2 = -1283
	fmt.Println(num1, num2)
}