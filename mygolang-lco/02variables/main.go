package main

import "fmt"

//No var style only works in methods or functions
// Token := "aeBjkuIOPuK5678Polervbn"

const Token string = "aeBjkuIOPuK5678Polervbn"

func main() {
	var username string = "mani"
	fmt.Println(username)
	fmt.Printf("username is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("isLoggedIn is of type: %T \n", smallVal)

	var smallFloat float32 = 255.12345678909978645312
	fmt.Println(smallFloat)
	fmt.Printf("isLoggedIn is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.12345678909978645312
	fmt.Println(bigFloat)
	fmt.Printf("isLoggedIn is of type: %T \n", bigFloat)

	var defaultValue1 string
	fmt.Println(defaultValue1)
	fmt.Printf("defaultValue1 is of type: %T \n", defaultValue1)

	var defaultValue2 bool
	fmt.Println(defaultValue2)
	fmt.Printf("defaultValue2 is of type: %T \n", defaultValue2)

	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("website is of type: %T \n", website)
	//website = 3

	numberOfUsers := 30000
	fmt.Println(numberOfUsers)
	fmt.Printf("numberOfUsers is of type: %T \n", numberOfUsers)

	fmt.Println(Token)
	fmt.Printf("Token is of type: %T \n", Token)
}
