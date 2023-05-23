package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza Shop!")
	fmt.Println("Please provide your rating: ")
	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	fmt.Println("Thanks for rating, ", val)
	numRating, err := strconv.ParseFloat(val, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your rating, ", numRating+1)
	}

}
