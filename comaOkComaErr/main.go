package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome To Washroom"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Exprience Rate:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks For Rating is : ", input)
}