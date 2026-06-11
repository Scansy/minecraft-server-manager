package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
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
		printMenuItems(menuItems)
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

// Prints the menu items
func printMenuItems(items []MenuItem) {
	for i, item := range items {
		fmt.Printf("%v. %v\n", i+1, item.title)
	} 
}

// Creates the menu items
func createMenuItems() []MenuItem {
	var menuItems []MenuItem = make([]MenuItem, 0, 10)
	menuItems = append(menuItems, createStartServerMenu()) 
	menuItems = append(menuItems, createStopServerMenu())	
	menuItems = append(menuItems, MenuItem{title: "Server status", action: func(){}})	
	menuItems = append(menuItems, MenuItem{title: "Quit", action: func(){
		fmt.Println("Bye!")
		os.Exit(0)
		}})	

	return menuItems
}

// Creates the Start server menu
func createStartServerMenu() MenuItem {
	title := "Start server"
	action := func() {
		serverProcess := exec.Command("java", "-Xmx4G", "-Xms4G", "-jar", "server.jar", "nogui")	
		serverProcess.Dir = "./server"
		serverProcess.Stderr = os.Stderr
		serverProcess.Stdout = os.Stdout
		err := serverProcess.Start()
		if err != nil {
			printError("Failed to start server.")
			return
		}
		pid := serverProcess.Process.Pid
		os.WriteFile("./pid.txt", []byte(fmt.Sprintf("%v", pid)), 0644)
	}

	return MenuItem{title, action}
}

// Creates the Stop server menu
func createStopServerMenu() MenuItem {
	title := "Stop server"
	action := func() {
		pidByte, err := os.ReadFile("./pid.txt")
		if err != nil {
			printError("Couldn't find pid.txt")
		}

		pid, err := strconv.Atoi(string(pidByte[:]))
		if err != nil {
			printError("Couldn't convert PID into int")
		}
		
		serverProcess, err := os.FindProcess(pid)
		if err != nil {
			printError("Couldn't find server process")
		}

		err = serverProcess.Signal(syscall.SIGTERM)
		if err != nil {
			printError("Couldn't SIGTERM the server")
		}
		serverProcess.Wait()
	}
	
	return MenuItem{title, action}
}

// Validates if the input falls within min and max, inclusive 
func validateUserChoice(input, min, max int) bool {
	return input >= min && input <= max
}

// Quickly prints to stderr
func printError(msg string) {
	fmt.Fprintln(os.Stderr, "Error: ", msg)
}