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

	// Ask the player what crop they want to plant
	cropName, symbol, err := AskWhatToPlant()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ask where to plant the crop
	row, col, err := AskWhereToPlant()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Plant the crop
	player.PlantCrop(row, col, Crop{Name: cropName, Symbol: symbol, FullyGrown: false})

	// Save the player data after planting
	SavePlayer(player)

	// Display the updated player information
	fmt.Println("\n---Current Resources---")
	for resource, amount := range player.Resources {
		fmt.Printf("%s: %d\n", resource, amount)
	}

	// Display the player's garden (plot)
	fmt.Println("\n---Current Garden---")
	player.Plot.printGarden()

	fmt.Println("Game saved!")

	// Garden Simulation
	// garden := CreatePlot(5, 5)
	// potato := Crop{Name: "Flower", Symbol: "🥔", fullyGrown: true}
	// mango := Crop{Name: "Mango", Symbol: "🥭", fullyGrown: false}
	// mango2 := Crop{Name: "Mango", Symbol: "🥭", fullyGrown: true}
	// garden.Plant(0, 0, &potato)
	// garden.Plant(2, 2, &mango)
	// garden.Plant(2, 3, &mango2)
	// garden.printGarden()
}

// WHAT TO PLANT? - This will be Elaine's part about the store.
func AskWhatToPlant() (string, string, error) {
	var cropName string
	fmt.Println("What crop would you like to plant?")
	fmt.Println("Available crops: Carrot, Potato, Flower")
	fmt.Scanln(&cropName)

	cropName = strings.ToLower(cropName)
	// Validate crop choice
	var symbol string
	switch cropName {
	case "carrot":
		symbol = "🥕"
	case "potato":
		symbol = "🥔"
	case "flower":
		symbol = "🌸"
	default:
		return "", "", fmt.Errorf("invalid crop choice")
	}
	return cropName, symbol, nil
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
	return player
}
