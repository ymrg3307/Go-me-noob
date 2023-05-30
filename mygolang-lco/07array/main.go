package main

import "fmt"

func main() {
	var fruitList [3]string
	fruitList[0] = "mango"
	fruitList[2] = "apple"

	fmt.Println("fruitList is : ", fruitList)
	fmt.Println("What is value at index 1 as it is not initialized? : ", fruitList[1])
	fmt.Println("fruitList is of size: ", len(fruitList))

	var veggyList = [3]string{"potato", "onion", "beans"}
	fmt.Println("veggyList is : ", veggyList)
	fmt.Println("What is value at index 1? : ", veggyList[1])
	fmt.Println("veggyList is of size: ", len(veggyList))
}
