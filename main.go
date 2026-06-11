package main

import (
	"fmt"
	"os"
)

type MenuItem struct {
	title string
	action func()
}

func main() {
	// Setup
	var menuItems []MenuItem = createMenuItems()

	// Variables
	var numberOfChoices int = len(menuItems) 

	fmt.Println("Welcome to the Minecraft Server Manager!")
	fmt.Println("Current server: ")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Start server\n2. Stop server\n3. Server status\n4. Quit")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			printError("Please enter a number")
		} else if !validateUserChoice(choice, 1, numberOfChoices) {
			printError(fmt.Sprintf("Please enter between 1 and %v", numberOfChoices))
		} else {
			choice-- 
			chosenMenuItem := menuItems[choice]
			chosenMenuItem.action()
		}

	}
}

// Creates the menu items
func createMenuItems() []MenuItem {
	var menuItems []MenuItem = make([]MenuItem, 0, 10)
	menuItems = append(menuItems, MenuItem{title: "Start server", action: func(){}})	
	menuItems = append(menuItems, MenuItem{title: "Stop server", action: func(){}})	
	menuItems = append(menuItems, MenuItem{title: "Server status", action: func(){}})	
	menuItems = append(menuItems, MenuItem{title: "Quit", action: func(){
		fmt.Println("Bye!")
		os.Exit(0)
		}})	

	return menuItems
}

// Validates if the input falls within min and max, inclusive 
func validateUserChoice(input, min, max int) bool {
	return input >= min && input <= max
}

// Quickly prints to stderr
func printError(msg string) {
	fmt.Fprintln(os.Stderr, "Error: ", msg)
}