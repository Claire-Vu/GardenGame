package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// MAIN GAME
func main() {
	ClearConsole()
	fmt.Println("Welcome to the Farming Simulation Game!")
	fmt.Println("Are you a new player or continuing? (Type 'new' or 'continue')")
	var playerType string
	fmt.Scanln(&playerType) // taking input from user
	playerType = strings.ToLower(playerType)

	// LOADS IN PLAYER
	var player Player // variable of type Player to store player info
	if playerType == "new" {
		player = HandleNewPlayer()
		ClearConsole()
		fmt.Println("\n---Current Status---")
		player.DisplayInfo()

	} else if playerType == "continue" {
		player = HandleExistingPlayer()
		ClearConsole()
		fmt.Printf("\nWelcome back, %s!\n", player.Username)
		fmt.Println("\n---Current Status---")
		player.DisplayInfo()

	} else {
		fmt.Println("Invalid input.")
		return
	}

	// RUNS GAME
	var gameRunning bool = true
	for gameRunning {
		// prints command options
		player.printMenu()

		// Error Message
		var errMessage error = nil

		// Ask what user wants to do
		var choiceStr string
		fmt.Print("Enter your choice (1-6): ")
		fmt.Scanln(&choiceStr)

		// If user doesn't enter an int
		choice, errChoice := strconv.Atoi(choiceStr)
		for errChoice != nil {
			fmt.Println("Invalid command, please enter an integer.")
			fmt.Print("Enter your choice (1-6): ")
			fmt.Scanln(&choiceStr)
			choice, errChoice = strconv.Atoi(choiceStr)
		}

		// Validate the input
		if choice < 1 || choice > 6 {
			errMessage = fmt.Errorf("invalid command, please choose a valid option between 1 and 6")
		}

		// PlANT COMMAND
		if choice == 1 {
			// Ask the player what crop they want to plant
			cropName, err := AskWhatToPlant(&player)
			for err != nil {
				if err.Error() == "no crops are available to plant. Please restock your seeds" {
					fmt.Println(strings.ToUpper(err.Error()))
					// Return to menu
					break
				}
				// If there's any other error, print it and ask again
				fmt.Println(err)
				cropName, err = AskWhatToPlant(&player)
			}

			// Player has NO CROPS available, BACK TO THE MENU!!!
			if err != nil && err.Error() == "no crops are available to plant. Please restock your seeds" {
				continue
			}

			// Ask where to plant the crop
			row, col, err := player.AskWhereToPlant()
			for err != nil {
				fmt.Println(err)
				row, col, err = player.AskWhereToPlant()
			}

			// Plant the crop -- CROP OBJECT (YAY!)
			crop, err := getCropObject(cropName)
			if err != nil { // If crop is not one of available crops
				fmt.Println(err)
				return
			}

			// Handles error where cannot plant because there is
			// something already there and/or if the seed doesn't exist
			errPlot := player.PlantCrop(row, col, crop)
			if errPlot != nil {
				errMessage = errPlot
			}

		}

		// HARVEST COMMAND
		if choice == 2 {
			player.HarvestAll()
		}

		// REMOVE COMMAND
		if choice == 3 {
			var rowStr, colStr string

			// convert strings to integers
			fmt.Print("Enter the row: ")
			fmt.Scanln(&rowStr)
			rowInt, errRow := strconv.Atoi(rowStr)
			for errRow != nil {
				fmt.Println("Invalid Integer value, please enter an integer.")
				fmt.Print("Enter the row: ")
				fmt.Scanln(&rowStr)
				rowInt, errRow = strconv.Atoi(rowStr)
			}

			// convert strings to integers
			fmt.Print("Enter the col: ")
			fmt.Scanln(&colStr)
			colInt, errRow := strconv.Atoi(colStr)
			for errRow != nil {
				fmt.Println("Invalid Integer value, please enter an integer.")
				fmt.Print("Enter the col: ")
				fmt.Scanln(&colStr)
				rowInt, errRow = strconv.Atoi(colStr)
			}

			// Attempts to removes the Item
			errPlot := player.Plot.removeItem(rowInt, colInt) // returns error if no item at location
			if errPlot != nil {
				errMessage = errPlot
			}
		}

		// SHOP
		if choice == 4 {
			// keeps running until player exits shop
			inShop := player.StoreFront()
			for inShop != "Exit" {
				inShop = player.StoreFront()
			}
		}

		// END DAY
		if choice == 5 {
			// All crops become a day older
			player.Plot.updateCrops()
			// day increaes by 1
			player.Day += 1
		}

		// EXIT
		if choice == 6 {
			fmt.Println("Exiting the game...")
			os.Exit(0)
		}

		// Automatically grows the player's plot when they reach the required points
		player.updatePlot()
		// Saves the player data after each action
		SavePlayer(player)
		ClearConsole()

		// Displays error message from invalid actions
		if errMessage != nil {
			fmt.Println(strings.ToUpper(errMessage.Error()))

			// Waits 2 seconds before clearing console so user's can read error
			fmt.Println("Loading game...")
			time.Sleep(2 * time.Second)
			ClearConsole()
		}

		fmt.Println("Game saved!")

		// Display the updated player information
		fmt.Println("\n---Current Status---")
		player.DisplayInfo()

	}

}

// PRINTS AVAILABLE USER ACTIONS
func (p *Player) printMenu() {
	fmt.Println()
	fmt.Println("Current Plot:")
	fmt.Println()
	p.Plot.printGarden()
	fmt.Println()
	fmt.Println("GARDEN OPTIONS:")
	fmt.Println("(1 - PLANT) (2 - HARVEST) (3 - REMOVE) (4 - SHOP) (5 - END DAY) (6 - EXIT)")
}

// WHAT TO PLANT?
func AskWhatToPlant(player *Player) (string, error) {
	var cropName string
	var allSeedCount int

	// Check if there are available crops
	for _, count := range player.SeedStorage {
		if count > 0 {
			allSeedCount++
		}
	}

	// If there are available crops, display them
	if allSeedCount > 0 {
		fmt.Println("Available crops to plant:")
		for crop, count := range player.SeedStorage {
			if count > 0 {
				fmt.Printf("  %s: %d seed(s)\n", crop, count)
			}
		}
	} else {
		// If no crops are available, return the error
		return "", fmt.Errorf("no crops are available to plant. Please restock your seeds")
	}

	fmt.Println("Or type 'exit' to quit the game.")
	fmt.Scanln(&cropName)

	// Handle the exit case
	if strings.ToLower(cropName) == "exit" {
		fmt.Println("Exiting the game...")
		os.Exit(0)
	}

	cropName = strings.ToLower(cropName)

	// Verify the crop exists and player has seeds
	if count, exists := player.SeedStorage[cropName]; !exists {
		return "", fmt.Errorf("invalid crop name: %s", cropName)
	} else if count <= 0 {
		return "", fmt.Errorf("you don't have any %s seeds left", cropName)
	}

	return cropName, nil
}

// WHERE TO PLANT?
func (p *Player) AskWhereToPlant() (int, int, error) {
	var rowStr, colStr string
	fmt.Println("Enter the row and column (e.g., 0 1) where you want to plant the crop:")

	_, err := fmt.Scanln(&rowStr, &colStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input format, please enter row and column (e.g., 0 1)")
	}

	// convert strings to integers
	rowInt, err := strconv.Atoi(rowStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid row: %s is not a valid integer", rowStr)
	}

	colInt, err := strconv.Atoi(colStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid column: %s is not a valid integer", colStr)
	}

	// Ensure the input is within bounds
	if rowInt < 0 || rowInt >= p.Plot.Rows || colInt < 0 || colInt >= p.Plot.Cols {
		return 0, 0, fmt.Errorf("invalid row or column, must be between 0 and %d", p.Plot.Rows-1)
	}

	return rowInt, colInt, nil
}

// FOR NEW PLAYER
func HandleNewPlayer() Player {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	player := CreateNewPlayer(name, 5, 5) // START WITH 5X5
	SavePlayer(player)
	return player
}

// CONTINUE PLAYER
func HandleExistingPlayer() Player {
	fmt.Println("Enter your username:")
	var username string
	fmt.Scanln(&username)
	player, err := LoadPlayer(username)
	if err != nil {
		fmt.Println("Could not find player.")
		os.Exit(1) // Exit with status code 1 (indicating an error)
	}
	return player
}

// ClearConsole clears the terminal screen based on the operating system.
func ClearConsole() {
	// Checks the operating system
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// Unix system
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
