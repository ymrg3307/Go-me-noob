package main

import (
	"fmt"
	"sort"
)

func main() {
	// Fruits List
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of Fruitlist is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Dragonfruit")
	fmt.Println("Fruitlist is : ", fruitList)

	fruitList = append(fruitList[:2])
	fmt.Println("Fruitlist is : ", fruitList)

	fruitList = append(fruitList, "WTH!!!")
	fmt.Println("Fruitlist is : ", fruitList)
	fmt.Println("Fruitlist is : ", fruitList[:5])
	fmt.Println("")

	//High Scores
	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 100
	highScores[2] = 345
	highScores[3] = 265
	// highScores[4] = 777 //throws error as index out range

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"c++", "c", "JS", "c#", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
