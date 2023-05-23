package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza: ")

	//comma ok or error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", strings.TrimSpace(input))
	fmt.Printf("Type of this rating is %T \n", input)
}
