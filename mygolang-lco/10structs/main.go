package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; no super or parent

	mani := User{"maniratna", true, 23}

	fmt.Println(mani)
	fmt.Printf("mani details are: %+v\n", mani)
	fmt.Printf("Name is %v and my alive status is %v \n", mani.Name, mani.Status)
}

type User struct {
	Name   string
	Status bool
	Age    int
}
