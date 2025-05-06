package main

import "fmt"

func main() {
	var username string = "Mehedi Hasan Shamim"
	fmt.Println(username)
	fmt.Printf("Variable Type Of : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable Type Of : %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable Type Of : %T \n", smallValue)

	var smallFloat float32 = 255.4444444444444444
	fmt.Println(smallFloat)
	fmt.Printf("Variable Type Of : %T \n", smallFloat)

	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("Variable Type Of : %T \n", anotherVariable)

	numberCount := 4000.99
	fmt.Println(numberCount)

}