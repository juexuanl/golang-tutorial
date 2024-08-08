package main //if the package is main, it tells the GO compiler this code should be compiled into executable program at the end

import "fmt"

func main() {
	//Println
	name := "Zoey"
	age := 36
	fmt.Println("My name is", name, "my age is", age)

	// Printf(formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.3f points! \n", 2.3349)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
