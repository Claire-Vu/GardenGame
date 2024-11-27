package main

import (
	"bufio"
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

		// Scanner to prompt for input
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt the user for action
		var choiceStr string
		var choice int
		fmt.Print("Enter your choice (1-6): ")
		if scanner.Scan() {
			choiceStr = strings.TrimSpace(scanner.Text())
			inputVal, errChoice := strconv.Atoi(choiceStr)
			if errChoice != nil {
				choice = 0
			} else {
				choice = inputVal
			}
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
				fmt.Println()
				fmt.Println(strings.ToUpper(err.Error()))
				row, col, err = player.AskWhereToPlant()
			}

			// Plant the crop -- CROP OBJECT (YAY!)
			crop, err := getCropObject(cropName)
			if err != nil { // If crop is not one of available crops
				fmt.Println()
				fmt.Println(strings.ToUpper(err.Error()))
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
			var rowInt, colInt int

			// Validates row input is an integer
			fmt.Print("Enter the row: ")
			if scanner.Scan() {
				rowStr = strings.TrimSpace(scanner.Text())
				// If the input cannot be converted to a string then set rowInt
				// to -1 (invalid input) and set error message accordingly
				inputVal, errStr := strconv.Atoi(rowStr)
				if errStr != nil {
					rowInt = -1
					errMessage = fmt.Errorf("invalid row input")
					// If valid prompt for col
				} else {
					rowInt = inputVal
					// Validates col input is an integer
					fmt.Print("Enter the col: ")
					if scanner.Scan() {
						colStr = strings.TrimSpace(scanner.Text())
						// If the input cannot be converted to a string then set colInt
						// to -1 (invalid input)
						inputVal, errStr := strconv.Atoi(colStr)
						if errStr != nil {
							colInt = -1
							errMessage = fmt.Errorf("invalid column input")
						} else {
							colInt = inputVal
						}
					}
				}
			}

			// If  a valid input for row or col then attempts to removeItem
			if rowInt != -1 && colInt != -1 {
				// Attempts to removes the Item
				errPlot := player.Plot.removeItem(rowInt, colInt) // returns error if no item at location
				if errPlot != nil {
					errMessage = errPlot
				}
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

	// need new scanner for reading user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cropName = scanner.Text()

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
	// Need a new scanner - this take the whole string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the row and column (e.g., 0 1) where you want to plant the crop:")
	scanner.Scan()
	input := scanner.Text()
	values := strings.Fields(input) // need to split this into row and col

	// Input should have exactly two values
	if len(values) != 2 {
		return 0, 0, fmt.Errorf("invalid input format, please enter row and column (e.g., 0 1)")
	}

	// Convert strings to integers
	rowInt, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid row: %s is not a valid integer", values[0])
	}

	colInt, err := strconv.Atoi(values[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid column: %s is not a valid integer", values[1])
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
