package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Minecraft Server Manager!")
	fmt.Println("Current server: ")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Start server\n2. Stop server\n3. Server status")
		var choice int
		_, error := fmt.Scan(&choice)
		if error != nil {
			fmt.Fprintln(os.Stderr, "Error: please enter a number.")
		}
	}
}