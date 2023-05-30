package main

import "fmt"

func main() {
	var ptr1 *int
	fmt.Println("ptr1 is : ", ptr1)

	myNum := 25
	ptr1 = &myNum

	fmt.Println("memory address of myNum: ", ptr1)
	fmt.Println("value at the above memory address: ", *ptr1)

	*ptr1 = *ptr1 + 2
	fmt.Println("modified myNum : ", myNum)
}
