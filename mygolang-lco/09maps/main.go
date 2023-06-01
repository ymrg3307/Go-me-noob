package main

import "fmt"

func main() {
	//make(map[string]string) also works
	languages := map[string]string{}
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	fmt.Println("Languages list is: ")
	for _, value := range languages {
		fmt.Println(value)
	}

}
