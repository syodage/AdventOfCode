package main

import "fmt"

func main() {
	fmt.Println("Hello AoC 2021!")
	if err := Day01(); err != nil {
		fmt.Printf("Error returned: %v", err)
	}
}
