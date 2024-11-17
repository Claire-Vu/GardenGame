package main

import (
	"fmt"
	"os"
	"strings"
)

// MAIN GAME
func main() {
	fmt.Println("Welcome to the Farming Simulation Game!")
	fmt.Println("Are you a new player or continuing? (Type 'new' or 'continue')")
	var playerType string
	fmt.Scanln(&playerType) // taking input from user
	playerType = strings.ToLower(playerType)

	var player Player // variable of type Player to store player info

	if playerType == "new" {
		player = HandleNewPlayer()

	} else if playerType == "continue" {
		player = HandleExistingPlayer()

	} else {
		fmt.Println("Invalid input.")
		return
	}

	var gameRunning bool = true
	for gameRunning {
		// prints command options
		player.printMenu()

		// Ask what user wants to do
		var choice int
		fmt.Print("Enter your choice (1-6): ")
		fmt.Scanln(&choice)

		// Validate the input
		if choice < 1 || choice > 6 {
			fmt.Println("Invalid command. Please choose a valid option between 1 and 6.")
		}

		// PlANT COMMAND
		if choice == 1 {
			// Ask the player what crop they want to plant
			cropName, symbol, err := AskWhatToPlant(&player)
			for err != nil {
				fmt.Println(err)
				cropName, symbol, err = AskWhatToPlant(&player)
			}
			// Ask where to plant the crop
			row, col, err := AskWhereToPlant()
			for err != nil {
				fmt.Println(err)
				row, col, err = AskWhereToPlant()
			}
			// Plant the crop
			player.PlantCrop(row, col, Crop{Name: cropName, Symbol: symbol, FullyGrown: false})

		}
		// HARVEST COMMAND
		if choice == 2 {
			player.HarvestAll()
		}
		// REMOVE COMMAND
		if choice == 3 {
			var row, col int
			fmt.Print("Enter the row: ")
			fmt.Scan(&row)
			fmt.Print("Enter the col: ")
			fmt.Scan(&col)
			player.Plot.removeItem(row, col)
		}
		// SHOP
		if choice == 4 {
			// GO TO SHOP (PRINT SHOP MENU AND COMMANDS)
		}
		// END DAY
		if choice == 5 {
			player.Plot.updateCrops()
			player.Day += 1
			gameRunning = false

		}
		// EXIT
		if choice == 6 {
			fmt.Println("Exiting the game...")
			os.Exit(0)
		}

		// Saves the player data after each action
		SavePlayer(player)

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

// WHAT TO PLANT? - This will be Elaine's part about the store.
func AskWhatToPlant(player *Player) (string, string, error) {
	var cropName string
	fmt.Println("What crop would you like to plant?")
	fmt.Println("Available crops: carrot, potato, corn, pumpkin, garlic")
	fmt.Println("Or type 'exit' to quit the game.")
	fmt.Scanln(&cropName)

	// Handle the exit case
	if strings.ToLower(cropName) == "exit" {
		fmt.Println("Exiting the game...")
		os.Exit(0)
	}

	cropName = strings.ToLower(cropName)

	// Map of available crops and their symbols
	cropData := map[string]struct {
		symbol string
	}{
		"carrot":  {"🥕"},
		"potato":  {"🥔"},
		"corn":    {"🌽"},
		"pumpkin": {"🎃"},
		"garlic":  {"🧄"},
	}

	// Validate crop choice
	data, exists := cropData[cropName]
	if !exists {
		return "", "", fmt.Errorf("Invalid crop choice")
	}

	if player.SeedStorage[cropName] <= 0 {
		return "", "", fmt.Errorf("You don't have any %s seeds left.", cropName)
	}

	return cropName, data.symbol, nil
}

// WHERE TO PLANT?
func AskWhereToPlant() (int, int, error) {
	var row, col int
	fmt.Println("Enter the row and column (e.g., 0 1) where you want to plant the crop:")
	fmt.Scanln(&row, &col)

	// Ensure the input is within bounds
	if row < 0 || row > 4 || col < 0 || col > 4 {
		return 0, 0, fmt.Errorf("invalid row or column, must be between 0 and 4")
	}

	return row, col, nil
}

// FOR NEW PLAYER
func HandleNewPlayer() Player {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	player := CreateNewPlayer(name, 5, 5) // START WITH 5X5
	SavePlayer(player)
	fmt.Println("\n---Current Status---")
	player.DisplayInfo()
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
	fmt.Printf("\nWelcome back, %s!\n", player.Username)
	fmt.Println("\n---Current Status---")
	player.DisplayInfo()
	return player
}
